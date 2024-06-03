package global

{{- if .HasGlobal }}

import "xtt/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}