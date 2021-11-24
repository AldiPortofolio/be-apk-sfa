package kafka

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
)

var (
	kafkaBrokerUrl string
	kafkaTopic     string
	kafkaClient    string
	kafkaTimeout   int
	maxmsgbyte     int
)

// init ..
func init() {
	kafkaBrokerUrl = beego.AppConfig.DefaultString("kafka.logging.brokers", "13.250.26.210:9092,13.250.26.210:9093,13.250.26.210:9094")
	kafkaTopic = beego.AppConfig.DefaultString("kafka.logging.topics", "ottofin-logging")
	kafkaClient = beego.AppConfig.DefaultString("kafka.logging.clientid", "ottofin-logging-sub")
	kafkaTimeout = beego.AppConfig.DefaultInt("kafka.logging.produce.timeout", 10)
	maxmsgbyte = beego.AppConfig.DefaultInt("kafka.logging.produce.maxmsgbyte", 50000000)
}

// GetConectionKafka ..
func GetConectionKafka(brokerList string, config *sarama.Config) KafkaProducer {
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	return KafkaProducer{
		Connection: producer,
		ErrRes:     err,
	}
}

// GetKafkaBroker ..
func GetKafkaBroker() string {
	return kafkaBrokerUrl
}

// GetTopic ..
func GetTopic() string {
	return kafkaTopic
}

// GetConnKakfaProcedure ..
func GetConnKakfaProcedure() KafkaProducer {
	producer, err := sarama.NewSyncProducer(strings.Split(GetKafkaBroker(), ","), GetConfigKafka())
	return KafkaProducer{
		Connection: producer,
		ErrRes:     err,
	}
}

// GetConfigKafka ..
func GetConfigKafka() *sarama.Config {

	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	configKafka := sarama.NewConfig()
	configKafka.Producer.Return.Successes = true
	configKafka.Producer.Partitioner = sarama.NewRandomPartitioner
	configKafka.ClientID = kafkaClient
	configKafka.Producer.MaxMessageBytes = maxmsgbyte
	configKafka.Producer.Timeout = 5 * time.Second
	configKafka.Net.DialTimeout = 2 * time.Second
	configKafka.Net.ReadTimeout = 5 * time.Second
	configKafka.Net.WriteTimeout = 5 * time.Second

	return configKafka
}
