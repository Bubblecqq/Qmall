package logic

import (
	"QMall/seckill/cmd/rpc/internal/svc"
	"QMall/seckill/cmd/rpc/pb"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type SaveSecKillStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveSecKillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSecKillStockLogic {
	return &SaveSecKillStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveSecKillStockLogic) SaveSecKillStock(in *pb.SaveSecKillStockReq) (*pb.SaveSecKillStockResp, error) {
	cacheKey := fmt.Sprintf("secKillStock:%d", in.ProductsId)
	fmt.Printf("正在查询秒杀库存信息>>>>>当前请求库存Id：%v\n", in.ProductsId)
	lockKey := fmt.Sprintf("secKillStock_lock:%d", in.ProductsId)
	locked, err := l.svcCtx.Redis.SetNX(l.ctx, lockKey, "locked", 10*time.Second).Result()
	if err != nil {
		return &pb.SaveSecKillStockResp{}, fmt.Errorf("获取分布式锁时出错：%v", err)
	}
	if !locked {
		return &pb.SaveSecKillStockResp{}, fmt.Errorf("获取分布式锁失败，请稍后尝试！")
	}
	defer l.svcCtx.Redis.Del(l.ctx, lockKey)

	stock, err := l.svcCtx.SecKillRepository.GetSecKillStockWithCache2(l.ctx, cacheKey)
	if err != nil {
		if err.Error() != "cache miss" {
			fmt.Printf("获取缓存库存信息出错：%v\n", err)
			return &pb.SaveSecKillStockResp{}, err
		}
		// 从数据库查询然后扣减
		// 缓存未命中，从数据库查询
		fmt.Printf("当前缓存未命中，正在查询MYSQL....\n")
		secKillStock, err := l.svcCtx.SecKillRepository.GetSecKillStockByProductsId(in.ProductsId)
		if err != nil {
			fmt.Printf("从数据库获取库存信息出错：%v\n", err)
			return &pb.SaveSecKillStockResp{}, err
		}
		if in.Quantity > secKillStock.Stock {
			return &pb.SaveSecKillStockResp{}, errors.New("当前秒杀数量超过库存数！")
		}
		newStock := secKillStock.Stock - in.Quantity
		if newStock < 0 {
			return &pb.SaveSecKillStockResp{}, errors.New("当前秒杀数量超过库存数！")
		}
		err = l.svcCtx.SecKillRepository.SetSecKillStockWithCacheV3(l.ctx, cacheKey, newStock)
		if err != nil {
			// 记录缓存设置失败的日志，但不影响返回结果
			fmt.Printf("设置秒杀库存信息缓存失败：%v\n", err)
		}
		stock = newStock
		if err = l.updateStockAndHandleError(in.ProductsId, in.Quantity, cacheKey); err != nil {
			return &pb.SaveSecKillStockResp{}, err
		}
	} else {
		// 缓存命中
		success, err := l.svcCtx.SecKillRepository.DecreaseStockWithLockV3(l.ctx, cacheKey, in.ProductsId, in.Quantity)
		if err != nil {
			return &pb.SaveSecKillStockResp{}, err
		}
		if !success {
			return &pb.SaveSecKillStockResp{}, errors.New("库存不足，无法完成预扣")
		}
		if err = l.updateStockAndHandleError(in.ProductsId, in.Quantity, cacheKey); err != nil {
			return &pb.SaveSecKillStockResp{}, err
		}
	}
	// 获取更新后的库存
	stock, err = l.svcCtx.SecKillRepository.GetSecKillStockWithCache2(l.ctx, cacheKey)
	if err != nil {
		return &pb.SaveSecKillStockResp{}, err
	}
	//
	err = l.svcCtx.SecKillRepository.UpdateSecKillUserQuotaFromStock(in.UserId, in.ProductsId, in.Quantity, in.Limit)
	if err != nil {
		return &pb.SaveSecKillStockResp{}, err
	}
	fmt.Printf("预扣逻辑缓存更新成功！....\n")
	return &pb.SaveSecKillStockResp{
		Stock: stock,
	}, nil
}

func (l *SaveSecKillStockLogic) updateStockAndHandleError(productsId, quantity int64, cacheKey string) error {
	// 需要更新mysql
	err := l.svcCtx.SecKillRepository.UpdateSecKillStockByProductsId(productsId, quantity)
	if err != nil {
		// 这里进行回滚操作，比如将 Redis 中的库存加回去
		rollbackErr := l.svcCtx.SecKillRepository.SecKillStockRollbackStockInCache(l.ctx, cacheKey, quantity)
		if rollbackErr != nil {
			fmt.Printf("回滚 Redis 库存失败：%v\n", rollbackErr)
		}
		fmt.Printf("更新 MySQL 库存失败：%v\n", err)
		return err
	}
	// 再更新缓存
	newStock, err := l.svcCtx.SecKillRepository.GetSecKillStockByProductsId(productsId)
	if err != nil {
		fmt.Printf("更新缓存时从数据库获取库存信息失败：%v\n", err)
		return err
	}
	err = l.svcCtx.SecKillRepository.SetSecKillStockWithCacheV3(l.ctx, cacheKey, newStock.Stock)
	if err != nil {
		fmt.Printf("更新缓存失败：%v\n", err)
		return err
	}
	return nil
}

//func (l *SaveSecKillStockLogic) SaveSecKillStock(in *pb.SaveSecKillStockReq) (*pb.SaveSecKillStockResp, error) {
//
//	cacheKey := fmt.Sprintf("secKillStock:%d", in.ProductsId)
//	fmt.Printf("正在查询秒杀库存信息>>>>>当前请求库存Id：%v\n", in.ProductsId)
//
//	var secKillStock *model.SecKillStock
//	secKillStock, err := l.svcCtx.SecKillRepository.GetSecKillStockWithCache(l.ctx, cacheKey)
//	if err != nil {
//		if err.Error() != "cache miss" {
//			return &pb.SaveSecKillStockResp{}, err
//		}
//		// 缓存未命中，从数据库查询
//		fmt.Printf("当前缓存未命中，正在查询MYSQL....\n")
//		secKillStock, err = l.svcCtx.SecKillRepository.GetSecKillStockByProductsId(in.ProductsId)
//		if err != nil {
//			return &pb.SaveSecKillStockResp{}, err
//		}
//		// 将从数据库查询到的结果存入缓存
//		err = l.svcCtx.SecKillRepository.SetSecKillStockWithCache(l.ctx, cacheKey, secKillStock)
//		if err != nil {
//			// 记录缓存设置失败的日志，但不影响返回结果
//			fmt.Printf("设置秒杀库存信息缓存失败：%v\n", err)
//		}
//		fmt.Printf("设置秒杀库存信息缓存成功！\n")
//	}
//	return &pb.SaveSecKillStockResp{
//		SecKillStock: convert.ModelSecKillStockConvertPb(secKillStock),
//	}, nil
//}
