package internal

import "os/exec"

func GitCommit(message string) error {
	commitCommand := exec.Command("git", "commit", "-m", message)
	return commitCommand.Run()
}
