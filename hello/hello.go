package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message, _ := greetings.Hello("Gladys")
    fmt.Println(message)
}
