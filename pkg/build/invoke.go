// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package build

import (
	"bufio"
	"bytes"
	"errors"
	"k8s.io/klog/v2"
	"os/exec"
	"slices"
	"unicode/utf8"
)

var (
	buildCommand         = "kustomize"
	buildCommandArgsBase = []string{"build"}
)

func KustomizeBuildCommand(path string, flags []string) (string, error) {
	cmd := exec.Command(buildCommand, slices.Concat(buildCommandArgsBase, []string{path}, flags)...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	var stdoutBuffer bytes.Buffer
	stdoutScanner := bufio.NewScanner(stdoutPipe)
	go func() {
		for stdoutScanner.Scan() {
			line := stdoutScanner.Text()
			stdoutBuffer.WriteString(line + "\n")
		}
	}()

	stderrScanner := bufio.NewScanner(stderrPipe)
	for stderrScanner.Scan() {
		line := stderrScanner.Text()
		klog.Infof("[%s] %s\n", buildCommand, line)
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	stdoutBytes := stdoutBuffer.Bytes()
	if utf8.Valid(stdoutBytes) {
		return string(stdoutBytes), nil
	} else {
		return "", errors.New("stdout was not valid UTF-8")
	}
}
