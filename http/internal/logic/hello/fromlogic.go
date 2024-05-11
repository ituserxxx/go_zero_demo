package hello

import (
	"context"

	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FromLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFromLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FromLogic {
	return &FromLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FromLogic) From(req *types.FromReq) (resp *types.FromRes, err error) {
	// todo: add your logic here and delete this line
	resp = &types.FromRes{Message: req.Name}
	return
}
