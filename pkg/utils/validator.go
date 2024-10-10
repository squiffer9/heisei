package utils

import (
	"regexp"
	"unicode/utf8"
)

// ValidateUsername checks if the username is valid
func ValidateUsername(username string) bool {
	// Permit only 3-20 alphanumeric characters and underscores for the username
	pattern := `^[a-zA-Z0-9_]{3,20}$`
	matched, _ := regexp.MatchString(pattern, username)
	return matched
}

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) bool {
	// Validate a simple email address
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// ValidateThreadTitle checks if the thread title is valid
func ValidateThreadTitle(title string) bool {
	// Validate the thread title length
	return utf8.RuneCountInString(title) >= 3 && utf8.RuneCountInString(title) <= 100
}

// ValidatePostContent checks if the post content is valid
func ValidatePostContent(content string) bool {
	// Validate the post content length
	return utf8.RuneCountInString(content) >= 1 && utf8.RuneCountInString(content) <= 10000
}

// SanitizeInput removes or escapes potentially harmful characters
func SanitizeInput(input string) string {
	// Sanitize the input by replacing < and > with HTML entities
	input = regexp.MustCompile(`<`).ReplaceAllString(input, "&lt;")
	input = regexp.MustCompile(`>`).ReplaceAllString(input, "&gt;")
	return input
}
