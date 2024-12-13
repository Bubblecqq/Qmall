package shoppingCart

import (
	"QMall/product/cmd/rpc/product/product"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"
)

func (l *UpdateCartLogic) GetProductById(in *pb.UpdateCartReq) (*product.GetProductByIdResp, error) {
	// 根据服务名称和配置信息发现服务实例
	// "consul://192.168.23.233:8500/product.rpc"
	//target, _ := zrpc.NewClientWithTarget(l.svcCtx.Config.ShoppingCartRpcConf.Target)
	// 创建远程服务客户端
	//client := product.NewProductZrpcClient(target)
	// 进行远程调用，例如调用远程服务的某个方法
	p, err := l.svcCtx.ProductRPC.GetProduct(l.ctx, &product.GetProductByIdReq{
		Id: in.ProductId,
	})
	if err != nil {
		return p, err
	}
	return p, nil
}

func (l *UpdateCartLogic) GetProductSkuByIdCall(in *pb.UpdateCartReq) (*product.GetProductSkuByIdResp, error) {

	p, err := l.svcCtx.ProductRPC.GetProductSku(l.ctx, &product.GetProductSkuByIdReq{
		Id: in.ProductSkuId,
	})
	if err != nil {
		return p, err
	}
	return p, nil
}
