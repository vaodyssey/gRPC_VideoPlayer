package server

import (
	"context"
	pb "gRPC_BE/videoExtractor"
	"log"
	"net"

	"google.golang.org/grpc"
)

// var (
// 	// tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
// 	certFile   = flag.String("cert_file", "", "The TLS cert file")
// 	keyFile    = flag.String("key_file", "", "The TLS key file")
// 	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
// 	port       = flag.Int("port", 50051, "The server port")
// )

type videoStreamServiceServer struct {
	pb.UnimplementedVideoStreamServiceServer
}

func (s *videoStreamServiceServer) GetVideoBuffer(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{OutputVideo: []byte("ahihi")}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	videoStreamServiceServer := &videoStreamServiceServer{}
	pb.RegisterVideoStreamServiceServer(s, videoStreamServiceServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
