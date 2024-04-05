package controllers

import "gorm.io/gorm"

type Server struct {
	db *gorm.DB
}

// NewServer creates a new server
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}
