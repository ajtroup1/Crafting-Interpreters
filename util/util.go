package util

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
)

// Color functions to wrap text in specific color codes
func RedText(text string) string {
	return Red + text + Reset
}

func GreenText(text string) string {
	return Green + text + Reset
}

func YellowText(text string) string {
	return Yellow + text + Reset
}