package hello

import (
	"context"
	db2 "go_zero_demo/db"
	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.DelHelloReq) (resp *types.DelHelloRes, err error) {
	resp = &types.DelHelloRes{}
	err = db2.DB.Model(db2.Hello{}).Where("id", req.Id).Delete(db2.Hello{}).Error

	return
}
