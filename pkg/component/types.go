package component

import "sigs.k8s.io/controller-runtime/pkg/client"

//todo change to crd definition communication by grpc

func Register(name string) {

}

//Workload 表示具体承载组件的工作负载的类型，可能的值： Deployment StatefulSet CloneSet 等等只要实现对应的接口即可
type Workload interface {
	Name() string
	DesiredReplicas() int32
	AvailableReplicas() int32
	Core() client.Object
	Exist() bool
}

//Define 表示一种组件类型的定义用以支持多种部署模式的组件
type Define interface {
	Name() string
	Description() string
	WorkLoad() Workload
	Install() error
	Template() string
	Parameters() map[string]interface{}
	DryRun() error
	Version() string
	Healthy() bool
}
