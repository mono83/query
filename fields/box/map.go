package box

import "github.com/mono83/query"

func fieldNameMap(fs []query.Field) map[string]query.Field {
	m := make(map[string]query.Field, len(fs))
	for _, f := range fs {
		m[f.Name()] = f
	}
	return m
}
