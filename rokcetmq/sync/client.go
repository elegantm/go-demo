package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
	"os"
	"strconv"
	"time"
)

var sleepTime = 1 * time.Second
var pullConsumer rocketmq.PullConsumer

const (
	topic             = "test-topic"
	consumerGroupName = "testPullGroup"
	tag               = "testPull"
	namespace         = "ns"
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
	//topic := "test"

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
	//simpleConsumer()

	pullConsumerNew()

	go func() {

		for {
			pull()
		}
	}()

	sig := make(chan os.Signal)
	<-sig
	err := pullConsumer.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}

func simpleConsumer() {
	sig := make(chan os.Signal)
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.0.60:9876"})),
	)
	err := c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
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

func pullConsumerNew() {
	c, _ := rocketmq.NewPullConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNameServer([]string{"192.168.0.60:9876"}),
		consumer.WithMaxReconsumeTimes(2),
	)
	pullConsumer = c
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: tag,
	}
	err := pullConsumer.Subscribe(topic, selector)
	if err != nil {
		log.Fatalf("fail to Subscribe: %v", err)
	}
	err = pullConsumer.Start()
	if err != nil {
		log.Fatalf("fail to Start: %v", err)
	}
	const refreshPersistOffsetDuration = time.Second * 5
	timer := time.NewTimer(refreshPersistOffsetDuration)
	go func() {
		for ; true; <-timer.C {
			err = pullConsumer.PersistOffset(context.TODO(), topic)
			if err != nil {
				log.Printf("[pullConsumer.PersistOffset] err=%v", err)
			}
			timer.Reset(refreshPersistOffsetDuration)
		}
	}()

}

func pull() {
	resp, err := pullConsumer.Pull(context.TODO(), 1)
	if err != nil {
		log.Printf("[pull error] err=%v", err)
		time.Sleep(sleepTime)
		return
	}
	switch resp.Status {
	case primitive.PullFound:
		log.Println("enter pull found +++++++++++++++++++++++++++++++")
		log.Printf("[pull message successfully] MinOffset:%d, MaxOffset:%d, nextOffset: %d, len:%d\n", resp.MinOffset, resp.MaxOffset, resp.NextBeginOffset, len(resp.GetMessages()))
		var queue *primitive.MessageQueue
		if len(resp.GetMessages()) <= 0 {
			return
		}
		for _, msg := range resp.GetMessageExts() {
			// todo LOGIC CODE HERE
			queue = msg.Queue
			//log.Println(msg.Queue, msg.QueueOffset, msg.GetKeys(), msg.MsgId, string(msg.Body))
			log.Println(msg)
		}
		// update offset
		err = pullConsumer.UpdateOffset(queue, resp.NextBeginOffset)
		if err != nil {
			log.Printf("[pullConsumer.UpdateOffset] err=%v", err)
		}

	case primitive.PullNoNewMsg, primitive.PullNoMsgMatched:
		log.Printf("[no pull message]   next = %d\n", resp.NextBeginOffset)
		time.Sleep(sleepTime)
		return
	case primitive.PullBrokerTimeout:
		log.Printf("[pull broker timeout]  next = %d\n", resp.NextBeginOffset)

		time.Sleep(sleepTime)
		return
	case primitive.PullOffsetIllegal:
		log.Printf("[pull offset illegal] next = %d\n", resp.NextBeginOffset)
		return
	default:
		log.Printf("[pull error]  next = %d\n", resp.NextBeginOffset)
	}
}
