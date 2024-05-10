// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package build

import (
	"github.com/kerria-dev/kerria/pkg/resources"
	"math"
	"path/filepath"
)

type LockfileDiff struct {
	Delete []*resources.BuildStatus
	Check  []*resources.BuildStatus
	Create []*resources.BuildStatus
}

// CheckDifference finds the intersection and difference of the resources.Repository and resources.Lockfile based on the
// source path.
func CheckDifference(repository *resources.Repository, lockfile *resources.Lockfile) (lockfileDiff *LockfileDiff) {
	lockfileDiff = &LockfileDiff{}

	// set of all source paths
	sourcePaths := make(map[string]struct{})
	for _, source := range repository.Sources {
		for _, sourcePath := range source.Paths {
			sourcePaths[sourcePath] = struct{}{}
		}
	}

	// set of source paths found in the lockfile
	sourcePathsUsed := make(map[string]struct{})

	maxID := math.MinInt
	// find intersection and lockfile difference
	for _, lockBuildStatus := range lockfile.Builds {
		_, exists := sourcePaths[lockBuildStatus.SourcePath]
		if exists {
			lockfileDiff.Check = append(lockfileDiff.Check, lockBuildStatus)
			sourcePathsUsed[lockBuildStatus.SourcePath] = struct{}{}
		} else {
			lockfileDiff.Delete = append(lockfileDiff.Delete, lockBuildStatus)
		}
		// additionally, we keep track of the highest ID, so we can continue counting when creating new builds
		maxID = max(maxID, lockBuildStatus.ID)
	}

	// find repository difference
	for _, source := range repository.Sources {
		for _, sourcePath := range source.Paths {
			_, exists := sourcePathsUsed[sourcePath]
			if !exists {
				maxID++
				lockfileDiff.Create = append(lockfileDiff.Create,
					&resources.BuildStatus{
						ID:             maxID,
						SourceHashType: lockfile.DefaultHash,
						SourcePath:     sourcePath,
						BuildHashType:  lockfile.DefaultHash,
						BuildPath:      filepath.ToSlash(filepath.Join(repository.BuildPath, sourcePath)),
					})
			}
		}
	}

	return
}
