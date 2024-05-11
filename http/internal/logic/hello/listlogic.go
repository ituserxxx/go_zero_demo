package hello

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	db2 "go_zero_demo/db"
	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.HelloListReq) (resp *types.HelloListResp, err error) {
	resp = &types.HelloListResp{
		List:  make([]*types.HelloListItem, 0),
		Total: 0,
	}
	orm := db2.DB.Model(db2.Hello{})
	orm.Count(&resp.Total)
	orm.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&resp.List)
	return
}
