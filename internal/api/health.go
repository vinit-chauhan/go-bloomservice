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

package api

import "github.com/gofiber/fiber/v2"

func HealthHandler(c *fiber.Ctx) error {
	// This is a simple health check endpoint that returns a 200 OK status
	// with a message indicating that the service is running.
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Bloom service is running",
	})
}

func HealthRouter(router fiber.Router) {
	router.Get("/health", HealthHandler)
}
