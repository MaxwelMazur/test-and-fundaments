package iteration

func repeat(character string) (characterFull string) {
	for i := 0; i < 4; i++ {
		characterFull += character
	}
	return
}
