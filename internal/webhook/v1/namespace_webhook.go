/*
Copyright 2024.

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

package v1

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

//+kubebuilder:webhook:path=/validate-core-v1-namespace,mutating=false,failurePolicy=fail,sideEffects=None,groups=core,resources=namespaces,verbs=create;update;delete,versions=v1,name=vnamespace.kb.io,admissionReviewVersions=v1

type NamespaceValidator struct {
	Client client.Client
}

func (v *NamespaceValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	log := logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{})

	log.WithField("namespace", req.Namespace).WithField("user", req.UserInfo.Username).WithField("action", req.Operation).Info("namespace action")
	return admission.Allowed("logged action")
}
