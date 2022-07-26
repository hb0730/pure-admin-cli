package template

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"pure-admin-cli/constants"
)

type Command struct {
	ProjectName string
	Template    constants.GitTemplate
	Tag         string
	LocalPath   string
	Force       bool
}

func Run(cmd *cobra.Command, cmdModel Command) error {
	cmd.Println("Start clone pure template ...")
	// 规范化路径
	localPath := filepath.Clean(cmdModel.LocalPath)
	localPath = filepath.Join(localPath, cmdModel.ProjectName)
	isSuccess, err := createDir(cmd, localPath, cmdModel.Force)
	if err != nil {
		return err
	}
	if !isSuccess {
		return nil
	}
	err = clone(cmd, localPath, cmdModel.Tag, cmdModel.Template)
	if err == nil {
		cmd.Println(fmt.Sprintf("Successfully created project %s", cmdModel.ProjectName))
	}
	return err
}
func createDir(cmd *cobra.Command, localPath string, isForce bool) (bool, error) {
	if len(localPath) == 0 {
		return false, errors.New("local path missing")
	}
	info, err := os.Stat(localPath)
	if err != nil {
		//使用os.IsNotExist()判断为true,说明文件或文件夹不存在
		if os.IsNotExist(err) {
			return true, nil
		}
	}
	if !info.IsDir() {
		return false, errors.New("local path is not dir")
	}
	if isForce {
		cmd.Println("Cleaning dir ...")
		err := os.RemoveAll(localPath)
		if err != nil {
			//fmt.Printf("clean local dir error")
			return false, errors.New("clean local dir error")
		}
		cmd.Println("Clean dir success ...")
	}
	return true, nil
}
func clone(cmd *cobra.Command, localPath, version string, template constants.GitTemplate) (err error) {
	cmd.Println(fmt.Sprintf("Clone template: %s  version: %s  branch: %s local path: %s",
		template.DownloadUrl, version, template.Branch, localPath))
	// 是克隆branch还是tag
	gitOptions := &git.CloneOptions{
		URL:          template.DownloadUrl,
		Progress:     cmd.OutOrStdout(),
		SingleBranch: true,
	}
	if template.IsBranch {
		gitOptions.ReferenceName = plumbing.NewBranchReferenceName(template.Branch)
	} else {
		if constants.DefaultVersion == version {
			gitOptions.ReferenceName = plumbing.NewBranchReferenceName(template.Branch)
		} else {
			gitOptions.ReferenceName = plumbing.NewTagReferenceName(version)
		}
	}
	_, err = git.PlainClone(localPath, false, gitOptions)
	return
}
