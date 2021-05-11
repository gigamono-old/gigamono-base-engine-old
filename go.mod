module github.com/gigamono/gigamono-document-engine

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/gigamono/gigamono v0.0.0-20210505114150-59066b8b3792
	github.com/gigamono/gigamono-workflow-engine v0.0.0-20210429165056-56fff37e8f30
	github.com/gin-gonic/gin v1.7.1
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/soheilhy/cmux v0.1.4
	github.com/vektah/gqlparser/v2 v2.2.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/grpc v1.34.0
)

replace github.com/gigamono/gigamono v0.0.0-20210505114150-59066b8b3792 => ../gigamono
