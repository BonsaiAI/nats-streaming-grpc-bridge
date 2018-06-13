// Copyright (c) 2009 The Go Authors. All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
// 
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
// * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package units

import (
	"errors"
	"fmt"
	"strings"
)

var (
	siUnits = []string{"", "K", "M", "G", "T", "P", "E"}
)

func ToString(n int64, scale int64, suffix, baseSuffix string) string {
	mn := len(siUnits)
	out := make([]string, mn)
	for i, m := range siUnits {
		if n%scale != 0 || i == 0 && n == 0 {
			s := suffix
			if i == 0 {
				s = baseSuffix
			}
			out[mn-1-i] = fmt.Sprintf("%d%s%s", n%scale, m, s)
		}
		n /= scale
		if n == 0 {
			break
		}
	}
	return strings.Join(out, "")
}

// Below code ripped straight from http://golang.org/src/pkg/time/format.go?s=33392:33438#L1123
var errLeadingInt = errors.New("units: bad [0-9]*") // never printed

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x int64, rem string, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x >= (1<<63-10)/10 {
			// overflow
			return 0, "", errLeadingInt
		}
		x = x*10 + int64(c) - '0'
	}
	return x, s[i:], nil
}

func ParseUnit(s string, unitMap map[string]float64) (int64, error) {
	// [-+]?([0-9]*(\.[0-9]*)?[a-z]+)+
	orig := s
	f := float64(0)
	neg := false

	// Consume [-+]?
	if s != "" {
		c := s[0]
		if c == '-' || c == '+' {
			neg = c == '-'
			s = s[1:]
		}
	}
	// Special case: if all that is left is "0", this is zero.
	if s == "0" {
		return 0, nil
	}
	if s == "" {
		return 0, errors.New("units: invalid " + orig)
	}
	for s != "" {
		g := float64(0) // this element of the sequence

		var x int64
		var err error

		// The next character must be [0-9.]
		if !(s[0] == '.' || ('0' <= s[0] && s[0] <= '9')) {
			return 0, errors.New("units: invalid " + orig)
		}
		// Consume [0-9]*
		pl := len(s)
		x, s, err = leadingInt(s)
		if err != nil {
			return 0, errors.New("units: invalid " + orig)
		}
		g = float64(x)
		pre := pl != len(s) // whether we consumed anything before a period

		// Consume (\.[0-9]*)?
		post := false
		if s != "" && s[0] == '.' {
			s = s[1:]
			pl := len(s)
			x, s, err = leadingInt(s)
			if err != nil {
				return 0, errors.New("units: invalid " + orig)
			}
			scale := 1.0
			for n := pl - len(s); n > 0; n-- {
				scale *= 10
			}
			g += float64(x) / scale
			post = pl != len(s)
		}
		if !pre && !post {
			// no digits (e.g. ".s" or "-.s")
			return 0, errors.New("units: invalid " + orig)
		}

		// Consume unit.
		i := 0
		for ; i < len(s); i++ {
			c := s[i]
			if c == '.' || ('0' <= c && c <= '9') {
				break
			}
		}
		u := s[:i]
		s = s[i:]
		unit, ok := unitMap[u]
		if !ok {
			return 0, errors.New("units: unknown unit " + u + " in " + orig)
		}

		f += g * unit
	}

	if neg {
		f = -f
	}
	if f < float64(-1<<63) || f > float64(1<<63-1) {
		return 0, errors.New("units: overflow parsing unit")
	}
	return int64(f), nil
}
