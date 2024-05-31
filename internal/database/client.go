package database

import "github.com/cufee/aftermath/internal/database/prisma/db"

type Client struct {
	prisma *db.PrismaClient
}

func NewClient() (*Client, error) {
	prisma := db.NewClient()
	err := prisma.Connect()
	if err != nil {
		return nil, err
	}
	return &Client{prisma: prisma}, nil
}
