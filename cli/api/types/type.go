package types

type Request struct {
	Path                  string
	Method                string
	AuthorizationRequired bool
	Header                any
	Body                  any
}
