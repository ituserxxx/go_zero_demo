package hello

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"http/db"
	"http/internal/svc"
	"http/internal/types"
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
	orm := db.DB.Model(db.Hello{})
	orm.Count(&resp.Total)
	orm.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize).Find(&resp.List)
	return
}
