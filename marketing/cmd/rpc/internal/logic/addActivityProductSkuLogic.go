package logic

import (
	"context"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductSkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductSkuLogic {
	return &AddActivityProductSkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityProductSkuLogic) AddActivityProductSku(in *pb.AddActivityProductSkuReq) (*pb.AddActivityProductSkuResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddActivityProductSkuResp{}, nil
}
