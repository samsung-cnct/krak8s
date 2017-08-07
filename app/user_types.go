// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application User Types
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// chartPostBody user type.
type chartPostBody struct {
	// Chart name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Chart's registry
	Registry *string `form:"registry,omitempty" json:"registry,omitempty" xml:"registry,omitempty"`
	// Chart config --set argument string
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
	// Chart version string
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Validate validates the chartPostBody type instance.
func (ut *chartPostBody) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// Publicize creates ChartPostBody from chartPostBody
func (ut *chartPostBody) Publicize() *ChartPostBody {
	var pub ChartPostBody
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Registry != nil {
		pub.Registry = ut.Registry
	}
	if ut.Set != nil {
		pub.Set = ut.Set
	}
	if ut.Version != nil {
		pub.Version = *ut.Version
	}
	return &pub
}

// ChartPostBody user type.
type ChartPostBody struct {
	// Chart name
	Name string `form:"name" json:"name" xml:"name"`
	// Chart's registry
	Registry *string `form:"registry,omitempty" json:"registry,omitempty" xml:"registry,omitempty"`
	// Chart config --set argument string
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
	// Chart version string
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the ChartPostBody type instance.
func (ut *ChartPostBody) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// cluterPostBody user type.
type cluterPostBody struct {
	// The number of real nodes in the pool
	NodePoolSize *int `form:"nodePoolSize,omitempty" json:"nodePoolSize,omitempty" xml:"nodePoolSize,omitempty"`
}

// Finalize sets the default values for cluterPostBody type instance.
func (ut *cluterPostBody) Finalize() {
	var defaultNodePoolSize = 3
	if ut.NodePoolSize == nil {
		ut.NodePoolSize = &defaultNodePoolSize
	}
}

// Validate validates the cluterPostBody type instance.
func (ut *cluterPostBody) Validate() (err error) {
	if ut.NodePoolSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "nodePoolSize"))
	}
	if ut.NodePoolSize != nil {
		if *ut.NodePoolSize < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, *ut.NodePoolSize, 1, true))
		}
	}
	if ut.NodePoolSize != nil {
		if *ut.NodePoolSize > 11 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, *ut.NodePoolSize, 11, false))
		}
	}
	return
}

// Publicize creates CluterPostBody from cluterPostBody
func (ut *cluterPostBody) Publicize() *CluterPostBody {
	var pub CluterPostBody
	if ut.NodePoolSize != nil {
		pub.NodePoolSize = *ut.NodePoolSize
	}
	return &pub
}

// CluterPostBody user type.
type CluterPostBody struct {
	// The number of real nodes in the pool
	NodePoolSize int `form:"nodePoolSize" json:"nodePoolSize" xml:"nodePoolSize"`
}

// Validate validates the CluterPostBody type instance.
func (ut *CluterPostBody) Validate() (err error) {
	if ut.NodePoolSize < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, ut.NodePoolSize, 1, true))
	}
	if ut.NodePoolSize > 11 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, ut.NodePoolSize, 11, false))
	}
	return
}

// mongoPostBody user type.
type mongoPostBody struct {
	// Appplication Registry Identifier
	Application *string `form:"application,omitempty" json:"application,omitempty" xml:"application,omitempty"`
	// Appplication Version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Finalize sets the default values for mongoPostBody type instance.
func (ut *mongoPostBody) Finalize() {
	var defaultApplication = "quay.io/samsung_cnct/mongodb-replicaset"
	if ut.Application == nil {
		ut.Application = &defaultApplication
	}
	var defaultVersion = "v1.2.0"
	if ut.Version == nil {
		ut.Version = &defaultVersion
	}
}

// Validate validates the mongoPostBody type instance.
func (ut *mongoPostBody) Validate() (err error) {
	if ut.Application == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// Publicize creates MongoPostBody from mongoPostBody
func (ut *mongoPostBody) Publicize() *MongoPostBody {
	var pub MongoPostBody
	if ut.Application != nil {
		pub.Application = *ut.Application
	}
	if ut.Version != nil {
		pub.Version = *ut.Version
	}
	return &pub
}

// MongoPostBody user type.
type MongoPostBody struct {
	// Appplication Registry Identifier
	Application string `form:"application" json:"application" xml:"application"`
	// Appplication Version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the MongoPostBody type instance.
func (ut *MongoPostBody) Validate() (err error) {
	if ut.Application == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if ut.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}
