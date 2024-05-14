package fizzbuzz

func numbers(length int, modificator func (int) string) []string {
    return make([]string, length)
}