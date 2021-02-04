/*
Copyright 2020-2021 The Tekton Authors
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

package triggers

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
	"github.com/tektoncd/dashboard/pkg/controllers/utils"
	logging "github.com/tektoncd/dashboard/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
)

func NewClusterTriggerBindingController(sharedInformerFactory dynamicinformer.DynamicSharedInformerFactory) {
	logging.Log.Debug("In NewClusterTriggerBindingController")

	gvr := schema.GroupVersionResource{
		Group:    "triggers.tekton.dev",
		Version:  "v1alpha1",
		Resource: "clustertriggerbindings",
	}

	utils.NewController(
		"ClusterTriggerBinding",
		sharedInformerFactory.ForResource(gvr).Informer(),
		broadcaster.ClusterTriggerBindingCreated,
		broadcaster.ClusterTriggerBindingUpdated,
		broadcaster.ClusterTriggerBindingDeleted,
		nil,
	)
}