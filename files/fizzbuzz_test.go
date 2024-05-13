package fizzbuzz

import (
    "fmt"
    "testing"
)

type testResult struct {
    argument int
    expected string
}

var testData = []testResults{
    {argument: 2, expected: "2"},    
}

func TestFizzBuzz(t *testing.T) {
    for _, result := testData {
        actual := FizzBuzz(result.argument)
        if actual != result.expected {
            t.Error(fmt.Sprintf("Expected %#v while actual is %#v", result.expected, actual))
        }
    }
}
