
package main

import (
    "fmt"
)

func fibonacci() {
    var toggle bool = false
    var num1 int = 1
    var num2 int = 1
    fmt.Println(num1)
    fmt.Println(num2)

    for i := 0; i < 10; i++ {
        if toggle {
            num1 = num1 + num2
            fmt.Println(num1)
            toggle = false
        } else {
            num2 = num1 + num2
            fmt.Println(num2)
            toggle = true
        }
    } 
}
