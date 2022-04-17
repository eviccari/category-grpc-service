package utils

// PanicIfError - Stop service execution if critical error occurs
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}