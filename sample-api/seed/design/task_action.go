package design

import (
	"fmt"

	"github.com/podhmo/gos/openapigen"
)

// helper function for create Pagination[T]
func pagination(b *openapigen.Builder, typ openapigen.Type) *openapigen.Object {
	return b.Object(
		b.Field("total", b.Int()),
		b.Field("more", b.Bool()),
		b.Field("cursor", b.String()).Required(false),
		b.Field("data", typ),
	).
		Title(fmt.Sprintf("Pagination[%s]", typ))
}

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
		b.Output(pagination(b, b.Array(Task))),
	)
)
