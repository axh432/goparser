package model

import "fmt"

type mPackage struct {
	Name string
}

func (p *mPackage) toString() string {
	return fmt.Sprintf("package %s", p.Name)
}