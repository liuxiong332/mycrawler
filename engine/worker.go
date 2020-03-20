package engine

import (
	"crawler/parser"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func WorkRequest(req parser.RequestInfo) []parser.RequestInfo {
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

type WorkerNotifier interface {
	WorkerReady(chan parser.RequestInfo)
}

type WorkerRunner interface {
	Run()
}

type Worker struct {
	InputChan chan parser.RequestInfo
	OutChan   chan<- []parser.RequestInfo
	Notifier  WorkerNotifier
}

func NewWorker(notifier WorkerNotifier, outChan chan<- []parser.RequestInfo) WorkerRunner {
	w := Worker{}
	w.InputChan = make(chan parser.RequestInfo)
	w.Notifier = notifier
	w.OutChan = outChan
	return &w
}

func (w *Worker) Run() {
	for {
		w.Notifier.WorkerReady(w.InputChan)
		req := <-w.InputChan
		outRes := WorkRequest(req)
		w.OutChan <- outRes
	}
}
