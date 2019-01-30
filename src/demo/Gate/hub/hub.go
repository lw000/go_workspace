package lwhub

import (
	"Gate/utilty"
	"container/list"
	"sync"
)

type HandlerFunc func(data []byte)

type Key struct {
	mid uint16
	sid uint16
}

type HandlerItem struct {
	f       HandlerFunc
	eventId uint32
}

type Handler struct {
	m  sync.RWMutex
	ls *list.List
	f  HandlerFunc
}

type Hub struct {
	m sync.Map
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Register(mid, sid uint16, f HandlerFunc) {
	k := Key{mid: mid, sid: sid}
	v, ok := h.m.Load(k)
	if ok {
		hd := v.(*Handler)
		hd.Add(f)
	} else {
		hd := NewHandler()
		hd.f = f
		hd.Add(f)
		h.m.Store(k, hd)
	}
}

func (h *Hub) DistMessage(mid, sid uint16, eventId uint32, data []byte) {
	k := Key{mid: mid, sid: mid}
	h.m.Range(func(key, value interface{}) bool {
		if key.(Key) == k {
			hd := value.(*Handler)
			item := hd.Get(eventId)
			if item != nil {
				item.f(data)
			}
			hd.Remove(eventId)
			return true
		}
		return false
	})
}

func NewHandler() *Handler {
	return &Handler{
		ls: list.New(),
	}
}

func (h *Handler) Add(f HandlerFunc) uint32 {
	h.m.Lock()
	defer h.m.Unlock()
	eventId := lwutilty.HashCode(lwutilty.UUID())
	item := &HandlerItem{eventId: eventId, f: f}
	h.ls.PushBack(item)
	return eventId
}

func (h *Handler) Get(eventId uint32) *HandlerItem {
	h.m.Lock()
	defer h.m.Unlock()

	var ev *HandlerItem
	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev = e.Value.(*HandlerItem)
		if ev.eventId == eventId {
			return ev
		}
	}
	return ev
}

func (h *Handler) Remove(eventId uint32) {
	h.m.Lock()
	defer h.m.Unlock()
	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*HandlerItem)
		if ev.eventId == eventId {
			h.ls.Remove(e)
			break
		}
	}
}
