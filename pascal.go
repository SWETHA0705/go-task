package main

import (
    "fmt"
)

func pas(n int) int {
    var s int
    for s = 1; n > 1; n-- {
        s = s * n
    }
    return s
}

func pascal(i, j int) int {
    return pas(i) / (pas(i-j) * pas(j))
}
func main() {
    var i int
    n := 5
    for i = 0; i <= n; i++ {
        for j := 0; j <= n-i; j++ {
            fmt.Printf("  ")
        }

        for j := 0; j <= i; j++ {
            fmt.Printf(" %3d", pascal(i, j))
        }
        fmt.Println()
    }
    
}
