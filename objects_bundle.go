package datatypes

import "sigs.k8s.io/controller-runtime/pkg/client"

type ObjectsBundle struct {
	Objects        []client.Object `json:"objects"`
	DeletedObjects []client.Object `json:"deletedObjects"`
}

func NewObjectBundle() *ObjectsBundle {
	return &ObjectsBundle {
		Objects:        make([]client.Object, 0),
		DeletedObjects: make([]client.Object, 0),
	}
}

func (bundle *ObjectsBundle) AddObject(object client.Object) {
	bundle.Objects = append(bundle.Objects, object)
}

func (bundle *ObjectsBundle) AddDeletedObject(object client.Object) {
	bundle.DeletedObjects = append(bundle.DeletedObjects, object)
}
