package traffic

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Kuadrant/multicluster-gateway-controller/pkg/apis/v1alpha1"
)

type CreateOrUpdateTraffic func(ctx context.Context, i Interface) error
type DeleteTraffic func(ctx context.Context, i Interface) error

type Interface interface {
	runtime.Object
	metav1.Object
	GetKind() string
	GetHosts() []string
	GetCacheKey() string
	GetNamespaceName() types.NamespacedName
	AddTLS(host string, secret *corev1.Secret)
	HasTLS() bool
	RemoveTLS(host []string)
	GetSpec() interface{}
	GetDNSTargets(ctx context.Context) ([]v1alpha1.Target, error)
	ExposesOwnController() bool
}
