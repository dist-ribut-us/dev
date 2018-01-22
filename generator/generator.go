package generator

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"
	"text/template"
)

type GeneratorFile struct {
	Package string
	Imports []string
	TSMaps  []ThreadsafeMap
}

var generatorTemplate = template.Must(template.New("generator").Parse(`package {{.Package}}

{{if .Imports}}import ({{range $i, $import := .Imports}}
  "{{$import}}"{{end}}
){{end}}
{{range $i, $tsMap := .TSMaps}}{{$tsMap.Generate}}{{end}}
`))

func (g *GeneratorFile) Generate() string {
	g.checkImports()
	buf := &bytes.Buffer{}
	err := generatorTemplate.Execute(buf, g)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *GeneratorFile) checkImports() {
	imps := make(map[string]bool)
	if len(g.TSMaps) > 0 {
		imps["sync"] = true
	}
	for _, imp := range g.Imports {
		imps[imp] = true
	}
	g.Imports = nil
	for imp := range imps {
		g.Imports = append(g.Imports, imp)
	}
	sort.Slice(g.Imports, func(i, j int) bool { return g.Imports[i] < g.Imports[j] })
}

func Read(r io.Reader) (*GeneratorFile, error) {
	dec := json.NewDecoder(r)
	var g GeneratorFile
	err := dec.Decode(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}
