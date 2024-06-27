// go:build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cufee/aftermath/internal/database"
)

func main() {
	// db, err := newDatabaseClientFromEnv()
	// if err != nil {
	// 	panic(err)
	// }
	// _ = db

	err := ScrapeNewsImages()
	if err != nil {
		panic(err)
	}
}

func newDatabaseClientFromEnv() (database.Client, error) {
	err := os.MkdirAll(os.Getenv("DATABASE_PATH"), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("os#MkdirAll failed %w", err)
	}

	client, err := database.NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")))
	if err != nil {

		return nil, fmt.Errorf("database#NewClient failed %w", err)
	}

	return client, nil
}
