# controller-tools
Helpers for writing kubernetes controllers


## Examples
### Integration Tests

```golang
package something_test

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/johnhoman/controller-tools/manager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var cfg *rest.Config
var testEnv *envtest.Environment

var _ = BeforeSuite(func() {
    testEnv = &envtest.Environment{}
    var err error
    cfg, err = testEnv.Start()
    Expect(err).ToNot(HaveOccurred())
    Expect(cfg).ToNot(BeNil())
})

var _ = AfterSuite(func() {
    Expect(testEnv.Stop()).To(Succeed())
})

var _ = Describe("TestCase", func() {
    var mgr manager.IntegrationTest
    BeforeEach(func() {
        mgr = manager.IntegrationTestBuilder().
            WithScheme(scheme.Scheme).
            Complete(cfg)
        mgr.StartManager()
    })
    AfterEach(func() {
        mgr.StopManager()
    })
    It("Create a configmap", func() {
        cm := &corev1.ConfigMap{}
        cm.SetName("test-configmap-1")

        // mgr is set up with a namespaced client and a context
        mgr.Expect().Create(cm).To(Succeed())

        // To wait for the configmap to become available (e.g. asynchronous)
        cm = &corev1.ConfigMap{}
        mgr.Eventually().Get(types.NamespacedName{Name: "test-configmap-1"}, cm).Should(Succeed())
        Expect(cm.GetName()).To(Equal("test-configmap-1"))
    })
}) 
```