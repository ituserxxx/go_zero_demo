package hello

import (
	"context"

	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RpcListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRpcListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcListLogic {
	return &RpcListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RpcListLogic) RpcList(req *types.HelloListReq) (resp *types.HelloRpcListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
