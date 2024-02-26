package internal

// Checks for errors
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
