package minornumber

func Less(numbers ...int) int {
	minor := numbers[0]
	for _, v := range numbers {

		if minor > v {
			minor = v
		}
	}
	return minor
}
