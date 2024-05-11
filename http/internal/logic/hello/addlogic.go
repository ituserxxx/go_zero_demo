package hello

import (
	"context"
	"http/db"
	"time"

	"http/internal/svc"
	"http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddHelloReq) (resp *types.AddHelloRes, err error) {
	resp = &types.AddHelloRes{}
	newHello := db.Hello{
		Name:       req.Name,
		CreateTime: time.Now().Unix(),
	}
	err = db.DB.Model(db.Hello{}).Create(&newHello).Error

	resp.Id = newHello.Id
	return
}
