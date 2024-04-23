using System.ComponentModel.DataAnnotations;
using System.Diagnostics.Metrics;
using System.Text.RegularExpressions;

namespace gRPC_FE.Validation
{
    public class IsVideoTimeValid : ValidationAttribute
    {
        private string validVideoTimeRegex = "^(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d$";
        public IsVideoTimeValid()
        {
        }

        public override bool IsValid(object value)
        {
            string val = value as string; ;
            return Regex.IsMatch(val, validVideoTimeRegex);
        }
    }
}
