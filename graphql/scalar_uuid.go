package graphql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUIDScalar(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.String()))
	})
}

func UnmarshalUUIDScalar(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.Parse(v)
	case *string:
		if v == nil {
			return uuid.Nil, nil
		}

		return uuid.Parse(*v)
	default:
		return uuid.Nil, fmt.Errorf("%T is not a uuid", v)
	}
}
