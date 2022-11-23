// https://github.com/segmentio/kafka-go/blob/main/examples/consumer-logger/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func main() {
	// get kafka reader using environment variables.
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	groupID := os.Getenv("groupID")

	proxy := os.Getenv("proxy")

	// fix: in k8s environment, katana's proxy option must be like a normal domain, so it can not be k8s service name.
	// we can try to treat XRAY_PROXY_SERVICE_PORT as http proxy
	if os.Getenv("XRAY_PROXY_SERVICE_PORT") != "" {
		proxy = os.Getenv("XRAY_PROXY_SERVICE_PORT")
		proxy = strings.ReplaceAll(proxy, "tcp://", "http://")
	}

	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		cmd := "/usr/local/bin/katana -v -u " + string(m.Value) + " -proxy " + proxy
		fmt.Println(cmd)
		output, err := exec.Command("/bin/sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println(err)
		}
		println(string(output))
	}
}
