package engine

import (
	"crawler/parser"
	"log"
)

type Scheduler struct {
	reqChan    <-chan parser.RequestInfo
	workerChan chan chan parser.RequestInfo
}

func NewScheduler(reqChan <-chan parser.RequestInfo) *Scheduler {
	s := Scheduler{}
	s.reqChan = reqChan
	s.workerChan = make(chan chan parser.RequestInfo)
	return &s
}

func (s *Scheduler) WorkerReady(r chan parser.RequestInfo) {
	log.Printf("Scheduler worker ready")
	s.workerChan <- r
}

func (s *Scheduler) Run() {
	reqQueue := make([]parser.RequestInfo, 0)
	var workChanQueue []chan parser.RequestInfo

	log.Printf("Scheduler start run")
	for {
		var req parser.RequestInfo
		var curWorkChan chan parser.RequestInfo

		log.Printf("Scheduler Check queue %d, %d\n", len(reqQueue), len(workChanQueue))
		if len(reqQueue) > 0 && len(workChanQueue) > 0 {
			req = reqQueue[0]
			curWorkChan = workChanQueue[0]
		}

		select {
		case newReq := <-s.reqChan:
			reqQueue = append(reqQueue, newReq)
			log.Printf("Scheduler Get requests, %d\n", len(reqQueue))
		case curChan := <-s.workerChan:
			workChanQueue = append(workChanQueue, curChan)
			log.Printf("Scheduler Chan ready, chan queue %d\n", len(workChanQueue))
		case curWorkChan <- req:
			reqQueue = reqQueue[1:]
			workChanQueue = workChanQueue[1:]
			log.Printf("Scheduler Consume one req, %d, %d\n", len(reqQueue), len(workChanQueue))
		}
	}
}
