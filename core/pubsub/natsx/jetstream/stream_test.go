package jetstream

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"strconv"
	"testing"
	"time"
)

const (
	streamName     = "ORDERS"
	streamSubjects = "ORDERS.*"
	subjectName    = "ORDERS.created"
)

type Order struct {
	OrderID    int    `json:"order_id"`
	CustomerID string `json:"customer_id"`
	Status     string `json:"status"`
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestJetStream(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	// 创建JetStreamContext
	js, err := nc.JetStream()
	checkErr(t, err)
	// 创建stream流
	err = createStream(js)
	// checkErr(t, err)
	// 通过发布消息创建订单
	err = createOrder(js)
	checkErr(t, err)
	// go TestPull(t)
	go TestSubscribe(t)
	time.Sleep(10 * time.Second)
}

// createOrder 以 "ORDERS.created"主题发布事件流
func createOrder(js nats.JetStreamContext) error {
	var order Order
	for i := 1; i <= 10; i++ {
		order = Order{
			OrderID:    i,
			CustomerID: "Cust-" + strconv.Itoa(i),
			Status:     "created",
		}
		orderJSON, _ := json.Marshal(order)
		_, err := js.Publish(subjectName, orderJSON)
		if err != nil {
			return err
		}
	}
	return nil
}

// createStream 使用JetStreamContext创建流
func createStream(js nats.JetStreamContext) error {
	// Check if the ORDERS ctx already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating ctx %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	subSubjectName = "ORDERS.created"
	pubSubjectName = "ORDERS.approved"
)

func TestPull(t *testing.T) {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create Pull based consumer with maximum 128 inflight.
	// PullMaxWaiting defines the max inflight pull requests.
	sub, _ := js.PullSubscribe(subSubjectName, "order-review", nats.PullMaxWaiting(128))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			_ = msg.Ack()
			var order Order
			err := json.Unmarshal(msg.Data, &order)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("order-review service")
			log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
			reviewOrder(js, order)
		}
	}
}

// reviewOrder reviews the order and publishes ORDERS.approved event
func reviewOrder(js nats.JetStreamContext, order Order) {
	// Changing the Order status
	order.Status = "approved"
	orderJSON, _ := json.Marshal(order)
	_, err := js.Publish(pubSubjectName, orderJSON)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Order with OrderID:%d has been %s\n", order.OrderID, order.Status)
}

func TestSubscribe(t *testing.T) {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("ORDERS.*", func(msg *nats.Msg) {
		msg.Ack()
		var order Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("monitor service subscribes from subject:%s\n", msg.Subject)
		log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
	}, nats.Durable("monitor"), nats.ManualAck())

	runtime.Goexit()
}
