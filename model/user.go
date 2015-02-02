package model

type User struct {
	Id    int
	Name  string
	Token string
}

func GetUserIdByName(uname string) (int, error) {
	return 0, nil
}
