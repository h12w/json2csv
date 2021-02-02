package main

import (
	"sort"
	"strings"
)

type M = map[string]interface{}
type Path []string

func walk(path Path, parent interface{}, fn func(Path, interface{})) {
	fn(path, parent)
	m, ok := parent.(M)
	if !ok {
		return
	}

	keys := make([]string, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		walk(append(path, key), m[key], fn)
	}
}

func (p Path) Value(m M) interface{} {
	if len(p) == 0 {
		return m
	}

	parents, child := p[:len(p)-1], p[len(p)-1]
	var ok bool
	for _, key := range parents {
		m, ok = m[key].(M)
		if !ok {
			return nil
		}
	}
	return m[child]
}

func (p Path) String() string {
	return strings.Join(p, ".")
}
