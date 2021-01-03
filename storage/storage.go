package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/CristianPadilla/BDwithGoPostgreSQL/pkg/product"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

//NewPostgresDB for make connection only once
func NewPostgresDB() {
	once.Do(func() {
		//here function is gonna be executed once
		var err error
		connStr := "postgres://postgres:root@localhost:5432/godb?sslmode=disable"
		//arguments are: nombre del driver and cadena de conexion(url), buscar en documentacion del driver
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("can't connect to database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)

		}
		fmt.Println("connected to postgres")
	})
}

//Pool returns a unique instance of db
func Pool() *sql.DB {
	return db
}

//StringToNull serves to verify if a string is empty or null (it's different) manejo de nulos
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

//StringToNull serves to verify if a string is empty or null (it's different) manejo de nulos
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

//scanner created like helper for scan a row of product
type scanner interface {
	Scan(dest ...interface{}) error
}

// to consult for a row at products table
func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observatiosNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observatiosNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return nil, err
	}
	m.Observations = observatiosNull.String
	m.UpdatedAt = updatedAtNull.Time
	return m, nil
}
