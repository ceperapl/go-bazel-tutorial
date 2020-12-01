package greeting

import "fmt"

// GetGreeting does the obvious
func GetGreeting(who string) string {
	return fmt.Sprintf("Hello %s\n", who)
}
