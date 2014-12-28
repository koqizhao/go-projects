package main

import ("os"; "log"; "fmt"; "path/filepath")
import "extensions/strext"

var digitParts [10][7]string = [10][7]string {
    {
        "  000  ",
        " 0   0 ",
        "0     0",
        "0     0",
        "0     0",
        " 0   0 ",
        "  000  ",
    },
    {
        " 1 ",
        "11 ",
        " 1 ",
        " 1 ",
        " 1 ",
        " 1 ",
        "111",
    },
    {
        " 222 ",
        "2   2",
        "   2 ",
        "  2  ",
        " 2   ",
        "2    ",
        "22222",
    },
    {
        "  333 ",
        " 3   3",
        "     3",
        "    33",
        "     3",
        " 3   3",
        "333   ",
    },
    {
        "   4  ",
        "  44  ",
        " 4 4  ",
        "4  4  ",
        "444444",
        "   4  ",
        "   4  ",
    },
    {
        "55555",
        "5    ",
        "5    ",
        " 555 ",
        "    5",
        "5   5",
        " 555 ",
    },
    {
        "666  ",
        "6    ",
        "6    ",
        "66666",
        "6   6",
        "6   6",
        "66666",
    },
    {
        "77777",
        "    7",
        "   7 ",
        "  7  ",
        " 7   ",
        "7    ",
        "7    ",
    },
    {
        " 888 ",
        "8   8",
        "8   8",
        " 888 ",
        "8   8",
        "8   8",
        " 888 ",
    },
    {
        " 9999",
        "9   9",
        "9   9",
        " 9999",
        "    9",
        "    9",
        "    9",
    },
}

func main() {
    if len(os.Args) == 1 {
        fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }

    stringOfDigits := os.Args[1]
    digitCount := len(stringOfDigits)
    digits := make([]int, digitCount)
    for i := 0; i < digitCount; i++ {
        digit := stringOfDigits[i] - '0'
        if digit < 0 || digit > 9 {
            log.Fatal("Invalid whole number")
        }

        digits[i] = int(digit)
    }

    const DIGIT_SEPARATOR = " "
    const FIRST_LAST_LINE_DIGIT_SEPARATOR = "*"
    const SEPARATOR_SPAN = 2

    lineLength := 0
    for _, v := range digits {
        if lineLength > 0 {
            lineLength += SEPARATOR_SPAN
        }
        lineLength += len(digitParts[v][0])
    }
    first_last_line := strext.Mul(FIRST_LAST_LINE_DIGIT_SEPARATOR, lineLength)

    fmt.Println(first_last_line)

    digitSeparator := strext.Mul(DIGIT_SEPARATOR, SEPARATOR_SPAN)
    for i, l := 0, len(digitParts[0]); i < l; i++ {
        line := ""
        for _, v := range digits {
            if line != "" {
                line += digitSeparator
            }
            line += digitParts[v][i]
        }
        fmt.Println(line)
    }

    fmt.Println(first_last_line)
}
