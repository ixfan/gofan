package generate

import (
	"io/ioutil"
	"strings"
)

func GoModule() string {
	files, _ := ioutil.ReadFile("./go.mod")
	lines := strings.Split(string(files), "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(strings.ReplaceAll(lines[0], "module", ""))
	} else {
		panic(any("未获取到go mod文件内容"))
	}
}
