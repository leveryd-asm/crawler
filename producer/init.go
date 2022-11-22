package main

import "os"

// get kafka writer using environment variables.
var (
	assetEndpoint = os.Getenv("assetEndpoint")
	kafkaURL      = os.Getenv("kafkaURL")
	topic         = os.Getenv("topic")
)
