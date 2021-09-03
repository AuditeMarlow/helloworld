// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/helloworld.proto

package helloworld

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Helloworld service

func NewHelloworldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Helloworld service

type HelloworldService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (Helloworld_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Helloworld_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (Helloworld_BidiStreamService, error)
}

type helloworldService struct {
	c    client.Client
	name string
}

func NewHelloworldService(name string, c client.Client) HelloworldService {
	return &helloworldService{
		c:    c,
		name: name,
	}
}

func (c *helloworldService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Helloworld.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldService) ClientStream(ctx context.Context, opts ...client.CallOption) (Helloworld_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &helloworldServiceClientStream{stream}, nil
}

type Helloworld_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type helloworldServiceClientStream struct {
	stream client.Stream
}

func (x *helloworldServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *helloworldService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (Helloworld_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "Helloworld.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &helloworldServiceServerStream{stream}, nil
}

type Helloworld_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type helloworldServiceServerStream struct {
	stream client.Stream
}

func (x *helloworldServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloworldService) BidiStream(ctx context.Context, opts ...client.CallOption) (Helloworld_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "Helloworld.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &helloworldServiceBidiStream{stream}, nil
}

type Helloworld_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type helloworldServiceBidiStream struct {
	stream client.Stream
}

func (x *helloworldServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *helloworldServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Helloworld service

type HelloworldHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, Helloworld_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, Helloworld_ServerStreamStream) error
	BidiStream(context.Context, Helloworld_BidiStreamStream) error
}

func RegisterHelloworldHandler(s server.Server, hdlr HelloworldHandler, opts ...server.HandlerOption) error {
	type helloworld interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type Helloworld struct {
		helloworld
	}
	h := &helloworldHandler{hdlr}
	return s.Handle(s.NewHandler(&Helloworld{h}, opts...))
}

type helloworldHandler struct {
	HelloworldHandler
}

func (h *helloworldHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.HelloworldHandler.Call(ctx, in, out)
}

func (h *helloworldHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.HelloworldHandler.ClientStream(ctx, &helloworldClientStreamStream{stream})
}

type Helloworld_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type helloworldClientStreamStream struct {
	stream server.Stream
}

func (x *helloworldClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *helloworldHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HelloworldHandler.ServerStream(ctx, m, &helloworldServerStreamStream{stream})
}

type Helloworld_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type helloworldServerStreamStream struct {
	stream server.Stream
}

func (x *helloworldServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *helloworldHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.HelloworldHandler.BidiStream(ctx, &helloworldBidiStreamStream{stream})
}

type Helloworld_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type helloworldBidiStreamStream struct {
	stream server.Stream
}

func (x *helloworldBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *helloworldBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *helloworldBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *helloworldBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *helloworldBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *helloworldBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
