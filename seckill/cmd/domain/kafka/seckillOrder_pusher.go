package kafka

import (
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/rpc/pb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type ISecKillOrderPusher interface {
	PusherWaitForPaymentSecKillOrder(ctx context.Context, in *pb.IncreaseSecKillOrderReq) (*model.SecKillOrder, error)
}

type SecKillOrderPusher struct {
	KqPusher *kafka.Writer
}

func (s *SecKillOrderPusher) PusherWaitForPaymentSecKillOrder(ctx context.Context, in *pb.IncreaseSecKillOrderReq) (*model.SecKillOrder, error) {
	now := time.Now()
	secKillOrder := &model.SecKillOrder{
		CreateTime:  now,
		Seller:      in.Seller,
		Buyer:       in.Buyer,
		Price:       in.Price,
		ProductsId:  in.ProductsId,
		ProductsNum: in.ProductsNum,
		OrderNum:    in.OrderNum,
		Status:      1,
		UpdateTime:  &now,
	}
	// 将订单信息序列化为 JSON
	orderBytes, err := json.Marshal(secKillOrder)
	if err != nil {
		fmt.Errorf("订单信息序列化失败: %v", err)
		return nil, err
	}

	err = s.KqPusher.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("waitForPayment"),
			Value: orderBytes,
		},
	)
	fmt.Printf("发送第订单中....\n")
	//fmt.Printf("当前订单详情：%v\n", secKillOrder)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return nil, err
	}
	time.Sleep(1 * time.Second)
	return secKillOrder, nil
}

func NewSecKillOrderPusher(Brokers []string, Topic string) ISecKillOrderPusher {
	return &SecKillOrderPusher{
		KqPusher: &kafka.Writer{
			Addr:         kafka.TCP(Brokers...),
			Topic:        Topic,
			RequiredAcks: kafka.RequireAll,
			Async:        true,
			Balancer:     &kafka.LeastBytes{},
			Completion: func(messages []kafka.Message, err error) {
				if err != nil {
					log.Fatal("failed to write messages:", err)
				}
				fmt.Printf("Messages %v sent successfully!\n", messages)
			},
			AllowAutoTopicCreation: true,
		},
	}
}
