// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package processor

import (
	"encoding/hex"
	"github.com/kerria-dev/kerria/pkg/build"
	"github.com/kerria-dev/kerria/pkg/resources"
	"time"
)

// Protocol is a semantic version string
type Protocol string

const Protocol_0_1_0 Protocol = "0.1.0"

type RepositoryMessage struct {
	Protocol    Protocol        `json:"protocol"`
	Name        string          `json:"name"`
	BuildPath   string          `json:"buildPath"`
	Processor   Processor       `json:"processor"`
	Sources     []Source        `json:"sources"`
	Discoveries DiscoveryStatus `json:"discoveries"`
}

type Processor struct {
	Name       string      `json:"name"`
	Stage      string      `json:"stage"`
	Properties interface{} `json:"properties"`
}

type Source struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DiscoveryStatus struct {
	Delete  []Discovery `json:"delete"`
	Check   []Discovery `json:"check"`
	Rebuilt []Discovery `json:"rebuilt"`
	Create  []Discovery `json:"create"`
}

type Discovery struct {
	Source         int               `json:"source"`
	SourcePath     string            `json:"sourcePath"`
	SourceHash     string            `json:"sourceHash"`
	SourceHashType string            `json:"sourceHashType"`
	BuildPath      string            `json:"buildPath"`
	BuildTimestamp string            `json:"buildTimestamp"`
	BuildHash      string            `json:"buildHash"`
	BuildHashType  string            `json:"buildHashType"`
	Name           string            `json:"name"`
	Namespace      string            `json:"namespace"`
	Labels         map[string]string `json:"labels"`
	Annotations    map[string]string `json:"annotations"`
}

func NewRepositoryMessage(
	repository *resources.Repository,
	intersection *build.DiscoveryIntersection,
) *RepositoryMessage {
	message := &RepositoryMessage{}
	message.Protocol = Protocol_0_1_0
	message.Name = repository.Name
	message.BuildPath = repository.BuildPath
	for _, repoSource := range repository.Sources {
		message.Sources = append(message.Sources, Source{
			ID:   repoSource.ID,
			Name: repoSource.Name})
	}
	for _, pair := range intersection.Delete {
		message.Discoveries.Delete = append(message.Discoveries.Delete, *newDiscovery(&pair))
	}
	for _, pair := range intersection.Check {
		message.Discoveries.Check = append(message.Discoveries.Check, *newDiscovery(&pair))
	}
	for _, pair := range intersection.Create {
		message.Discoveries.Create = append(message.Discoveries.Create, *newDiscovery(&pair))
	}
	return message
}

func (message *RepositoryMessage) WithProcessor(processor *resources.Processor) {
	message.Processor.Name = processor.Name
	message.Processor.Stage = resources.ProcessorStagesReverse[processor.Stage]
	message.Processor.Properties = processor.Properties
}

func (message *RepositoryMessage) WithRebuilt(rebuilt []build.PairedDiscovery) {
	for _, pair := range rebuilt {
		message.Discoveries.Rebuilt = append(message.Discoveries.Rebuilt, *newDiscovery(&pair))
	}
}

func newDiscovery(pair *build.PairedDiscovery) *Discovery {
	return &Discovery{
		Source:         pair.Source.ID,
		SourcePath:     pair.BuildStatus.SourcePath,
		SourceHash:     hex.EncodeToString(pair.BuildStatus.SourceHash),
		SourceHashType: resources.HashAlgorithmsReverse[pair.BuildStatus.SourceHashType],
		BuildPath:      pair.BuildStatus.BuildPath,
		BuildTimestamp: pair.BuildStatus.Timestamp.Format(time.RFC3339),
		BuildHash:      hex.EncodeToString(pair.BuildStatus.BuildHash),
		BuildHashType:  resources.HashAlgorithmsReverse[pair.BuildStatus.BuildHashType],
		Name:           pair.Discovery.Name,
		Namespace:      pair.Discovery.Namespace,
		Labels:         pair.Discovery.Labels,
		Annotations:    pair.Discovery.Annotations,
	}
}
