namespace gRPC_FE.Requests
{
    public class VideoBufferRequest
    {
        public int StartTime { get; set; }  
        public int Duration { get; set; }
        public byte[] VideoBytes { get; set; } 
    }
}
