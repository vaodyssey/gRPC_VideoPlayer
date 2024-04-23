using Grpc.Net.Client;
using Microsoft.AspNetCore.Mvc;
using gRPC_FE;
using gRPC_FE.Requests;
using gRPC_FE.Constants;
using Google.Protobuf;
using gRPC_FE.Utils;

namespace gRPC_FE.Client
{
    public class VideoClient
    {
        private VideoRequest _videoRequest;
        public async Task<byte[]> GetVideoBuffer(VideoRequest videoRequest)
        {
            _videoRequest = videoRequest;
            using var channel = GetGrpcChannel();
            var client = new VideoStreamService.VideoStreamServiceClient(channel);
            var reply = await client.GetVideoBufferAsync(new Request()
            {
                StartTimeSeconds = await GetStartTimeSeconds(),
                DurationSeconds = await GetDurationSeconds(),
                InputVideo = ByteString.CopyFrom(videoRequest.VideoBytes),
            });
            return reply.OutputVideo.ToArray();
        }
        private GrpcChannel GetGrpcChannel()
        {
            return GrpcChannel.ForAddress("http://localhost:3333", new GrpcChannelOptions
            {
                MaxReceiveMessageSize = gRPCStats.MAX_RECEIVE_MESSAGE_SIZE,
                MaxSendMessageSize = gRPCStats.MAX_SEND_MESSAGE_SIZE
            });
        }
        
        private Task<int > GetStartTimeSeconds() {
            return Task.Run(() =>
            {
                return TimeUtils.TimeStringToInt(_videoRequest.StartTime);
            });
        }
        private Task<int> GetDurationSeconds()
        {
            return Task.Run(() =>
            {
                return TimeUtils.TimeStringToInt(_videoRequest.EndTime) - TimeUtils.TimeStringToInt(_videoRequest.StartTime);
                });
        }
    }
}
