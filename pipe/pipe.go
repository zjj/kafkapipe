package pipe

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var msgChan = make(chan *kafka.Message, 1024)

func writeToKafka(cfg *kafka.ConfigMap, topic string) error {
	p, err := kafka.NewProducer(cfg)
	if err != nil {
		return err
	}

	for msg := range msgChan {
		err = p.Produce(
			&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          msg.Value,
			},
			nil,
		)
		if err != nil {
			log.Printf("Failed to produce message %s", err.Error())
		}
	}
	return nil
}

func readFromKafka(cfg *kafka.ConfigMap, topic string) error {
	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	for {
		ev := c.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			select {
			case msgChan <- e:
			default:
				log.Println("Failed to put message to msgChan, it's full")
			}
		case kafka.Error:
			log.Printf("Failed to consumer Error: %v\n", e)
		default:
		}
	}
}

func Start(fromCfg *kafka.ConfigMap, topicFrom string, toCfg *kafka.ConfigMap, topicTo string) {
	go func() {
		readFromKafka(fromCfg, topicFrom)
	}()
	writeToKafka(toCfg, topicTo)
}
