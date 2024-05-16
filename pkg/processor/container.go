// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kerria-dev/kerria/pkg/resources"
	"k8s.io/klog/v2"
	"os/exec"
	"os/user"
	"path/filepath"
)

var (
	dockerCommand = "docker"
)

func DockerCommand(processor *resources.Processor, message *RepositoryMessage) error {
	args, err := DockerArgs(processor)
	if err != nil {
		return err
	}
	cmd := exec.Command(dockerCommand, args...)

	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	defer stdinPipe.Close()
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdoutPipe.Close()
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderrPipe.Close()

	stdoutScanner := bufio.NewScanner(stdoutPipe)
	go func() {
		for stdoutScanner.Scan() {
			_ = stdoutScanner.Text()
		}
	}()

	stderrScanner := bufio.NewScanner(stderrPipe)
	go func() {
		for stderrScanner.Scan() {
			line := stderrScanner.Text()
			klog.Infof("[%s] %s\n", processor.Name, line)
		}
	}()

	encoder := json.NewEncoder(stdinPipe)
	err = encoder.Encode(message)
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

const (
	dockerNetworkNone = "none"
	dockerNetworkHost = "host"
)

func DockerArgs(processor *resources.Processor) ([]string, error) {
	network := dockerNetworkNone
	if processor.Network {
		network = dockerNetworkHost
	}
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}
	UIDGID := fmt.Sprintf("%s:%s", currentUser.Uid, currentUser.Gid)
	args := []string{"run",
		"--rm",
		"--interactive",
		"--attach", "STDIN",
		"--attach", "STDOUT",
		"--attach", "STDERR",
		"--network", network,
		"--user", UIDGID,
		"--security-opt=no-new-privileges",
	}

	for _, storageMount := range processor.StorageMounts {
		absPath, err := filepath.Abs(storageMount.Source)
		if err != nil {
			return nil, err
		}
		mode := ""
		if storageMount.ReadWriteMode {
			mode = ",readonly"
		}
		args = append(args, "--mount", fmt.Sprintf(
			"type=%s,source=%s,target=%s%s",
			storageMount.MountType, absPath, storageMount.Destination, mode))
	}

	for _, env := range processor.Env {
		args = append(args, "--env", env)
	}

	args = append(args, processor.Image)

	return args, nil
}
