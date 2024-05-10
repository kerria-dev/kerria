// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package kerria_dev

import (
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1"
	"reflect"
)

// GroupName is the group name used in this package
const (
	GroupName = "kerria.dev"

	APIVersionV1Alpha1 = GroupName + "/" + v1alpha1.Version
)

var (
	Types = map[string]map[string]reflect.Type{
		APIVersionV1Alpha1: v1alpha1.Types,
	}
	Defaulters = map[string]map[string]reflect.Value{
		APIVersionV1Alpha1: v1alpha1.Defaulters,
	}
)
