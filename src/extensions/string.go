package extensions

import "strings"

func IsWhitespace(s string) bool {
    return strings.Trim(s, " \t\n\r") == ""
}
