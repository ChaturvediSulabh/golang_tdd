package iteration

const count = 5

func repeat(char string) (newChar string) {
	for i := 0; i < count; i++ {
		newChar += char
	}
	return
}
