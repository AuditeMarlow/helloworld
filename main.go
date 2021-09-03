package main

import (
	"github.com/auditemarlow/helloworld/handler"
	pb "github.com/auditemarlow/helloworld/proto"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
