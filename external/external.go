package external

import (
	"fmt"
	"os/exec"
	"os/user"
	"strings"
)

func IsInstalled(programName string) bool {
	_, err := exec.LookPath(programName)
	return err == nil
}

func Execute(programName string, args ...string) (string, error) {
	cmd := exec.Command(programName, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func GetUsername() string {
	name, err := Execute("git", "config", "user.name")
	if err == nil {
		return strings.TrimSpace(name)
	}

	user, err := user.Current()
	if err == nil {
		return strings.TrimSpace(user.Username)
	}

	fmt.Println("Make sure to replace <your-name> with your actual name.")
	return "<your-name>"
}
