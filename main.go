package main

import (
	"net"
	"os"
	pb "wordfilter/proto"

	log "github.com/Sirupsen/logrus"
	_ "github.com/gonet2/libs/statsd-pprof"
	"google.golang.org/grpc"
)

const (
	_port = ":50002"
)

func main() {
	// 监听
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	log.Info("listening on ", lis.Addr())

	// 注册服务
	s := grpc.NewServer()
	ins := &server{}
	ins.init()
	pb.RegisterWordFilterServiceServer(s, ins)

	// 开始服务
	s.Serve(lis)
}
