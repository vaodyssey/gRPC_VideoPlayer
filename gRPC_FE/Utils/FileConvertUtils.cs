using gRPC_FE.Constants;
using Microsoft.AspNetCore.Components.Forms;
using Microsoft.AspNetCore.Mvc.Formatters;

namespace gRPC_FE.Utils
{
    public static class FileConvertUtils
    {
        public static async Task<byte[]> ConvertIBrowserFileToByteArray(IBrowserFile browserFile)
        {

            var fileStream = browserFile.OpenReadStream(maxAllowedSize:VideoInfo.MAX_VIDEO_SIZE);
            using (var resultStream = new MemoryStream())
            {
                await fileStream.CopyToAsync(resultStream);
                return resultStream.ToArray();
            }
        }
        public static Task<string> ConvertByteArrayToBase64(string mediaType, byte[] data)
        {
            return Task.Run(() =>
            {
                var result = $"data:{mediaType};base64,{Convert.ToBase64String(data)}";
                return result;
            });
        }
    }
}
