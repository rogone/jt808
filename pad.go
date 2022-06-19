/*================================================================
*   Copyright (C) 2022 * Ltd. All rights reserved.
*
*   File name   : pad.go
*   Author      : rogone
*   Email       : rogone@163.com
*   Created date: 2022-06-19 13:18:55
*   Description :
*
*===============================================================*/
package jt808

type Padder interface {
	Pad(buf []byte, c byte, n int) []byte
}

var (
	DummyPadder dummyPadder
	HeadPadder  headPadder
	TailPadder  tailPadder
)

type dummyPadder struct{}

func (dummyPadder) Pad(buf []byte, c byte, n int) []byte {
	return buf
}

type headPadder struct{}

func (*headPadder) Pad(buf []byte, c byte, n int) []byte {
	if n <= 0 {
		return buf
	}

	ret := make([]byte, 0, len(buf)+n)
	for i := 0; i < n; i++ {
		ret = append(ret, c)
	}
	ret = append(ret, buf...)
	return ret
}

type tailPadder struct{}

func (*tailPadder) Pad(buf []byte, c byte, n int) []byte {
	if n <= 0 {
		return buf
	}

	for i := 0; i < n; i++ {
		buf = append(buf, c)
	}
	return buf
}
