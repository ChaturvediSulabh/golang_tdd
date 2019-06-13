package iteration

func repeat(char string, count int) (newChar string) {
	for i := 0; i < count; i++ {
		newChar += char
	}
	return
}
