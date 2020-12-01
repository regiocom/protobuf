package typedeclimport

import subpkg "github.com/regiocom/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
