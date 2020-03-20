package main

import (
	"crawler/engine"
	"crawler/parser"
	"log"
)

func NewWorker(notifier engine.WorkerNotifier, outChan chan<- []parser.RequestInfo) engine.WorkerRunner {
	distWorker := DistWorker{}
	distWorker.Worker = engine.NewWorker(notifier, outChan).(*engine.Worker)
	distWorker.Conn = NewConnector()
	return &distWorker
}

type DistWorker struct {
	Worker *engine.Worker
	Conn   *Connector
}

func (w *DistWorker) Run() {
	defer w.Conn.close()
	for {
		w.Worker.Notifier.WorkerReady(w.Worker.InputChan)
		req := <-w.Worker.InputChan
		outRes := w.Conn.Process(req)
		log.Printf("Get Response", outRes)
		w.Worker.OutChan <- outRes
	}
}
