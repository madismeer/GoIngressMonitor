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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	networkv1 "github.com/madismeer/GoIngressMonitor/api/v1"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=network.example.com,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=network.example.com,resources=ingresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=network.example.com,resources=ingresses/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Ingress object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var customIngress networkv1.Ingress

	if err := r.Get(ctx, req.NamespacedName, &customIngress); err != nil {
		log.Error(err, "unable to fetch CustomIngress", "NamespacedName", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Corrected log statement
	log.Info("Host value", "host", customIngress.Spec.Host)
	fmt.Println("Host value for assignment is:", customIngress.Spec.Host)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkv1.Ingress{}).
		Complete(r)
}
