package types

type Request struct {
	Path                  string
	Method                string
	AuthorizationRequired bool
	UrlParams             []string
	Header                map[string]string
	QueryParams           map[string]string
	Body                  map[string]string
}
