/*
Copyright 2023.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logger "sigs.k8s.io/controller-runtime/pkg/log"

	pachydermv1alpha1 "github.com/alexanderjophus/pachyderm-operator/api/v1alpha1"
	pachydermclient "github.com/pachyderm/pachyderm/v2/src/client"
	"github.com/pachyderm/pachyderm/v2/src/pps"
)

// PipelineReconciler reconciles a Pipeline object
type PipelineReconciler struct {
	client.Client
	Scheme          *runtime.Scheme
	PachydermClient *pachydermclient.APIClient
}

//+kubebuilder:rbac:groups=pachyderm.dejophus.dev,resources=pipelines,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pachyderm.dejophus.dev,resources=pipelines/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pachyderm.dejophus.dev,resources=pipelines/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=replicationcontrollers,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *PipelineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	pachyderm := &pachydermv1alpha1.Pipeline{}
	if err := r.Get(ctx, req.NamespacedName, pachyderm); err != nil {
		log.Error(err, "unable to fetch Pachyderm")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	spec := pachyderm.Spec

	if err := r.PachydermClient.CreateProjectPipeline(
		spec.Project,
		spec.Name,
		spec.Transform.Image,
		spec.Transform.Cmd,
		spec.Transform.StdIn,
		&pps.ParallelismSpec{Constant: 1},
		&pps.Input{
			Pfs: &pps.PFSInput{
				Glob: spec.Input.Pfs.Glob,
				Repo: spec.Input.Pfs.Repo,
			},
		},
		"master",
		false,
	); err != nil {
		log.Error(err, "unable to create pipeline")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil

}

// SetupWithManager sets up the controller with the Manager.
func (r *PipelineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pachydermv1alpha1.Pipeline{}).
		// Owns(&corev1.ReplicationController{}). // need to figure out permissions
		Complete(r)
}
