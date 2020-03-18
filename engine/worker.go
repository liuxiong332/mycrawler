package engine

import (
	"crawler/parser"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func workRequest(req parser.RequestInfo) []parser.RequestInfo {
	log, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}
	defer log.Sync()

	<-time.Tick(200 * time.Millisecond)
	resp, err := http.Get(req.Url)
	if err != nil {
		log.Error("Failed to Get Url", zap.String("Url", req.Url))
		return nil
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body ", zap.Error(err))
		return nil
	}
	res, err := req.Parser(bytes)
	if err != nil {
		log.Error("Failed to parse source", zap.Error(err))
		return nil
	}

	for _, m := range res.Payload {
		log.Info("Get Payload", zap.Reflect("payload", m))
		//if res, ok := m.(parser.PersonInfo); ok {
		//	log.Info("Get Payload", zap.Reflect("payload", res))
		//}
	}
	return res.Requests
}

type Worker struct {
	inputChan chan parser.RequestInfo
	outChan   chan<- []parser.RequestInfo
	notifier  WorkerNotifier
}

func NewWorker(notifier WorkerNotifier, outChan chan<- []parser.RequestInfo) *Worker {
	w := Worker{}
	w.inputChan = make(chan parser.RequestInfo)
	w.notifier = notifier
	w.outChan = outChan
	return &w
}

type WorkerNotifier interface {
	WorkerReady(chan parser.RequestInfo)
}

func (w *Worker) Run() {
	for {
		w.notifier.WorkerReady(w.inputChan)
		req := <-w.inputChan
		outRes := workRequest(req)
		w.outChan <- outRes
	}
}
