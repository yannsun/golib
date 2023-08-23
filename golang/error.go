package golang

// PanicError 出错时进行panic
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
