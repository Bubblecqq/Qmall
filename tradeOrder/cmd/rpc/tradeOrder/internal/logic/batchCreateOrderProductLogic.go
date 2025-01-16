package logic

import (
	"QMall/tradeOrder/cmd/domain/model"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/internal/svc"
	"QMall/tradeOrder/cmd/rpc/tradeOrder/pb"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateOrderProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateOrderProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateOrderProductLogic {
	return &BatchCreateOrderProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchCreateOrderProductLogic) BatchCreateOrderProduct(in *pb.AddProductOrderReq) (*pb.AddProductOrderResp, error) {
	var orderProductList []*model.TradeOrderProduct
	for i := 0; i < len(in.ProductOrder); i++ {
		//orderProductList = append(orderProductList, model.TradeOrderProduct{
		//	CreateTime:      time.Now(),
		//	UserId:          in.ProductOrder[i].UserId,
		//	ProductName:     in.ProductOrder[i].ProductName,
		//	ProductImageUrl: in.ProductOrder[i].ProductImageUrl,
		//	ProductId:       in.ProductOrder[i].ProductId,
		//	ProductSkuId:    in.ProductOrder[i].ProductSkuId,
		//	ProductPrice:    in.ProductOrder[i].ProductPrice,
		//	Quantity:        in.ProductOrder[i].Quantity,
		//	OrderId:         in.ProductOrder[i].OrderId,
		//	SkuDescribe:     in.ProductOrder[i].SkuDescribe,
		//	DetailStatus:    in.ProductOrder[i].DetailStatus,
		//	ActivityType:    in.ProductOrder[i].ActivityType,
		//	ActivityId:      in.ProductOrder[i].ActivityId,
		//})
		orderProductList = append(orderProductList, pb.PbProductOrderConvertModel(in.ProductOrder[i]))
	}
	fmt.Printf("正在批量添加订单商品....\n")
	err := l.svcCtx.TradeOrderProductRepository.BatchCreateOrderProduct(orderProductList)
	if err != nil {
		fmt.Printf("批量添加订单商品失败！原因见：%v\n", err)
		return &pb.AddProductOrderResp{}, err
	}
	return &pb.AddProductOrderResp{}, nil
}
