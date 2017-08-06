// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "krak8s": Application Contexts
//
// Command:
// $ goagen
// --design=krak8s/design
// --out=$(GOPATH)/src/krak8s
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreateGoaMongoContext provides the goa_mongo create action context.
type CreateGoaMongoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *MongoPostBody
}

// NewCreateGoaMongoContext parses the incoming request URL and body, performs validations and creates the
// context used by the goa_mongo controller create action.
func NewCreateGoaMongoContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateGoaMongoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateGoaMongoContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateGoaMongoContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// DeleteGoaMongoContext provides the goa_mongo delete action context.
type DeleteGoaMongoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Ns   string
	User string
}

// NewDeleteGoaMongoContext parses the incoming request URL and body, performs validations and creates the
// context used by the goa_mongo controller delete action.
func NewDeleteGoaMongoContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteGoaMongoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteGoaMongoContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramNs := req.Params["ns"]
	if len(paramNs) > 0 {
		rawNs := paramNs[0]
		rctx.Ns = rawNs
	}
	paramUser := req.Params["user"]
	if len(paramUser) > 0 {
		rawUser := paramUser[0]
		rctx.User = rawUser
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteGoaMongoContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// ReadGoaMongoContext provides the goa_mongo read action context.
type ReadGoaMongoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Ns   string
	User string
}

// NewReadGoaMongoContext parses the incoming request URL and body, performs validations and creates the
// context used by the goa_mongo controller read action.
func NewReadGoaMongoContext(ctx context.Context, r *http.Request, service *goa.Service) (*ReadGoaMongoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ReadGoaMongoContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramNs := req.Params["ns"]
	if len(paramNs) > 0 {
		rawNs := paramNs[0]
		rctx.Ns = rawNs
	}
	paramUser := req.Params["user"]
	if len(paramUser) > 0 {
		rawUser := paramUser[0]
		rctx.User = rawUser
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ReadGoaMongoContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
