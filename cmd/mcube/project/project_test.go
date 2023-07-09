package project_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yoas0bi/micro-toolkit/cmd/mcube/project"
)

func TestSaveFile(t *testing.T) {
	should := assert.New(t)

	p := project.Project{
		PKG:  "test",
		Name: "test",
	}

	err := p.SaveFile(project.PROJECT_SETTING_FILE_PATH)
	should.NoError(err)
}
