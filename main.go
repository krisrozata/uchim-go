package main

import "fmt"

func main() {
    var n int
    fmt.Print("Въведи число: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        fmt.Println("Това не е число")
        return
    }

    if n <= 0 {
        fmt.Println("Грешка: числото трябва да е положително")
        return
    }

    for i := 1; i <= n; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
