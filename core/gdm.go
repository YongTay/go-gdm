package core

import (
	"encoding/xml"
	"fmt"
	"go-gdm/model"
	"go-gdm/utils"
)

type Gdm struct{}

func (g *Gdm) resolveDependence(dep model.Dependence) {
	name := dep.Package
	version := dep.Version
	if version == "" {
		data, err := utils.ExecCommand("go", "get", name)
		utils.Check(err)
		fmt.Println(string(data))
		return
	}
	data, err := utils.ExecCommand("go", "get", "-u", name+"@"+version)
	utils.Check(err)
	fmt.Println(string(data))
}

func Run() {

	g := &Gdm{}

	p := model.Project{}
	p.Workspace = utils.GetWorkspace()

	utils.ExecCommand("go", "mod", "tidy")

	data := utils.ReadConfigFile(p.Workspace)
	// xml解析
	err := xml.Unmarshal(data, &p)
	utils.Check(err)

	deps := p.Dependencies.Dependence
	depsLen := len(deps)
	if depsLen == 0 {
		return
	}
	for i := 0; i < depsLen; i++ {
		dep := deps[i]
		g.resolveDependence(dep)
	}

	fmt.Println("success!")
}
