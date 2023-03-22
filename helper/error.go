package helper

func ErrorPanic(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

