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

package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinit-chauhan/go-bloomservice/internal/api/v1/handlers"
)

func BloomRouter(router fiber.Router) {
	// Add item to the Bloom filter
	// Body:
	// {
	//   "item": "string" // item to add to the Bloom filter
	// }
	router.Post("/add", handlers.AddHandler)

	// Check if an item is in the Bloom filter
	// Body:
	// {
	//   "item": "string" // item to check in the Bloom filter
	// }
	router.Post("/exists", handlers.CheckHandler)

	// Get statistics about the Bloom filter
	// Returns:
	// {
	//   "size": 123456, // size of the Bloom filter in bits
	//   "num_hash_functions": 5, // number of hash functions used
	//   "num_items": 1000, // number of items added to the Bloom filter
	// }
	router.Get("/stats", handlers.StatsHandler)

	// Reset the Bloom filter
	// This endpoint clears the Bloom filter, resetting it to its initial state.
	router.Delete("/reset", handlers.ResetHandler)
}
