package main

import "fmt"



func printValue(value any) {
    switch v := value.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Integer:", v)
    case bool:
        fmt.Println("Boolean:", v)
    default:
        fmt.Println("Unknown type")
    }
}


func main() {
    
    printValue("test")
    printValue(12)
    
}