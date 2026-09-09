package main

// docker run -d --name jaeger -e COLLECTOR_OTLP_ENABLED=true -p 16686:16686 -p 4318:4318 jaegertracing/all-in-one:latest
// http://localhost:16686
// curl http://localhost:8080/products/1
// curl http://localhost:8080/checkout 

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gorm.io/gorm"
)

// راه اندازی ارسال کننده داده ها به Jaeger
func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-api-service"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ProductRepository struct{ db *gorm.DB }

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	return &product, err
}

type ProductService struct{ repo *ProductRepository }

func (s *ProductService) GetProduct(ctx context.Context, id string) (*Product, error) {
	tr := otel.Tracer("service-tracer")
	ctx, span := tr.Start(ctx, "ProductService.GetProduct")
	defer span.End()

	time.Sleep(50 * time.Millisecond)

	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// اگر دیتابیس ارور داد، ثبتش می‌کنیم
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return product, nil
}

// --- متد جدید با احتمال ۵۰٪ ارور شبیه‌سازی شده ---
func (s *ProductService) Checkout(ctx context.Context) error {
	tr := otel.Tracer("service-tracer")
	ctx, span := tr.Start(ctx, "ProductService.Checkout")
	defer span.End()

	// ۱. ابتدا یک کوئری سالم به دیتابیس می‌زنیم
	_, _ = s.repo.GetByID(ctx, "1")

	// ۲. یک معطلی شبیه‌سازی شده
	time.Sleep(80 * time.Millisecond)

	// ۳. شبیه‌سازی ارور رندوم (۵۰ درصد احتمال شکست)
	if rand.Intn(2) == 0 {
		err := errors.New("payment gateway timeout: insufficient funds or network error")
		
		// ثبت ارور روی Span برای قرمز شدن در Jaeger
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		
		return err
	}

	return nil
}

type ProductHandler struct{ service *ProductService }

func (h *ProductHandler) GetProductHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	product, err := h.service.GetProduct(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// --- هندلر جدید برای تست ارور ---
func (h *ProductHandler) CheckoutHandler(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.service.Checkout(ctx)
	if err != nil {
		// ارسال خطای 500 به کلاینت
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "checkout completed successfully",
	})
}

func main() {
	ctx := context.Background()

	// اتصال به Jaeger
	tp, err := initTracer(ctx)
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(ctx)

	// دیتابیس sqlite
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.Use(otelgorm.NewPlugin())
	db.AutoMigrate(&Product{})
	db.Create(&Product{ID: 1, Name: "Laptop", Price: 1200})

	repo := &ProductRepository{db: db}
	service := &ProductService{repo: repo}
	handler := &ProductHandler{service: service}

	r := gin.Default()
	r.Use(otelgin.Middleware("my-api-service"))

	r.GET("/products/:id", handler.GetProductHandler)
	r.GET("/checkout", handler.CheckoutHandler) // <--- مسیر جدید

	r.Run(":8080")
}