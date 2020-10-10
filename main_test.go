package main

import (
	"context"
	pb "github.com/fjlyx97/short_url/proto"
	"google.golang.org/grpc"
	"log"
	"sync"
	"testing"
)

const (
	addr = "localhost:8002"
)

func Test_server_SetShortUrl(t *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewShortUrlServiceClient(conn)
	w := sync.WaitGroup{}

	threads := 100
	w.Add(threads)
	for i := 0; i < threads; i++ {
		go func() {
			r, err := c.SetShortUrl(context.Background(), &pb.SetUrlReq{
				Url: "www.baidu.com",
			})
			if err != nil {
				t.Error(err)
			}
			t.Log(r)
			t.Log(r.GetCode())
			t.Log(r.GetShortUrl())
			w.Done()
		}()
	}
	w.Wait()
}
