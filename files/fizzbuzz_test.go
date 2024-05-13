package fizzbuzz

import ("testing")

func TestFizzBuzz(t *testing.T) {
    expected := "2"
    actual := FizzBuzz(2)
    
    if actual != expected {
        t.Error("Expected", expected, "while actual", actual)
    }
}
