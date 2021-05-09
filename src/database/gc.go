package database

func GC() {
	for _, v := range DB {
		v.Close()
	}
}
