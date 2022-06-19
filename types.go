/*================================================================
*   Copyright (C) 2022 * Ltd. All rights reserved.
*
*   File name   : types.go
*   Author      : rogone
*   Email       : rogone@163.com
*   Created date: 2022-06-18 20:48:40
*   Description :
*
*===============================================================*/
package jt808

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Int interface {
	~uint8 | ~uint16 | ~uint32
	Read(io.Reader) error
	Write(io.Writer) error
}

type V interface {
	~uint8 | ~uint16 | ~uint32
}

//type Reader interface {
//~Int | ~LV | ~TLV
//Read(r io.Reader) error
//}

//type Writer interface {
//~Int | ~LV | ~TLV
//Write(w io.Writer) error
//}

type BYTE uint8

func (b *BYTE) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, b)
}

func (b BYTE) Write(r io.Writer) error {
	return binary.Write(r, binary.BigEndian, b)
}

type WORD uint16

func (w *WORD) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, w)
}

func (w WORD) Write(r io.Writer) error {
	return binary.Write(r, binary.BigEndian, w)
}

type DWORD uint32

func (d *DWORD) Read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, d)
}

func (d DWORD) Write(r io.Writer) error {
	return binary.Write(r, binary.BigEndian, d)
}

type LByte[L Int] struct {
	l L
	v []byte
}

func (lv *LByte[L]) Read(r io.Reader) error {
	err := lv.l.Read(r)
	if err != nil {
		return err
	}

	lv.v = make([]byte, int(lv.l))
	err = readLen(r, lv.v, int(lv.l))
	if err != nil {
		return err
	}
	return nil
}

func (lv *LByte[L]) Write(w io.Writer) error {
	err := lv.l.Write(w)
	if err != nil {
		return err
	}

	err = writeLen(w, lv.v, int(lv.l))
	if err != nil {
		return err
	}
	return nil
}

type LV[L, V Int] struct {
	l L
	v []V
}

func (lv *LV[L, V]) Read(r io.Reader) error {
	err := lv.l.Read(r)
	if err != nil {
		return err
	}

	lv.v = make([]V, int(lv.l))
	for i := range lv.v {
		err = lv.v[i].Read(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (lv *LV[L, V]) Write(w io.Writer) error {
	err := lv.l.Write(w)
	if err != nil {
		return err
	}

	for i := range lv.v {
		err = lv.v[i].Write(w)
		if err != nil {
			return err
		}
	}
	return nil
}

type TLV[T, L, V Int] struct {
	t  T
	lv LV[L, V]
}

func (tlv *TLV[T, L, V]) Read(r io.Reader) error {
	err := tlv.t.Read(r)
	if err != nil {
		return err
	}

	err = tlv.lv.Read(r)
	if err != nil {
		return err
	}
	return nil
}

func (tlv *TLV[T, L, V]) Write(w io.Writer) error {
	err := tlv.t.Write(w)
	if err != nil {
		return err
	}

	err = tlv.lv.Write(w)
	if err != nil {
		return err
	}
	return nil
}

type FixLenBytes struct {
	l int
	v []byte
}

func (flb *FixLenBytes) SetLen(l int) {
	flb.l = l
}

func (flb *FixLenBytes) Set(buf []byte) {
	flb.l = len(buf)
	flb.v = buf
}

// SetLen before Read
func (flb *FixLenBytes) Read(r io.Reader) error {
	if flb.v == nil {
		flb.v = make([]byte, flb.l)
	}

	n, err := r.Read(flb.v)
	if err != nil {
		return err
	}

	if n != flb.l {
		return fmt.Errorf("want:%d, read:%d", flb.l, n)
	}
	return nil
}

func (flb *FixLenBytes) Write(w io.Writer) error {
	n, err := w.Write(flb.v)
	if err != nil {
		return err
	}

	if n != flb.l {
		return fmt.Errorf("want:%d, write:%d", flb.l, n)
	}
	return nil
}
