package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // Command to install Go packages
    var cmd = exec.Command("go", "install")
    
    // Outputting the command's status
    stdout, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    output := string(stdout)

    // Formatting the output to match the Go code format
    formattedCode := `package main

import (
    "fmt"
    "os/exec"
)

func main() {

    var cmd = exec.Command("go", "install")

    stdout, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    output = string(stdout)

    // Formatting the output to match the Go code format
    formattedCode = `package main

import (
    "fmt"
    "os/exec"
)

func main() {

    var cmd = exec.Command("go", "install")

    stdout, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    output = string(stdout)

    // Formatting the output to match the Go code format
    formattedCode = `

package main

import (
    "fmt"
    "os/exec"
)

func main() {

    var cmd = exec.Command("go", "install")

    stdout, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    output = string(stdout)

}

