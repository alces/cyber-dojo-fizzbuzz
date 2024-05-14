package fizzbuzz

import (
    "fmt"
)

func FizzBuzz(num int) string {
    result := ""
 
    if num%3 == 0 {
        result += "Fizz"
    }
    
    if num%5 == 0 {
        result += "Buzz"
    }
    
    if result == "" {
        result = fmt.Sprintf("%d", num)
    }
    
    return result
}
