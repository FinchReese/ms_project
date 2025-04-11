type {{.StructName}} struct{ {{range $index,$field := .FieldInfos}}
    {{$field.Field}}  {{$field.Type}}{{end}}
}