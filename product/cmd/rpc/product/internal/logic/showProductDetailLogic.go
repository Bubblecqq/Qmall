package logic

import (
	"QMall/product/cmd/domain/model"
	"context"
	"fmt"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShowProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowProductDetailLogic {
	return &ShowProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShowProductDetailLogic) ShowProductDetail(in *pb.ShowProductDetailReq) (*pb.ShowProductDetailResp, error) {

	productDetail, err := l.svcCtx.ProductRepository.ShowProductDetail(in.Id)

	if err != nil {
		return &pb.ShowProductDetailResp{}, err
	}
	fmt.Println(">>>>>>>>>>>>", productDetail)
	fmt.Println("isEnable>", productDetail.IsEnabled)
	return &pb.ShowProductDetailResp{
		ProductDetail: Product_detailModel_Convert_PbModel(productDetail),
	}, nil
}

func Product_detailModel_Convert_PbModel(productDetail *model.Product) *pb.DetailProduct {

	return &pb.DetailProduct{
		Id:                productDetail.Id,
		Name:              productDetail.Name,
		ProductType:       productDetail.ProductType,
		CategoryId:        productDetail.CategoryId,
		StartingPrice:     productDetail.StartingPrice,
		TotalStock:        productDetail.TotalStock,
		MainPicture:       productDetail.MainPicture,
		RemoteAreaPostage: productDetail.RemoteAreaPostage,
		SingleBuyLimit:    productDetail.SingleBuyLimit,
		Remark:            productDetail.Remark,
		Detail:            productDetail.Detail.Detail,
		PictureList:       GetPictureList(productDetail),
		IsEnable:          productDetail.IsEnabled,
	}
}

func GetPictureList(productDetail *model.Product) []string {
	var pictureList []string
	n := len(productDetail.PictureList)
	for i := 0; i < n; i++ {
		pictureList = append(pictureList, productDetail.PictureList[i].Picture)
	}
	return pictureList
}
