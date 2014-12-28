package main

import ("os"; "log"; "fmt"; "path/filepath")
import "extensions"

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

    const SEPERATOR = " "
    const SEPERATOR_SPAN = 2
    for i, l := 0, len(digitParts[0]); i < l; i++ {
        line := ""
        for _, v := range digits {
            if line != "" {
                line += extensions.MultiplyString(SEPERATOR, SEPERATOR_SPAN)
            }
            line += digitParts[v][i]
        }
        if line == "" {
            break
        }
        fmt.Println(line)
    }
}
