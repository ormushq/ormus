package event

type ID string
type Properties map[string]interface{}
type CustomData map[string]string // TODO: there are multiple places where segment in javascript allow user to post additional custom properties like here but as is GO a typed programming language I don't know a way to handle it, but to add a dictionary field like this
