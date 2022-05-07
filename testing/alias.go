package testing

import (
	"testing"

	"github.com/johnhoman/controller-tools/internal/testing/crud"
	"github.com/johnhoman/controller-tools/testing/suite"
)

type (
	Suite   = suite.Suite
	EnvTest = suite.EnvTest
	T       = testing.T
)

var RunSuite = suite.Run
var Create = crud.Create

var (
	_ = RunSuite
)
