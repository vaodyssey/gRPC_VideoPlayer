using gRPC_FE.Validation;
using System.ComponentModel.DataAnnotations;

namespace gRPC_FE.Requests
{
    public class VideoRequest
    {
        [Required(ErrorMessage = "You must input the start time.")]
        [IsVideoTimeValid]
        public string StartTime { get; set; }
        [Required(ErrorMessage = "You must input the end time.")]
        [IsVideoTimeValid]
        public string EndTime { get; set; } 
        public byte[] VideoBytes { get; set; } 
    }
}
