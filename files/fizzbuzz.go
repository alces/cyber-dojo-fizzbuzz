package fizzbuzz

import (
    "strconv"
)

func FizzBuzz(num int) (result string) {
    if num%3 == 0 {
        result += "Fizz"
    }
    
    if num%5 == 0 {
        result += "Buzz"
    }
    
    if result == "" {
        result = strconv.Itoa(num)
    }
    
    return result
}
