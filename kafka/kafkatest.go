package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"strings"
)

type Client struct {
	asyncProducer sarama.AsyncProducer
	syncProducer  sarama.SyncProducer
	consumer      *cluster.Consumer
}

func NewKafkaClient() *Client {
	return &Client{
		syncProducer:  kafkaSyncProducer,
		asyncProducer: kafkaAsyncProducer,
		consumer:      kafkaConsumer,
	}
}

var kafkaSyncProducer sarama.SyncProducer
var kafkaAsyncProducer sarama.AsyncProducer
var kafkaConsumer *cluster.Consumer

func Init() {
	brokerListStr := "192.168.243.149:9092"
	fmt.Printf("Init kafka addr:%s\n", brokerListStr)
	brokerList := strings.Split(brokerListStr, ",")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Compression = sarama.CompressionGZIP
	config.Producer.MaxMessageBytes = 1024 * 1024 * 5

	asyncProducer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalf("kafka connect error:%v", err.Error())
	}
	kafkaAsyncProducer = asyncProducer
	syncProducer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalf("kafka connect error:%v", err.Error())
	}
	kafkaSyncProducer = syncProducer

	consumerConfig := cluster.NewConfig()
	groupId := "test_group_id_001"
	topicName1 := "test1"
	consumer, err := cluster.NewConsumer(brokerList, groupId, []string{topicName1}, consumerConfig)
	if err != nil {
		log.Fatalf("kafka connect error:%v", err.Error())
	}
	kafkaConsumer = consumer
}

func (kc *Client) AsyncSend(topic string, msg string) (int32, int64, error) {
	p := kc.asyncProducer
	p.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	select {
	case res := <-p.Successes():
		return res.Partition, res.Offset, nil
	case err := <-p.Errors():
		fmt.Print("Produced message failure: ", err)
		return 0, 0, err
	}
}

func (kc *Client) SyncSend(topic string, msg string) (int32, int64, error) {

	partition, offset, err := kc.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	})
	if nil != err {
		fmt.Println(err)
	}
	return partition, offset, err
}

func (kc *Client) RecMsg2(f func(message *sarama.ConsumerMessage)) {
	consumer := kc.consumer
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
				go f(msg)
			}
		}
	}
}
