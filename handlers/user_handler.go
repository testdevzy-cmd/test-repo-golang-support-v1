package handlers

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/services"
	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) HandleCreateUser(id int, name, email string, age int) {
	user, err := h.userService.CreateUser(id, name, email, age)
	if err != nil {
		fmt.Printf("Handler Error: %v\n", err)
		return
	}

	if utils.ValidateEmail(user.EmailAddress) {
		greeting := utils.FormatGreeting(user.FullName, "Welcome")
		fmt.Println(greeting)
		fmt.Printf("User %s created successfully.\n", user.FullName)
	} else {
		fmt.Println("Invalid email provided for user.")
	}
}

func (h *UserHandler) HandleGetReport(id int) {
	report := h.userService.PrintUserReport(id)
	fmt.Println("Generated Report:")
	fmt.Println(report)
}
