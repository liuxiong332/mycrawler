package main

import (
	"context"
	"crawler/engine"
	"crawler/parser"
	"crawler/rpc"
	"log"
	"net"
	"reflect"

	"github.com/golang/protobuf/ptypes"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct {
	rpc.UnimplementedCrawlerWorkServer
}

func convertReqInfo(req *rpc.WorkerRequest) parser.RequestInfo {
	target := parser.RequestInfo{}

	target.Url = req.Url
	switch req.Type {
	case "Person":
		target.Parser = parser.ParsePersonRes
	case "PersonBrief":
		target.Parser = parser.ParsePersonBriefRes
	case "Region":
		target.Parser = parser.ParseRegionRes
	}
	return target
}

func convertResult(res []parser.RequestInfo) *rpc.WorkerResult {
	retRes := rpc.WorkerResult{}

	payload := rpc.PersonPayload{}
	payloadRes, err := ptypes.MarshalAny(&payload)
	if err != nil {
		log.Printf("Error %v\n", err)
	}

	retRes.Payload = append(retRes.Payload, payloadRes)
	for _, m := range res {
		req := rpc.RequestInfo{}
		reflect.Copy(reflect.ValueOf(req), reflect.ValueOf(m))
		retRes.Requests = append(retRes.Requests, &req)
	}
	return &retRes
}

func (s *Server) Process(ctx context.Context, req *rpc.WorkerRequest) (*rpc.WorkerResult, error) {
	return convertResult(engine.WorkRequest(convertReqInfo(req))), nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rpc.RegisterCrawlerWorkServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
