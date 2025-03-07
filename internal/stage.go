package internal

import "os/exec"

func HasStagedChanges() bool {
	stageCmd := exec.Command("git", "diff", "--cached", "--name-only", "--exit-code")
	err := stageCmd.Run()
	return err != nil
}
