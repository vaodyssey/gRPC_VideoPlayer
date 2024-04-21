using Grpc.Net.Client;
using Microsoft.AspNetCore.Mvc;
using gRPC_FE;
using gRPC_FE.Requests;
using gRPC_FE.Constants;

namespace gRPC_FE.Client
{
    public class VideoClient
    {
        public async Task<byte[]> GetVideoBuffer(VideoBufferRequest videoBufferRequest)
        {
            using var channel = GetGrpcChannel();
            var client = new VideoStreamService.VideoStreamServiceClient(channel);
            var reply = await client.GetVideoBufferAsync(new Request()
            {
                StartTimeSeconds = videoBufferRequest.StartTime,
                DurationSeconds = videoBufferRequest.Duration
            });
            return reply.VideoBuffer.ToArray(); 
        }
        private GrpcChannel GetGrpcChannel()
        {
            return GrpcChannel.ForAddress("http://localhost:3333",new GrpcChannelOptions {
                MaxReceiveMessageSize = gRPCStats.MAX_RECEIVE_MESSAGE_SIZE, 
                MaxSendMessageSize = gRPCStats.MAX_SEND_MESSAGE_SIZE
            });
        }
    }
}
