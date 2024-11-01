/*
Copyright 2024 llmos.ai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/llmos-ai/llmos-gpu-stack/pkg/apis/gpustack.llmos.ai/v1"
	scheme "github.com/llmos-ai/llmos-gpu-stack/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// GPUDevicesGetter has a method to return a GPUDeviceInterface.
// A group's client should implement this interface.
type GPUDevicesGetter interface {
	GPUDevices() GPUDeviceInterface
}

// GPUDeviceInterface has methods to work with GPUDevice resources.
type GPUDeviceInterface interface {
	Create(ctx context.Context, gPUDevice *v1.GPUDevice, opts metav1.CreateOptions) (*v1.GPUDevice, error)
	Update(ctx context.Context, gPUDevice *v1.GPUDevice, opts metav1.UpdateOptions) (*v1.GPUDevice, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, gPUDevice *v1.GPUDevice, opts metav1.UpdateOptions) (*v1.GPUDevice, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GPUDevice, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GPUDeviceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GPUDevice, err error)
	GPUDeviceExpansion
}

// gPUDevices implements GPUDeviceInterface
type gPUDevices struct {
	*gentype.ClientWithList[*v1.GPUDevice, *v1.GPUDeviceList]
}

// newGPUDevices returns a GPUDevices
func newGPUDevices(c *GpustackV1Client) *gPUDevices {
	return &gPUDevices{
		gentype.NewClientWithList[*v1.GPUDevice, *v1.GPUDeviceList](
			"gpudevices",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.GPUDevice { return &v1.GPUDevice{} },
			func() *v1.GPUDeviceList { return &v1.GPUDeviceList{} }),
	}
}
