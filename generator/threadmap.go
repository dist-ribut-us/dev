package generator

import (
	"bytes"
	"text/template"
)

type ThreadsafeMap struct {
	Key    string
	Val    string
	Name   string
	Export bool
}

var theadsafeMapTemplate = template.Must(template.New("threadsafe-map").Parse(`
type {{.Name}} struct {
	Map map[{{.Key}}]{{.Val}}
	sync.RWMutex
}

func {{if .Export}}New{{else}}new{{end}}{{.Name}}() *{{.Name}} {
	return &{{.Name}}{
		Map: make(map[{{.Key}}]{{.Val}}),
	}
}

func (t *{{.Name}}) {{if .Export}}Get{{else}}get{{end}}(key {{.Key}}) ({{.Val}}, bool) {
	t.RLock()
	k, b := t.Map[key]
	t.RUnlock()
	return k, b
}

func (t *{{.Name}}) {{if .Export}}Set{{else}}set{{end}}(key {{.Key}}, val {{.Val}}) {
	t.Lock()
	t.Map[key] = val
	t.Unlock()
}

func (t *{{.Name}}) {{if .Export}}Delete{{else}}delete{{end}}(keys ...{{.Key}}) {
	t.Lock()
	for _, key := range keys {
		delete(t.Map, key)
	}
	t.Unlock()
}
`))

func (t *ThreadsafeMap) Generate() string {
	buf := &bytes.Buffer{}
	err := theadsafeMapTemplate.Execute(buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
