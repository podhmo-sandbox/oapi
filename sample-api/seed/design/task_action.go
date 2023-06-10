package design

import "github.com/podhmo/gos/openapigen"

var (
	Task = openapigen.Define("Task", b.Object(
		b.Field("name", b.Reference(Name)),
		b.Field("done", b.Bool()),
		b.Field("createdAt", b.ReferenceByName("DateTime")),
	))
)

var (
	ListTask = b.Action("ListTask",
		b.Input(b.Param("sort", b.String().Enum([]string{"createdAt", "-createdAt"})).AsQuery()),
		b.Output(b.Array(Task)),
	)
)
