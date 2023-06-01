package telnet

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Conn struct {
	conn net.Conn
	in   io.Reader
	out  io.Writer
}

func (c *Conn) Run() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	chErr := make(chan error)
	go c.Send(chErr)
	go c.Receive(chErr)
	select {
	case <-done:
		log.Println("Program finished")
	case err := <-chErr:
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *Conn) Receive(chErr chan error) {
	_, err := io.Copy(c.out, c.conn)
	chErr <- err
}

func (c *Conn) Send(chErr chan error) {
	_, err := io.Copy(c.conn, c.in)
	chErr <- err
}

func NewConn(in io.Reader, out io.Writer, config *Config) (*Conn, error) {
	conn, err := connect(config)
	if err != nil {
		return nil, err
	}
	return &Conn{
		conn: conn,
		in:   in,
		out:  out,
	}, nil
}

func connect(config *Config) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", config.Host+":"+config.Port, config.Timeout)
	if err != nil {
		time.Sleep(config.Timeout)
		return nil, err
	}
	return conn, nil
}

type Config struct {
	Host    string
	Port    string
	Timeout time.Duration
}
