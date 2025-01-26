package logic

import (
	"context"

	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndDeductQuotaAndStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndDeductQuotaAndStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndDeductQuotaAndStockLogic {
	return &CheckAndDeductQuotaAndStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAndDeductQuotaAndStockLogic) CheckAndDeductQuotaAndStock(in *pb.CheckAndDeductQuotaAndStockReq) (*pb.CheckAndDeductQuotaAndStockResp, error) {

	return &pb.CheckAndDeductQuotaAndStockResp{}, nil
}
