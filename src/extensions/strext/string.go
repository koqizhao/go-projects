package strext

import "strings"
import "fmt"
import "unicode"
import "unicode/utf8"
import "bytes"

func Contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}

func IsWhitespace(s string) bool {
	return strings.Trim(s, " \t\n\r") == ""
}

func Mul(s string, count int) string {
	if count == 0 {
		return ""
	}

	return strings.Repeat(s, count)
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

func PadLeft(number, width int, char rune) string {
	return pad(number, width, char, true)
}

func PadRight(number, width int, char rune) string {
	return pad(number, width, char, false)
}

func SimplifyWhitespace(s string) string {
	var buffer bytes.Buffer
	skip := true
	for _, char := range s {
		if unicode.IsSpace(char) {
			if !skip {
				buffer.WriteRune(' ')
				skip = true
			}
		} else {
			buffer.WriteRune(char)
			skip = false
		}
	}
	s = buffer.String()
	if skip && len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

func pad(number, width int, char rune, left bool) string {
	s := ToString(number)
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		if left {
			s = strings.Repeat(string(char), gap) + s
		} else {
			s += strings.Repeat(string(char), gap)
		}
	}

	return s
}
