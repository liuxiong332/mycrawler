package main

import (
	"context"
	"crawler/dist_engine/util"
	"crawler/parser"
	"crawler/rpc"
	"log"
	"time"

	"google.golang.org/grpc"
)

type Connector struct {
	conn    *grpc.ClientConn
	client  rpc.CrawlerWorkClient
	context context.Context
}

func NewConnector() *Connector {
	connector := Connector{}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	connector.conn = conn
	//defer conn.Close()
	connector.client = rpc.NewCrawlerWorkClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Hour)
	connector.context = ctx
	return &connector
}

func convertReq(reqInfo parser.RequestInfo) *rpc.WorkerRequest {

	return &rpc.WorkerRequest{
		Url:  reqInfo.Url,
		Type: util.ConvertReqParser(reqInfo.Parser),
	}
}

func convertRes(result *rpc.WorkerResult) []parser.RequestInfo {
	var reqs []parser.RequestInfo
	for _, r := range result.Requests {
		reqs = append(reqs, parser.RequestInfo{
			Url:    r.Url,
			Parser: util.ConvertReqType(r.Type),
		})
	}
	return reqs
}

func (connector *Connector) Process(reqInfo parser.RequestInfo) []parser.RequestInfo {
	r, err := connector.client.Process(connector.context, convertReq(reqInfo))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return nil
	}
	return convertRes(r)
}

func (connector *Connector) close() {
	connector.close()
}
