/*
Copyright 2018 The Knative Authors

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

package build

import (
	"github.com/knative/build/pkg/apis/build/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type BuildSpec struct{}

type BuildSpecOpts struct {
	GitURL      string
	GitRevision string

	ServiceAccountName string

	Image string
}

func (s BuildSpec) Build(opts BuildSpecOpts) v1alpha1.BuildSpec {
	spec := v1alpha1.BuildSpec{
		ServiceAccountName: opts.ServiceAccountName,
		Source: &v1alpha1.SourceSpec{
			Git: &v1alpha1.GitSourceSpec{
				Url:      opts.GitURL,
				Revision: opts.GitRevision,
			},
		},
		Steps: []corev1.Container{
			{
				Name:  "build-and-push",
				Image: "gcr.io/kaniko-project/executor",
				Args: []string{
					"--dockerfile=/workspace/Dockerfile",
					"--destination=" + opts.Image,
				},
			},
		},
	}

	return spec
}
