package fizzbuzz

import (
    "fmt"
    "testing"
)

type testResult struct {
    argument int
    expected string
}

var testData = []testResult{
    {argument: 2, expected: "2"},
    {argument: 3, expected: "Fizz"},
    {argument: 11, expected: "11"},
}

func TestFizzBuzz(t *testing.T) {
    for _, result := range testData {
        actual := FizzBuzz(result.argument)
        if actual != result.expected {
            t.Error(fmt.Sprintf("Expected %#v while actual is %#v", result.expected, actual))
        }
    }
}
