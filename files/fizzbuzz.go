package fizzbuzz

import (
    "fmt"
)

func FizzBuzz(num int) string {
    if num%3 == 0 {
        return "Fizz"
    }
    
    if num%5 == 0 {
        return "Buzz"
    }
    
    return fmt.Sprintf("%d", num)
}
