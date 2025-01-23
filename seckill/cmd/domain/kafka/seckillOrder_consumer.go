package kafka

import (
	"QMall/common"
	"QMall/seckill/cmd/domain/model"
	"QMall/seckill/cmd/domain/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type ISecKillOrderConsumer interface {
	ReaderSecKillOrder(ctx context.Context)
}

type SecKillOrderConsumer struct {
	KqReader          *kafka.Reader
	SecKillRepository repository.ISecKillRepository
}

func NewSecKillOrderConsumer(config *ConsumerConfig) ISecKillOrderConsumer {
	//user := "root"
	//pass := "000000"
	//host := "192.168.23.233:3306"
	//database := "mall_seckill"
	//charset := "utf8mb4"
	//
	//redisHost := "192.168.23.233:6379"
	//redisPass := "123456"
	dsn := fmt.Sprintf(common.MysqlConnect, config.MysqlConfig.User, config.MysqlConfig.Pass, config.MysqlConfig.Host, config.MysqlConfig.Database, config.MysqlConfig.Charset)

	directory := mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
	db, err := gorm.Open(directory, &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:         config.RedisConfig.Host,
		Password:     config.RedisConfig.Pass,
		PoolSize:     30,
		MinIdleConns: 30,
	})
	return &SecKillOrderConsumer{
		KqReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  config.KafkaConfig.Brokers,
			Topic:    config.KafkaConfig.Topic,
			GroupID:  config.KafkaConfig.GroupId,
			MinBytes: 10e3,
			MaxBytes: 10e6,
			//CommitInterval: 1 * time.Second, // 自动提交偏移量的时间间隔
		}),
		SecKillRepository: repository.NewSecKillRepository(db, redisClient),
	}
}

func (r *SecKillOrderConsumer) ReaderSecKillOrder(ctx context.Context) {
	// 创建一个 Kafka 阅读器
	//reader := kafka.NewReader(kafka.ReaderConfig{
	//	Brokers:  []string{"192.168.23.233:9092"},
	//	Topic:    "test5",
	//	GroupID:  "user1",
	//	MinBytes: 10e3,
	//	MaxBytes: 10e6,
	//	//CommitInterval: 1 * time.Second, // 自动提交偏移量的时间间隔
	//})
	// 设置偏移量为最早的消息
	defer r.KqReader.Close()

	fmt.Println("开始读取秒杀订单信息...")
	for {
		// 从 Kafka 读取消息
		message, err := r.KqReader.ReadMessage(ctx)
		if err != nil {
			log.Printf("读取消息时出错: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		// 解析消息内容
		//var order map[string]interface{}
		order := &model.SecKillOrder{}
		if err := json.Unmarshal(message.Value, &order); err != nil {
			log.Printf("解析消息时出错: %v", err)
			continue
		}

		// 打印接收到的订单信息
		fmt.Println("====================================================================")
		fmt.Printf("接收到秒杀订单信息: %v\n", order)

		// 在这里可以添加将订单信息存储到 MySQL 或其他存储的逻辑
		_, err = r.SecKillRepository.IncreaseSecKillOrderWithKafka(order)
		if err != nil {
			fmt.Printf("待支付秒杀订单信息：%v插入失败！失败原因：%v\n", order, err)
		}
		fmt.Printf("待支付秒杀订单信息：%v插入成功！\n", order)

	}
}
