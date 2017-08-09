// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": project Resource Client
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateProjectPayload is the project create action payload.
type CreateProjectPayload struct {
	// name of project
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateProjectPath computes a request path to the create action of project.
func CreateProjectPath() string {

	return fmt.Sprintf("/v1/projects")
}

// Create a new project entry
func (c *Client) CreateProject(ctx context.Context, path string, payload *CreateProjectPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateProjectRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateProjectRequest create the request corresponding to the create action endpoint of the project resource.
func (c *Client) NewCreateProjectRequest(ctx context.Context, path string, payload *CreateProjectPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeleteProjectPath computes a request path to the delete action of project.
func DeleteProjectPath(projectid string) string {
	param0 := projectid

	return fmt.Sprintf("/v1/projects/%s", param0)
}

// DeleteProject makes a request to the delete action endpoint of the project resource
func (c *Client) DeleteProject(ctx context.Context, path string, project string) (*http.Response, error) {
	req, err := c.NewDeleteProjectRequest(ctx, path, project)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteProjectRequest create the request corresponding to the delete action endpoint of the project resource.
func (c *Client) NewDeleteProjectRequest(ctx context.Context, path string, project string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("project", project)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetProjectPath computes a request path to the get action of project.
func GetProjectPath(projectid string) string {
	param0 := projectid

	return fmt.Sprintf("/v1/projects/%s", param0)
}

// Retrieve project with given id.
func (c *Client) GetProject(ctx context.Context, path string, project string) (*http.Response, error) {
	req, err := c.NewGetProjectRequest(ctx, path, project)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetProjectRequest create the request corresponding to the get action endpoint of the project resource.
func (c *Client) NewGetProjectRequest(ctx context.Context, path string, project string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("project", project)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListProjectPath computes a request path to the list action of project.
func ListProjectPath() string {

	return fmt.Sprintf("/v1/projects")
}

// Retrieve all projects.
func (c *Client) ListProject(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListProjectRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListProjectRequest create the request corresponding to the list action endpoint of the project resource.
func (c *Client) NewListProjectRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
