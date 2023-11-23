package model

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// uint を marshal / unmarchal する
//
// https://gqlgen.com/reference/scalars/
// https://qiita.com/3104k/items/caf17633d4926aee8a84

func MarshalUint64(t uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.FormatUint(t, 10))
		if err != nil {
			return
		}
	})
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch t := v.(type) {
	case string:
		return strconv.ParseUint(t, 10, 64)
	case int:
		return uint64(t), nil
	case int64:
		return uint64(t), nil
	case json.Number:
		i, err := t.Int64()
		return uint64(i), err
	case float64:
		return uint64(t), nil
	}

	return 0, fmt.Errorf("unable to unmarshal uint64: %#v %T", v, v)
}
