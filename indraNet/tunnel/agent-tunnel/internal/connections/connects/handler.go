package connects

import (
	"agent-tunnel/proto"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"
	protobuf "google.golang.org/protobuf/proto"
)

func (c *TunnelClient) handleMessage(ctx context.Context, msg *proto.Envelope) error {
	switch m := msg.Message.(type) {
	case *proto.Envelope_Request:
		return c.handleConnect(ctx, m.Request)
	case *proto.Envelope_Data:
		return c.handleStream(ctx, m.Data)
	case *proto.Envelope_Close:
		return c.handleClose(ctx, m.Close)
	case *proto.Envelope_Control:
		log.Printf("[TUNNEL] control msg: %v", msg.Message)
	default:
		return fmt.Errorf("unknown message type: %T", m)
	}

	return nil
}

func (c *TunnelClient) handleConnect(ctx context.Context, req *proto.TunnelRequest) error {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		return fmt.Errorf("dial [ERROR]: %w", err)
	}

	c.mu.Lock()
	c.Streams[req.Id] = conn
	defer c.mu.Unlock()

	if len(req.Body) > 0 {
		conn.Write(req.Body)
	}

	// Create a new TCP connection for the stream
	go c.pipeToLocal(req.Id, conn)

	log.Printf("Stream %s connected [INFO]", req.Id)
	return nil
}

func (c *TunnelClient) handleStream(ctx context.Context, chunk *proto.TunnelData) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if conn, ok := c.Streams[chunk.Id]; ok {
		if _, err := conn.Write(chunk.Chunk); err != nil {
			return fmt.Errorf("write to stream %s: %w", chunk.Id, err)
		}
	} else {
		log.Printf("Received data for unknown stream ID: %s", chunk.Id)
	}
	return nil
}

func (c *TunnelClient) handleClose(ctx context.Context, closeMsg *proto.TunnelClose) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if conn, ok := c.Streams[closeMsg.Id]; ok {
		conn.Close()
		delete(c.Streams, closeMsg.Id)
	}
	return nil
}

func (c *TunnelClient) pipeToLocal(id string, conn net.Conn) error {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != net.ErrClosed {
				log.Printf("Error reading from stream %s: %v", id, err)
			}
			return err
		}

		dataMsg := &proto.TunnelData{
			Id:    id,
			Chunk: buf[:n],
		}

		env := &proto.Envelope{
			Message: &proto.Envelope_Data{Data: dataMsg},
		}

		b, _ := protobuf.Marshal(env)
		c.conn.WriteMessage(websocket.BinaryMessage, b)

	}
}
