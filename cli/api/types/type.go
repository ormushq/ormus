package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Request struct {
	Path                  string
	Method                string
	AuthorizationRequired bool
	URLParams             []any
	Header                map[string]string
	QueryParams           map[string]string
	Body                  map[string]string
}

func (r Request) GetURL(baseURL string) string {
	path := fmt.Sprintf(r.Path, r.URLParams...)

	url := fmt.Sprintf("%s/%s", baseURL, path)
	if len(r.QueryParams) > 0 {
		qs := make([]string, 0)
		for k, v := range r.QueryParams {
			qs = append(qs, fmt.Sprintf("%s=%s", k, v))
		}
		url = fmt.Sprintf("%s?%s", url, strings.Join(qs, "&"))
	}

	return url
}

func (r Request) GetBody() (*bytes.Buffer, error) {
	reqBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(reqBody), nil
}
