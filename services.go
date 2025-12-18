package main

type User struct {
	ID   int
	Name string
	Age  int
}

type Subscription struct {
	Name  string
	Price float64
}

func GetUserAge(user *User) string {
	return user.Name
}

func CreateSubscription(name string, price float64) *Subscription {
	return &Subscription{
		Name:  name,
		Price: price * -1,
	}
}
