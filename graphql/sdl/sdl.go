package sdl

import (
	"io/ioutil"
	"os"
	"path"
)

// todo: 后续考虑由外面传入
func GetSchemaString() (content string, err error) {
	wd, _ := os.Getwd()
	file, err := os.Open(path.Join(wd, "graphql", "sdl", "graphql.graphql"))
	if err != nil {
		return
	}

	contentByte, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	content = string(contentByte)

	return
}
