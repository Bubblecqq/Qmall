package logic

import (
	"context"

	"QMall/marketing/cmd/rpc/internal/svc"
	"QMall/marketing/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddActivityProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityProductLogic {
	return &AddActivityProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddActivityProductLogic) AddActivityProduct(in *pb.AddActivityProductReq) (*pb.AddActivityProductResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddActivityProductResp{}, nil
}
