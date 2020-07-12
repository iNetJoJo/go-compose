package goCompose

import (
	"bytes"
	"html/template"
)

type composeVersion string

const V1 = "1"
const V2 = "2"
const V3 = "3"

type Composer struct {
	Version  composeVersion `json:"version,omitempty"`
	Services []*Service `json:"services,omitempty"`
	Volumes  []*Volume `json:"volumes,omitempty"`
	Networks []*Network `json:"networks,omitempty"`
}

func NewComposer(version composeVersion) *Composer {
	return &Composer{
		Version:  version,
		Services: []*Service{},
		Volumes:  []*Volume{},
		Networks: []*Network{},
	}
}

func (c *Composer) Build() ([]byte, error) {
	tmpl, err := template.New("composer").Parse(composeTmpl)
	if err != nil {
		return nil, err
	}

	var writer bytes.Buffer
	err = tmpl.Execute(&writer, c)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}
