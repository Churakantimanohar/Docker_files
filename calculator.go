package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    for {
        // Read input from the user
        fmt.Print("Enter calculation (or 'exit' to quit): ")
        text, _ := reader.ReadString('\n')
        text = strings.TrimSpace(text)

        if text == "exit" {
            break
        }

        // Split the input into parts
        parts := strings.Split(text, " ")
        if len(parts) != 3 {
            fmt.Println("Invalid input. Try again.")
            continue
        }

        // Convert operands to integers
        left, err := strconv.Atoi(parts[0])
        if err != nil {
            fmt.Println("Invalid input. Try again.")
            continue
        }
        right, err := strconv.Atoi(parts[2])
        if err != nil {
            fmt.Println("Invalid input. Try again.")
            continue
        }

        // Perform calculation
        var result int
        switch parts[1] {
        case "+":
            result = left + right
        case "-":
            result = left - right
        case "*":
            result = left * right
        case "/":
            if right == 0 {
                fmt.Println("Division by zero is not allowed.")
                continue
            }
            result = left / right
        default:
            fmt.Println("Invalid operator. Try again.")
            continue
        }

        // Print the result
        fmt.Printf("Result: %d\n", result)
    }
}

