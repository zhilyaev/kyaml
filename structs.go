package main

import (
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
)

type UniqName string

type Config struct {
	// Collections
	Deployments  map[UniqName]apps.Deployment
	DaemonSets   map[UniqName]apps.DaemonSet
	StatefulSets map[UniqName]apps.StatefulSet

	// Specs
	Pod       core.Pod
	Container core.Container

	// Sugar
	Envs        map[string]string
	Annotations map[string]string
	Labels      map[string]string
	Image       Image
}

type Image struct {
	tag        string
	registry   string
	repository string
}
