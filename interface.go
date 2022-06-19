/*================================================================
*   Copyright (C) 2022 * Ltd. All rights reserved.
*
*   File name   : interface.go
*   Author      : rogone
*   Email       : rogone@163.com
*   Created date: 2022-06-19 12:24:12
*   Description :
*
*===============================================================*/
package jt808

import (
	"io"
)

type Reader interface {
	//~Int | ~LV | ~TLV
	Read(r io.Reader) error
}

type Writer interface {
	//~Int | ~LV | ~TLV
	Write(w io.Writer) error
}
