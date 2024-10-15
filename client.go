package mrpc

import (
	"errors"
	"github/archaeoraptor/mrpc/codec"
	"io"
	"sync"
)

type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call
}

func (call *Call) done() {
	call.Done <- call
}

type Client struct {
	cc       codec.Codec
	opt      *Option
	sending  sync.Mutex
	header   codec.Header
	mu       sync.Mutex
	seq      uint64
	pending  map[uint64]*Call
	closing  bool // user has called close
	shutdown bool // server has told us to stop
}

// Close implements io.Closer.
func (c *Client) Close() error {
	panic("unimplemented")
}

var _ io.Closer = (*Client)(nil)

var ErrShutdiwb = errors.New("connection is shut down")
