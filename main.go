package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dec := json.NewDecoder(os.Stdin)
	var header Header
	w := csv.NewWriter(os.Stdout)
	for {
		m := make(map[string]interface{})
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if header == nil {
			header = extractHeader(m)
			if err := w.Write(header.Strings()); err != nil {
				return err
			}
		}
		if err := w.Write(header.Values(m)); err != nil {
			return err
		}
	}
	return nil
}
