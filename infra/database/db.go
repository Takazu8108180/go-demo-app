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
	Client *ent.Client
}

func NewDB() (*DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// Configure ent
	dbConn := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("USERNAME"), os.Getenv("DATABASE"), os.Getenv("USERPASS"))
	client, err := ent.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: client,
	}, nil
}

// Run the auto migration tool.
func (d *DB) AutoMigrate() error {
	if err := d.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return nil
}

func (d *DB) CloseDB() error {
	d.Client.Close()
	return nil
}
