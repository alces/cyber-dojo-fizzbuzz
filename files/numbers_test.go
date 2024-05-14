package fizzbuzz

import (
    "fmt"
    "testing"
)

func TestNumbersLen(t *testing.T) {
    expected := 10
    modificator := func(_ int) string {
        return ""
    }

    if actual := len(numbers(expected, modificator)); actual != expected {
        t.Error(fmt.Sprintf("Expected %v while actual %v", expected, actual))
    }
}