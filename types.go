package go_socketio_redis_adapter

import "sync"

// request types.
const (
	roomLenReqType   = "0"
	clearRoomReqType = "1"
	allRoomReqType   = "2"
)

// request structs.
type roomLenRequest struct {
	RequestType string
	RequestID   string
	Room        string
	numSub      int        `json:"-"`
	msgCount    int        `json:"-"`
	connections int        `json:"-"`
	mutex       sync.Mutex `json:"-"`
	done        chan bool  `json:"-"`
}

// response struct.
type roomLenResponse struct {
	RequestType string
	RequestID   string
	Connections int
}

type clearRoomRequest struct {
	RequestType string
	RequestID   string
	Room        string
	UUID        string
}

type allRoomRequest struct {
	RequestType string
	RequestID   string
	rooms       map[string]bool `json:"-"`
	numSub      int             `json:"-"`
	msgCount    int             `json:"-"`
	mutex       sync.Mutex      `json:"-"`
	done        chan bool       `json:"-"`
}

type allRoomResponse struct {
	RequestType string
	RequestID   string
	Rooms       []string
}
