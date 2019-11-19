package utils

import "fmt"

func MessageWithError(errType string, err error) string {
	return fmt.Sprintf(`{"error_type": "%s", "message": "%s"}`, errType, err)
}
