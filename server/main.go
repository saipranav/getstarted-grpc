package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/saipranav/getstarted-grpc-go/store"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8070")
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	osCh := make(chan os.Signal, 2)
	signal.Notify(osCh, os.Interrupt, syscall.SIGTERM)

	grpcServer := grpc.NewServer()
	store.RegisterStoreServer(grpcServer, &storeServer{})

	go stopGrpcServer(ctx, grpcServer, osCh)

	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}

func stopGrpcServer(ctx context.Context, s *grpc.Server, ch chan os.Signal) {
	select {
	case <-ctx.Done():
	case <-ch:
		s.GracefulStop()
	}
}

type storeServer struct {
}

func (s *storeServer) Save(ctx context.Context, in *store.Entity) (*store.EntityResponse, error) {
	fmt.Println("saved", in)
	res := store.EntityResponse{Entity: in}
	return &res, nil
}

func (s *storeServer) Restore(ctx context.Context, in *store.EntityRequest) (*store.EntityResponse, error) {
	fmt.Println("restored", in)
	res := store.EntityResponse{Entity: &store.Entity{Id: in.Id}}
	return &res, nil
}
