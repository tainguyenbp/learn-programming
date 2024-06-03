
package main

import "fmt"

func main() {
    // This line will be executed immediately.
    defer fmt.Println("This runs after main")
    
    // This line will be executed immediately.
    fmt.Println("Main ended")
    // After the surrounding function returns, 
    // the deferred function will execute.
}
