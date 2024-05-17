### lab1 golang basic - generates a random string
vim generates_random_string.go
```
package main

import (
    "fmt"
    "math/rand"
)

var letters = []rune("m2uH8vZpA6fTNW9RuZsJ8r7gzk2UEL4Gjh42Q8R9f3T5yCgqmxbfHWDPYkYGSNuC")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func main() {
    fmt.Println(randSeq(6))
}
```
