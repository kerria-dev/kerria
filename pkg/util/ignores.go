// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package util

import "github.com/ianlewis/go-gitignore"

const (
	ignoreFileKerria = ".kerriaignore"
	ignoreFileGit    = ".gitignore"
)

var (
	// In order of priority, first is highest
	defaultIgnoreFiles = []string{
		ignoreFileKerria, ignoreFileGit,
	}
)

type Ignores []gitignore.GitIgnore
type Matches []gitignore.Match

func NewIgnores(baseDir string) (ignores Ignores, err error) {
	for _, ignoreFile := range defaultIgnoreFiles {
		var ignore gitignore.GitIgnore
		ignore, err = gitignore.NewRepositoryWithFile(baseDir, ignoreFile)
		if err != nil {
			return
		}
		ignores = append(ignores, ignore)
	}
	return
}

func (ignores Ignores) Absolute(path string, dir bool) (matches Matches) {
	for _, ignore := range ignores {
		matches = append(matches, ignore.Absolute(path, dir))
	}
	return
}

func (matches Matches) Ignored() bool {
	for _, match := range matches {
		if match != nil {
			if match.Ignore() {
				return true
			} else if match.Include() {
				return false
			}
		}
	}
	return false
}
