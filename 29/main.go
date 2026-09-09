package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // CORS bypass
	},
}

// ساختار Hub برای مدیریت تمام کلاینت‌های متصل
type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) Broadcast(message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for conn := range h.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error broadcasting to a client, closing:", err)
			conn.Close()
			delete(h.clients, conn)
		}
	}
}

func (h *Hub) Register(conn *websocket.Conn) {
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
	fmt.Printf("New tab connected! Total connected tabs: %d\n", len(h.clients))
}

func (h *Hub) Unregister(conn *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, conn)
	h.mu.Unlock()
	fmt.Printf("Tab disconnected. Total remaining tabs: %d\n", len(h.clients))
}

func main() {
	r := gin.Default()
	hub := NewHub()

	// مسیر وب‌سوکت در جین
	r.GET("/ws", func(c *gin.Context) {
		// ۱. تبدیل HTTP به WebSocket با استفاده از Gin Context
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Upgrade error:", err)
			return
		}

		// ۲. ثبت اتصال
		hub.Register(conn)

		// ۳. پاکسازی هنگام قطع اتصال
		defer func() {
			hub.Unregister(conn)
			conn.Close()
		}()

		// ۴. حلقه دریافت پیام (Blocking)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				// قطع اتصال عادی یا بستن تب
				break
			}

			fmt.Printf("Message received from a tab: %s\n", message)

			// ۵. پخش پیام به تمام تب‌ها
			hub.Broadcast(message)
		}
	})

	fmt.Println("Gin Broadcast WebSocket Server running on :8080")
	r.Run(":8080")
}
