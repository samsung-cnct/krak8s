// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application Media Types
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
	"unicode/utf8"
)

// Application representation type (default view)
//
// Identifier: application/app+json; view=default
type App struct {
	// Application name
	Name   string `form:"name" json:"name" xml:"name"`
	Status *struct {
		// Last deployment time
		DeployedAt time.Time `form:"deployed_at" json:"deployed_at" xml:"deployed_at"`
		// Application specific notification / statuses / notes (if any)
		Notes *string `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
		// Deployment state
		State string `form:"state" json:"state" xml:"state"`
	} `form:"status" json:"status" xml:"status"`
	// Application version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the App media type instance.
func (mt *App) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	if mt.Status == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "status"))
	}
	if mt.Status != nil {

		if mt.Status.State == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.status`, "state"))
		}
		if !(mt.Status.State == "UNKNOWN" || mt.Status.State == "DEPLOYED" || mt.Status.State == "DELETED" || mt.Status.State == "SUPERSEDED" || mt.Status.State == "FAILED" || mt.Status.State == "DELETING") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status.state`, mt.Status.State, []interface{}{"UNKNOWN", "DEPLOYED", "DELETED", "SUPERSEDED", "FAILED", "DELETING"}))
		}
	}
	return
}

// DecodeApp decodes the App instance encoded in resp body.
func (c *Client) DecodeApp(resp *http.Response) (*App, error) {
	var decoded App
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AppCollection is the media type for an array of App (default view)
//
// Identifier: application/app+json; type=collection; view=default
type AppCollection []*App

// Validate validates the AppCollection media type instance.
func (mt AppCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeAppCollection decodes the AppCollection instance encoded in resp body.
func (c *Client) DecodeAppCollection(resp *http.Response) (AppCollection, error) {
	var decoded AppCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Cluster resource representation type (default view)
//
// Identifier: application/cluster+json; view=default
type Cluster struct {
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// Requested node pool size
	NodePoolSize int `form:"nodePoolSize" json:"nodePoolSize" xml:"nodePoolSize"`
	// Lifecycle state
	State string `form:"state" json:"state" xml:"state"`
}

// Validate validates the Cluster media type instance.
func (mt *Cluster) Validate() (err error) {

	if mt.State == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "state"))
	}
	if !(mt.State == "create_requested" || mt.State == "starting" || mt.State == "active" || mt.State == "delete_requested" || mt.State == "deleting" || mt.State == "deleted") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.state`, mt.State, []interface{}{"create_requested", "starting", "active", "delete_requested", "deleting", "deleted"}))
	}
	return
}

// DecodeCluster decodes the Cluster instance encoded in resp body.
func (c *Client) DecodeCluster(resp *http.Response) (*Cluster, error) {
	var decoded Cluster
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// MongoDB ReplicaSet instance representation type (default view)
//
// Identifier: application/mongo+json; view=default
type Mongo struct {
	// Application registry identifier
	Application string `form:"application" json:"application" xml:"application"`
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// Application version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the Mongo media type instance.
func (mt *Mongo) Validate() (err error) {
	if mt.Application == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "application"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}

	return
}

// DecodeMongo decodes the Mongo instance encoded in resp body.
func (c *Client) DecodeMongo(resp *http.Response) (*Mongo, error) {
	var decoded Mongo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Users and tennants of the system are represented as the type Project (default view)
//
// Identifier: application/namespace+json; view=default
type Namespace struct {
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// API href of the namespace
	Href string `form:"href" json:"href" xml:"href"`
	// namespace name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Namespace media type instance.
func (mt *Namespace) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	return
}

// Users and tennants of the system are represented as the type Project (link view)
//
// Identifier: application/namespace+json; view=link
type NamespaceLink struct {
	// API href of the namespace
	Href string `form:"href" json:"href" xml:"href"`
	// namespace name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the NamespaceLink media type instance.
func (mt *NamespaceLink) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	return
}

// DecodeNamespace decodes the Namespace instance encoded in resp body.
func (c *Client) DecodeNamespace(resp *http.Response) (*Namespace, error) {
	var decoded Namespace
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeNamespaceLink decodes the NamespaceLink instance encoded in resp body.
func (c *Client) DecodeNamespaceLink(resp *http.Response) (*NamespaceLink, error) {
	var decoded NamespaceLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// NamespaceCollection is the media type for an array of Namespace (default view)
//
// Identifier: application/namespace+json; type=collection; view=default
type NamespaceCollection []*Namespace

// Validate validates the NamespaceCollection media type instance.
func (mt NamespaceCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// NamespaceCollection is the media type for an array of Namespace (link view)
//
// Identifier: application/namespace+json; type=collection; view=link
type NamespaceLinkCollection []*NamespaceLink

// Validate validates the NamespaceLinkCollection media type instance.
func (mt NamespaceLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeNamespaceCollection decodes the NamespaceCollection instance encoded in resp body.
func (c *Client) DecodeNamespaceCollection(resp *http.Response) (NamespaceCollection, error) {
	var decoded NamespaceCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeNamespaceLinkCollection decodes the NamespaceLinkCollection instance encoded in resp body.
func (c *Client) DecodeNamespaceLinkCollection(resp *http.Response) (NamespaceLinkCollection, error) {
	var decoded NamespaceLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Users and tennants of the system are represented as the type Project (default view)
//
// Identifier: application/project+json; view=default
type Project struct {
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// API href of project
	Href string `form:"href" json:"href" xml:"href"`
	// identity of project
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the Project media type instance.
func (mt *Project) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	if utf8.RuneCountInString(mt.ID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 2, true))
	}
	return
}

// Users and tennants of the system are represented as the type Project (link view)
//
// Identifier: application/project+json; view=link
type ProjectLink struct {
	// API href of project
	Href string `form:"href" json:"href" xml:"href"`
	// identity of project
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the ProjectLink media type instance.
func (mt *ProjectLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if utf8.RuneCountInString(mt.ID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 2, true))
	}
	return
}

// DecodeProject decodes the Project instance encoded in resp body.
func (c *Client) DecodeProject(resp *http.Response) (*Project, error) {
	var decoded Project
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeProjectLink decodes the ProjectLink instance encoded in resp body.
func (c *Client) DecodeProjectLink(resp *http.Response) (*ProjectLink, error) {
	var decoded ProjectLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ProjectCollection is the media type for an array of Project (default view)
//
// Identifier: application/project+json; type=collection; view=default
type ProjectCollection []*Project

// Validate validates the ProjectCollection media type instance.
func (mt ProjectCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ProjectCollection is the media type for an array of Project (link view)
//
// Identifier: application/project+json; type=collection; view=link
type ProjectLinkCollection []*ProjectLink

// Validate validates the ProjectLinkCollection media type instance.
func (mt ProjectLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeProjectCollection decodes the ProjectCollection instance encoded in resp body.
func (c *Client) DecodeProjectCollection(resp *http.Response) (ProjectCollection, error) {
	var decoded ProjectCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeProjectLinkCollection decodes the ProjectLinkCollection instance encoded in resp body.
func (c *Client) DecodeProjectLinkCollection(resp *http.Response) (ProjectLinkCollection, error) {
	var decoded ProjectLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
