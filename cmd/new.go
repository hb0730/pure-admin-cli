package cmd

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/gcli/v3"
	"pure-admin-cli/constants"
	"pure-admin-cli/template"
	"strings"
)

var (
	newCmd = &gcli.Command{
		Name: "new",
		Desc: "create a new project",
		Func: func(c *gcli.Command, args []string) error {
			return executeShell(c)
		},
	}
	qs = []*survey.Question{
		{
			Name:   "name",
			Prompt: &survey.Input{Message: "Project name:"},
			Validate: func(ans interface{}) error {
				if ans == nil {
					return errors.New("value is required")
				}
				if len(strings.TrimSpace(ans.(string))) == 0 {
					return errors.New("value is required")
				}
				return nil
			},
		},
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "Choose pure template:",
				Options: constants.TemplateArray,
				Default: constants.TemplateArray[0],
			},
		},
		{
			Name: "repo",
			Prompt: &survey.Select{
				Message: "Choose pure repo:",
				Options: constants.TemplateRepoArray,
				Default: constants.TemplateRepoArray[0],
			},
		},
		{
			Name: "version",
			Prompt: &survey.Input{
				Message: "Pure version:",
				Default: "last",
			},
			Validate: func(ans interface{}) error {
				if ans == nil {
					return errors.New("value is required")
				}
				if len(strings.TrimSpace(ans.(string))) == 0 {
					return errors.New("value is required")
				}
				return nil
			},
		},
		{
			Name: "localPath",
			Prompt: &survey.Input{
				Message: "Create project local path:",
				Default: constants.DefaultLocalPath,
			},
			Validate: func(ans interface{}) error {
				if ans == nil {
					return errors.New("value is required")
				}
				if len(strings.TrimSpace(ans.(string))) == 0 {
					return errors.New("value is required")
				}
				return nil
			},
		},
		{
			Name: "force",
			Prompt: &survey.Confirm{
				Message: "Overwrite target directory",
				Default: false,
			},
		},
	}
)

func executeShell(cmd *gcli.Command) error {
	answers := constants.Answers{}
	err := survey.Ask(qs, &answers, survey.WithShowCursor(true))
	if err != nil {
		return err
	}
	return template.Run(cmd, answers.Convert())
}
