package logic

import (
	"QMall/product/cmd/domain/model"
	"context"

	"QMall/product/cmd/rpc/product/internal/svc"
	"QMall/product/cmd/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageIndexLogic {
	return &PageIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PageIndexLogic) PageIndex(in *pb.PageReq) (*pb.PageResp, error) {

	page, err := l.svcCtx.ProductRepository.Page(in.Length, in.PageIndex)

	if err != nil {
		return &pb.PageResp{}, nil
	}
	pbPageProductList := make([]*pb.Product, len(page))
	for i := 0; i < len(pbPageProductList); i++ {
		pbPageProductList[i] = model.ProductModelConvertPbProduct(&page[i])
	}

	return &pb.PageResp{
		PageProductList: pbPageProductList,
	}, nil
}
