package operation

const (
	min = iota
	max
	average
	sort
	wordCount
)

func Operations() map[string]int {
	return map[string]int{
		"min":        min,
		"max":        max,
		"average":    average,
		"sort":       sort,
		"word_count": wordCount,
	}
}

func IsDefined(op string) bool {
	if _, ok := Operations()[op]; ok {
		return true
	}

	return false
}
