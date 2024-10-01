package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/sarama"
)

type sendEmailConsumerGroup struct{}

var SendEmailKafkaProducer sarama.SyncProducer
var SendEmailKafkaConsumer sarama.ConsumerGroup

func init() {
	// SendEmailKafkaProducer = kafka.InitSynProducer(configs.Data().Kafka.Brokers)
	// SendEmailKafkaConsumer = kafka.InitConsumerGroup(configs.Data().Kafka.Brokers, configs.Data().Kafka.GroupID)
}
func (m sendEmailConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m sendEmailConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

type SendReqInKafka struct {
	EmailAddr  string
	EmailID    uint64
	SecretCode string
}

func (m sendEmailConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		req := SendReqInKafka{}
		json.Unmarshal(msg.Value, &req)
		fmt.Println(req)

		subject := "Welcome to wzw's E-mall"
		// TODO: 改为config
		verifyUrl := fmt.Sprintf("http://localhost:8888/api/v1/verify_email?email_id=%d&secret_code=%s",
			req.EmailID, req.SecretCode)
		content := fmt.Sprintf(`Hello!<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, verifyUrl)
		to := []string{req.EmailAddr}

		err := NewGmailSender("e-mall", os.Getenv("EMAIL_ADDR"), os.Getenv("EMAIL_KEY")).SendEmail(subject, content, to, nil, nil, nil)
		if err != nil {
			fmt.Printf("failed to send verify email: %s", err)
		}
		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

var group sendEmailConsumerGroup

func Consumer4SendEmail() {
	for {
		fmt.Println("for loop!")
		err := SendEmailKafkaConsumer.Consume(context.Background(), []string{"email"}, group)

		if err != nil {
			break
		}

	}

	_ = SendEmailKafkaConsumer.Close()
}
