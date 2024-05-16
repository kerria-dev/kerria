// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Kerria

package resources

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	krapi "github.com/kerria-dev/kerria/pkg/apis/kerria.dev"
	"github.com/kerria-dev/kerria/pkg/apis/kerria.dev/v1alpha1"
	"github.com/kerria-dev/kerria/pkg/openapi"
	"hash"
	"os"
	"reflect"
	kyyaml "sigs.k8s.io/kustomize/kyaml/yaml"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
	"time"
)

type HashAlgorithm int

const (
	MD5 HashAlgorithm = iota
	SHA1
	SHA256
	SHA512

	LockFile = "kerria.lock.yaml"
	LockKind = "Lockfile"
)

var (
	HashAlgorithms = map[string]HashAlgorithm{
		"md5":    MD5,
		"sha1":   SHA1,
		"sha256": SHA256,
		"sha512": SHA512,
	}
	HashAlgorithmsReverse map[HashAlgorithm]string
)

type Lockfile struct {
	Name        string
	DefaultHash HashAlgorithm
	Builds      []*BuildStatus
}

type BuildStatus struct {
	ID             int
	Timestamp      time.Time
	SourceHash     []byte
	SourceHashType HashAlgorithm
	SourcePath     string
	BuildHash      []byte
	BuildHashType  HashAlgorithm
	BuildPath      string
}

// LockfileFromAPI converts the latest API into the internal representation
func LockfileFromAPI(apiLock *v1alpha1.Lockfile) (lockfile *Lockfile, err error) {
	lockfile = &Lockfile{}
	lockfile.Name = apiLock.Name
	lockfile.DefaultHash = HashAlgorithms[apiLock.Spec.DefaultHash]
	for idx, apiBuildStatus := range apiLock.Status.Builds {

		buildStatus := BuildStatus{}
		buildStatus.ID = idx
		buildStatus.Timestamp, err = time.Parse(time.RFC3339, apiBuildStatus.Timestamp)
		if err != nil {
			return
		}
		buildStatus.SourceHash, err = hex.DecodeString(apiBuildStatus.SourceHash)
		if err != nil {
			return
		}
		buildStatus.SourceHashType = HashAlgorithms[string(apiBuildStatus.SourceHashType)]
		buildStatus.SourcePath = apiBuildStatus.SourcePath
		buildStatus.BuildHash, err = hex.DecodeString(apiBuildStatus.BuildHash)
		if err != nil {
			return
		}
		buildStatus.BuildHashType = HashAlgorithms[string(apiBuildStatus.BuildHashType)]
		buildStatus.BuildPath = apiBuildStatus.BuildPath

		lockfile.Builds = append(lockfile.Builds, &buildStatus)
	}
	return
}

func (lockfile *Lockfile) AsAPI() (apiLock *v1alpha1.Lockfile) {
	apiLock = &v1alpha1.Lockfile{}
	apiLock.APIVersion = krapi.APIVersionV1Alpha1
	apiLock.Kind = LockKind
	apiLock.Name = lockfile.Name
	apiLock.Spec.DefaultHash = HashAlgorithmsReverse[lockfile.DefaultHash]
	for _, buildStatus := range lockfile.Builds {
		apiBuildStatus := v1alpha1.BuildStatus{
			Timestamp:      buildStatus.Timestamp.Format(time.RFC3339),
			SourceHash:     hex.EncodeToString(buildStatus.SourceHash),
			SourceHashType: v1alpha1.HashAlgorithm(HashAlgorithmsReverse[buildStatus.SourceHashType]),
			SourcePath:     buildStatus.SourcePath,
			BuildHash:      hex.EncodeToString(buildStatus.BuildHash),
			BuildHashType:  v1alpha1.HashAlgorithm(HashAlgorithmsReverse[buildStatus.BuildHashType]),
			BuildPath:      buildStatus.BuildPath,
		}
		apiLock.Status.Builds = append(apiLock.Status.Builds, apiBuildStatus)
	}
	return
}

func (lockfile *Lockfile) Write() error {
	return lockfile.WriteWithPath(LockFile)
}

func (lockfile *Lockfile) WriteWithPath(path string) error {
	apiLock := lockfile.AsAPI()
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	encoder.DefaultSeqIndent()
	encoder.SetIndent(2)
	err = encoder.Encode(apiLock)
	if err != nil {
		return err
	}
	return nil
}

func Hasher(algorithm HashAlgorithm) (hash.Hash, error) {
	switch algorithm {
	case MD5:
		return md5.New(), nil
	case SHA1:
		return sha1.New(), nil
	case SHA256:
		return sha256.New(), nil
	case SHA512:
		return sha512.New(), nil
	default:
		return nil, errors.New("not a supported hashing algorithm")
	}
}

func ReadLockfile() (*Lockfile, error) {
	return ReadLockfileWithPath(LockFile)
}

func ReadLockfileWithPath(path string) (lockfile *Lockfile, err error) {
	var rnode *kyyaml.RNode
	rnode, err = kyyaml.ReadFile(path)
	if err != nil {
		return
	}
	err = rnode.DeAnchor()
	if err != nil {
		return
	}
	apiVersion := rnode.GetApiVersion()
	kind := rnode.GetKind()
	if apiVersion != krapi.APIVersionV1Alpha1 {
		err = fmt.Errorf("unsupported apiVersion for Lockfile %s", apiVersion)
		return
	}
	if kind != LockKind {
		err = fmt.Errorf("incorrect kind %s is not Lockfile", kind)
		return
	}
	typeMeta := kyyaml.TypeMeta{
		APIVersion: apiVersion,
		Kind:       kind,
	}
	var value reflect.Value
	value, err = openapi.DecodeValidatedDefaulted(rnode, typeMeta)
	if err != nil {
		panic(err)
	}
	v1alpha1Lock := value.Interface().(*v1alpha1.Lockfile)
	lockfile, err = LockfileFromAPI(v1alpha1Lock)
	return
}

func init() {
	HashAlgorithmsReverse = make(map[HashAlgorithm]string, len(HashAlgorithms))
	for key, value := range HashAlgorithms {
		HashAlgorithmsReverse[value] = key
	}
}
