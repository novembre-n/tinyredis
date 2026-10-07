package server

import (
	"bufio"
	"io"
	"log"
	"net"
)

// Server is a tinyredis TCP server.
type Server struct {
	Addr string
}

// New creates a Server that will listen on addr (e.g. ":6379").
func New(addr string) *Server {
	return &Server{Addr: addr}
}

// ListenAndServe accepts connections until the listener fails.
// It blocks forever, so call it from the main goroutine.
func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	log.Printf("tinyredis listening on %s", s.Addr)

	for {
		conn, err := listener.Accept() // blocks until a client connects
		if err != nil {
			return err
		}
		go s.handleConn(conn) // one goroutine per client
	}
}

// handleConn echoes every line a client sends, until they disconnect.
func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("read error from %s: %v", conn.RemoteAddr(), err)
			}
			return
		}
		if _, err := conn.Write([]byte(line)); err != nil {
			log.Printf("write error to %s: %v", conn.RemoteAddr(), err)
			return
		}
	}
}