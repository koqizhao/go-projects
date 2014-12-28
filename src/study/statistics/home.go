package main

import (
    "net/http"
    "fmt"
    "strconv"
    "strings"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
    var numbersString, errorClass, errorMessage string
    var count int
    var mean, median float64
    if err := request.ParseForm(); err != nil {
        errorMessage = err.Error()
    } else {
        count, mean, median, numbersString, errorMessage = getStatistics(request)
    }

    if errorMessage != "" {
        errorClass = "show"
    }

    fmt.Fprint(writer, _HTML_TOP)
    fmt.Fprintf(writer, _HTML_FORM, numbersString, errorClass, errorMessage)
    fmt.Fprintf(writer, _HTML_RESULT_TABLE, numbersString, count, mean, median)
    fmt.Fprint(writer, _HTML_BOTTOM)
}

func getStatistics(request *http.Request) (count int, mean, median float64, numbersString, errorMessage string) {
    if data, found := request.Form["numbers"]; found {
        numbersString = data[0]
    } else {
        return count, mean, median, numbersString, errorMessage
    }

    numbersString = strings.Replace(numbersString, ",", " ", -1)
    var numbers []float64
    for _, field := range strings.Fields(numbersString) {
        if v, err := strconv.ParseFloat(field, 64); err == nil {
            numbers = append(numbers, v)
        }
    }

    count = len(numbers)
    if count == 0 {
        return count, mean, median, numbersString, errorMessage
    }

    stats := NewStatistics(numbers)
    return count, stats.mean, stats.median, numbersString, errorMessage
}
