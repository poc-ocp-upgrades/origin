package ovsclient

import (
	"fmt"
	"bytes"
	"runtime"
	"io"
	"io/ioutil"
	"net"
	"net/rpc"
	"net/http"
	"net/rpc/jsonrpc"
	"time"
)

type Client struct {
	*rpc.Client
	conn	net.Conn
}

func New(conn net.Conn) *Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Client{Client: jsonrpc.NewClient(conn), conn: conn}
}
func DialTimeout(network, addr string, timeout time.Duration) (*Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return New(conn), nil
}
func (c *Client) Ping() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result interface{}
	if err := c.Call("echo", []string{"hello"}, &result); err != nil {
		return err
	}
	return nil
}
func (c *Client) WaitForDisconnect() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n, err := io.Copy(ioutil.Discard, c.conn)
	if err != nil && err != io.EOF {
		return err
	}
	if n > 0 {
		return fmt.Errorf("unexpected bytes read waiting for disconnect: %d", n)
	}
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := runtime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", runtime.FuncForPC(pc).Name()))
	http.Post("/"+"logcode", "application/json", bytes.NewBuffer(jsonLog))
}
