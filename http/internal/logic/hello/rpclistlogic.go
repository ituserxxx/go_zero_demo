package hello

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"
	"go_zero_demo/rpc/hello/helloclient"
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
	resp = &types.HelloRpcListResp{
		List:  make([]*types.HelloListItem, 0),
		Total: 0,
	}
	arr, err := l.svcCtx.HelloRpc.UserList(l.ctx, &helloclient.UserListReq{
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	for _, item := range arr.List {
		var info *types.HelloListItem
		_ = copier.Copy(&info, item)
		resp.List = append(resp.List, info)
	}
	resp.Total = arr.Total
	return
}
