package engine

import (
	"crawler/parser"
	"log"
	"os"
)

const InitialUrl = "http://www.zhenai.com/zhenghun"

func RunEngine() {
	runQueue := []parser.RequestInfo{
		{
			Url:    InitialUrl,
			Parser: parser.ParseRegionRes,
		},
	}

	reqChan := make(chan parser.RequestInfo)
	workerChan := make(chan []parser.RequestInfo)

	scheduler := NewScheduler(reqChan)
	go scheduler.Run()
	// 启动20个 worker goroutine 去开始爬虫任务
	for i := 0; i < 20; i += 1 {
		worker := NewWorker(scheduler, workerChan)
		go worker.Run()
	}

	var sendCount, recvCount = 0, 0
	for {
		var curReq parser.RequestInfo
		var inChan chan parser.RequestInfo
		if len(runQueue) > 0 {
			curReq = runQueue[0]
			inChan = reqChan
		}

		select {
		case inChan <- curReq:
			runQueue = runQueue[1:]
			sendCount += 1
		case reqs := <-workerChan:
			runQueue = append(runQueue, reqs...)
			recvCount += 1

			if sendCount == recvCount && len(runQueue) == 0 {
				log.Printf("Break the loop, send count %d, recv count %d", sendCount, recvCount)
				os.Exit(0)
			}
		}
	}
}
