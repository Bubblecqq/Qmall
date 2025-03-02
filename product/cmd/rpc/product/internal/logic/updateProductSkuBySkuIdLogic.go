package logic

import (
	"QMall/product/cmd/domain/model"
	"context"
	"fmt"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductSkuBySkuIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductSkuBySkuIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSkuBySkuIdLogic {
	return &UpdateProductSkuBySkuIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductSkuBySkuIdLogic) UpdateProductSkuBySkuId(in *pb.UpdateProductSkuBySkuIdReq) (*pb.UpdateProductSkuBySkuIdResp, error) {

	fmt.Printf("正在更新库存中>>>>>>库存Id：%v\n", in.SkuId)

	skuId, err := l.svcCtx.ProductRepository.UpdateProductSkuBySkuId(in)

	if err != nil {
		return &pb.UpdateProductSkuBySkuIdResp{}, err
	}

	return &pb.UpdateProductSkuBySkuIdResp{
		ProductSku: model.ProductSkuModelConvertPbProductSku(skuId),
	}, nil
}
