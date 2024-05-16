package main

import (
    "os"
    "text/template"
)

const xmlTemplate = `<colors>
    {{- range . }}
    <color>{{ . }}</color>
    {{- end }}
</colors>
`

func main() {
    colors := []string{"RED", "BLUE", "GREEN", "YELLOW"} // Add your new color here
    tmpl, err := template.New("xml").Parse(xmlTemplate)
    if err != nil {
        panic(err)
    }
    err = tmpl.Execute(os.Stdout, colors)
    if err != nil {
        panic(err)
    }
}
