package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2 = "<!DOCTYPE html>\n<html lang=\"kr\">\n  <head>\n    <meta charset=\"utf-8\">\n    <title>안녕</title>\n  </head>\n  <body>\n    <h1>{{ .title }}</h1>\n  </body>\n</html>"
var _Assets59a02c018f806788fc3516b59ecab68b0a08b8ef = "<!DOCTYPE html>\n<html lang=\"kr\">\n  <head>\n    <meta charset=\"utf-8\">\n    <title>안녕 gin</title>\n  </head>\n  <body>\n    <form method=\"post\">\n      <input name=\"name\" value=\"{{ .person.Name }}\" />\n      <input name=\"age\" value=\"{{ .person.Age }}\" />\n      <button>save</button>\n      <pre>\n        {{ .person.Name }}, {{ .person.Age }}\n      </pre>\n    </form>\n  </body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"templates"}, "/templates": []string{"index.html", "form1.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1739843409, 1739843409408337732),
		Data:     nil,
	}, "/templates": &assets.File{
		Path:     "/templates",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1739842026, 1739842026766875398),
		Data:     nil,
	}, "/templates/index.html": &assets.File{
		Path:     "/templates/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1739841780, 1739841780718675113),
		Data:     []byte(_Assets9f0ae851af9adb08e3242917d7d52dc270c7bae2),
	}, "/templates/form1.html": &assets.File{
		Path:     "/templates/form1.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1739842598, 1739842598061462929),
		Data:     []byte(_Assets59a02c018f806788fc3516b59ecab68b0a08b8ef),
	}}, "")
