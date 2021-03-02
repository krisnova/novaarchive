package filesystem

import (
	"os"
	"os/user"
	"strings"

	"github.com/kris-nova/logger"
)

func Home() string {
	home := os.Getenv("HOME")
	if strings.Contains(home, "root") {
		return "/root"
	}
	usr, err := user.Current()
	if err != nil {
		logger.Warning("unable to find user: %v", err)
		return ""
	}
	return usr.HomeDir
}
