package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"test_demo/api/internal/svc"
	"test_demo/api/internal/types"
	"test_demo/common/constants"
	"test_demo/db/model"
	"time"
)

type UserOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrderLogic {
	return &UserOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOrderLogic) UserOrder(req *types.UserOrderRequest) (resp *types.UserOrderResponse, err error) {
	// todo: add your logic here and delete this line
	//  记录缓存
	l.svcCtx.DemoRedis.Set(l.ctx, fmt.Sprintf("%d_%d", req.UserId, req.ProductId), req.Action, time.Second*100)
	//  取出缓存并打印
	cacheValue, _ := l.svcCtx.DemoRedis.Get(l.ctx, fmt.Sprintf("%d_%d", req.UserId, req.ProductId)).Result()
	fmt.Println("缓存值为: ", cacheValue)
	// 查询用户和商品是否存在
	var userRow model.User
	if findUserErr := l.svcCtx.DemoDB.Where("user_id = ?", req.UserId).First(&userRow).Error; findUserErr != nil {
		l.Logger.Errorf("查询用户失败, err: %v", findUserErr)
		return nil, fmt.Errorf("用户不存在")
	}
	var productRow model.Product
	if findProductErr := l.svcCtx.DemoDB.Where("product_id = ?", req.ProductId).First(&productRow).Error; findProductErr != nil {
		l.Logger.Errorf("查询商品失败, err: %v", findProductErr)
		return nil, fmt.Errorf("商品不存在")
	}
	//  判断库存是否足够
	if productRow.StockQuantity < int64(req.OrderAmount) {
		return nil, fmt.Errorf("库存不足")
	}
	switch req.Action {
	case constants.PLACE_ORDER:
		//  下单，开启事务
		tx := l.svcCtx.DemoDB.Begin()
		newOrderValue := model.OrderTable{
			UserID:      int64(req.UserId),
			ProductID:   int64(req.ProductId),
			OrderTime:   time.Now().Unix(),
			OrderStatus: 1,
			OrderAmount: int64(req.OrderAmount),
		}
		if createOrderErr := tx.Create(&newOrderValue).Error; createOrderErr != nil {
			tx.Rollback()
			l.Logger.Errorf("创建订单失败, err: %v", createOrderErr)
			return nil, fmt.Errorf("创建订单失败")
		}
		//  更新商品库存
		if updateProductErr := tx.Model(model.Product{}).
			Where("product_id = ?", req.ProductId).
			Update("stock_quantity", productRow.StockQuantity-int64(req.OrderAmount)).
			Error; updateProductErr != nil {
			tx.Rollback()
			l.Logger.Errorf("更新商品库存失败, err: %v", updateProductErr)
			return nil, fmt.Errorf("更新商品库存失败")
		}
		tx.Commit()
		return &types.UserOrderResponse{
			Code: 200,
			Msg:  "下单成功",
		}, nil
		return resp, nil
	case constants.CANCEL_ORDER:
		//	查询订单
		var orderRow model.OrderTable
		if findOrderErr := l.svcCtx.DemoDB.Where("order_id = ?", req.OrderId).First(&orderRow).Error; findOrderErr != nil {
			l.Logger.Errorf("查询订单失败, err: %v", findOrderErr)
			return nil, fmt.Errorf("订单不存在")
		}
		//  取消订单，开启事务
		tx := l.svcCtx.DemoDB.Begin()
		if updateOrderErr := tx.Model(model.OrderTable{}).
			Where("order_id = ?", req.OrderId).
			Update("order_status", 2).
			Error; updateOrderErr != nil {
			tx.Rollback()
			l.Logger.Errorf("更新订单状态失败, err: %v", updateOrderErr)
			return nil, fmt.Errorf("更新订单状态失败")
		}
	}
	return
}
