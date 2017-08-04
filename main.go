package main

import (
	"fmt"
)

// Greet returns a greeting message for a name
func greet(name *string) string {
	return fmt.Sprintf("Hello %s!", *name)
}


func main() {
        var name string
        fmt.Print("What is your name?: ")
        _, err := fmt.Scanf("%s", &name)
        if err != nil {
                fmt.Println("Could not get name")
        }
        fmt.Println(greet(&name))
}

