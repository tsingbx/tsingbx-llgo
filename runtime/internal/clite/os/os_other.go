//go:build !darwin

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

package os

import (
	_ "unsafe"

	c "github.com/goplus/llgo/runtime/internal/clite"
)

const (
	LLGoFiles   = "_os/os.c"
	LLGoPackage = "link"
)

const (
	PATH_MAX = 4096
)

type (
	ModeT uint32
	UidT  uint32
	GidT  uint32
	OffT  int64
	DevT  uint64
)

//go:linkname Clearenv C.clearenv
func Clearenv() c.Int
