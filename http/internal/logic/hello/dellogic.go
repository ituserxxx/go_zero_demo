package hello

import (
	"context"
	"http/db"

	"http/internal/svc"
	"http/internal/types"

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
	err = db.DB.Model(db.Hello{}).Where("id", req.Id).Delete(db.Hello{}).Error

	return
}
