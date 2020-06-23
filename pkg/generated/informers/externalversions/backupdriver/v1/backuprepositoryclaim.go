/*
Copyright the Velero contributors.

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

package v1

import (
	time "time"

	backupdriverv1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/backupdriver/v1"
	versioned "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/generated/listers/backupdriver/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BackupRepositoryClaimInformer provides access to a shared informer and lister for
// BackupRepositoryClaims.
type BackupRepositoryClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BackupRepositoryClaimLister
}

type backupRepositoryClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBackupRepositoryClaimInformer constructs a new informer for BackupRepositoryClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBackupRepositoryClaimInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBackupRepositoryClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBackupRepositoryClaimInformer constructs a new informer for BackupRepositoryClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBackupRepositoryClaimInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackupdriverV1().BackupRepositoryClaims(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BackupdriverV1().BackupRepositoryClaims(namespace).Watch(options)
			},
		},
		&backupdriverv1.BackupRepositoryClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *backupRepositoryClaimInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBackupRepositoryClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *backupRepositoryClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&backupdriverv1.BackupRepositoryClaim{}, f.defaultInformer)
}

func (f *backupRepositoryClaimInformer) Lister() v1.BackupRepositoryClaimLister {
	return v1.NewBackupRepositoryClaimLister(f.Informer().GetIndexer())
}
