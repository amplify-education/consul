// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package catalogv2beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Service within kubernetes types, where deepcopy-gen is used.
func (in *Service) DeepCopyInto(out *Service) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service. Required by controller-gen.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Service. Required by controller-gen.
func (in *Service) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServicePort within kubernetes types, where deepcopy-gen is used.
func (in *ServicePort) DeepCopyInto(out *ServicePort) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePort. Required by controller-gen.
func (in *ServicePort) DeepCopy() *ServicePort {
	if in == nil {
		return nil
	}
	out := new(ServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServicePort. Required by controller-gen.
func (in *ServicePort) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}