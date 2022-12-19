package server

import (
	"context"
	"github.com/lukinairina90/crud_audit_log/pkg/domain/audit"
)

//type UnsafeAuditServiceServer interface {
//	mustEmbedUnimplementedAuditServiceServer()
//}

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	service AuditService
	audit.UnimplementedAuditServiceServer
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	err := h.service.Insert(ctx, req)

	return &audit.Empty{}, err
}

//func (h *AuditServer) mustEmbedUnimplementedAuditServiceServer() {}
