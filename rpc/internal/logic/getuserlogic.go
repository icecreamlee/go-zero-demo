package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/rpc/internal/svc"
	"go-zero-demo/rpc/types/user"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	return &user.UserResponse{
		Id:   "1",
		Name: "test",
	}, nil
}
