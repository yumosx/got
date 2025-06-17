package errx

func Recover(fn func(err any)) func() {
	return func() {
		if err := recover(); err != nil {
			fn(err)
		}
	}
}
