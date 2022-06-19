/*================================================================
*   Copyright (C) 2022 * Ltd. All rights reserved.
*
*   File name   : utils.go
*   Author      : rogone
*   Email       : rogone@163.com
*   Created date: 2022-06-18 23:38:30
*   Description :
*
*===============================================================*/
package jt808

import (
	"fmt"
	"io"
	"unicode"
)

func bytes2bcd(s []byte) ([]byte, error) {
	if len(s)%2 == 1 {
		return nil, fmt.Errorf("need 2*n chars")
	}

	ret := make([]byte, len(s)/2)
	for i := range ret {
		hi, lo := s[2*i], s[2*i+1]
		if !unicode.IsDigit(rune(hi)) || !unicode.IsDigit(rune(lo)) {
			return nil, fmt.Errorf("%s is not a number", s)
		}
		ret[i] = (hi-'0')<<4 | ((lo - '0') & 0xff)
	}
	return ret, nil
}

func bcd2bytes(bcd []byte) ([]byte, error) {
	ret := make([]byte, len(bcd)*2)

	for i, v := range bcd {
		ret[2*i] = (v >> 4) + '0'
		ret[2*i+1] = (v & 0xff) + '0'
	}
	return ret, nil
}

func readLen(r io.Reader, buf []byte, size int) error {
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	if n != size {
		return fmt.Errorf("want:%d, read:%d", size, n)
	}
	return nil
}

func writeLen(w io.Writer, buf []byte, size int) error {
	n, err := w.Write(buf)
	if err != nil {
		return err
	}

	if n != size {
		return fmt.Errorf("want:%d, write:%d", size, n)
	}
	return nil
}
