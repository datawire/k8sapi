package k8sapi

import (
	argoRollouts "github.com/datawire/argo-rollouts-go-client/pkg/client/clientset/versioned"
	typedArgoRollouts "github.com/datawire/argo-rollouts-go-client/pkg/client/clientset/versioned/typed/rollouts/v1alpha1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	admissionregistrationv1alpha1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	internalv1alpha1 "k8s.io/client-go/kubernetes/typed/apiserverinternal/v1alpha1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	authenticationv1 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	authenticationv1alpha1 "k8s.io/client-go/kubernetes/typed/authentication/v1alpha1"
	authenticationv1beta1 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	authorizationv1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	authorizationv1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	autoscalingv1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	autoscalingv2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	autoscalingv2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	certificatesv1 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	certificatesv1alpha1 "k8s.io/client-go/kubernetes/typed/certificates/v1alpha1"
	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coordinationv1 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	coordinationv1alpha1 "k8s.io/client-go/kubernetes/typed/coordination/v1alpha1"
	coordinationv1beta1 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	discoveryv1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	discoveryv1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	eventsv1 "k8s.io/client-go/kubernetes/typed/events/v1"
	eventsv1beta1 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	flowcontrolv1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1"
	flowcontrolv1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
	flowcontrolv1beta2 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	flowcontrolv1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
	networkingv1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1alpha1 "k8s.io/client-go/kubernetes/typed/networking/v1alpha1"
	networkingv1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	nodev1 "k8s.io/client-go/kubernetes/typed/node/v1"
	nodev1alpha1 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	nodev1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	policyv1 "k8s.io/client-go/kubernetes/typed/policy/v1"
	policyv1beta1 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rbacv1alpha1 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	rbacv1beta1 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	resourcev1alpha3 "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
	schedulingv1 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	schedulingv1alpha1 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	schedulingv1beta1 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	storagev1 "k8s.io/client-go/kubernetes/typed/storage/v1"
	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	storagev1beta1 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	storagemigrationv1alpha1 "k8s.io/client-go/kubernetes/typed/storagemigration/v1alpha1"
)

type JoinedClientSetInterface interface {
	kubernetes.Interface
	argoRollouts.Interface
}

type joinedClientSetInterface struct {
	ki  kubernetes.Interface
	ari argoRollouts.Interface
}

func NewJoinedClientSetInterface(ki kubernetes.Interface, ari argoRollouts.Interface) JoinedClientSetInterface {
	return &joinedClientSetInterface{
		ki:  ki,
		ari: ari,
	}
}

func (s *joinedClientSetInterface) Discovery() discovery.DiscoveryInterface {
	return s.ki.Discovery()
}

func (s *joinedClientSetInterface) AdmissionregistrationV1() admissionregistrationv1.AdmissionregistrationV1Interface {
	return s.ki.AdmissionregistrationV1()
}

func (s *joinedClientSetInterface) AdmissionregistrationV1alpha1() admissionregistrationv1alpha1.AdmissionregistrationV1alpha1Interface {
	return s.ki.AdmissionregistrationV1alpha1()
}

func (s *joinedClientSetInterface) AdmissionregistrationV1beta1() admissionregistrationv1beta1.AdmissionregistrationV1beta1Interface {
	return s.ki.AdmissionregistrationV1beta1()
}

func (s *joinedClientSetInterface) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	return s.ki.InternalV1alpha1()
}

func (s *joinedClientSetInterface) AppsV1() appsv1.AppsV1Interface {
	return s.ki.AppsV1()
}

func (s *joinedClientSetInterface) AppsV1beta1() appsv1beta1.AppsV1beta1Interface {
	return s.ki.AppsV1beta1()
}

func (s *joinedClientSetInterface) AppsV1beta2() appsv1beta2.AppsV1beta2Interface {
	return s.ki.AppsV1beta2()
}

func (s *joinedClientSetInterface) AuthenticationV1() authenticationv1.AuthenticationV1Interface {
	return s.ki.AuthenticationV1()
}

func (s *joinedClientSetInterface) AuthenticationV1alpha1() authenticationv1alpha1.AuthenticationV1alpha1Interface {
	return s.ki.AuthenticationV1alpha1()
}

func (s *joinedClientSetInterface) AuthenticationV1beta1() authenticationv1beta1.AuthenticationV1beta1Interface {
	return s.ki.AuthenticationV1beta1()
}

func (s *joinedClientSetInterface) AuthorizationV1() authorizationv1.AuthorizationV1Interface {
	return s.ki.AuthorizationV1()
}

func (s *joinedClientSetInterface) AuthorizationV1beta1() authorizationv1beta1.AuthorizationV1beta1Interface {
	return s.ki.AuthorizationV1beta1()
}

func (s *joinedClientSetInterface) AutoscalingV1() autoscalingv1.AutoscalingV1Interface {
	return s.ki.AutoscalingV1()
}

func (s *joinedClientSetInterface) AutoscalingV2() autoscalingv2.AutoscalingV2Interface {
	return s.ki.AutoscalingV2()
}

func (s *joinedClientSetInterface) AutoscalingV2beta1() autoscalingv2beta1.AutoscalingV2beta1Interface {
	return s.ki.AutoscalingV2beta1()
}

func (s *joinedClientSetInterface) AutoscalingV2beta2() autoscalingv2beta2.AutoscalingV2beta2Interface {
	return s.ki.AutoscalingV2beta2()
}

func (s *joinedClientSetInterface) BatchV1() batchv1.BatchV1Interface {
	return s.ki.BatchV1()
}

func (s *joinedClientSetInterface) BatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return s.ki.BatchV1beta1()
}

