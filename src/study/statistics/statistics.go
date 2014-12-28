package main

import (
    "sort"
    "extensions/mathext"
    )

type statistics struct {
    numbers []float64
    mean float64
    median float64
}

func NewStatistics(numbers []float64) *statistics {
    var stats statistics
    stats.numbers = numbers
    if l := len(numbers); l > 0 {
        sort.Float64s(numbers)
        stats.mean = mathext.Avg(numbers)
        stats.median = mathext.Median(numbers)
    }
    return &stats
}
