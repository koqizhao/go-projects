package strext

import "strings"
import "fmt"

func IsWhitespace(s string) bool {
    return strings.Trim(s, " \t\n\r") == ""
}

func Mul(s string, count int) string {
    if count == 0 {
        return ""
    }

    var result = ""
    for i := 0; i < count; i++ {
        result += s
    }

    return result
}

func ToString(data interface{}) string {
    return fmt.Sprint(data)
}

func Join(separator string, data ...interface{}) string {
    result := ""
    for i, l := 0, len(data); i < l; i++ {
        if result != "" {
            result += separator
        }
        result += ToString(data[i])
    }
    return result
}
