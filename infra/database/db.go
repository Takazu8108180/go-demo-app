package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/takazu8108180/go-demo-app/ent"
)

type DB struct {
	db *ent.Client
}

func NewDB() *DB {
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

	return &DB{
		db: client,
	}
}

// Run the auto migration tool.
func (d *DB) AutoMigrate() error {
	if err := d.db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return nil
}

func (d *DB) CloseDB() error {
	d.db.Close()
	return nil
}
