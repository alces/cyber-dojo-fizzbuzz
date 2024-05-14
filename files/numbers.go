package fizzbuzz

func numbers(length int, modificator func (int) string) []string {
    result := make([]string, length)

    for i := 0; i < length; i++ {
        result[i] = modificator(i + 1)
    }
    
    return result
}