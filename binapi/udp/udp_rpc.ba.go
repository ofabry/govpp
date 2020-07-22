// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package udp

import (
	"context"
	"fmt"
	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
	"io"
)

// RPCService defines RPC service  udp.
type RPCService interface {
	UDPEncapAdd(ctx context.Context, in *UDPEncapAdd) (*UDPEncapAddReply, error)
	UDPEncapDel(ctx context.Context, in *UDPEncapDel) (*UDPEncapDelReply, error)
	UDPEncapDump(ctx context.Context, in *UDPEncapDump) (RPCService_UDPEncapDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) UDPEncapAdd(ctx context.Context, in *UDPEncapAdd) (*UDPEncapAddReply, error) {
	out := new(UDPEncapAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UDPEncapDel(ctx context.Context, in *UDPEncapDel) (*UDPEncapDelReply, error) {
	out := new(UDPEncapDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UDPEncapDump(ctx context.Context, in *UDPEncapDump) (RPCService_UDPEncapDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UDPEncapDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UDPEncapDumpClient interface {
	Recv() (*UDPEncapDetails, error)
	api.Stream
}

type serviceClient_UDPEncapDumpClient struct {
	api.Stream
}

func (c *serviceClient_UDPEncapDumpClient) Recv() (*UDPEncapDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UDPEncapDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}