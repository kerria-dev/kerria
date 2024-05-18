// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package util

import (
	"fmt"
	"os"
	"path/filepath"
)

var gitRepositoryRootCache string

func GitRepositoryRoot(path string) (string, error) {
	if gitRepositoryRootCache != "" {
		return gitRepositoryRootCache, nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	for {
		// Check if current path contains a .git/
		fileInfo, err := os.Stat(filepath.Join(absPath, ".git"))
		if err == nil && fileInfo.IsDir() {
			gitRepositoryRootCache = absPath
			return absPath, nil
		}
		// Otherwise traverse to parent
		parent := filepath.Dir(absPath)
		if parent == absPath {
			return "", fmt.Errorf("no Git repository found for path %s", path)
		}
		absPath = parent
	}
}
