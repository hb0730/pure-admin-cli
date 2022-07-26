package cmd

import (
	"errors"
	"github.com/gookit/gcli/v3"
	"github.com/manifoldco/promptui"
	"pure-admin-cli/constants"
	"pure-admin-cli/template"
	"strconv"
)

var (
	NewCmd = &gcli.Command{
		Name: "new",
		Desc: "create a new project",
		Func: func(c *gcli.Command, args []string) error {
			return executeShell(c)
		},
	}
)

func executeShell(cmd *gcli.Command) error {
	cmdModel := template.Command{}
	projectNamePrompt := promptui.Prompt{
		Label: "project name",
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New("project name missing")
			}
			return nil
		},
	}
	projectName, err := projectNamePrompt.Run()
	if err != nil {
		return err
	}
	cmdModel.ProjectName = projectName
	templatePrompt := promptui.Select{
		Label: "template",
		Items: constants.TemplateArray,
	}
	_, projectTemplate, err := templatePrompt.Run()
	if err != nil {
		return err
	}
	templateRepoPrompt := promptui.Select{
		Label: "template repo",
		Items: constants.TemplateRepoArray,
	}
	repoIndex, _, err := templateRepoPrompt.Run()
	if err != nil {
		return err
	}
	cmdModel.Template = constants.TemplateRepoMapped[projectTemplate][repoIndex]
	templateVersionPrompt := promptui.Prompt{
		Label:   "template version, default: " + constants.DefaultVersion,
		Default: constants.DefaultVersion,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New("template version missing")
			}
			return nil
		},
	}
	version, err := templateVersionPrompt.Run()
	if err != nil {
		return err
	}
	cmdModel.Tag = version
	forcePrompt := promptui.Select{
		Label: "overwrite target directory if it exists",
		Items: []bool{false, true},
	}
	_, force, err := forcePrompt.Run()
	if err != nil {
		return err
	}
	isForce, err := strconv.ParseBool(force)
	if err != nil {
		return err
	}
	cmdModel.Force = isForce

	pathPrompt := promptui.Prompt{
		Label:   "create project path, default: " + constants.DefaultLocalPath,
		Default: constants.DefaultLocalPath,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New("project local path missing")
			}
			return nil
		},
	}
	localPathStr, err := pathPrompt.Run()
	if err != nil {
		return err
	}
	cmdModel.LocalPath = localPathStr
	return template.Run(cmd, cmdModel)
}