func (s *joinedClientSetInterface) CertificatesV1() certificatesv1.CertificatesV1Interface {
	return s.ki.CertificatesV1()
}

func (s *joinedClientSetInterface) CertificatesV1beta1() certificatesv1beta1.CertificatesV1beta1Interface {
	return s.ki.CertificatesV1beta1()
}

func (s *joinedClientSetInterface) CertificatesV1alpha1() certificatesv1alpha1.CertificatesV1alpha1Interface {
	return s.ki.CertificatesV1alpha1()
}

func (s *joinedClientSetInterface) CoordinationV1alpha1() coordinationv1alpha1.CoordinationV1alpha1Interface {
	return s.ki.CoordinationV1alpha1()
}

func (s *joinedClientSetInterface) CoordinationV1beta1() coordinationv1beta1.CoordinationV1beta1Interface {
	return s.ki.CoordinationV1beta1()
}

func (s *joinedClientSetInterface) CoordinationV1() coordinationv1.CoordinationV1Interface {
	return s.ki.CoordinationV1()
}

func (s *joinedClientSetInterface) CoreV1() v1.CoreV1Interface {
	return s.ki.CoreV1()
}

func (s *joinedClientSetInterface) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	return s.ki.DiscoveryV1()
}

func (s *joinedClientSetInterface) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	return s.ki.DiscoveryV1beta1()
}

func (s *joinedClientSetInterface) EventsV1() eventsv1.EventsV1Interface {
	return s.ki.EventsV1()
}

func (s *joinedClientSetInterface) EventsV1beta1() eventsv1beta1.EventsV1beta1Interface {
	return s.ki.EventsV1beta1()
}

func (s *joinedClientSetInterface) ExtensionsV1beta1() extensionsv1beta1.ExtensionsV1beta1Interface {
	return s.ki.ExtensionsV1beta1()
}

func (s *joinedClientSetInterface) FlowcontrolV1() flowcontrolv1.FlowcontrolV1Interface {
	return s.ki.FlowcontrolV1()
}

func (s *joinedClientSetInterface) FlowcontrolV1beta1() flowcontrolv1beta1.FlowcontrolV1beta1Interface {
	return s.ki.FlowcontrolV1beta1()
}

func (s *joinedClientSetInterface) FlowcontrolV1beta2() flowcontrolv1beta2.FlowcontrolV1beta2Interface {
	return s.ki.FlowcontrolV1beta2()
}

func (s *joinedClientSetInterface) FlowcontrolV1beta3() flowcontrolv1beta3.FlowcontrolV1beta3Interface {
	return s.ki.FlowcontrolV1beta3()
}

func (s *joinedClientSetInterface) NetworkingV1() networkingv1.NetworkingV1Interface {
	return s.ki.NetworkingV1()
}

func (s *joinedClientSetInterface) NetworkingV1alpha1() networkingv1alpha1.NetworkingV1alpha1Interface {
	return s.ki.NetworkingV1alpha1()
}

func (s *joinedClientSetInterface) NetworkingV1beta1() networkingv1beta1.NetworkingV1beta1Interface {
	return s.ki.NetworkingV1beta1()
}

func (s *joinedClientSetInterface) NodeV1() nodev1.NodeV1Interface {
	return s.ki.NodeV1()
}

func (s *joinedClientSetInterface) NodeV1alpha1() nodev1alpha1.NodeV1alpha1Interface {
	return s.ki.NodeV1alpha1()
}

func (s *joinedClientSetInterface) NodeV1beta1() nodev1beta1.NodeV1beta1Interface {
	return s.ki.NodeV1beta1()
}

func (s *joinedClientSetInterface) PolicyV1() policyv1.PolicyV1Interface {
	return s.ki.PolicyV1()
}

func (s *joinedClientSetInterface) PolicyV1beta1() policyv1beta1.PolicyV1beta1Interface {
	return s.ki.PolicyV1beta1()
}

func (s *joinedClientSetInterface) RbacV1() rbacv1.RbacV1Interface {
	return s.ki.RbacV1()
}

func (s *joinedClientSetInterface) RbacV1beta1() rbacv1beta1.RbacV1beta1Interface {
	return s.ki.RbacV1beta1()
}

func (s *joinedClientSetInterface) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return s.ki.RbacV1alpha1()
}

func (s *joinedClientSetInterface) ResourceV1alpha3() resourcev1alpha3.ResourceV1alpha3Interface {
	return s.ki.ResourceV1alpha3()
}

func (s *joinedClientSetInterface) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return s.ki.SchedulingV1alpha1()
}

func (s *joinedClientSetInterface) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return s.ki.SchedulingV1beta1()
}

func (s *joinedClientSetInterface) SchedulingV1() schedulingv1.SchedulingV1Interface {
	return s.ki.SchedulingV1()
}

func (s *joinedClientSetInterface) StorageV1beta1() storagev1beta1.StorageV1beta1Interface {
	return s.ki.StorageV1beta1()
}

func (s *joinedClientSetInterface) StorageV1() storagev1.StorageV1Interface {
	return s.ki.StorageV1()
}

func (s *joinedClientSetInterface) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return s.ki.StorageV1alpha1()
}

func (s *joinedClientSetInterface) StoragemigrationV1alpha1() storagemigrationv1alpha1.StoragemigrationV1alpha1Interface {
	return s.ki.StoragemigrationV1alpha1()
}

func (s *joinedClientSetInterface) ArgoprojV1alpha1() typedArgoRollouts.ArgoprojV1alpha1Interface {
	return s.ari.ArgoprojV1alpha1()
}
