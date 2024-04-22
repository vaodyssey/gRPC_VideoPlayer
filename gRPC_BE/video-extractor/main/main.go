package main

import (
	"context"
	sgp "gRPC_BE/subGenProto"
	"gRPC_BE/utils"
	ve "gRPC_BE/videoExtractor"
	"log"
	"net"
	"os"
	"os/exec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var sgpClient sgp.SubtitleGeneratorClient

type videoStreamServiceServer struct {
	ve.UnimplementedVideoStreamServiceServer
}

func (s *videoStreamServiceServer) GetVideoBuffer(ctx context.Context, request *ve.Request) (*ve.Response, error) {
	inputPath := "../demo_videos/demo.mp4"
	outputPath := "../demo_videos/output.mp4"
	RemoveExistingOutput(outputPath)
	startTime := utils.SecondsToTimeString(int(request.StartTimeSeconds), "15:04:05")
	endTime := utils.GetEndTime(int(request.StartTimeSeconds), int(request.DurationSeconds), "15:04:05")
	videoBytes := request.InputVideo
	err := os.WriteFile(inputPath, videoBytes, 0644)
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
	maxSizeOption := grpc.MaxCallRecvMsgSize(1024 * 1024 * 100)
	responseVideo, err := sgpClient.Generate(context.Background(), &sgp.InputVideo{
		Video: videoSection,
	}, maxSizeOption)
	if err != nil {
		log.Fatalf("failed to add subtitle to this video. %v", err)
	}
	return &ve.Response{OutputVideo: responseVideo.Video}, nil

}
func RemoveExistingOutput(outputPath string) {

	if _, err := os.Stat(outputPath); err == nil {
		err := os.Remove(outputPath)
		if err != nil {
			log.Fatalf("Error removing file: %v", err) //print error if file is not removed
		}
	}

}

func InitializeVideoExtractorServer() {
	lis, err := net.Listen("tcp", ":3333")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 100),
		grpc.MaxSendMsgSize(1024 * 1024 * 100), // Set max receive message size to 10MB
	}
	s := grpc.NewServer(opts...)
	videoStreamServiceServer := &videoStreamServiceServer{}
	ve.RegisterVideoStreamServiceServer(s, videoStreamServiceServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func InitializeSubtitleGenerationServer() sgp.SubtitleGeneratorClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error dialing to Subtitle Generation Server: %v", err) //print error if file is not removed
	}
	return sgp.NewSubtitleGeneratorClient(conn)
}

func main() {
	sgpClient = InitializeSubtitleGenerationServer()
	InitializeVideoExtractorServer()
}
