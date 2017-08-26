package melody

import (
	//	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/thinkboy/log4go"
)

type handleMessageFunc func(*Session, []byte)
type handleErrorFunc func(*Session, error)
type handleSessionFunc func(*Hub, *Session)
type filterFunc func(*Session) bool

// Melody implements a websocket manager.
type Melody struct {
	Config               *Config
	Upgrader             *websocket.Upgrader
	messageHandler       handleMessageFunc
	messageHandlerBinary handleMessageFunc
	errorHandler         handleErrorFunc
	connectHandler       handleSessionFunc
	disconnectHandler    handleSessionFunc
	pongHandler          handleSessionFunc
	//	hub                  *Hub
	hubs     map[string]*Hub
	hubsLock sync.RWMutex
}

// New creates a new melody instance with default Upgrader and Config.
func New() *Melody {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	//	hub := newHub()

	//	go hub.run()

	return &Melody{
		Config:               newConfig(),
		Upgrader:             upgrader,
		messageHandler:       func(*Session, []byte) {},
		messageHandlerBinary: func(*Session, []byte) {},
		errorHandler:         func(*Session, error) {},
		connectHandler:       func(*Hub, *Session) {},
		disconnectHandler:    func(*Hub, *Session) {},
		pongHandler:          func(*Hub, *Session) {},
		//hub:                  hub,
		hubs:     make(map[string](*Hub)),
		hubsLock: sync.RWMutex{},
	}
}

func (m *Melody) getHub(name string) *Hub {
	var hub *Hub
	var present bool

	m.hubsLock.RLock()
	if hub, present = m.hubs[name]; !present {
		m.hubsLock.RUnlock()
		m.hubsLock.Lock()
		if hub, present = m.hubs[name]; !present {
			hub = newHub()
			go hub.run()
			m.hubs[name] = hub
		}
		m.hubsLock.Unlock()
	} else {
		m.hubsLock.RUnlock()
	}

	return hub
}

// HandleConnect fires fn when a session connects.
func (m *Melody) HandleConnect(fn func(*Hub, *Session)) {
	m.connectHandler = fn
}

// HandleDisconnect fires fn when a session disconnects.
func (m *Melody) HandleDisconnect(fn func(*Hub, *Session)) {
	m.disconnectHandler = fn
}

// HandlePong fires fn when a pong is received from a session.
func (m *Melody) HandlePong(fn func(*Hub, *Session)) {
	m.pongHandler = fn
}

// HandleMessage fires fn when a text message comes in.
func (m *Melody) HandleMessage(fn func(*Session, []byte)) {
	m.messageHandler = fn
}

// HandleMessageBinary fires fn when a binary message comes in.
func (m *Melody) HandleMessageBinary(fn func(*Session, []byte)) {
	m.messageHandlerBinary = fn
}

// HandleError fires fn when a session has an error.
func (m *Melody) HandleError(fn func(*Session, error)) {
	m.errorHandler = fn
}

// HandleRequest upgrades http requests to websocket connections and dispatches them to be handled by the melody instance.
func (m *Melody) HandleRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := m.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		m.errorHandler(nil, err)
		return
	}

	// 检查参数
	rid := r.FormValue("rid")
	if len(rid) <= 0 {
		log.Error("rid empty")
		conn.Close()
		return
	}
	//	log.Info("HandleRequest rid = %v", rid)

	uid := r.FormValue("uid")
	if len(uid) <= 0 {
		log.Error("uid empty")
		conn.Close()
		return
	}
	//	log.Info("HandleRequest uid = %v", uid)

	extra := r.FormValue("extra")
	if len(extra) <= 0 {
		log.Error("extra empty")
		conn.Close()
		return
	}
	//	log.Info("HandleRequest extra = %v", extra)

	session := &Session{
		Request: r,
		conn:    conn,
		output:  make(chan *envelope, m.Config.MessageBufferSize),
		melody:  m,
		Rid:     rid,
		Uid:     uid,
		Extra:   extra,
	}

	hub := m.getHub(rid)

	//m.hub.register <- session

	hub.register <- session

	go m.connectHandler(hub, session)

	go session.writePump()

	session.readPump(hub)

	//	log.Debug("HandleRequest[%d] rid: %v, uid: %v", 1, rid, uid)

	if hub.open {
		hub.unregister <- session
	}

	go m.disconnectHandler(hub, session)

	//	log.Debug("HandleRequest[%d] rid: %v, uid: %v", 2, rid, uid)
}

// Broadcasts a text message to all hubs
func (m *Melody) BroadcastAll(msg []byte) {
	message := &envelope{t: websocket.TextMessage, msg: msg}
	m.hubsLock.Lock()
	defer m.hubsLock.Unlock()
	for _, hub := range m.hubs {
		if hub.size() > 0 {
			hub.broadcast <- message
		}
	}
}

// Broadcast broadcasts a text message to all sessions.
func (m *Melody) Broadcast(name string, msg []byte) {
	message := &envelope{t: websocket.TextMessage, msg: msg}
	hub := m.getHub(name)
	if hub.size() > 0 {
		hub.broadcast <- message
	}
}

// BroadcastFilter broadcasts a text message to all sessions that fn returns true for.
func (m *Melody) BroadcastFilterAll(msg []byte, fn func(*Session) bool) {
	message := &envelope{t: websocket.TextMessage, msg: msg, filter: fn}
	m.hubsLock.Lock()
	defer m.hubsLock.Unlock()
	for _, hub := range m.hubs {
		if hub.size() > 0 {
			hub.broadcast <- message
		}
	}
}

// BroadcastFilter broadcasts a text message to all sessions that fn returns true for.
func (m *Melody) BroadcastFilter(name string, msg []byte, fn func(*Session) bool) {
	message := &envelope{t: websocket.TextMessage, msg: msg, filter: fn}
	hub := m.getHub(name)
	if hub.size() > 0 {
		hub.broadcast <- message
	}
}

// BroadcastOthers broadcasts a text message to all sessions except session s.
func (m *Melody) BroadcastOthers(name string, msg []byte, s *Session) {
	m.BroadcastFilter(name, msg, func(q *Session) bool {
		return s != q
	})
}

// BroadcastBinary broadcasts a binary message to all sessions.
func (m *Melody) BroadcastBinary(name string, msg []byte) {
	message := &envelope{t: websocket.BinaryMessage, msg: msg}
	hub := m.getHub(name)
	if hub.size() > 0 {
		hub.broadcast <- message
	}
}

// BroadcastBinaryFilter broadcasts a binary message to all sessions that fn returns true for.
func (m *Melody) BroadcastBinaryFilter(name string, msg []byte, fn func(*Session) bool) {
	message := &envelope{t: websocket.BinaryMessage, msg: msg, filter: fn}
	hub := m.getHub(name)
	if hub.size() > 0 {
		hub.broadcast <- message
	}
}

// BroadcastBinaryOthers broadcasts a binary message to all sessions except session s.
func (m *Melody) BroadcastBinaryOthers(name string, msg []byte, s *Session) {
	m.BroadcastBinaryFilter(name, msg, func(q *Session) bool {
		return s != q
	})
}

// Close closes the melody instance and all connected sessions.
func (m *Melody) Close() {
	m.hubsLock.Lock()
	defer m.hubsLock.Unlock()

	for name, hub := range m.hubs {
		delete(m.hubs, name)
		hub.exit <- true
	}
}

func (m *Melody) GetOnlines(rid string, limit int) []interface{} {
	hub := m.getHub(rid)
	if hub != nil {
		return hub.GetOnlines(limit)
	} else {
		return make([]interface{}, 0, 0)
	}
}
