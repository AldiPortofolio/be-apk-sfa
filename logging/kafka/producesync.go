package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

// KafkaProducer ..
type KafkaProducer struct {
	//Req        *pb.PulseRequest
	Connection sarama.SyncProducer
	ErrRes     error
	KeyEncoder sarama.Encoder
	Value      sarama.Encoder
	Topic      string
	Done       bool
}

// Process ...
func (kafkareq *KafkaProducer) Process() {
	//fmt.Println("Kafka PRocedure")
	//wg := &sync.WaitGroup{}
	//wg.Add(1)
	//	doneMsg := make(chan bool)
	kafkareq.Done = kafkareq.Send()
	//wg.Done()
}

// Send .. Message to Kafka (SyncProducer) ...
func (kafkareq *KafkaProducer) Send() bool {
	producer := kafkareq.Connection

	producerMsg := &sarama.ProducerMessage{
		Topic: kafkareq.Topic,
		Key:   kafkareq.KeyEncoder,
		Value: kafkareq.Value,
	}

	err := sendProducer(producer, producerMsg)
	if err != nil {
		fmt.Println("Error :", err)
		return false
	}

	return true
}

// sendProducer ...
func sendProducer(producer sarama.SyncProducer, produceMessage *sarama.ProducerMessage) error {
	defer func() {
		if err := producer.Close(); err != nil {

			logs.Error("Producer Can't Close : ", err)
		}
	}()

	_, _, err := producer.SendMessage(produceMessage)
	//fmt.Println("Partion ",partition)
	//if partition != 0 {
	//	return errors.New("Unexpected partition")
	//}
	if err != nil {

		logs.Error("Producer Send Message : ", err)
		return err
	}
	return nil

}
