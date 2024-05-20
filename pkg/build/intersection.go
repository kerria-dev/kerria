// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package build

import (
	"github.com/kerria-dev/kerria/pkg/resources"
	"math"
	"path/filepath"
)

type DiscoveryIntersection struct {
	Delete []PairedDiscovery
	Check  []PairedDiscovery
	Create []PairedDiscovery
}

// PairedDiscovery matches a source resources.Discovery together with its resources.BuildStatus.
// This information is used when sending metadata about the state of the repository to a processor
type PairedDiscovery struct {
	Source      *resources.Source
	Discovery   *resources.Discovery
	BuildStatus *resources.BuildStatus
}

// Intersect finds the intersection and difference of the resources.Repository and resources.Lockfile based on the
// source path.
func Intersect(repository *resources.Repository, lockfile *resources.Lockfile) *DiscoveryIntersection {
	intersection := &DiscoveryIntersection{}

	// map of all discovery paths
	discoPaths := make(map[string]PairedDiscovery)
	for _, source := range repository.Sources {
		for _, discovery := range source.Discoveries {
			discoPaths[discovery.Path] = PairedDiscovery{
				Source:    source,
				Discovery: discovery,
			}
		}
	}

	// set of discovery paths found in the lockfile
	discoPathsUsed := make(map[string]struct{})

	maxID := math.MinInt
	// find intersection and lockfile difference
	for _, lockBuildStatus := range lockfile.Builds {
		pair, exists := discoPaths[lockBuildStatus.DiscoPath]
		if exists {
			pair.BuildStatus = lockBuildStatus
			intersection.Check = append(intersection.Check, pair)
			discoPathsUsed[lockBuildStatus.DiscoPath] = struct{}{}
		} else {
			intersection.Delete = append(intersection.Delete, PairedDiscovery{
				BuildStatus: lockBuildStatus,
			})
		}
		// additionally, we keep track of the highest ID, so we can continue counting when creating new builds
		maxID = max(maxID, lockBuildStatus.ID)
	}

	// find repository difference
	for _, source := range repository.Sources {
		for _, discovery := range source.Discoveries {
			_, exists := discoPathsUsed[discovery.Path]
			if !exists {
				maxID++
				intersection.Create = append(intersection.Create, PairedDiscovery{
					Source:    source,
					Discovery: discovery,
					BuildStatus: &resources.BuildStatus{
						ID:            maxID,
						DiscoHashType: lockfile.DefaultHash,
						DiscoPath:     discovery.Path,
						BuildHashType: lockfile.DefaultHash,
						BuildPath:     filepath.ToSlash(filepath.Join(repository.BuildPath, discovery.Path)),
					}})
			}
		}
	}

	return intersection
}
