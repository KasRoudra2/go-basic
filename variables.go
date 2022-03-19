package main 

import "fmt" 

func main(){
    var i int // Integer Type

    var x float32 //32 Byte Float

    var y string // String type
    
    // x is defined but no value is asigned, so it will be outputted as blank
    
    i = 15
    
    y = "Hi"

    var m = 10 // Type can be omitted if a initial value is given in declaration

    n := 10 // Shorthand variable declaration in go

    const c = "String" // Constants in go must be declared with a initial value
    fmt.Println(i, x, y, m, n, c)
}