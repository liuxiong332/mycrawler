package main

import (
	"crawler/engine"
)

func main() {
	engine.RunEngine(engine.NewWorker)
}
