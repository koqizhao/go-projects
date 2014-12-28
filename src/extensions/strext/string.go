package strext

import "strings"

func IsWhitespace(s string) bool {
    return strings.Trim(s, " \t\n\r") == ""
}

func Multiply(s string, count int) string {
    if count == 0 {
        return ""
    }

    var result = ""
    for i := 0; i < count; i++ {
        result += s
    }

    return result
}
