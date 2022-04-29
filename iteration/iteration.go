package iteration

func Repeat(input string, number int) string {
	ret := ""
	for i := 0; i < number; i++ {
		ret += input
	}
	return ret
}
