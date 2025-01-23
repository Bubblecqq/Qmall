package kafka

type Mysql struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

type Redis struct {
	Host string `json:"host"`
	Pass string `json:"pass"`
}

type Kafka struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
	GroupId string   `json:"group_id"`
}

type ConsumerConfig struct {
	MysqlConfig Mysql
	RedisConfig Redis
	KafkaConfig Kafka
}
