package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type temporary interface {
	Temporary() bool
}

type message struct {
	conn net.Conn
	data string
}

type client struct {
	name   string
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	wg     sync.WaitGroup
	room   *Room
}

func newClient(conn net.Conn, r *Room, name string) *client {

	c := client{
		name:   name,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		room:   r,
		conn:   conn,
	}

	c.wg.Add(1)
	go c.read()

	return &c
}

func (c *client) read() {
	for {

		line, err := c.reader.ReadString('\n')

		if err == nil {

			m := message{
				data: line,
				conn: c.conn,
			}

			c.room.outgoing <- m
			continue
		}

		if e, is := err.(temporary); is {
			if !e.Temporary() {
				c.wg.Done()
				return
			}
		}

		if err == io.EOF {
			fmt.Println("No connection available")
			c.wg.Done()
			return
		}

	}
}

func (c *client) write(m message) {
	msg := fmt.Sprintf("%s %s", c.name, m.data)
	c.writer.WriteString(msg)
	c.writer.Flush()
}

func (c *client) drop() {
	c.conn.Close()
	c.wg.Wait()
}

type Room struct {
	lisener  net.Listener
	outgoing chan message
	joining  chan net.Conn
	shutdown chan struct{}
	wg       sync.WaitGroup
	clients  []*client
}

func New() *Room {

	room := Room{
		outgoing: make(chan message),
		joining:  make(chan net.Conn),
		shutdown: make(chan struct{}),
	}
	room.start()

	return &room
}

func (r *Room) start() {

	r.wg.Add(2)

	go func() {
		for {

			select {
			case msg := <-r.outgoing:
				r.sendGroupMessage(msg)
			case conn := <-r.joining:
				r.join(conn)
			case <-r.shutdown:
				r.wg.Done()
				return
			}
		}
	}()

	go func() {
		var err error
		r.lisener, err = net.Listen("tcp", ":6666")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Listening port 6666")

		for {

			conn, err := r.lisener.Accept()

			if e, is := err.(temporary); is {
				if !e.Temporary() {
					r.wg.Done()
					return
				}
			}

			r.joining <- conn
		}
	}()

}

func (r *Room) Close() {
	r.lisener.Close()
	close(r.shutdown)
	r.wg.Wait()

	for _, c := range r.clients {
		c.drop()
	}

}

func (r *Room) sendGroupMessage(m message) {

	fmt.Printf("r.clients = %+v\n", m)
	for _, c := range r.clients {
		if c.conn != m.conn {
			//write the mssage here
			fmt.Println("write here")
			c.write(m)
		}
	}

}

func (r *Room) join(c net.Conn) {
	name := fmt.Sprintf("Conn %d", len(r.clients))
	client := newClient(c, r, name)
	r.clients = append(r.clients, client)
}
