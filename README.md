json2csv
========

A CLI tool reading [JSON Lines](https://jsonlines.org/) from stdin, flattening it and writing to stdout in CSV format. The CSV header is extracted in alphabetic order from the first JSON object, where nested field names are flattened with dot syntax.

Usage
-----

```bash
go get h12.io/json2csv
cat xxx.jsonl | json2csv
```
