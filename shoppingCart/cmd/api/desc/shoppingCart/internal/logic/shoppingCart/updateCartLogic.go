package shoppingCart

import (
	"QMall/product/cmd/domain/model"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/svc"
	"QMall/shoppingCart/cmd/api/desc/shoppingCart/internal/types"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/shoppingcart"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改购物车
func NewUpdateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartLogic {
	return &UpdateCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartLogic) UpdateCart(req *types.UpdateShoppingCartReq) (resp *types.UpdateShoppingCartResp, err error) {

	updateCartReq := &shoppingcart.UpdateCartReq{
		Id:                 req.Id,
		UserId:             req.UserId,
		ProductId:          req.ProductId,
		ProductSkuId:       req.ProductSkuId,
		ProductName:        req.ProductName,
		ProductMainPicture: req.ProductMainPicture,
		IsDelete:           req.IsDeleted,
		//UpdateTime:         timestamppb.New(time.Now()),
	}

	_, err = l.svcCtx.ShoppingCartRpcConf.UpdateCart(l.ctx, updateCartReq)
	productById, err := l.GetProductById(updateCartReq)
	product := productById.Product

	productSkuById, err := l.GetProductSkuByIdCall(updateCartReq)
	productSku := productSkuById.ProductSku
	resp = new(types.UpdateShoppingCartResp)
	resp.Id = req.Id
	resp.UserId = req.UserId
	resp.ProductId = req.ProductId
	resp.ProductSkuId = req.ProductSkuId
	resp.ProductName = req.ProductName
	resp.Product = types.ProductModelConvertTypes(model.PbProductModelConvert(product))
	resp.ProductSku = types.ProductSkuModelConvertTypes(model.PbProductSkuModelConvert(productSku))
	return
}
