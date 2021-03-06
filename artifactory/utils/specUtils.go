package utils

import (
	"encoding/json"
	"github.com/jfrogdev/jfrog-cli-go/utils/io/fileutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
	"strconv"
)

const (
	WILDCARD SpecType = "wildcard"
	SIMPLE SpecType = "simple"
	AQL SpecType = "aql"
)

type Aql struct {
	ItemsFind map[string]interface{} `json:"items.find"`
	Sort map[string]interface{} `json:"sort"`
	Limit int `json:"limit"`
}

type File struct {
	Pattern     string
	Target      string
	Props       string
	Recursive   string
	Flat        string
	Regexp      string
	Aql         Aql
	Build       string
	IncludeDirs string
}

type SpecFiles struct {
	Files []File
}

func (spec *SpecFiles) Get(index int) *File {
	if index < len(spec.Files) {
		return &spec.Files[index]
	}
	return new(File)
}


func CreateSpecFromFile(specFilePath string) (spec *SpecFiles, err error) {
	spec = new(SpecFiles)
	content, err := fileutils.ReadFile(specFilePath)
	if cliutils.CheckError(err) != nil {
		return
	}

	err = json.Unmarshal(content, spec)
	if cliutils.CheckError(err) != nil {
		return
	}
	return
}

func CreateSpec(pattern, target, props, build string, recursive, flat, regexp, includeDirs bool) (spec *SpecFiles) {
	spec = &SpecFiles{
		Files: []File{
			{
				Pattern:     pattern,
				Target:      target,
				Props:       props,
				Build:       build,
				Recursive:   strconv.FormatBool(recursive),
				Flat:        strconv.FormatBool(flat),
				Regexp:      strconv.FormatBool(regexp),
				IncludeDirs: strconv.FormatBool(includeDirs),
			},
		},
	}
	return spec
}

func (file File) GetSpecType() (specType SpecType) {
	switch {
	case file.Pattern != "" && (IsWildcardPattern(file.Pattern) || file.Build != ""):
		specType = WILDCARD
	case file.Pattern != "":
		specType = SIMPLE
	case file.Aql.ItemsFind != nil :
		specType = AQL
	}
	return specType
}

func (file File) IsIncludeDirs() bool {
	return file.IncludeDirs == "true"
}

type SpecType string