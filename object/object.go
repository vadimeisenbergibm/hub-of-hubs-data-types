package object

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Object interface {
	metav1.Object
	runtime.Object
}

type HubOfHubsObject struct {
	Object  Object `json:"object"`
	Deleted bool   `json:"deleted"`
}
