package controller

import (
	"github.com/jarifibrahim/my-operator/dependant-operator/pkg/controller/dependantservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, dependantservice.Add)
}
