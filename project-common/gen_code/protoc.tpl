message {{.MsgName}} { {{range $index,$field := .FieldInfos}}
    {{$field.Type}} {{$field.Field}} = {{Add $index 1}}; {{end}}
}