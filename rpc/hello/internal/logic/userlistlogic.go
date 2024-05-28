package logic

import (
	"context"
	db2 "go_zero_demo/db"

	"go_zero_demo/rpc/hello/hello"
	"go_zero_demo/rpc/hello/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *hello.UserListReq) (*hello.UserListResp, error) {
	var resp = &hello.UserListResp{}
	orm := db2.DB.Model(db2.Hello{})
	orm.Count(&resp.Total)
	orm.Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&resp.List)
	return resp, nil
}
