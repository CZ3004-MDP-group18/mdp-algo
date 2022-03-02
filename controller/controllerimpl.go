package controller

import (
	"encoding/json"
	"fmt"
	"mdp_algo/common"
	"net"
)

type controllerImpl struct {
	conn net.Conn
}

func NewController() Controller {
	fmt.Println("Dialing...", remoteAddress)
	conn, err := net.Dial("tcp", remoteAddress)
	if err != nil {
		panic("Connection failed")
	}
	return &controllerImpl{
		conn: conn,
	}
}

func send(conn net.Conn, object interface{}) {
	buffer, err := json.Marshal(object)
	if err != nil {
		panic("Marshal failed")
	}
	_, err = conn.Write(buffer)
	fmt.Println(conn.RemoteAddr(), "TCP Sent:", string(buffer))
	if err != nil {
		panic("TCP Write failed")
	}
	return
}
func recv(conn net.Conn, objectPtr interface{}) error {
	fmt.Println(conn.RemoteAddr(), "TCP Waiting for message...")
	buffer := make([]byte, 1024)
	size, err := conn.Read(buffer)
	fmt.Println(conn.RemoteAddr(), "TCP Received:", string(buffer))
	if err != nil {
		panic("TCP Read failed")
	}
	return json.Unmarshal(buffer[:size], objectPtr)
}

func (c *controllerImpl) Move(move string) {
	send(c.conn, MovePayload{
		Move: move,
	})
}

func (c *controllerImpl) Sensor() common.SensorPayload {
	payload := common.SensorPayload{}
	err := recv(c.conn, &payload)
	if err != nil {
		panic(err)
	}

	return payload
}
