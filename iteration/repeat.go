package repeat

func Repeat(input string, amount int) (repeated string) {
	for i := 0; i < amount; i++ {
		repeated += input
	}

	return
}
