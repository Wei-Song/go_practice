package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type myserver struct{}

func (server myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {
	case "/":
		fmt.Fprintf(w, "get server")
	case "/test":
		fmt.Fprintf(w, "test")
	default:
		fmt.Fprintf(w, "unknow http")
	}

}

func signalf() error {

	signCh := make(chan os.Signal)

	signal.Notify(signCh)
	s := <-signCh
	fmt.Println("catch system signal", s)
	switch s {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
		return errors.New("find return signal,exit")
		//
	default:
		fmt.Println("other signal")
	}

	return nil
}

func main() {

	group, errctx := errgroup.WithContext(context.Background())

	var s myserver

	se := http.Server{
		Handler: s,
		Addr:    ":9090",
	}
	http.Handle("/", s)
	group.Go(func() error {
		defer fmt.Println("g1 return")
		return se.ListenAndServe()
	})

	group.Go(func() error {
		select {
		case <-errctx.Done():
			fmt.Println("g2 return")
			return se.Shutdown(errctx)
		}
		return nil
	})

	group.Go(func() error {
		err := signalf()
		if err != nil {
			fmt.Println("g3 return")
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("all goroutine are dead get errors:", err)
	}
}
