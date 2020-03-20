package main

import "crawler/engine"

const (
	address = "localhost:50051"
)

func main() {
	engine.RunEngine(NewWorker)
}
