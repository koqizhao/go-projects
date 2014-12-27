package main

import (
    "bufio"
    "os"
    "fmt"
    "math"
    "runtime"
    )

type polar struct {
    radius float64
    theta float64
}

type cartesian struct {
    x float64
    y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, or %s to quit."
const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

func init() {
    if runtime.GOOS == "windows" {
        prompt = fmt.Sprintf(prompt, "[Ctrl+C]")
    } else {
        prompt = fmt.Sprintf(prompt, "[Ctrl+D]")
    }
}

func main() {
    questions := make(chan polar)
    defer close(questions)
    answers := make(chan cartesian)
    defer close(answers)

    go reply(questions, answers)
    ask(questions, answers)
}

func ask(questions chan polar, answers chan cartesian) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println(prompt)
    for {
        fmt.Printf("Radius and angle: ")
        line, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err)
            break
        }

        var radius, theta float64
        if _, err := fmt.Sscanf(line, "%f %f", &radius, &theta); err != nil {
            fmt.Fprintln(os.Stderr, err)
            fmt.Fprintln(os.Stderr, "invalid input")
            continue
        }

        questions <- polar { radius, theta }
        answer := <-answers

        fmt.Printf(result, radius, theta, answer.x, answer.y)
    }
}

func reply(questions chan polar, answers chan cartesian) {
    go func() {
        for {
            question := <-questions
            theta := question.theta * math.Pi / 180.0
            answers <- cartesian{
                x: question.radius * math.Cos(theta),
                y: question.radius * math.Sin(theta),
            }
        }
    }()
}
