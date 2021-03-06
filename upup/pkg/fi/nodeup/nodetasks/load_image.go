/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nodetasks

import (
	"fmt"
	"github.com/golang/glog"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/nodeup/cloudinit"
	"k8s.io/kops/upup/pkg/fi/nodeup/local"
	"k8s.io/kops/upup/pkg/fi/utils"
	"k8s.io/kops/util/pkg/hashing"
	"os/exec"
	"path"
	"strings"
)

// LoadImageTask is responsible for downloading a docker image
type LoadImageTask struct {
	Source string
	Hash   string
}

var _ fi.Task = &LoadImageTask{}

func (t *LoadImageTask) String() string {
	return fmt.Sprintf("LoadImageTask: %s", t.Source)
}

func (e *LoadImageTask) Find(c *fi.Context) (*LoadImageTask, error) {
	glog.Warningf("LoadImageTask checking if image present not yet implemented")
	return nil, nil
}

func (e *LoadImageTask) Run(c *fi.Context) error {
	return fi.DefaultDeltaRunMethod(e, c)
}

func (_ *LoadImageTask) CheckChanges(a, e, changes *LoadImageTask) error {
	return nil
}

func (_ *LoadImageTask) RenderLocal(t *local.LocalTarget, a, e, changes *LoadImageTask) error {
	hash, err := hashing.FromString(e.Hash)
	if err != nil {
		return err
	}

	url := e.Source

	localFile := path.Join(t.CacheDir, hash.String()+"_"+utils.SanitizeString(url))
	_, err = fi.DownloadURL(url, localFile, hash)
	if err != nil {
		return err
	}

	// Load the image into docker
	args := []string{"docker", "load", "-i", localFile}
	human := strings.Join(args, " ")

	glog.Infof("running command %s", human)
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error loading docker image with '%s': %v: %s", human, err, string(output))
	}

	return nil
}

func (_ *LoadImageTask) RenderCloudInit(t *cloudinit.CloudInitTarget, a, e, changes *LoadImageTask) error {
	return fmt.Errorf("LoadImageTask::RenderCloudInit not implemented")
}
