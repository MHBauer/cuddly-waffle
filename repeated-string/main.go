/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
// how many times does a occur in the string.
numAinString := int64(strings.Count(s, "a"))
// how many times does the string length fit into n
base := n /int64(len(s));
rem := n % int64(len(s));

// for the remainder, count a's

return base * numAinString + int64(strings.Count(s[:rem], "a"))
}

