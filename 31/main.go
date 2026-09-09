package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/glebarez/sqlite" // Pure-Go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	// "github.com/libtnb/sqlite" // Pure-Go SQLite driver, checkout https://github.com/libtnb/sqlite for details
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQLite database: %v", err)
	}

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Failed to create Casbin adapter: %v", err)
	}

	m, err := model.NewModelFromFile("model.conf")
	if err != nil {
		log.Fatalf("Failed to parse model string: %v", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatalf("Failed to load policy from DB: %v", err)
	}

	// _, _ = e.AddPolicy("admin", "main.exe", "iran", "in", "allow")
	// _, _ = e.AddPolicy("admin", "web.exe", "iran", "out", "allow")
	// _, _ = e.AddPolicy("10.10.10.10", "main.exe", "iran", "in", "deny")

	// _, _ = e.AddRoleForUser("97.97.97.97", "admin")
	// _, _ = e.AddRoleForUser("10.10.10.10", "admin")

	allowed, err := e.Enforce("97.97.97.97", "main.exe", "iran", "in")
	if err != nil {
		log.Fatalf("Enforce error: %v", err)
	}

	fmt.Println("==========================================")
	if allowed {
		fmt.Println("Result: ALLOWED (Fetched and evaluated via DB policy)")
	} else {
		fmt.Println("Result: DENIED")
	}
	fmt.Println("==========================================")
}
