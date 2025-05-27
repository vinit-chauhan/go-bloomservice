// Copyright 2025 Vinit Chauhan

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/vinit-chauhan/go-bloomservice/internal/bloom"
)

func main() {
	filter := bloom.New(100_000, 5)

	filter.Add("hello")
	filter.Add("world")

	fmt.Println(filter.Exists("hello")) // true
	fmt.Println(filter.Exists("world")) // true
	fmt.Println(filter.Exists("foo"))   // false

	filter.Clear()
}
