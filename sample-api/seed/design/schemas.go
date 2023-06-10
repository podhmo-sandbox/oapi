package design

import "github.com/podhmo/gos/openapigen"

var Builder = openapigen.NewBuilder(openapigen.DefaultConfig()) // export
var b = Builder

// Error
var (
	Error = openapigen.Define("Error", b.Object(
		b.Field("message", b.String()),
	)).Doc("default error")
)

// primitives
var (
	Name     = openapigen.Define("Name", b.String().MinLength(1)).Doc("name of something")
	DateTime = openapigen.Define("DateTime", b.String().Format("date-time"))
)
