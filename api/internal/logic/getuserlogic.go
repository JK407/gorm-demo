package logic

import (
	"context"
	"fmt"
	"test_demo/api/internal/svc"
	"test_demo/api/internal/types"
	"test_demo/db/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	// todo: add your logic here and delete this line
	var userDBRow model.User
	if findUserErr := l.svcCtx.DemoDB.Where("user_id = ?", req.UserId).First(&userDBRow).Error; findUserErr != nil {
		l.Logger.Errorf("查询用户失败, err: %v", findUserErr)
		return nil, findUserErr
	}
	fmt.Println("DB查询到的用户信息为: ", userDBRow)

	//u := query.User
	//user, queryUserErr := u.WithContext(l.ctx).First()
	//if queryUserErr != nil {
	//	l.Logger.Errorf("查询用户失败, err: %v", queryUserErr)
	//	return nil, queryUserErr
	//
	//}
	//fmt.Println("query查询到的用户信息为: ", user)
	return
}
