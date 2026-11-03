package main

import "fmt"



func printValue(value any) {
    text, ok := value.(int)

    if ok {
        fmt.Println("Integer:", text)
    } else {
        fmt.Println("Not a Integer")
    }
}
func main() {
    
    printValue("test")
    printValue(12)
    
}