package handlers

import (
	"fmt"
	"os/exec"
)

func connection2(s Server) {
	cmd := exec.Command(
		"ssh",
		"-t",
		fmt.Sprintf("%s@%s", s.User, s.Address),
		"uname -a",
	)
	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}

func deploy(t string, b string) {

}
