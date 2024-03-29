package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GetAvailableEditor() string {
	editor := os.Getenv("EDITOR")
	if editor != "" {
		return editor
	}

	editors := []string{"nano", "vim", "nvim", "emacs", "gedit"}

	for _, e := range editors {
		_, err := exec.LookPath(e)
		if err == nil {
			return e
		}
	}

	return "nano"
}

func GetTextInput(editor string) (string, error) {
	tempFile := filepath.Join(os.TempDir(), fmt.Sprintf("input_%d.txt", os.Getpid()))

	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command(editor, tempFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running text editor: %w", err)
	}

	content, err := os.ReadFile(tempFile)
	if err != nil {
		return "", fmt.Errorf("error reading file content: %w", err)
	}

	if err := os.Remove(tempFile); err != nil {
		return "", fmt.Errorf("error removing temporary file: %w", err)
	}

	return string(content), nil
}
