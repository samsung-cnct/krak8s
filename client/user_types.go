// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application User Types
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
)

// applicationPostBody user type.
type applicationPostBody struct {
	// Application name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Application's registry
	Registry *string `form:"registry,omitempty" json:"registry,omitempty" xml:"registry,omitempty"`
	// Application config --set argument string
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
	// Application version string
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Validate validates the applicationPostBody type instance.
func (ut *applicationPostBody) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	return
}

// Publicize creates ApplicationPostBody from applicationPostBody
func (ut *applicationPostBody) Publicize() *ApplicationPostBody {
	var pub ApplicationPostBody
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

// ApplicationPostBody user type.
type ApplicationPostBody struct {
	// Application name
	Name string `form:"name" json:"name" xml:"name"`
	// Application's registry
	Registry *string `form:"registry,omitempty" json:"registry,omitempty" xml:"registry,omitempty"`
	// Application config --set argument string
	Set *string `form:"set,omitempty" json:"set,omitempty" xml:"set,omitempty"`
	// Application version string
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the ApplicationPostBody type instance.
func (ut *ApplicationPostBody) Validate() (err error) {
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
	// The number of worker nodes in the projects resource pool
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
		if *ut.NodePoolSize < 3 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, *ut.NodePoolSize, 3, true))
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
	// The number of worker nodes in the projects resource pool
	NodePoolSize int `form:"nodePoolSize" json:"nodePoolSize" xml:"nodePoolSize"`
}

// Validate validates the CluterPostBody type instance.
func (ut *CluterPostBody) Validate() (err error) {
	if ut.NodePoolSize < 3 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, ut.NodePoolSize, 3, true))
	}
	if ut.NodePoolSize > 11 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.nodePoolSize`, ut.NodePoolSize, 11, false))
	}
	return
}
