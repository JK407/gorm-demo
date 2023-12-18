package logic

import (
	"context"
	"fmt"
	"test_demo/db/model"

	"test_demo/rpc/internal/svc"
	"test_demo/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *order.GetUserRequest) (*order.GetUserResponse, error) {
	// todo: add your logic here and delete this line
	var userDBRow model.User
	if findUserErr := l.svcCtx.DemoDB.Where("user_id = ?", in.UserId).First(&userDBRow).Error; findUserErr != nil {
		l.Logger.Errorf("查询用户失败, err: %v", findUserErr)
		return nil, findUserErr
	}
	fmt.Println("DB查询到的用户信息为: ", userDBRow)
	return &order.GetUserResponse{}, nil
}
