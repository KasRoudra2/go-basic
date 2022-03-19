package main

import "fmt"

func main() {
    var x int // constants cannot be changed, so you can't use constant in storing input values. We can use shorthand :=
    
    fmt.Println("Enter your age: ")
    fmt.Scanln(&x) // &x is the memory address of x where the input is stored 
    if x>18 { // No parentheses required in if block
        fmt.Println("Adult")
    } else { // Else should be in this line
        fmt.Println("Young")
    }
}