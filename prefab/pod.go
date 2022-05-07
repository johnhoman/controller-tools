package prefab

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func setSpec(obj *corev1.Pod, spec Map) {
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})

	if b, err := json.Marshal(spec); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal(b, &m2); err != nil {
			panic(err)
		}
	}
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(b, &m1); err != nil {
		panic(err)
	}
	m1["spec"] = m2
	b, err = json.Marshal(m1)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, obj); err != nil {
		panic(err)
	}
}

func SetSpec(spec Map) func(client.Object) {
	return func(obj client.Object) { setSpec(obj.(*corev1.Pod), spec) }
}

func Nginx() func(client.Object) {
	return func(obj client.Object) {
		setSpec(obj.(*corev1.Pod), Map{
			"containers": List{
				Map{
					"name":  "nginx",
					"image": "nginx:latest",
				},
			},
		})
	}
}

func NewPod(opts ...func(client.Object)) client.Object {
	pod := &corev1.Pod{}
	opts = append(opts, RandomName())
	for _, opt := range opts {
		opt(pod)
	}
	return pod
}
