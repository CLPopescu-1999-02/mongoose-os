// Code generated by clubbygen.
// GENERATED FILE DO NOT EDIT
// +build !clubby_strict

package swupdate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"cesanta.com/clubby"
	"cesanta.com/clubby/endpoint"
	"cesanta.com/clubby/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/trace"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://cesanta.com/mg_rpc/service/v1/SWUpdate"

type ListSectionsResult struct {
	Section  *string `json:"section,omitempty"`
	Version  *string `json:"version,omitempty"`
	Writable *bool   `json:"writable,omitempty"`
}

type UpdateArgs struct {
	Blob           *string `json:"blob,omitempty"`
	Blob_type      *string `json:"blob_type,omitempty"`
	Blob_url       *string `json:"blob_url,omitempty"`
	Commit_timeout *int64  `json:"commit_timeout,omitempty"`
	Section        *string `json:"section,omitempty"`
	Version        *string `json:"version,omitempty"`
}

type Service interface {
	Commit(ctx context.Context) error
	ListSections(ctx context.Context) ([]ListSectionsResult, error)
	Revert(ctx context.Context) error
	Update(ctx context.Context, args *UpdateArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
	TraceCall(context.Context, string, *frame.Command) (context.Context, trace.Trace, func(*error))
}

func NewClient(i Instance, addr string) Service {
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Commit(pctx context.Context) (err error) {
	cmd := &frame.Command{
		Cmd: "/v1/SWUpdate.Commit",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) ListSections(pctx context.Context) (res []ListSectionsResult, err error) {
	cmd := &frame.Command{
		Cmd: "/v1/SWUpdate.ListSections",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	var r []ListSectionsResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Revert(pctx context.Context) (err error) {
	cmd := &frame.Command{
		Cmd: "/v1/SWUpdate.Revert",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) Update(pctx context.Context, args *UpdateArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "/v1/SWUpdate.Update",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func RegisterService(i *clubby.Instance, impl Service) error {
	s := &_Server{impl}
	i.RegisterCommandHandler("/v1/SWUpdate.Commit", s.Commit)
	i.RegisterCommandHandler("/v1/SWUpdate.ListSections", s.ListSections)
	i.RegisterCommandHandler("/v1/SWUpdate.Revert", s.Revert)
	i.RegisterCommandHandler("/v1/SWUpdate.Update", s.Update)
	i.RegisterService(ServiceID, _ServiceDefinition)
	return nil
}

type _Server struct {
	impl Service
}

func (s *_Server) Commit(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	return nil, s.impl.Commit(ctx)
}

func (s *_Server) ListSections(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	return s.impl.ListSections(ctx)
}

func (s *_Server) Revert(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	return nil, s.impl.Revert(ctx)
}

func (s *_Server) Update(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args UpdateArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Update(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "doc": "SWUpdate service provides a way to update device's software.",
  "methods": {
    "Commit": {
      "doc": "Commit a previously initiated update."
    },
    "ListSections": {
      "doc": "Returns a list of components of the device's software. Each section is updated individually.",
      "result": {
        "items": {
          "properties": {
            "section": {
              "type": "string"
            },
            "version": {
              "type": "string"
            },
            "writable": {
              "type": "boolean"
            }
          },
          "type": "object"
        },
        "type": "array"
      }
    },
    "Revert": {
      "doc": "Revert a previously initiated update."
    },
    "Update": {
      "args": {
        "blob": {
          "doc": "Image as a string, if appropriate.",
          "type": "string"
        },
        "blob_type": {
          "doc": "Type of the blob. Valid values: manifest, zip.",
          "type": "string"
        },
        "blob_url": {
          "doc": "URL pointing to the image if it's too big to fit in the ` + "`" + `blob` + "`" + `.",
          "type": "string"
        },
        "commit_timeout": {
          "doc": "Normally update is committed if firmware init succeeds, If timeout is set and non-zero, the update will require an explicit commit. If the specified time expires without a commit, update is rolled back.",
          "type": "integer"
        },
        "section": {
          "doc": "Name of the section to update.",
          "type": "string"
        },
        "version": {
          "doc": "Optional version of the new image.",
          "type": "string"
        }
      },
      "doc": "Instructs the device to update a given section."
    }
  },
  "name": "/v1/SWUpdate",
  "namespace": "http://cesanta.com/mg_rpc/service",
  "visibility": "private"
}`))
