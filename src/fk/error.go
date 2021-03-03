package fk

func Check(v error) {
	if v != nil {
		panic(v)
	}
	return
}
