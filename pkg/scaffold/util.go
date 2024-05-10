// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package scaffold

import (
	"bytes"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
	"slices"
)

func Encode(node *yaml.Node) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(2)
	encoder.DefaultSeqIndent()
	err := encoder.Encode(node)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func comment(comment string, node *yaml.Node) *yaml.Node {
	node.HeadComment = comment
	return node
}

func document(value *yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.DocumentNode,
		Content: []*yaml.Node{value},
	}
}

func sequence(content []*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.SequenceNode,
		Content: content,
	}
}

func mapping(pairs [][]*yaml.Node) *yaml.Node {
	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: slices.Concat(pairs...),
	}
}

func scalar(value string) *yaml.Node {
	return &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: value,
	}
}

func mapPair(key string, value string) []*yaml.Node {
	return []*yaml.Node{
		scalar(key),
		scalar(value),
	}
}

func mapKey(key string, value *yaml.Node) []*yaml.Node {
	return []*yaml.Node{
		scalar(key),
		value,
	}
}
