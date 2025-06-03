package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vinit-chauhan/go-bloomservice/internal/bloom"
	"github.com/vinit-chauhan/go-bloomservice/internal/server"
)

type OperationType int

const (
	OpInsert OperationType = iota
	OpLookup
	OpReset
)

type TestStep struct {
	Name               string
	Operation          OperationType
	Item               string
	ExpectedStatusCode int
}

type TestScenario struct {
	Name  string
	Steps []TestStep
}

const TestAPIVersion = "v1"

func getEndpoint(op OperationType) string {
	switch op {
	case OpInsert:
		return "/api/" + TestAPIVersion + "/add"
	case OpLookup:
		return "/api/" + TestAPIVersion + "/exists"
	case OpReset:
		return "/api/" + TestAPIVersion + "/reset"
	default:
		return ""
	}
}

func TestBloomE2EGrouped(t *testing.T) {
	scenarios := []TestScenario{
		{
			Name: "Insert, Lookup, Reset, Lookup again",
			Steps: []TestStep{
				{"Insert hello", OpInsert, "hello", http.StatusCreated},
				{"Lookup hello", OpLookup, "hello", http.StatusOK},
				{"Reset filter", OpReset, "", http.StatusOK},
				{"Lookup hello after reset", OpLookup, "hello", http.StatusNotFound},
			},
		},
		{
			Name: "Multiple inserts and lookups",
			Steps: []TestStep{
				{"Insert foo", OpInsert, "foo", http.StatusCreated},
				{"Insert bar", OpInsert, "bar", http.StatusCreated},
				{"Lookup foo", OpLookup, "foo", http.StatusOK},
				{"Lookup bar", OpLookup, "bar", http.StatusOK},
				{"Lookup baz", OpLookup, "baz", http.StatusNotFound},
			},
		},
	}

	for _, scenario := range scenarios {
		scenario := scenario // capture
		t.Run(scenario.Name, func(t *testing.T) {
			bloom.Init(10000, 0.01) // fresh filter per scenario
			app := server.StartServer()

			for _, step := range scenario.Steps {
				t.Logf("Running step: %s", step.Name)

				var reqBody io.Reader
				if step.Operation != OpReset {
					body, _ := json.Marshal(map[string]string{"item": step.Item})
					reqBody = bytes.NewReader(body)
				}

				var req *http.Request
				switch step.Operation {
				case OpInsert:
					req = httptest.NewRequest("POST", getEndpoint(OpInsert), reqBody)
				case OpLookup:
					req = httptest.NewRequest("POST", getEndpoint(OpLookup), reqBody)
				case OpReset:
					req = httptest.NewRequest("DELETE", getEndpoint(OpReset), nil)
				default:
					t.Fatalf("Invalid operation: %v", step.Operation)
				}
				req.Header.Set("Content-Type", "application/json")

				resp, err := app.Test(req)
				if err != nil {
					t.Fatalf("Failed request: %v", err)
				}
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)
				t.Logf("Response: %s", body)

				if resp.StatusCode != step.ExpectedStatusCode {
					t.Errorf("Step %q: expected %d, got %d", step.Name, step.ExpectedStatusCode, resp.StatusCode)
				}
			}
		})
	}
}
