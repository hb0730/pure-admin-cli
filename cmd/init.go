package cmd

import (
	"errors"
	"fmt"
	"github.com/gookit/gcli/v3"
	"pure-admin-cli/constants"
	"pure-admin-cli/template"
)

var (
	_answers = &constants.Answers{}
	initCmd  = &gcli.Command{
		Name: "init",
		Desc: "create a new project",
		Func: func(c *gcli.Command, args []string) error {
			arg := c.Arg("name")
			if arg == nil {
				return errors.New("project name is required")
			}
			_answers.Name = arg.String()
			return template.Run(c, _answers.Convert())
		},
	}
)

func init() {
	initCmd.StrVar(&_answers.Template, &gcli.FlagMeta{
		Name:   "template",
		Shorts: []string{"t"},
		Desc:   "pure template",
		DefVal: constants.TemplateArray[0],
		Validator: func(val string) error {
			if len(val) == 0 {
				return errors.New("template is required")
			}
			if constants.Contains(constants.TemplateArray, val) == -1 {
				return errors.New(fmt.Sprintf("template  supports: %s", constants.TemplateArray))
			}
			return nil
		},
	})
	initCmd.StrVar(&_answers.Repo, &gcli.FlagMeta{
		Name:   "repo",
		Shorts: []string{"r"},
		Desc:   "pure template repo",
		DefVal: constants.TemplateRepoArray[0],
		Validator: func(val string) error {
			if len(val) == 0 {
				return errors.New("template repo is required")
			}
			if constants.Contains(constants.TemplateRepoArray, val) == -1 {
				return errors.New(fmt.Sprintf("template repo supports: %s", constants.TemplateRepoArray))
			}
			return nil
		},
	})
	initCmd.StrVar(&_answers.Version, &gcli.FlagMeta{
		Name:   "version",
		Shorts: []string{"v"},
		Desc:   "pure template version",
		DefVal: constants.DefaultVersion,
		Validator: func(val string) error {
			if len(val) == 0 {
				return errors.New("template version is required")
			}
			return nil
		},
	})
	initCmd.StrVar(&_answers.LocalPath, &gcli.FlagMeta{
		Name:   "path",
		Shorts: []string{"p"},
		Desc:   "project local path",
		DefVal: constants.DefaultLocalPath,
		Validator: func(val string) error {
			if len(val) == 0 {
				return errors.New("project local path is required")
			}
			return nil
		},
	})
	initCmd.BoolVar(&_answers.Force, &gcli.FlagMeta{
		Name:   "force",
		Shorts: []string{"f"},
		Desc:   "Overwrite target directory",
		DefVal: false,
		Validator: func(val string) error {
			if len(val) == 0 {
				return errors.New("force is required")
			}
			return nil
		},
	})
	//initCmd.Str("force", "f", "false", "Overwrite target directory")
	initCmd.BindArg(&gcli.Argument{
		Name:     "name",
		Desc:     "project name",
		Required: true,
		Validator: func(val interface{}) (interface{}, error) {
			if val == nil {
				return nil, errors.New("project name is required")
			}
			if len(val.(string)) == 0 {
				return nil, errors.New("project name is required")
			}
			return val, nil
		},
	})
}
