package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Dependence struct {
	Package string `xml:"package"`
	Url     string `xml:"url"`
}

type Dependencies struct {
	Dependence []Dependence `xml:"dependence"`
}

type Project struct {
	Dependencies Dependencies `xml:"dependencies"`
	Workspace    string       `xml:"workspace"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getGoEnvProperty(name string) string {
	cmd := exec.Command("go", "env", "|", "grep", name)
	output, err := cmd.Output()
	check(err)
	return strings.TrimSpace(string(output))
}

func getGlobalPath() (goroot, gopath string) {
	goroot = getGoEnvProperty("GOROOT")
	gopath = getGoEnvProperty("GOPATH")
	return
}

func isPackageExist(path string) bool {
	goroot, gopath := getGlobalPath()
	suffix := "/src/" + path
	_, err := os.Stat(goroot + suffix)
	if os.IsExist(err) {
		return true
	}
	_, err = os.Stat(gopath + suffix)
	if os.IsExist(err) {
		return true
	}
	return false
}

func doWebGoGet(dep Dependence, p Project) {

}

func doLocalGoGet(dep Dependence, p Project) {
	_, gopath := getGlobalPath()
	fmt.Println(gopath)
	i := strings.Index(dep.Url, "/")
	depDir := dep.Url[:i]
	targetDir := []string{
		gopath + "/src/" + depDir,
		gopath + "/pkg/mod/" + depDir,
		//goroot + "/src/" + depDir,
	}
	if isPackageExist(dep.Package) {
	} else {
		for i := 0; i < len(targetDir); i++ {
			os.Remove(targetDir[i])
			if '/' == dep.Url[0] {
				err := os.Symlink(dep.Url, targetDir[i])
				check(err)
			} else {
				err := os.Symlink(p.Workspace+"/"+depDir, targetDir[i])
				check(err)
			}
		}
	}
}

func doGoGet(dep Dependence, p Project) {
	if "http" == dep.Url[:4] {
		doWebGoGet(dep, p)
	} else {
		doLocalGoGet(dep, p)
	}
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println("当前执行路径：" + dir)
	p := Project{}
	p.Workspace = dir

	gdmFilename := p.Workspace + "/gdm.xml"
	data, err := os.ReadFile(gdmFilename)
	check(err)

	// xml解析
	err = xml.Unmarshal(data, &p)
	check(err)

	deps := p.Dependencies.Dependence
	depsLen := len(deps)
	if depsLen == 0 {
		return
	}
	for i := 0; i < depsLen; i++ {
		doGoGet(deps[i], p)
	}

	fmt.Println("success!")
}
