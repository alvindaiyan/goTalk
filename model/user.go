package model

type User struct {
	Id    int
	Name  string
	Token string
}

const (
	TABLE_NAME = "userinfo"
)

func GetUserIdByName(uname string) (int, error) {
	return 0, nil
}
