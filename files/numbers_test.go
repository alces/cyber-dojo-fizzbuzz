package fizzbuzz

import (
    "fmt"
    "testing"
)

func TestNumbersLen(t *testing.T) {
    expected := 10

    if actual := len(numbers(expected)); actual != expected {
        t.Error(fmt.Sprintf("Expected %v while actual %v", expected, actual))
    }
}