package main // Need to import in regular go files. A Go program starts running in the main package.

import "fmt" //import the fmt package for using Println function

func main(){ // main function declaration
    fmt.Println("Hello World") // Println is an exported function of package fmt
}