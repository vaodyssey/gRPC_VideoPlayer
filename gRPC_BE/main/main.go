package main

import (
	"context"
	pb "gRPC_BE/protobuf"
	"gRPC_BE/utils"
	"log"
	"net"
	"os"
	"os/exec"

	"google.golang.org/grpc"
)

type videoStreamServiceServer struct {
	pb.UnimplementedVideoStreamServiceServer
}

func (s *videoStreamServiceServer) GetVideoBuffer(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	inputPath := "../demo_videos/demo.mp4"
	outputPath := "../demo_videos/output.mp4"
	RemoveExistingOutput(outputPath)
	startTime := utils.SecondsToTimeString(int(request.StartTimeSeconds), "15:04:05")
	endTime := utils.GetEndTime(int(request.StartTimeSeconds), int(request.DurationSeconds), "15:04:05")
	cmd := exec.Command("ffmpeg",
		"-i", inputPath,
		"-ss", startTime,
		"-to", endTime,
		outputPath,
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatalf("ffmpeg command failed: %v", err)
	}

	// Read the output file into memory
	videoSection, err := os.ReadFile(outputPath)
	if err != nil {
		log.Fatalf("failed to read video section: %v", err)
	}
	return &pb.Response{VideoBuffer: videoSection}, nil

}
func RemoveExistingOutput(outputPath string) {

	if _, err := os.Stat(outputPath); err == nil {
		err := os.Remove(outputPath)
		if err != nil {
			log.Fatalf("Error removing file: %v", err) //print error if file is not removed
		}
	}

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
