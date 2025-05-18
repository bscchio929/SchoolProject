package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate a random number between 1 and 10
    randomNumber := rand.Intn(10) + 1

    // Format the output to include the current date and time
    formattedOutput := fmt.Sprintf("Date: %s, Time: %d:%02d", time.Now().Format("2006-01-02"), time.Now().Hour(), time.Now().Minute())

    // Print the random number in a human-readable format
    fmt.Println("The randomly generated number is:", randomNumber)
    fmt.Println(formattedOutput)
}
