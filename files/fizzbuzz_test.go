package fizzbuzz

import (
    "fmt"
    "testing"
)

type testResult struct {
    argument int
    expected string
}

testData = []testResults{
    {argument: 2, expected: "2"},    
}

func TestFizzBuzz(t *testing.T) {
    expected := "2"
    actual := FizzBuzz(2)
    
    if actual != expected {
        t.Error(fmt.Sprintf("Expected %#v while actual is %#v", expected, actual))
    }
}
