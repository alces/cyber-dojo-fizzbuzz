package fizzbuzz

import (
    "fmt"
    "testing"
)

func TestFizzBuzz(t *testing.T) {
    expected := "2"
    actual := FizzBuzz(2)
    
    if actual != expected {
        t.Error(fmt.Sprintf("Expected %#v while actual is %#v", expected, actual))
    }
}
