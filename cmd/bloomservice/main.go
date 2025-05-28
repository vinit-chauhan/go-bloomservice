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
	"github.com/vinit-chauhan/go-bloomservice/internal/bloom"
	"github.com/vinit-chauhan/go-bloomservice/internal/server"
)

func main() {
	estimatedKeys, falsePositiveRate := 100_000, 0.01
	bloom.Init(estimatedKeys, falsePositiveRate)

	app := server.StartServer()
	app.Listen(":8080")

	defer func() {
		if err := app.Shutdown(); err != nil {
			panic("Failed to shutdown server: " + err.Error())
		}
		bloom.Filter.Clear()
	}()
}
