package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func OsFileSeparator() string {
	return string(filepath.Separator)
}

func ExecCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	fmt.Println(strings.Join(cmd.Args, " "))
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%v: %v", err, string(stderr.Bytes()))
	}
	return stdout.Bytes(), nil
}
