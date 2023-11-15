package utils

import "fmt"

const (
	jq = "jq"
)

func JQ(s string) (string, error) {
	return ExecCmdRetOut("sh", "-c", fmt.Sprintf("echo '%s' | %s", s, jq))
}
