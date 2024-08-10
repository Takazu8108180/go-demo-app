package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/ent"
	h "github.com/takazu8108180/go-demo-app/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error Don't Read Env: %v", err)
	}

	// Configure ent
	dbConn := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("USERNAME"), os.Getenv("DATABASE"), os.Getenv("USERPASS"))
	client, err := ent.Open("postgres", dbConn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()

	r.GET("/", h.HelloHandler)
	r.POST("/article", h.PostArticleHandler)
	r.GET("/article/list", h.GetArticleListHandler)
	r.GET("/article/detail/:id", h.GetArticleDetailHandler)
	r.POST("/article/nice", h.PostNiceHandler)
	r.POST("/comment", h.PostCommentHandler)

	log.Println("Server start at port 8080")

	r.Run(":8080")
}
