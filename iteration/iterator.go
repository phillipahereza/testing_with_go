package iteration

// Repeat prints input 5 times
func Repeat(character string, repititions int) (repeated string) {
	for i := 0; i < repititions; i++ {
		repeated = repeated + character
	}
	return

}
