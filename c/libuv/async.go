/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package libuv

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// struct uv_async_t
type Async struct {
	Handle
	// On macOS arm64, sizeof uv_async_t is 128 bytes.
	// Handle is 92 bytes, so we need 36 bytes to fill the gap.
	// Maybe reserve more for future use.
	Unused [36]byte
}

// typedef void (*uv_async_cb)(uv_async_t* handle);
// llgo:type C
type AsyncCb func(*Async)

// int uv_async_init(uv_loop_t*, uv_async_t* async, uv_async_cb async_cb);
//
// llgo:link (*Loop).Async C.uv_async_init
func (loop *Loop) Async(a *Async, cb AsyncCb) c.Int {
	return 0
}

// int uv_async_send(uv_async_t* async);
//
// llgo:link (*Async).Send C.uv_async_send
func (a *Async) Send() c.Int {
	return 0
}
