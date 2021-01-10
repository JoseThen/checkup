package utils

// ErrorCheck ... utility function to help stream line error checking
func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
