package color

import "fmt"

func RedString(format string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", format)
}

func BoldRedString(format string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", format)
}

func GreenString(format string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", format)
}

func BoldGreenString(format string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", format)
}

func YellowString(format string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", format)
}
