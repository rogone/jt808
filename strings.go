/*================================================================
*   Copyright (C) 2022 * Ltd. All rights reserved.
*
*   File name   : strings.go
*   Author      : rogone
*   Email       : rogone@163.com
*   Created date: 2022-06-19 10:42:20
*   Description :
*
*===============================================================*/
package jt808

import (
	"fmt"
	"io"
	"time"
)

type StringN[L Int] LByte[L]

func (str *StringN[L]) Set(s string) {
	str.l = L(len(s))
	str.v = []byte(s)
}

func (str *StringN[L]) String() string {
	return string(str.v)
}

type ByteN[L Int] LByte[L]

func (str *ByteN[L]) Set(buf []byte) {
	str.l = L(len(buf))
	str.v = buf
}

func (str *ByteN[L]) Get() []byte {
	return str.v
}

type BCD struct {
	l int
	v []byte
	p Padder
	c byte
}

func (bcd *BCD) SetLen(l int) {
	bcd.l = l
}

func (bcd *BCD) SetPadder(p Padder) {
	bcd.p = p
}

func (bcd *BCD) SetPadderWith(p Padder, c byte) {
	bcd.p = p
	bcd.c = c
}

func (bcd *BCD) Set(buf []byte) error {
	var err error
	bcd.v, err = bytes2bcd(bcd.p.Pad(buf, bcd.c, bcd.l*2-len(buf)))
	if err != nil {
		return err
	}
	return nil
}

func (bcd *BCD) Get() []byte {
	ret, _ := bcd2bytes(bcd.v)
	return ret
}

func (bcd *BCD) Read(r io.Reader) error {
	if bcd.v == nil {
		bcd.v = make([]byte, bcd.l)
	}
	_, err := r.Read(bcd.v)
	if err != nil {
		return err
	}
	return nil
}

func (bcd *BCD) Write(w io.Writer) error {
	_, err := w.Write(bcd.v)
	if err != nil {
		return err
	}

	return nil
}

type BCDYYMMDDHHMMSS struct {
	t time.Time
}

func (bcd6 *BCDYYMMDDHHMMSS) Set(t time.Time) {
	bcd6.t = t
}

func (bcd6 *BCDYYMMDDHHMMSS) Get() time.Time {
	return bcd6.t
}

func (bcd6 *BCDYYMMDDHHMMSS) Read(r io.Reader) error {
	buf := make([]byte, 0, 6)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	if n != 6 {
		return fmt.Errorf("want:6, read:%d", n)
	}

	buf, _ = bcd2bytes(buf)
	bcd6.t, err = time.ParseInLocation("060102150405", string(buf), time.Local)
	if err != nil {
		return err
	}
	return nil
}

func (bcd6 *BCDYYMMDDHHMMSS) Write(w io.Writer) error {
	s := bcd6.t.Format("060102150405")
	buf, _ := bytes2bcd([]byte(s))
	n, err := w.Write(buf)
	if err != nil {
		return err
	}

	if n != 6 {
		return fmt.Errorf("want:6, write:%d", n)
	}
	return nil
}
