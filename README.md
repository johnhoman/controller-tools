# controller-tools
Helpers for writing kubernetes controllers


## Goals
* Provide boilerplate for Kubernetes tests such as resource retrieval 
* Provide custom matchers to assert on common type status's

## Design

``` mermaid
classDiagram
    Kind <|-- Pod
    Kind <|-- ServiceAccount
    
    class Pod{
    
    }
```

## Examples
### Integration Tests

```golang
package something_test

import (
    . "github.com/johnhoman/controller-tools"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


type MySuite struct {
    EnvTest
}

func (suite *MySuite) TestSomething() {
    suite.Nil(Create(suite.GetManager()))
}

```