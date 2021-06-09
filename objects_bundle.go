package datatypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ObjectsBundle struct {
	Objects        []interface{} `json:"objects"`
	DeletedObjects []interface{} `json:"deletedObjects"`
}

type Object interface {
	metav1.Object
	runtime.Object
}

func NewObjectBundle() *ObjectsBundle {
	return &ObjectsBundle {
		Objects:        make([]interface{}, 0),
		DeletedObjects: make([]interface{}, 0),
	}
}

func (bundle *ObjectsBundle) AddObject(object interface{}) {
	bundle.Objects = append(bundle.Objects, object)
}

func (bundle *ObjectsBundle) AddDeletedObject(object interface{}) {
	bundle.DeletedObjects = append(bundle.DeletedObjects, object)
}
