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

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinit-chauhan/go-bloomservice/internal/api"
	api_v1 "github.com/vinit-chauhan/go-bloomservice/internal/api/v1"
)

func StartServer() *fiber.App {
	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
			StrictRouting:         true,
			CaseSensitive:         true,
		},
	)

	app.Route("/", api.HealthRouter)
	app.Route("/api/v1", api_v1.BloomRouter)

	return app
}
