package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/spf13/viper"
	"github.com/wneessen/go-mail"
)

// --- 1. CONFIGURATION ---
const ENV_FILE = `C:\Users\MSI\Pictures\.env`

type EmailConfig struct {
	EmailFrom string `mapstructure:"email_from"`
	EmailTo   string `mapstructure:"email_to"`
	Password  string `mapstructure:"email_password"`
	Host      string `mapstructure:"smtp_host"`
	Port      int    `mapstructure:"smtp_port"`
}

func loadConfig() (*EmailConfig, error) {
	v := viper.New()
	v.AddConfigPath(ENV_FILE + `\..`)
	v.SetConfigType("env")
	v.SetConfigName(".env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read .env file: %w", err)
	}

	var cfg EmailConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

// --- 2. MAILER INTERFACE & IMPLEMENTATION ---

type Message struct {
	To       []string
	Subject  string
	TextBody string
	HTMLBody string
}

type Sender interface {
	Send(ctx context.Context, msg *Message) error
}

type GomailSender struct {
	cfg *EmailConfig
}

func NewGomailSender(cfg *EmailConfig) Sender {
	return &GomailSender{cfg: cfg}
}

func (s *GomailSender) Send(ctx context.Context, msg *Message) error {
	m := mail.NewMsg()

	if err := m.From(s.cfg.EmailFrom); err != nil {
		return fmt.Errorf("failed to set from address: %w", err)
	}

	if err := m.To(msg.To...); err != nil {
		return fmt.Errorf("failed to set recipient address: %w", err)
	}

	m.Subject(msg.Subject)

	if msg.TextBody != "" {
		m.SetBodyString(mail.TypeTextPlain, msg.TextBody)
	}

	if msg.HTMLBody != "" {
		m.AddAlternativeString(mail.TypeTextHTML, msg.HTMLBody)
	}

	m.EmbedFile("gomotorcycle.svg")
	// port, err := strconv.Atoi(s.cfg.Port)
	// if err != nil {
	// 	return fmt.Errorf("invalid smtp port '%s': %w", s.cfg.Port, err)
	// }
	port := s.cfg.Port

	client, err := mail.NewClient(
		s.cfg.Host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(s.cfg.EmailFrom),
		mail.WithPassword(s.cfg.Password),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return fmt.Errorf("failed to create mail client: %w", err)
	}

	if err := client.DialAndSendWithContext(ctx, m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// --- 3. DYNAMIC TEMPLATE & DATA ---

type OTPData struct {
	AppName          string
	OTPCode          string
	OTPExpireMinutes int
}

// --- 4. MAIN ENTRYPOINT ---

func main() {
	// 1. Load config from .env
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load environment: %v", err)
	}

	// 2. Initialize sender
	sender := NewGomailSender(cfg)

	// 3. Parse HTML template
	tmpl, err := template.ParseFiles("otp2.html")
	if err != nil {
		log.Fatalf("Failed to parse HTML template: %v", err)
	}

	otpData := OTPData{
		AppName:          "Cinema Tickets App",
		OTPCode:          "739401",
		OTPExpireMinutes: 5,
	}

	var htmlBuffer bytes.Buffer
	if err := tmpl.Execute(&htmlBuffer, otpData); err != nil {
		log.Fatalf("Failed to render HTML template: %v", err)
	}

	// 5. Construct email message
	targetEmail := cfg.EmailTo // Replace with recipient email
	msg := &Message{
		To:       []string{targetEmail},
		Subject:  fmt.Sprintf("%s is your verification code", otpData.OTPCode),
		TextBody: fmt.Sprintf("Your OTP code for %s is %s. It expires in %d minutes.", otpData.AppName, otpData.OTPCode, otpData.OTPExpireMinutes),
		HTMLBody: htmlBuffer.String(),
	}

	// 6. Send email
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Printf("Sending OTP email to %s...", targetEmail)
	if err := sender.Send(ctx, msg); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("OTP Email sent successfully!")
	time.Sleep(time.Second * 10)
}
