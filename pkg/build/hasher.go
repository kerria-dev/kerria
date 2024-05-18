// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package build

import (
	"github.com/kerria-dev/kerria/pkg/resources"
	"github.com/kerria-dev/kerria/pkg/util"
	"hash"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func DirectoryHash(algorithm resources.HashAlgorithm, baseDir string, relativeDir string) (digest []byte, err error) {
	baseDir, err = filepath.Abs(baseDir)
	if err != nil {
		return
	}
	var hasher hash.Hash
	hasher, err = resources.Hasher(algorithm)
	if err != nil {
		return
	}
	ignores, err := util.NewIgnores(baseDir)
	if err != nil {
		return
	}

	err = filepath.WalkDir(filepath.Join(baseDir, relativeDir), func(currentPath string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		var currentPathRel string
		currentPathRel, err = filepath.Rel(baseDir, currentPath)
		if err != nil {
			return err
		}
		matches := ignores.Absolute(currentPath, dirEntry.IsDir())
		ignored := matches.Ignored()
		if ignored && dirEntry.IsDir() {
			return filepath.SkipDir
		} else if !ignored && !dirEntry.IsDir() {
			// Only add hashes for filenames, as empty directories will not be tracked in VCS.
			// File path is normalized for differences between platforms
			currentPathRel = filepath.ToSlash(currentPathRel)
			hasher.Write([]byte(currentPathRel))
			var data []byte
			data, err = os.ReadFile(currentPath)
			if err != nil {
				return err
			}
			if utf8.Valid(data) {
				// Git will checkout with different line endings depending on the system.
				// These need to be normalized for consistent hashing.
				text := NormalizeLineEndings(string(data))
				hasher.Write([]byte(text))
			} else {
				hasher.Write(data)
			}
		}
		return nil
	})
	if err != nil {
		return
	}

	digest = hasher.Sum(nil)
	return
}

func NormalizeLineEndings(content string) string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	return content
}
