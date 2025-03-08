package internal

import "os/exec"

func GitCommit(message string) error {
	commitCommand := exec.Command("git", "commit", "-m", message)
	return commitCommand.Run()
}

func UndoLastCommit() {
	undoCmd := exec.Command("git", "reset", "--soft", "HEAD~1")
	err := undoCmd.Run()
	AbortOnError(err, "failed to undo last commit, please try again")
}
