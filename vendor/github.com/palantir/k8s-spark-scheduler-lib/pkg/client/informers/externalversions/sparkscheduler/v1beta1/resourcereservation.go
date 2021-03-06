/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	sparkschedulerv1beta1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/apis/sparkscheduler/v1beta1"
	versioned "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/clientset/versioned"
	internalinterfaces "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/palantir/k8s-spark-scheduler-lib/pkg/client/listers/sparkscheduler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceReservationInformer provides access to a shared informer and lister for
// ResourceReservations.
type ResourceReservationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ResourceReservationLister
}

type resourceReservationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceReservationInformer constructs a new informer for ResourceReservation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceReservationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceReservationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceReservationInformer constructs a new informer for ResourceReservation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceReservationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SparkschedulerV1beta1().ResourceReservations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SparkschedulerV1beta1().ResourceReservations(namespace).Watch(options)
			},
		},
		&sparkschedulerv1beta1.ResourceReservation{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceReservationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceReservationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceReservationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sparkschedulerv1beta1.ResourceReservation{}, f.defaultInformer)
}

func (f *resourceReservationInformer) Lister() v1beta1.ResourceReservationLister {
	return v1beta1.NewResourceReservationLister(f.Informer().GetIndexer())
}
