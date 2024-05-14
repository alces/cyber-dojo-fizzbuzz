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

func TestModifiedNumber(t *testing.T) {
    expected := "x1"
    modificator := func(a int) string {
        return fmt.Sprintf("x%d", a)
    }

    if actual := numbers(1, modificator)[0]; actual != expected {
        t.Error(fmt.Sprintf("Expected %#v while actual %#v", expected, actual))
    }
}

func TestAll(_ *testing.T) {
    for _, v := range numbers(100, FizzBuzz) {
        fmt.Println(v)
    }
}