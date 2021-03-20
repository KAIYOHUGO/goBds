package db

type User struct {
	Name          string
	Password      []byte
	OwnServerList []string
	Permission    int8
}
