package main

import "fmt"

func proveriChislo(n int) (string, error) {
    if n == 0 {
        return "", fmt.Errorf("нулата не е нито четна, нито нечетна")
    }
    if n%2 == 0 {
        return "четно", nil
    }
    if n%2 == 1 {
        return "нечетно", nil
    }
    return "", fmt.Errorf("числото %d е отрицателно", n)
}

func main() {
    var n int
    fmt.Print("Въведи число: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        fmt.Println("Това не е число")
        return
    }

    rez, err := proveriChislo(n)
    if err != nil {
        fmt.Println("Грешка:", err)
        return
    }
    fmt.Println(n, "е", rez)
}
