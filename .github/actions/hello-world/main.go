package main

import "fmt"
import "os"

func main() {
  // Access inputs as environment variables
  firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
  secondGreeting := os.Getenv("INPUT_SECONDGREETING")
  thirdGreeting := os.Getenv("INPUT_THIRDGREETING")

  // Use those inputs in the action logic
  fmt.Println("Hello " + firstGreeting)
  fmt.Println("Hello " + secondGreeting)

  // Sometimes inputs are not "required"
  if thirdGreeting != "" {
    fmt.Println("Hello " + thirdGreeting)
  }
}
