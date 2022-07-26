package cmd

import (
	"errors"
	"fmt"
	"github.com/gookit/gcli/v3"
	"pure-admin-cli/constants"
	"pure-admin-cli/template"
	"reflect"
)

var (
	projectName     string
	templateStr     string
	templateRepo    string
	templateVersion string
	projectForce    bool
	localPath       string
	InitCmd         = &gcli.Command{
		Name: "init",
		Desc: "create a new project",
		Config: func(c *gcli.Command) {
			c.StrVar(&templateStr, &gcli.FlagMeta{
				Name:   "template",
				Shorts: []string{"t"},
				DefVal: constants.TemplateArray[0],
				Desc:   fmt.Sprintf("pure templates: %s", constants.TemplateArray),
				Validator: func(val string) error {
					if len(val) == 0 {
						return errors.New("template missing")
					}
					if contains(constants.TemplateArray, val) == -1 {
						return errors.New(fmt.Sprintf("template  supports: %s", constants.TemplateArray))
					}
					return nil
				},
			})
			c.StrVar(&templateRepo, &gcli.FlagMeta{
				Name:   "repo",
				Shorts: []string{"r"},
				DefVal: constants.TemplateRepoArray[0],
				Desc:   fmt.Sprintf("pure repo urls: %s", constants.TemplateRepoArray),
				Validator: func(val string) error {
					if len(val) == 0 {
						return errors.New("template repo missing")
					}
					if contains(constants.TemplateRepoArray, val) == -1 {
						return errors.New(fmt.Sprintf("template  repo supports: %s", constants.TemplateRepoArray))
					}
					return nil
				},
			})
			c.StrVar(&templateVersion, &gcli.FlagMeta{
				Name:   "version",
				Shorts: []string{"v"},
				DefVal: constants.DefaultVersion,
				Desc:   "pure template tag",
				Validator: func(val string) error {
					if len(val) == 0 {
						return errors.New("template repo tag missing")
					}
					return nil
				},
			})
			c.BoolVar(&projectForce, &gcli.FlagMeta{
				Name:   "force",
				Shorts: []string{"f"},
				DefVal: constants.DefaultForce,
				Desc:   "overwrite target directory if it exists",
				Validator: func(val string) error {
					if len(val) == 0 {
						return errors.New("force missing")
					}
					return nil
				},
			})
			c.StrVar(&localPath, &gcli.FlagMeta{
				Name:   "path",
				Shorts: []string{"p"},
				DefVal: constants.DefaultLocalPath,
				Desc:   "create project local path",
				Validator: func(val string) error {
					if len(val) == 0 {
						return errors.New("project local path missing")
					}
					return nil
				},
			})
			c.AddArgument(&gcli.Argument{
				Name:     "name",
				Desc:     "project name",
				Required: true,
				Validator: func(val interface{}) (interface{}, error) {
					if val == nil {
						return nil, errors.New("project name missing")
					}
					if len(val.(string)) == 0 {
						return nil, errors.New("project name missing")
					}
					return val, nil
				},
			})
		},
		Func: func(c *gcli.Command, args []string) error {
			projectName = c.Arg("name").String()
			repoIndex := contains(constants.TemplateRepoArray, templateRepo)
			cmdModel := template.Command{
				ProjectName: projectName,
				Template:    constants.TemplateRepoMapped[templateStr][repoIndex],
				Tag:         templateVersion,
				LocalPath:   localPath,
				Force:       projectForce,
			}
			return template.Run(c, cmdModel)
		},
	}
)

func init() {
}
func contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}
