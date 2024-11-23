// Code generated by Kitex v0.10.1. DO NOT EDIT.

package rightservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	api "github.com/zlllgp/vegas/kitex_gen/api"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"QueryRight": kitex.NewMethodInfo(
		queryRightHandler,
		newQueryRightArgs,
		newQueryRightResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	rightServiceServiceInfo                = NewServiceInfo()
	rightServiceServiceInfoForClient       = NewServiceInfoForClient()
	rightServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return rightServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return rightServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return rightServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RightService"
	handlerType := (*api.RightService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.10.1",
		Extra:           extra,
	}
	return svcInfo
}

func queryRightHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.RightRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.RightService).QueryRight(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *QueryRightArgs:
		success, err := handler.(api.RightService).QueryRight(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryRightResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newQueryRightArgs() interface{} {
	return &QueryRightArgs{}
}

func newQueryRightResult() interface{} {
	return &QueryRightResult{}
}

type QueryRightArgs struct {
	Req *api.RightRequest
}

func (p *QueryRightArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *QueryRightArgs) Unmarshal(in []byte) error {
	msg := new(api.RightRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryRightArgs_Req_DEFAULT *api.RightRequest

func (p *QueryRightArgs) GetReq() *api.RightRequest {
	if !p.IsSetReq() {
		return QueryRightArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryRightArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QueryRightArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QueryRightResult struct {
	Success *api.RightResponse
}

var QueryRightResult_Success_DEFAULT *api.RightResponse

func (p *QueryRightResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *QueryRightResult) Unmarshal(in []byte) error {
	msg := new(api.RightResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryRightResult) GetSuccess() *api.RightResponse {
	if !p.IsSetSuccess() {
		return QueryRightResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryRightResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.RightResponse)
}

func (p *QueryRightResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QueryRightResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryRight(ctx context.Context, Req *api.RightRequest) (r *api.RightResponse, err error) {
	var _args QueryRightArgs
	_args.Req = Req
	var _result QueryRightResult
	if err = p.c.Call(ctx, "QueryRight", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
