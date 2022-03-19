package main

import "fmt"

func main() {
    i := 0
    for i<10 { // An alternative of while loop in other languages
        fmt.Println(i)
        i++
    }
    for j := 10; j<21; j++ { // No parentheses but simlilar to c
        fmt.Println(j)
    }
}