package main

import (
	"context"
	"fmt"
	"github.com/lukinairina90/crud_audit_log/internal/config"
	"github.com/lukinairina90/crud_audit_log/internal/repository"
	"github.com/lukinairina90/crud_audit_log/internal/server"
	"github.com/lukinairina90/crud_audit_log/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Сначала:
// go get google.golang.org/grpc
// go get google.golang.org/protobuf

// потом описываем audit.proto

// генерим audit.pb.go:
// protoc --go_out=pkg --go-grpc_out=../pkg/domain/audit proto/audit.proto

// потом устанавливаем драйвер монги:
// go get go.mongodb.org/mongo-driver/mongo

// Для конфигов скачиваем библиотеку:
// go get github.com/kelseyhightower/envconfig
// и делаем импорт в файле config.go:
// import "github.com/kelseyhightower/envconfig"

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
