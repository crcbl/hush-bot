package utils

import (
	"fmt"
	"os"
)

func GetDockerSecret(name string) string {
	b, e := os.ReadFile(fmt.Sprintf("/run/secrets/%s", name))
	if e != nil {
		return ""
	}
	return string(b)
}
