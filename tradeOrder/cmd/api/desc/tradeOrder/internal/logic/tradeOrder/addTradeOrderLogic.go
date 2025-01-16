package tradeOrder

import (
	"QMall/common"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/api/desc/tradeOrder/internal/types"
	"QMall/tradeOrder/cmd/domain/model"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/tradeorder"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type AddTradeOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddTradeOrderLogic 生成订单
func NewAddTradeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTradeOrderLogic {
	return &AddTradeOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTradeOrderLogic) AddTradeOrder(req *types.AddTradeOrderReq) (resp *types.AddTradeOrderResp, err error) {

	//cartsByUserId, err := l.svcCtx.ShoppingCartRPC.GetCartsByUserId(l.ctx, &shoppingcart.FindCartsByUserIdReq{
	//	UserId: req.UserId,
	//})
	//carts := cartsByUserId.Carts
	//var pbCarts []*pb.ShoppingCartSimple
	//for i := 0; i < len(carts); i++ {
	//	pbCarts = append(pbCarts, pb.ShoppingCartsConvert(carts[i]))
	//}
	carts, err := l.svcCtx.ShoppingCartRPC.ShowDetailShoppingCarts(l.ctx, &shoppingcart.ShowDetailShoppingCartsReq{
		UserId: req.UserId,
	})
	l.Info(fmt.Sprintf("lens>:%v", len(carts.ShoppingCartsProductInfo)))
	if err != nil || len(carts.ShoppingCartsProductInfo) <= 0 {
		fmt.Printf("当前用户ID：%v的购物车为空！错误见原因：%v\n", req.UserId, err)
		return
	}
	l.Info(fmt.Printf("当前用户ID：%v，购物车信息如下：%v\n", req.UserId, carts))
	// 计算购物车订单总价
	totalAmount := 0.0
	priceByUserId, err := l.svcCtx.ShoppingCartRPC.GetTotalPriceByUserId(l.ctx, &shoppingcart.GetTotalPriceByUserIdReq{
		UserId: req.UserId,
	})
	if err != nil {
		l.Error(fmt.Printf("当前用户ID：%v的购物车总价为 0 ！错误见原因：%v\n", req.UserId, err))
		return
	}
	totalAmount = priceByUserId.TotalPrice
	l.Info(fmt.Printf("当前用户ID：%v的购物车总价为 %v.\n", req.UserId, totalAmount))

	// 计算购物车订单总价
	shippingAmount := 0.0
	discountAmount := 0.0
	order := &model.TradeOrder{
		TotalAmount:    totalAmount,
		ShippingAmount: shippingAmount,
		DiscountAmount: discountAmount,
		PayAmount:      totalAmount + shippingAmount - discountAmount,
		OrderSource:    req.OrderSource,
		UserMessage:    req.UserMessage,
	}
	addTradeOrder := &tradeorder.AddTradeOrderReq{
		TradeOrder: pb.ModelTradeOrderConvertPb(order),
	}

	tradeOrder, err := l.svcCtx.TradeOrderRpcConf.AddTradeOrder(l.ctx, addTradeOrder)

	if err != nil {
		fmt.Printf("用户订单生成失败！原因见：%v\n", err)
		return
	}

	// 订单生成后需要清空购物车
	if tradeOrder.TradeOrder.OrderNo != "" {
		_, err = l.svcCtx.ShoppingCartRPC.DeleteCartsByUserId(l.ctx, &shoppingcart.DeleteCartsByUserIdReq{
			UserId: req.UserId,
		})
		if err != nil {
			fmt.Printf("用户id：%v的购物车清空失败，失败原因见：%v\n", req.UserId, err)
			return
		}
	}

	tradeOrderProduct := make([]*pb.ProductOrder, len(carts.ShoppingCartsProductInfo))
	for i := 0; i < len(tradeOrderProduct); i++ {
		tradeOrderProduct[i] = &pb.ProductOrder{
			UserId:          req.UserId,
			ProductSkuId:    carts.ShoppingCartsProductInfo[i].ProductSkuId,
			ProductId:       carts.ShoppingCartsProductInfo[i].ProductId,
			ProductName:     carts.ShoppingCartsProductInfo[i].ProductName,
			ProductPrice:    fmt.Sprintf("%v", carts.ShoppingCartsProductInfo[i].SellPrice),
			Quantity:        carts.ShoppingCartsProductInfo[i].Quantity,
			ProductImageUrl: carts.ShoppingCartsProductInfo[i].ProductMainPicture,
			OrderId:         tradeOrder.TradeOrder.Id,
			SkuDescribe:     carts.ShoppingCartsProductInfo[i].SkuDescribe,
			DetailStatus:    common.DetailNormal,
			ActivityType:    common.ActivityNormal,
			CreateUser:      req.UserId,
			CreateTime:      timestamppb.New(time.Now()),
			RealAmount:      fmt.Sprintf("%v", carts.ShoppingCartsProductInfo[i].SellPrice),
			RealPrice:       fmt.Sprintf("%v", carts.ShoppingCartsProductInfo[i].SellPrice),
		}
	}
	// 添加订单商品
	fmt.Println("正在生成订单商品.....")
	//pbCreateList := make([]*pb.ProductOrder, 0)
	//for i := 0; i < len(tradeOrderProduct); i++ {
	//	pbCreateList = append(pbCreateList, pb.ModelProductOrderConvertPb(&tradeOrderProduct[i]))
	//}
	fmt.Printf("当前用户Id：%v的商品有%v个！\n", req.UserId, len(tradeOrderProduct))
	_, err = l.svcCtx.TradeOrderRpcConf.BatchCreateOrderProduct(l.ctx, &tradeorder.AddProductOrderReq{
		ProductOrder: tradeOrderProduct,
	})
	if err != nil {
		l.Error(err)
		return
	}
	// 添加物流

	resp = new(types.AddTradeOrderResp)
	resp.TradeOrder = types.TradeOrderPbConvertTypes(tradeOrder.TradeOrder)
	return
}

//func (l *AddTradeOrderLogic) BatchCreateOrderProduct(tradeOrderProduct []model.TradeOrderProduct) {
//
//}
