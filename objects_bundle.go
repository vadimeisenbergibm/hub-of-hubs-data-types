package datatypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// represents k8s object
type Object interface {
	metav1.Object
	runtime.Object
}

type Bundle interface {
	AddObject(object Object)
	AddDeletedObject(object Object)
}

func NewObjectBundle() *ObjectsBundle {
	return &ObjectsBundle {
		Objects:        make([]Object, 0),
		DeletedObjects: make([]Object, 0),
	}
}

type ObjectsBundle struct {
	Objects        []Object `json:"objects"`
	DeletedObjects []Object `json:"deletedObjects"`
}
