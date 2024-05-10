// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package v1alpha1

import (
	"reflect"
)

const Version = "v1alpha1"
const KindRepository = "Repository"
const KindLockfile = "Lockfile"

var (
	Kinds = []string{
		KindRepository, KindLockfile,
	}

	Types = map[string]reflect.Type{
		KindRepository: reflect.TypeOf((*Repository)(nil)).Elem(),
		KindLockfile:   reflect.TypeOf((*Lockfile)(nil)).Elem(),
	}

	Defaulters = map[string]reflect.Value{
		KindRepository: reflect.ValueOf(SetObjectDefaults_Repository),
		KindLockfile:   reflect.ValueOf(SetObjectDefaults_Lockfile),
	}
)
