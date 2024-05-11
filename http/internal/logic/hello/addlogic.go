package hello

import (
	"context"
	db2 "go_zero_demo/db"
	"time"

	"go_zero_demo/http/internal/svc"
	"go_zero_demo/http/internal/types"

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
	newHello := db2.Hello{
		Name:       req.Name,
		CreateTime: time.Now().Unix(),
	}
	err = db2.DB.Model(db2.Hello{}).Create(&newHello).Error

	resp.Id = newHello.Id
	return
}
