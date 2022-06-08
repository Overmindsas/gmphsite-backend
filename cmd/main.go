package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Overmindsas/gmphsite-backend/internal/config"
	"github.com/Overmindsas/gmphsite-backend/internal/handlers"
	"github.com/Overmindsas/gmphsite-backend/internal/services"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()
	db, err := NewDb(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	srv := services.NewUserService(db)
	handler := handlers.NewUserHandler(srv)

	go func() {
		if err := handler.SrartServer(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

}

func NewDb(host, port, username, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
