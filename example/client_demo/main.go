package main

import (
	"log"
	"log/slog"

	"github.com/elleqt/go-rtmp"

	rtmpmsg "github.com/elleqt/go-rtmp/message"
)

const (
	chunkSize = 128
)

func main() {
	client, err := rtmp.Dial("rtmp", "localhost:1935", &rtmp.ConnConfig{
		Logger: slog.Default(),
	})
	if err != nil {
		log.Fatalf("Failed to dial: %+v", err)
	}
	defer client.Close()

	if err := client.Connect(nil); err != nil {
		log.Fatalf("Failed to connect: Err=%+v", err)
	}
	log.Println("connected")

	stream, err := client.CreateStream(nil, chunkSize)
	if err != nil {
		log.Fatalf("Failed to create stream: Err=%+v", err)
	}
	defer stream.Close()

	if err := stream.Publish(&rtmpmsg.NetStreamPublish{
		PublishingName: "testtesttesttest",
		PublishingType: "live",
	}); err != nil {
		log.Fatalf("Failed to send publish message: Err=%+v", err)
	}

	log.Println("stream created")
}
