package testing

import (
	"github.com/johnhoman/controller-tools/testing/internal/testing/crud"
	"testing"

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
