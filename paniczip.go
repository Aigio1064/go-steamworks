package steamworks

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
