package main

import (
	"fmt"
	pb "github.com/syanhaiD/envoy-grpc-prac/pkg/verify"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	os.Exit(exec())
}

type VerifyService struct{}

var biServers = map[int32]pb.Verify_BiDirectionalServer{}
var clientIDs = map[int32]struct{}{}
var mutexForMap sync.RWMutex

func exec() int {
	listenPort, err := net.Listen("tcp", ":20050")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	server := grpc.NewServer()
	svc := &VerifyService{}
	pb.RegisterVerifyServer(server, svc)
	err = server.Serve(listenPort)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (vs *VerifyService) BiDirectional(srv pb.Verify_BiDirectionalServer) error {
	ctx := srv.Context()

	var clientID int32
	mutexForMap.Lock()
	for i := int32(1); i < 25; i++ {
		if _, exist := clientIDs[i]; !exist {
			clientIDs[i] = struct{}{}
			clientID = i
			break
		}
	}
	biServers[clientID] = srv
	mutexForMap.Unlock()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			mutexForMap.Lock()
			delete(biServers, clientID)
			delete(clientIDs, clientID)
			mutexForMap.Unlock()
			fmt.Println(err)
			break
		}
		if err != nil {
			code := status.Code(err)
			if code.String() == "Canceled" {
				mutexForMap.Lock()
				delete(biServers, clientID)
				delete(clientIDs, clientID)
				mutexForMap.Unlock()
				fmt.Println(err)
				break
			} else {
				fmt.Println(err)
			}
			continue
		}

		resp := pb.BiResponse{
			Free1: req.Free1,
			P1: req.P1,
			P2: req.P2,
		}
		/*
		for _, serv := range biServers {
			go func(s pb.Verify_BiDirectionalServer) {
				if err := s.Send(&resp); err != nil {
					fmt.Println(err)
				}
			}(serv)
		}
		 */
		mutexForMap.RLock()
		for _, s := range biServers {
			if err := s.Send(&resp); err != nil {
				fmt.Println(err)
			}
		}
		mutexForMap.RUnlock()
	}

	return nil
}
