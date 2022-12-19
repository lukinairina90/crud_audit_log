package repository

import (
	"context"
	"github.com/lukinairina90/crud_audit_log/pkg/domain/audit"
	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Database
}

func NewAudit(db *mongo.Database) *Audit {
	return &Audit{db: db}
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := r.db.Collection("Logs").InsertOne(ctx, item)

	return err
}
