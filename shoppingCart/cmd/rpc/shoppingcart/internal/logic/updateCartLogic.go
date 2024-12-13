package logic

import (
	"QMall/product/cmd/rpc/product/product"
	"context"
	"github.com/zeromicro/go-zero/zrpc"

	"QMall/shoppingCart/cmd/rpc/shoppingcart/internal/svc"
	"QMall/shoppingCart/cmd/rpc/shoppingcart/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartLogic {
	return &UpdateCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartLogic) UpdateCart(in *pb.UpdateCartReq) (*pb.UpdateCartResp, error) {
	_, err := l.svcCtx.ShoppingCartRepository.UpdateShoppingCart(in)

	if err != nil {
		return &pb.UpdateCartResp{}, err
	}
	byId := TestGetProductById(in)
	l.Logger.Info(">>>>>>>>>>>>>>>Test Remote >>>>>>>", byId.Product)
	return &pb.UpdateCartResp{}, nil
}

func TestGetProductById(in *pb.UpdateCartReq) *product.GetProductByIdResp {
	// 根据服务名称和配置信息发现服务实例
	target, _ := zrpc.NewClientWithTarget("consul://192.168.23.233:8500/product.rpc")
	// 创建远程服务客户端
	client := product.NewProductZrpcClient(target)
	// 进行远程调用，例如调用远程服务的某个方法
	p, err := client.GetProduct(context.Background(), &product.GetProductByIdReq{
		Id: in.ProductId,
	})
	if err != nil {

	}
	return p
	// 处理远程调用返回的响应
	//...
}
