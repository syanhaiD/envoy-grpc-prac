package main

import (
	"context"
	"fmt"
	pb "github.com/syanhaiD/envoy-grpc-prac/pkg/verify"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

func main() {
	os.Exit(exec())
}

func exec() int {
	target := ":20050"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// create stream
	grpcClient := pb.NewVerifyClient(conn)
	var wg sync.WaitGroup
	clients := []string{"c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c10"}
	// clients = []string{"c1"}
	grpcConns := map[string]pb.Verify_BiDirectionalClient{}
	for _, clientName := range clients {
		grpcConns[clientName], err = grpcClient.BiDirectional(context.Background())
		if err != nil {
			fmt.Println(err)
		}
	}
	// recv
	/*
	for _, gConn := range grpcConns {
		go func(c pbfgl.Fgl_BiDirectionalClient) {
			for {
				resp, err := c.Recv()
				if err == io.EOF {
					fmt.Println("server EOF")
					break
				}
				if err != nil {
					fmt.Println("server EOF")
					break
				}
				fmt.Println(resp)
			}
		}(gConn)
	}
	*/
	procNum := 15 * 120
	for clientName, gConn := range grpcConns {
		wg.Add(1)
		go func(c pb.Verify_BiDirectionalClient, cn string) {
			for i := 0; i < procNum; i++ {
				req := pb.BiRequest{
					Free1: cn,
					P1: &pb.Packet1{
						Free1: float32(time.Now().Unix()),
						Free2: 0.1,
						Free3: 0.2,
					},
				}
				if err := c.Send(&req); err != nil {
					fmt.Println(err)
				}
				time.Sleep(time.Millisecond * 66)
			}
			wg.Done()
			if err := c.CloseSend(); err != nil {
				fmt.Println(err)
			}
		}(gConn, clientName)
	}
	wg.Wait()

	return 0
}

