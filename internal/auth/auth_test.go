package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	const authH = "Authorization"
	const keyPrefix = "ApiKey"
	const testKey = "a-test-key"

	t.Run("get-api-keys should", func(t *testing.T) {
		testcases := []struct {
			desc   string
			input  http.Header
			expErr bool
			expStr string
		}{
			{
				desc: "happy path",
				input: http.Header{
					authH: []string{fmt.Sprintf("%s %s", keyPrefix, testKey)},
				},
				expErr: false,
				expStr: testKey,
			},
			{
				desc: "sad path: incorrect header",
				input: http.Header{
					"nope": []string{fmt.Sprintf("%s %s", keyPrefix, testKey)},
				},
				expErr: true,
			},
		}

		for _, tc := range testcases {
			t.Run(tc.desc, func(t *testing.T) {
				// test
				key, err := GetAPIKey(tc.input)

				// assert
				if tc.expErr {
					if err == nil {
						t.Fatal("expected err and err was nil")
					}
					return
				}

				if err != nil {
					t.Fatalf("expected nil, but got: %+v\n", err)
				}
				if tc.expStr != key {
					t.Fatalf("expected %s, but got %s", tc.expStr, key)
				}
			})
		}
	})
}
