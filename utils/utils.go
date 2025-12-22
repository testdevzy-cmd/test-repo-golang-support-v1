package utils

import (
	"fmt"
	"strings"

	"github.com/test-repo/test-repo-golang-support-v1/models"
)

func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	if !strings.Contains(email, "@") {
		return false
	}
	return true
}

func ExtractUserInitials(u *models.User) string {
	if u == nil || u.FullName == "" {
		return "NA"
	}
	parts := strings.Split(u.FullName, " ")
	initials := ""
	for _, part := range parts {
		if len(part) > 0 {
			initials += strings.ToUpper(part[:1])
		}
	}
	return initials
}

func FormatGreeting(name, title string) string {
	if name == "" {
		return "Hello"
	}
	if title != "" {
		return fmt.Sprintf("Hello, %s %s!", title, name)
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func GenerateReport(names []string) string {
	if len(names) == 0 {
		return "No data to report"
	}
	var sb strings.Builder
	sb.WriteString("Report Summary:\n")
	for i, name := range names {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, name))
	}
	return sb.String()
}

func CalculateSum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func IsEmpty(str string) bool {
	return str == ""
}
