package fizzbuzz

import (
    "fmt"
    "testing"
)

var testData = []struct {
    argument int
    expected string
}{
    {argument: 2, expected: "2"},
    {argument: 3, expected: "Fizz"},
    {argument: 5, expected: "Buzz"},
    {argument: 6, expected: "Fizz"},
    {argument: 10, expected: "Buzz"},
    {argument: 13, expected: "13"},
}

func TestFizzBuzz(t *testing.T) {
    for _, res := range testData {        
        if actual := FizzBuzz(res.argument); actual != res.expected {
            t.Error(fmt.Sprintf("Expected %#v while actual is %#v", res.expected, actual))
        }
    }
}
