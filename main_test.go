package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintf(w,"Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && 
		err != http.ErrServerClosed {
			log.Printf("failed")
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}


func TestMainFunc(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	eg,ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})
	in := "message"

	rsp, err := http.Get("http://0.0.0.0:8000/" + in)
	if err != nil {
		t.Errorf("failed")
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("aaaaaa")
	}
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
	cancel()

	if err := eg.Wait(); err != nil{
		t.Fatal(err)
	}
}