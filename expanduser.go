package expanduser

import (
	"os/user"
	"strings"
)

func ExpandUser(data string) string {
	current, _ := user.Current()
	return strings.Replace(data, "~", current.HomeDir, 1)
}
