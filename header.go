package main

import "fmt"

type Header []Path

func extractHeader(m M) Header {
	var ps []Path
	walk(nil, m, func(path Path, v interface{}) {
		if _, isMap := v.(M); isMap {
			return
		}
		ps = append(ps, path)
	})
	return ps
}

func (h Header) Strings() []string {
	ss := make([]string, len(h))
	for i := range h {
		ss[i] = h[i].String()
	}
	return ss
}

func (h Header) Values(m M) []string {
	vs := make([]string, len(h))
	for i, path := range h {
		vs[i] = fmt.Sprint(path.Value(m))
	}
	return vs
}
