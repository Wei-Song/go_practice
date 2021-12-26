package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	v1 "go_practice/go_practice/src/week4/apis/shop/v1"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	service, err := InitUserService()
	if err != nil {
		log.Panicf("service init fail: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterUserServerServer(s, service)

	// 信号控制
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		<-ctx.Done()
		log.Println("shutting down server...")
		s.GracefulStop()
		return nil
	})

	g.Go(func() error {
		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			return errors.Wrap(err, "start server port :8080")
		}
		log.Println("grpc server will list :8080")
		return s.Serve(l)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	log.Printf("errgroup exiting: %+v\n", g.Wait())
}
