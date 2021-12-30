package model

import (
	"encoding/json"
	"github.com/apex/log"
	"github.com/gofiber/websocket/v2"
	"time"
)

type Response struct {
	Command   string        `json:"cmd"`
	Args      []interface{} `json:"args"`
	Warn      string        `json:"warn"`
	Success   bool          `json:"_s"`
	Timestamp int64         `json:"_ts"`
}

func NewResponse(command string, args ...interface{}) *Response {
	return &Response{
		Command:   command,
		Args:      args,
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
	}
}

//goland:noinspection GoUnusedExportedFunction
func NewResponseWithWarn(command, warn string, args ...interface{}) *Response {
	return &Response{
		Command:   command,
		Args:      args,
		Warn:      warn,
		Success:   true,
		Timestamp: time.Now().UnixMilli(),
	}
}

func NewResponseWithError(command, error string, args ...interface{}) *Response {
	return &Response{
		Command:   command,
		Args:      args,
		Warn:      error,
		Success:   false,
		Timestamp: time.Now().UnixMilli(),
	}
}

func (r *Response) Marshal() (res []byte) {
	var err error
	if res, err = json.Marshal(r); err != nil {
		log.WithError(err).Warn("cannot marshal response")
	}
	return
}

func (r *Response) String() string {
	return string(r.Marshal())
}

func (r *Response) Respond(conn *websocket.Conn) {
	if conn == nil {
		log.Warnf("tried to send '%s' to nil connection", r)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, r.Marshal()); err != nil {
		log.WithError(err).Warnf("[ws] cannot send %s to client", r.String())
	}
}
