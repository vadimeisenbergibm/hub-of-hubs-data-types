package bundle

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type Object interface {
	metav1.Object
	runtime.Object
}

type Bundle interface {
	AddObject(object Object)
	AddDeletedObject(uid types.UID)
}
