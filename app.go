// Copyright © 2021 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//   ███╗   ██╗ █████╗ ███╗   ███╗██╗
//   ████╗  ██║██╔══██╗████╗ ████║██║
//   ██╔██╗ ██║███████║██╔████╔██║██║
//   ██║╚██╗██║██╔══██║██║╚██╔╝██║██║
//   ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗
//   ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
//

package tgik

import (
	"context"

	"github.com/hexops/valast"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

// Version is the current release of your application.
var Version string = "0.0.1"

// App is a very important grown up business application.
type App struct {
	metav1.ObjectMeta
	description string
	objects     []runtime.Object
	// ----------------------------------
	// Add your configuration fields here
	// ----------------------------------
}

// NewApp will create a new instance of App.
//
// See https://github.com/naml-examples for more examples.
//
// This is where you pass in fields to your application (similar to Values.yaml)
// Example: func NewApp(name string, example string, something int) *App
func NewApp(name, description string) *App {
	return &App{
		description: description,
		ObjectMeta: metav1.ObjectMeta{
			Name:            name,
			ResourceVersion: Version,
		},
		// ----------------------------------
		// Add your configuration fields here
		// ----------------------------------
	}
}

func (a *App) Install(client *kubernetes.Clientset) error {
	var err error

	kubernetes_dashboardNamespace := &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: "kubernetes-dashboard"},
	}
	a.objects = append(a.objects, kubernetes_dashboardNamespace)

	if client != nil {
		_, err = client.CoreV1().Namespaces().Create(context.TODO(), kubernetes_dashboardNamespace, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboardServiceAccount := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
	}
	a.objects = append(a.objects, kubernetes_dashboardServiceAccount)

	if client != nil {
		_, err = client.CoreV1().ServiceAccounts("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboardServiceAccount, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboardService := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{corev1.ServicePort{
				Port:       443,
				TargetPort: intstr.IntOrString{IntVal: 8443},
			}},
			Selector: map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
	}
	a.objects = append(a.objects, kubernetes_dashboardService)

	if client != nil {
		_, err = client.CoreV1().Services("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboardService, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboard_certsSecret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard-certs",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Type: corev1.SecretType("Opaque"),
	}
	a.objects = append(a.objects, kubernetes_dashboard_certsSecret)

	if client != nil {
		_, err = client.CoreV1().Secrets("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboard_certsSecret, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboard_csrfSecret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard-csrf",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Data: map[string][]uint8{"csrf": []uint8{}},
		Type: corev1.SecretType("Opaque"),
	}
	a.objects = append(a.objects, kubernetes_dashboard_csrfSecret)

	if client != nil {
		_, err = client.CoreV1().Secrets("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboard_csrfSecret, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboard_key_holderSecret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard-key-holder",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Type: corev1.SecretType("Opaque"),
	}
	a.objects = append(a.objects, kubernetes_dashboard_key_holderSecret)

	if client != nil {
		_, err = client.CoreV1().Secrets("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboard_key_holderSecret, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboard_settingsConfigMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard-settings",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
	}
	a.objects = append(a.objects, kubernetes_dashboard_settingsConfigMap)

	if client != nil {
		_, err = client.CoreV1().ConfigMaps("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboard_settingsConfigMap, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboardRole := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: "rbac.authorization.k8s.io/rbacv1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Rules: []rbacv1.PolicyRule{
			rbacv1.PolicyRule{
				Verbs: []string{
					"get",
					"update",
					"delete",
				},
				APIGroups: []string{""},
				Resources: []string{"secrets"},
				ResourceNames: []string{
					"kubernetes-dashboard-key-holder",
					"kubernetes-dashboard-certs",
					"kubernetes-dashboard-csrf",
				},
			},
			rbacv1.PolicyRule{
				Verbs: []string{
					"get",
					"update",
				},
				APIGroups:     []string{""},
				Resources:     []string{"configmaps"},
				ResourceNames: []string{"kubernetes-dashboard-settings"},
			},
			rbacv1.PolicyRule{
				Verbs:     []string{"proxy"},
				APIGroups: []string{""},
				Resources: []string{"services"},
				ResourceNames: []string{
					"heapster",
					"dashboard-metrics-scraper",
				},
			},
			rbacv1.PolicyRule{
				Verbs:     []string{"get"},
				APIGroups: []string{""},
				Resources: []string{"services/proxy"},
				ResourceNames: []string{
					"heapster",
					"http:heapster:",
					"https:heapster:",
					"dashboard-metrics-scraper",
					"http:dashboard-metrics-scraper",
				},
			},
		},
	}
	a.objects = append(a.objects, kubernetes_dashboardRole)

	if client != nil {
		_, err = client.RbacV1().Roles("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboardRole, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboardClusterRole := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/rbacv1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   "kubernetes-dashboard",
			Labels: map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Rules: []rbacv1.PolicyRule{rbacv1.PolicyRule{
			Verbs: []string{
				"get",
				"list",
				"watch",
			},
			APIGroups: []string{"metrics.k8s.io"},
			Resources: []string{
				"pods",
				"nodes",
			},
		}},
	}
	a.objects = append(a.objects, kubernetes_dashboardClusterRole)

	if client != nil {
		_, err = client.RbacV1().ClusterRoles().Create(context.TODO(), kubernetes_dashboardClusterRole, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboardRoleBinding := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/rbacv1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Subjects: []rbacv1.Subject{rbacv1.Subject{
			Kind:      "ServiceAccount",
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
		}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "kubernetes-dashboard",
		},
	}
	a.objects = append(a.objects, kubernetes_dashboardRoleBinding)

	if client != nil {
		_, err = client.RbacV1().RoleBindings("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboardRoleBinding, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	kubernetes_dashboardClusterRoleBinding := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/rbacv1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: "kubernetes-dashboard"},
		Subjects: []rbacv1.Subject{rbacv1.Subject{
			Kind:      "ServiceAccount",
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
		}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "kubernetes-dashboard",
		},
	}
	a.objects = append(a.objects, kubernetes_dashboardClusterRoleBinding)

	if client != nil {
		_, err = client.RbacV1().ClusterRoleBindings().Create(context.TODO(), kubernetes_dashboardClusterRoleBinding, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	// Adding a deployment: "kubernetes-dashboard"
	kubernetes_dashboardDeployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/appsv1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes-dashboard",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "kubernetes-dashboard"},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: valast.Addr(int32(1)).(*int32),
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{
				"k8s-app": "kubernetes-dashboard",
			}},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"k8s-app": "kubernetes-dashboard"},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{
						corev1.Volume{
							Name: "kubernetes-dashboard-certs",
							VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{
								SecretName: "kubernetes-dashboard-certs",
							}},
						},
						corev1.Volume{
							Name:         "tmp-volume",
							VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
						},
					},
					Containers: []corev1.Container{corev1.Container{
						Name:  "kubernetes-dashboard",
						Image: "kubernetesui/dashboard:v2.3.1",
						Args: []string{
							"--auto-generate-certificates",
							"--namespace=kubernetes-dashboard",
						},
						Ports: []corev1.ContainerPort{corev1.ContainerPort{
							ContainerPort: 8443,
							Protocol:      corev1.Protocol("TCP"),
						}},
						VolumeMounts: []corev1.VolumeMount{
							corev1.VolumeMount{
								Name:      "kubernetes-dashboard-certs",
								MountPath: "/certs",
							},
							corev1.VolumeMount{
								Name:      "tmp-volume",
								MountPath: "/tmp",
							},
						},
						LivenessProbe: &corev1.Probe{
							Handler: corev1.Handler{HTTPGet: &corev1.HTTPGetAction{
								Path: "/",
								Port: intstr.IntOrString{
									IntVal: 8443,
								},
								Scheme: corev1.URIScheme("HTTPS"),
							}},
							InitialDelaySeconds: 30,
							TimeoutSeconds:      30,
						},
						ImagePullPolicy: corev1.PullPolicy("Always"),
						SecurityContext: &corev1.SecurityContext{
							RunAsUser:                valast.Addr(int64(1001)).(*int64),
							RunAsGroup:               valast.Addr(int64(2001)).(*int64),
							ReadOnlyRootFilesystem:   valast.Addr(true).(*bool),
							AllowPrivilegeEscalation: valast.Addr(false).(*bool),
						},
					}},
					NodeSelector:       map[string]string{"kubernetes.io/os": "linux"},
					ServiceAccountName: "kubernetes-dashboard",
					Tolerations: []corev1.Toleration{corev1.Toleration{
						Key:    "node-role.kubernetes.io/master",
						Effect: corev1.TaintEffect("NoSchedule"),
					}},
				},
			},
			RevisionHistoryLimit: valast.Addr(int32(10)).(*int32),
		},
	}
	a.objects = append(a.objects, kubernetes_dashboardDeployment)

	if client != nil {
		_, err = client.AppsV1().Deployments("kubernetes-dashboard").Create(context.TODO(), kubernetes_dashboardDeployment, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	dashboard_metrics_scraperService := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "corev1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dashboard-metrics-scraper",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "dashboard-metrics-scraper"},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{corev1.ServicePort{
				Port:       8000,
				TargetPort: intstr.IntOrString{IntVal: 8000},
			}},
			Selector: map[string]string{"k8s-app": "dashboard-metrics-scraper"},
		},
	}
	a.objects = append(a.objects, dashboard_metrics_scraperService)

	if client != nil {
		_, err = client.CoreV1().Services("kubernetes-dashboard").Create(context.TODO(), dashboard_metrics_scraperService, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	// Adding a deployment: "dashboard-metrics-scraper"
	dashboard_metrics_scraperDeployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/appsv1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dashboard-metrics-scraper",
			Namespace: "kubernetes-dashboard",
			Labels:    map[string]string{"k8s-app": "dashboard-metrics-scraper"},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: valast.Addr(int32(1)).(*int32),
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{
				"k8s-app": "dashboard-metrics-scraper",
			}},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      map[string]string{"k8s-app": "dashboard-metrics-scraper"},
					Annotations: map[string]string{"seccomp.security.alpha.kubernetes.io/pod": "runtime/default"},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{corev1.Volume{
						Name:         "tmp-volume",
						VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
					}},
					Containers: []corev1.Container{corev1.Container{
						Name:  "dashboard-metrics-scraper",
						Image: "kubernetesui/metrics-scraper:appsv1.0.6",
						Ports: []corev1.ContainerPort{corev1.ContainerPort{
							ContainerPort: 8000,
							Protocol:      corev1.Protocol("TCP"),
						}},
						VolumeMounts: []corev1.VolumeMount{corev1.VolumeMount{
							Name:      "tmp-volume",
							MountPath: "/tmp",
						}},
						LivenessProbe: &corev1.Probe{
							Handler: corev1.Handler{HTTPGet: &corev1.HTTPGetAction{
								Path: "/",
								Port: intstr.IntOrString{
									IntVal: 8000,
								},
								Scheme: corev1.URIScheme("HTTP"),
							}},
							InitialDelaySeconds: 30,
							TimeoutSeconds:      30,
						},
						SecurityContext: &corev1.SecurityContext{
							RunAsUser:                valast.Addr(int64(1001)).(*int64),
							RunAsGroup:               valast.Addr(int64(2001)).(*int64),
							ReadOnlyRootFilesystem:   valast.Addr(true).(*bool),
							AllowPrivilegeEscalation: valast.Addr(false).(*bool),
						},
					}},
					NodeSelector:       map[string]string{"kubernetes.io/os": "linux"},
					ServiceAccountName: "kubernetes-dashboard",
					Tolerations: []corev1.Toleration{corev1.Toleration{
						Key:    "node-role.kubernetes.io/master",
						Effect: corev1.TaintEffect("NoSchedule"),
					}},
				},
			},
			RevisionHistoryLimit: valast.Addr(int32(10)).(*int32),
		},
	}
	a.objects = append(a.objects, dashboard_metrics_scraperDeployment)

	if client != nil {
		_, err = client.AppsV1().Deployments("kubernetes-dashboard").Create(context.TODO(), dashboard_metrics_scraperDeployment, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}

	return err
}

func (a *App) Uninstall(client *kubernetes.Clientset) error {
	var err error

	if client != nil {
		err = client.CoreV1().Namespaces().Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().ServiceAccounts("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().Services("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().Secrets("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard-certs", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().Secrets("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard-csrf", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().Secrets("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard-key-holder", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().ConfigMaps("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard-settings", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.RbacV1().Roles("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.RbacV1().ClusterRoles().Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.RbacV1().RoleBindings("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.RbacV1().ClusterRoleBindings().Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.AppsV1().Deployments("kubernetes-dashboard").Delete(context.TODO(), "kubernetes-dashboard", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.CoreV1().Services("kubernetes-dashboard").Delete(context.TODO(), "dashboard-metrics-scraper", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	if client != nil {
		err = client.AppsV1().Deployments("kubernetes-dashboard").Delete(context.TODO(), "dashboard-metrics-scraper", metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	return err
}

func (a *App) Description() string {
	return a.description
}

func (a *App) Meta() *metav1.ObjectMeta {
	return &a.ObjectMeta
}

func (a *App) Objects() []runtime.Object {
	return a.objects
}
