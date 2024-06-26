// Code generated by Kitex v0.10.1. DO NOT EDIT.

package vegas

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	api "vegas/kitex_gen/api"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Draw": kitex.NewMethodInfo(
		drawHandler,
		newDrawArgs,
		newDrawResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	vegasServiceInfo                = NewServiceInfo()
	vegasServiceInfoForClient       = NewServiceInfoForClient()
	vegasServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return vegasServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return vegasServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return vegasServiceInfoForClient
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
	serviceName := "Vegas"
	handlerType := (*api.Vegas)(nil)
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

func drawHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(api.DrawRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(api.Vegas).Draw(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DrawArgs:
		success, err := handler.(api.Vegas).Draw(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DrawResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDrawArgs() interface{} {
	return &DrawArgs{}
}

func newDrawResult() interface{} {
	return &DrawResult{}
}

type DrawArgs struct {
	Req *api.DrawRequest
}

func (p *DrawArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(api.DrawRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DrawArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DrawArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DrawArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DrawArgs) Unmarshal(in []byte) error {
	msg := new(api.DrawRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DrawArgs_Req_DEFAULT *api.DrawRequest

func (p *DrawArgs) GetReq() *api.DrawRequest {
	if !p.IsSetReq() {
		return DrawArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DrawArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DrawArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DrawResult struct {
	Success *api.DrawResult
}

var DrawResult_Success_DEFAULT *api.DrawResult

func (p *DrawResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(api.DrawResult)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DrawResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DrawResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DrawResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DrawResult) Unmarshal(in []byte) error {
	msg := new(api.DrawResult)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DrawResult) GetSuccess() *api.DrawResult {
	if !p.IsSetSuccess() {
		return DrawResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DrawResult) SetSuccess(x interface{}) {
	p.Success = x.(*api.DrawResult)
}

func (p *DrawResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DrawResult) GetResult() interface{} {
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

func (p *kClient) Draw(ctx context.Context, Req *api.DrawRequest) (r *api.DrawResult, err error) {
	var _args DrawArgs
	_args.Req = Req
	var _result DrawResult
	if err = p.c.Call(ctx, "Draw", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
