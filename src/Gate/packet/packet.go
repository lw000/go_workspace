package lwpacket

import (
	"bytes"
	"encoding/binary"
)

type Packet struct {
	ver int32
	mid int32
	sid int32
	ext int32
	buf *bytes.Buffer
}

func NewPacket(mid, sid int32) *Packet {
	return &Packet{
		ver: 1,
		mid: mid,
		sid: sid,
		ext: 1,
		buf: bytes.NewBuffer(nil),
	}
}

func NewPacketWithData(data []byte) *Packet {
	return &Packet{
		buf: bytes.NewBuffer(data),
	}
}

func (p *Packet) writeHead() error {
	var err error
	err = binary.Write(p.buf, binary.LittleEndian, p.ver)
	if err != nil {
		return err
	}

	err = binary.Write(p.buf, binary.LittleEndian, p.mid)
	if err != nil {
		return err
	}

	err = binary.Write(p.buf, binary.LittleEndian, p.sid)
	if err != nil {
		return err
	}

	err = binary.Write(p.buf, binary.LittleEndian, p.ext)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) readHead() error {
	var err error
	err = binary.Read(p.buf, binary.LittleEndian, &p.ver)
	if err != nil {
		return err
	}

	err = binary.Read(p.buf, binary.LittleEndian, &p.mid)
	if err != nil {
		return err
	}

	err = binary.Read(p.buf, binary.LittleEndian, &p.sid)
	if err != nil {
		return err
	}

	err = binary.Read(p.buf, binary.LittleEndian, &p.ext)
	if err != nil {
		return err
	}

	return nil
}

func (p *Packet) Encode(data []byte) error {
	var err error
	err = p.writeHead()
	if err != nil {
		return err
	}

	if len(data) != 0 {
		var n int
		n, err = p.buf.Write(data)
		if err != nil {
			return err
		}

		if n < 0 {

		}
	}
	return nil
}

func (p *Packet) Decode() ([]byte, error) {
	var err error
	err = p.readHead()
	if err != nil {
		return nil, err
	}

	var data []byte
	var n int
	n, err = p.buf.Read(data)
	if err != nil {
		return nil, err
	}

	if n < 0 {

	}

	return data, nil
}

func (p *Packet) Mid() int32 {
	return p.mid
}

func (p *Packet) Sid() int32 {
	return p.sid
}

func (p *Packet) Ext() int32 {
	return p.ext
}

func (p *Packet) Bytes() []byte {
	return p.buf.Bytes()
}
