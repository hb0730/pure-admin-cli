package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
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
	initCmd         = &cobra.Command{
		Use:   "init",
		Short: "create a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdModel := template.Command{}
			value, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return errors.New("project name missing")
			}
			cmdModel.ProjectName = value

			templateVar, err := cmd.Flags().GetString("template")
			if err != nil {
				return err
			}
			if len(templateVar) == 0 {
				return errors.New("project template missing")
			}
			if contains(constants.TemplateArray, templateVar) == -1 {
				return errors.New(fmt.Sprintf("project template support: %s", constants.TemplateArray))
			}
			repo, err := cmd.Flags().GetString("repo")
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return errors.New("project template repo missing")
			}
			repoIndex := contains(constants.TemplateRepoArray, repo)
			if repoIndex == -1 {
				return errors.New(fmt.Sprintf("project template repo support: %s", constants.TemplateRepoArray))
			}
			cmdModel.Template = constants.TemplateRepoMapped[templateVar][repoIndex]

			value, err = cmd.Flags().GetString("version")
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return errors.New("project template version missing")
			}
			cmdModel.Tag = value
			force, err := cmd.Flags().GetBool("force")
			if err != nil {
				return err
			}
			cmdModel.Force = force
			value, err = cmd.Flags().GetString("path")
			if err != nil {
				return err
			}
			if len(value) == 0 {
				return errors.New("project local path missing")
			}
			cmdModel.LocalPath = value
			return template.Run(cmd, cmdModel)
		},
	}
)

func init() {
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "project name")
	initCmd.Flags().StringVarP(&templateStr, "template", "t", constants.TemplateArray[0], fmt.Sprintf("template: %s", constants.TemplateArray))
	initCmd.Flags().StringVarP(&templateRepo, "repo", "r", constants.TemplateRepoArray[0], fmt.Sprintf("template repo: %s", constants.TemplateRepoArray))
	initCmd.Flags().StringVarP(&templateVersion, "version", "v", constants.DefaultVersion, "template version")
	initCmd.Flags().BoolVarP(&projectForce, "force", "f", false, "overwrite target directory if it exists")
	initCmd.Flags().StringVarP(&localPath, "path", "p", constants.DefaultLocalPath, "create project local path")
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
