// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": mongo Resource Client
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

// CreateMongoPath computes a request path to the create action of mongo.
func CreateMongoPath(project string, ns string) string {
	param0 := project
	param1 := ns

	return fmt.Sprintf("/v1/projects/%s/ns/%s/mongo", param0, param1)
}

// Create a MongoDB
func (c *Client) CreateMongo(ctx context.Context, path string, payload *MongoPostBody, contentType string) (*http.Response, error) {
	req, err := c.NewCreateMongoRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateMongoRequest create the request corresponding to the create action endpoint of the mongo resource.
func (c *Client) NewCreateMongoRequest(ctx context.Context, path string, payload *MongoPostBody, contentType string) (*http.Request, error) {
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

// DeleteMongoPath computes a request path to the delete action of mongo.
func DeleteMongoPath(project string, ns string) string {
	param0 := project
	param1 := ns

	return fmt.Sprintf("/v1/projects/%s/ns/%s/mongo", param0, param1)
}

// Delete the MongoDB Deloyment
func (c *Client) DeleteMongo(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteMongoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteMongoRequest create the request corresponding to the delete action endpoint of the mongo resource.
func (c *Client) NewDeleteMongoRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetMongoPath computes a request path to the get action of mongo.
func GetMongoPath(project string, ns string) string {
	param0 := project
	param1 := ns

	return fmt.Sprintf("/v1/projects/%s/ns/%s/mongo", param0, param1)
}

// Get the status of the MongoDB Deloyment
func (c *Client) GetMongo(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetMongoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetMongoRequest create the request corresponding to the get action endpoint of the mongo resource.
func (c *Client) NewGetMongoRequest(ctx context.Context, path string) (*http.Request, error) {
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