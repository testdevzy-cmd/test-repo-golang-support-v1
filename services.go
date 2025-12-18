package main

type User struct {
	ID   int
	Name string
	Age  int
}

func GetUserAge(user *User) string {
	return user.Name
}
