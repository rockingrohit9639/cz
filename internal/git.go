package internal

import "os/exec"

func HasStagedChanges() bool {
	stageCmd := exec.Command("git", "diff", "--cached", "--name-only", "--exit-code")
	err := stageCmd.Run()
	return err != nil
}

func IsGitInitialized() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

func GitCommit(message string) error {
	commitCommand := exec.Command("git", "commit", "-m", message)
	return commitCommand.Run()
}

func UndoLastCommit() {
	undoCmd := exec.Command("git", "reset", "--soft", "HEAD~1")
	err := undoCmd.Run()
	AbortOnError(err, "failed to undo last commit, please try again")
}

func StageAll() {
	stageCmd := exec.Command("git", "add", ".")
	err := stageCmd.Run()
	AbortOnError(err, "failed to stage all changes")
}
