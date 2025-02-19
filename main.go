package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/infra/database"
	"github.com/takazu8108180/go-demo-app/infra/routers"
)

func main() {
	// Setup DB
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	defer func(db *database.DB) {
		db.CloseDB()
	}(db)

	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("failed auto migration: %v", err)
	}

	// Setup Router
	router := routers.NewRouter(gin.Default(), db)

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
