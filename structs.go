package main

import (
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type UniqName string

func (u UniqName) String() string {
	return string(u)

}

var Cfg Config

type Config struct {
	// Collections
	//Collections `json:",inline"`
	Deployments  map[UniqName]apps.Deployment  `json:"deployments,omitempty"`
	DaemonSets   map[UniqName]apps.DaemonSet   `json:"daemonSets,omitempty"`
	StatefulSets map[UniqName]apps.StatefulSet `json:"statefulSets,omitempty"`

	// Specs
	Container core.Container  `json:"container,omitempty"`
	Metadata  meta.ObjectMeta `json:"metadata,omitempty"`
}

type Image struct {
	tag        string
	registry   string
	repository string
}

type Collections struct {
	Deployments  map[UniqName]apps.Deployment  `json:"deployments,omitempty"`
	DaemonSets   map[UniqName]apps.DaemonSet   `json:"daemonSets,omitempty"`
	StatefulSets map[UniqName]apps.StatefulSet `json:"statefulSets,omitempty"`
}
