// Extensions for Protocol Buffers to create more go like structures.
//
// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
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

package fieldpath_test

import (
	"code.google.com/p/gogoprotobuf/fieldpath"
	"code.google.com/p/gogoprotobuf/proto"
	"code.google.com/p/gogoprotobuf/test"
	"fmt"
)

type handler struct {
}

func (this *handler) Int64(v int64) {
	fmt.Printf("%d\n", v)
}

func ExampleUnmarshaler() {
	a := &test.NinOptNative{
		Field4: proto.Int64(1234),
	}
	fp, err := fieldpath.NewInt64Path("test", "NinOptNative", test.ThetestDescription(), "Field4")
	if err != nil {
		panic(err)
	}
	buf, err := proto.Marshal(a)
	if err != nil {
		panic(err)
	}
	h := &handler{}
	u := fieldpath.NewInt64Unmarshaler(fp, h)
	err = u.Unmarshal(buf)
	if err != nil {
		panic(err)
	}
	// Output:
	// 1234
}
