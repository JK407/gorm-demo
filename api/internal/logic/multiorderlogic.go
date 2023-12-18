package logic

import (
	"context"
	"test_demo/api/internal/svc"
	"test_demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMultiOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiOrderLogic {
	return &MultiOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MultiOrderLogic) MultiOrder(req *types.MultiOrderRequest) (resp *types.MultiOrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
