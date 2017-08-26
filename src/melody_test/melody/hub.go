package melody

import (
	"strconv"
	"sync"
	"sync/atomic"

	log "github.com/thinkboy/log4go"
)

type Hub struct {
	sessions   map[*Session]bool
	broadcast  chan *envelope
	register   chan *Session
	unregister chan *Session
	exit       chan bool
	open       bool
	users      map[string](hubUser)
	usersLock  sync.RWMutex
	stats      hubStats
}

type hubUser struct {
	rid   string
	uid   string
	count int
	extra string
}

type hubStats struct {
	count      int64
	broadcasts int64
}

func newHub() *Hub {
	return &Hub{
		sessions:   make(map[*Session]bool),
		broadcast:  make(chan *envelope),
		register:   make(chan *Session),
		unregister: make(chan *Session),
		exit:       make(chan bool),
		open:       true,
		users:      make(map[string]hubUser),
		usersLock:  sync.RWMutex{},
	}
}

func (h *Hub) run() {
loop:
	for {
		select {
		case s := <-h.register:
			h.sessions[s] = true
			atomic.AddInt64(&h.stats.count, 1)
		case s := <-h.unregister:
			if _, ok := h.sessions[s]; ok {
				delete(h.sessions, s)
				s.conn.Close()
				close(s.output)
				atomic.AddInt64(&h.stats.count, -1)
			}
		case m := <-h.broadcast:
			for s := range h.sessions {
				if m.filter != nil {
					if m.filter(s) {
						s.writeMessage(m)
					}
				} else {
					s.writeMessage(m)
				}
				atomic.AddInt64(&h.stats.broadcasts, 1)
			}
		case <-h.exit:
			for s := range h.sessions {
				delete(h.sessions, s)
				s.conn.Close()
				close(s.output)
			}
			h.open = false
			break loop
		}
	}
}

func (h *Hub) size() int64 {
	return h.stats.count
}

// 用户加入房间(返回用户在房间的次数)
func (h *Hub) Join(uid string, extra string) (count int) {
	h.usersLock.Lock()
	defer h.usersLock.Unlock()

	var user hubUser
	var present bool

	if user, present = h.users[uid]; !present {
		user.uid = uid
		user.extra = extra
		user.count = 1

		h.users[uid] = user
	} else {
		user.count += 1
		user.extra = extra
		h.users[uid] = user
	}
	count = user.count

	log.Debug("join[%v]:(%v)", count, user)

	return count
}

// 用户加入房间(返回用户在房间的次数)
func (h *Hub) Leave(name string, extra string) (count int) {
	h.usersLock.Lock()
	defer h.usersLock.Unlock()

	var user hubUser
	var present bool
	if user, present = h.users[name]; !present {
		count = 0
	} else {
		user.count -= 1

		log.Debug("Leave[%v]:(%v)", user.count, user)

		if user.count <= 0 {
			delete(h.users, name)
		} else {
			h.users[name] = user
		}
		count = user.count
	}

	return count
}

// 获取房间的用户列表
func (h *Hub) GetOnlines(limit int) []interface{} {
	h.usersLock.RLock()
	defer h.usersLock.RUnlock()

	var count int = 0

	onlines := make([]interface{}, 0)
	var user hubUser
	for _, user = range h.users {
		v := make(map[string]string)
		v["uid"] = user.uid
		v["count"] = strconv.Itoa(user.count)
		v["extra"] = user.extra
		onlines = append(onlines, v)
		count++

		if count == limit {
			break
		}
	}
	return onlines
}
