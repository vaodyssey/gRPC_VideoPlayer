using Grpc.Net.Client;
using Microsoft.AspNetCore.Mvc;
using gRPC_FE;

namespace gRPC_FE.Client
{
    public class VideoClient
    {
        public async Task<byte[]> GetVideoBuffer()
        {
            using var channel = GetGrpcChannel();
            var client = new VideoStreamService.VideoStreamServiceClient(channel);
            var reply = await client.GetVideoBufferAsync(new Request()
            {
                StartTimeSeconds = 360,
                DurationSeconds = 15
            });
            return reply.VideoBuffer.ToArray(); 
        }
        private GrpcChannel GetGrpcChannel()
        {
            return GrpcChannel.ForAddress("http://localhost:3333",new GrpcChannelOptions {
                MaxReceiveMessageSize = 20 * 1024 * 1024, 
                MaxSendMessageSize = 20 * 1024 * 1024 
            });
        }
    }
}
