package main

import (
	"log"
	"context"
	"sync"

	pb "github.com/wangxi19/crawler/proto"
	"google.golang.org/grpc"
)

func doWork(path string) error {

	return nil
}

const (
	crawlerNumber int = 10
)

func main() {
	defer func() {
		if err := recover(); nil != err {
			log.Fatalf("[ERROR]: %v\n", err)
		}
	}()

	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if nil != err {
		log.Fatalf("[ERROR]: %v\n", err)
	}
	c := pb.NewDispatcherClient(conn)

	wg := sync.WaitGroup{}
	wg.Add(crawlerNumber)
	for i := 0; i < crawlerNumber; i++ {
		url, err := c.GetURL(context.Background(), &pb.Option{})
		if nil != err {
			log.Fatalf("[ERROR]: %v\n", err)
		}
		path := url.GetPath()
		go func() {
			defer wg.Done()
			doWork(path)
		}()
	}

	wg.Wait()
}