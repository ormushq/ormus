package projecthandler_test

import "testing"

func TestHandler_Create(t *testing.T) {
	testCases := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int

		expectedResponse string
		err              bool
	}{
		{
			name: "happy path",
		},
		{
			name: "validation error",
		},
		{
			name: "repo error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			// 2. execution
			// 3. assertion
		})
	}
}
