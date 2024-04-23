namespace gRPC_FE.Utils
{
    public static class TimeUtils
    {
        public static int TimeStringToInt(string timeString)
        {
            if (TimeSpan.TryParseExact(timeString, "hh\\:mm\\:ss", null, out TimeSpan timeSpan))
            {
                // Convert the TimeSpan to total seconds
                return (int)timeSpan.TotalSeconds;
            }
            else
            {
                // Handle invalid input
                return -1; // or throw an exception, return null, etc.
            }
        }
    }
}
