package goCompose

const composeTmpl = `version: '{{.Version}}'

services: {{ range $service := .Services }} 
  {{ $service.Name }}: {{ $length := len $service.Image }} {{ if ne $length 0 }}
    image: {{ $service.Image }}{{ else }}
	build: {{ $service.Build }} {{ end }} {{ $length := len $service.CGroupParent }} {{ if ne $length 0 }}
	cgroup_parent: {{ $service.CGroupParent }} {{ end }} {{ $length := len $service.Command }} {{ if ne $length 0 }} 
	command: {{ range $c := $service.Command }}{{ $c }} {{ end }} {{ end }}
    container_name: {{ $service.ContainerName }} {{ $length := len .DependsOn }} {{ if ne $length 0 }} 
	depends_on: {{ range $s := $service.DependsOn }}
  	  - {{ $s.Name }} {{ end }} {{ end }} {{ $length := len $service.Environment }} {{ if ne $length 0 }}
    environment: {{ range $key, $value := $service.Environment }} 
      - {{ $key }}= "{{ $value }}" {{ end }} {{ end }} {{ $length := len $service.Ports }} {{ if ne $length 0 }} 
    ports: {{ range $p := $service.Ports }} 
      - {{ $p.In }}:{{ $p.Out }} {{ end }} {{ end }} {{ $length := len $service.Volumes }} {{ if ne $length 0 }} 
    volumes: {{ range $v := $service.Volumes }}
      - {{ $v.Volume.Name }}:{{ $v.Path }} {{ end }} {{ end }} {{ $length := len $service.Networks }} {{ if ne $length 0 }}
    networks: {{ range $n := $service.Networks }}
      - {{ $n.Name }} {{ end }} {{ end }}
{{ end }}

{{ $length := len .Volumes }} {{ if ne $length 0 }}
volumes: {{ range $volume := .Volumes }}
  {{ $volume.Name }}: {{ end }}
{{ end }}

{{ $length := len .Networks }} {{ if ne $length 0 }}
networks: {{ range $network := .Networks }}
  {{ $network.Name }}:
    external: {{ $network.Config.External }}
    driver: {{ $network.Config.Driver }} {{ end }} {{ end }}
`

//
//
//services: {{ range $service := .Services }}
//	{{ $service.Name }}:
//		image: {{ $service.Image }}
//		container_name: {{ $service.ContainerName }}
//		environment: {{ range $key, $value := $service.Environment }}
//			- {{ $key }}: {{ $value }} {{ end }}
//		ports: {{ range $p := $service.Ports }}
//			- {{ $p.In }}:{{ $p.Out }} {{ end }}
//{{ end }}`