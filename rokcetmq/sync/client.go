package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
)

func SimpleRocketMQ() {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.0.60:9876"})),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}

	err = p.Start()
	if err != nil {
		fmt.Println("start producer error: ", err)
	}
	topic := "test"

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
		}
		res, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}

}

func main() {
	fmt.Println("enter rocketmq simple client")
	SimpleRocketMQ()
	simpleConsumer()
}

func simpleConsumer() {
	sig := make(chan os.Signal)
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.0.60:9876"})),
	)
	err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	<-sig
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
