//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Flux authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessLog) DeepCopyInto(out *AccessLog) {
	*out = *in
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(FileAccessLog)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessLog.
func (in *AccessLog) DeepCopy() *AccessLog {
	if in == nil {
		return nil
	}
	out := new(AccessLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	out.VirtualService = in.VirtualService
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMapServiceDiscovery) DeepCopyInto(out *CloudMapServiceDiscovery) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMapServiceDiscovery.
func (in *CloudMapServiceDiscovery) DeepCopy() *CloudMapServiceDiscovery {
	if in == nil {
		return nil
	}
	out := new(CloudMapServiceDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudMapServiceStatus) DeepCopyInto(out *CloudMapServiceStatus) {
	*out = *in
	if in.ServiceID != nil {
		in, out := &in.ServiceID, &out.ServiceID
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudMapServiceStatus.
func (in *CloudMapServiceStatus) DeepCopy() *CloudMapServiceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudMapServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnsServiceDiscovery) DeepCopyInto(out *DnsServiceDiscovery) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnsServiceDiscovery.
func (in *DnsServiceDiscovery) DeepCopy() *DnsServiceDiscovery {
	if in == nil {
		return nil
	}
	out := new(DnsServiceDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileAccessLog) DeepCopyInto(out *FileAccessLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileAccessLog.
func (in *FileAccessLog) DeepCopy() *FileAccessLog {
	if in == nil {
		return nil
	}
	out := new(FileAccessLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatchMethod) DeepCopyInto(out *HeaderMatchMethod) {
	*out = *in
	if in.Exact != nil {
		in, out := &in.Exact, &out.Exact
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(MatchRange)
		(*in).DeepCopyInto(*out)
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(string)
		**out = **in
	}
	if in.Suffix != nil {
		in, out := &in.Suffix, &out.Suffix
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatchMethod.
func (in *HeaderMatchMethod) DeepCopy() *HeaderMatchMethod {
	if in == nil {
		return nil
	}
	out := new(HeaderMatchMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckPolicy) DeepCopyInto(out *HealthCheckPolicy) {
	*out = *in
	if in.HealthyThreshold != nil {
		in, out := &in.HealthyThreshold, &out.HealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.IntervalMillis != nil {
		in, out := &in.IntervalMillis, &out.IntervalMillis
		*out = new(int64)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.TimeoutMillis != nil {
		in, out := &in.TimeoutMillis, &out.TimeoutMillis
		*out = new(int64)
		**out = **in
	}
	if in.UnhealthyThreshold != nil {
		in, out := &in.UnhealthyThreshold, &out.UnhealthyThreshold
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckPolicy.
func (in *HealthCheckPolicy) DeepCopy() *HealthCheckPolicy {
	if in == nil {
		return nil
	}
	out := new(HealthCheckPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpRetryPolicy) DeepCopyInto(out *HttpRetryPolicy) {
	*out = *in
	if in.PerRetryTimeoutMillis != nil {
		in, out := &in.PerRetryTimeoutMillis, &out.PerRetryTimeoutMillis
		*out = new(int64)
		**out = **in
	}
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(int64)
		**out = **in
	}
	if in.HttpRetryPolicyEvents != nil {
		in, out := &in.HttpRetryPolicyEvents, &out.HttpRetryPolicyEvents
		*out = make([]HttpRetryPolicyEvent, len(*in))
		copy(*out, *in)
	}
	if in.TcpRetryPolicyEvents != nil {
		in, out := &in.TcpRetryPolicyEvents, &out.TcpRetryPolicyEvents
		*out = make([]TcpRetryPolicyEvent, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpRetryPolicy.
func (in *HttpRetryPolicy) DeepCopy() *HttpRetryPolicy {
	if in == nil {
		return nil
	}
	out := new(HttpRetryPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpRoute) DeepCopyInto(out *HttpRoute) {
	*out = *in
	in.Match.DeepCopyInto(&out.Match)
	in.Action.DeepCopyInto(&out.Action)
	if in.RetryPolicy != nil {
		in, out := &in.RetryPolicy, &out.RetryPolicy
		*out = new(HttpRetryPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpRoute.
func (in *HttpRoute) DeepCopy() *HttpRoute {
	if in == nil {
		return nil
	}
	out := new(HttpRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpRouteAction) DeepCopyInto(out *HttpRouteAction) {
	*out = *in
	if in.WeightedTargets != nil {
		in, out := &in.WeightedTargets, &out.WeightedTargets
		*out = make([]WeightedTarget, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpRouteAction.
func (in *HttpRouteAction) DeepCopy() *HttpRouteAction {
	if in == nil {
		return nil
	}
	out := new(HttpRouteAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpRouteHeader) DeepCopyInto(out *HttpRouteHeader) {
	*out = *in
	if in.Invert != nil {
		in, out := &in.Invert, &out.Invert
		*out = new(bool)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = new(HeaderMatchMethod)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpRouteHeader.
func (in *HttpRouteHeader) DeepCopy() *HttpRouteHeader {
	if in == nil {
		return nil
	}
	out := new(HttpRouteHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpRouteMatch) DeepCopyInto(out *HttpRouteMatch) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HttpRouteHeader, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpRouteMatch.
func (in *HttpRouteMatch) DeepCopy() *HttpRouteMatch {
	if in == nil {
		return nil
	}
	out := new(HttpRouteMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	out.PortMapping = in.PortMapping
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	if in.AccessLog != nil {
		in, out := &in.AccessLog, &out.AccessLog
		*out = new(AccessLog)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchRange) DeepCopyInto(out *MatchRange) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(int64)
		**out = **in
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchRange.
func (in *MatchRange) DeepCopy() *MatchRange {
	if in == nil {
		return nil
	}
	out := new(MatchRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mesh) DeepCopyInto(out *Mesh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mesh.
func (in *Mesh) DeepCopy() *Mesh {
	if in == nil {
		return nil
	}
	out := new(Mesh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Mesh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshCondition) DeepCopyInto(out *MeshCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshCondition.
func (in *MeshCondition) DeepCopy() *MeshCondition {
	if in == nil {
		return nil
	}
	out := new(MeshCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshList) DeepCopyInto(out *MeshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Mesh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshList.
func (in *MeshList) DeepCopy() *MeshList {
	if in == nil {
		return nil
	}
	out := new(MeshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MeshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshSpec) DeepCopyInto(out *MeshSpec) {
	*out = *in
	if in.ServiceDiscoveryType != nil {
		in, out := &in.ServiceDiscoveryType, &out.ServiceDiscoveryType
		*out = new(MeshServiceDiscoveryType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshSpec.
func (in *MeshSpec) DeepCopy() *MeshSpec {
	if in == nil {
		return nil
	}
	out := new(MeshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeshStatus) DeepCopyInto(out *MeshStatus) {
	*out = *in
	if in.MeshArn != nil {
		in, out := &in.MeshArn, &out.MeshArn
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MeshCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeshStatus.
func (in *MeshStatus) DeepCopy() *MeshStatus {
	if in == nil {
		return nil
	}
	out := new(MeshStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortMapping) DeepCopyInto(out *PortMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortMapping.
func (in *PortMapping) DeepCopy() *PortMapping {
	if in == nil {
		return nil
	}
	out := new(PortMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
	if in.Http != nil {
		in, out := &in.Http, &out.Http
		*out = new(HttpRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.Tcp != nil {
		in, out := &in.Tcp, &out.Tcp
		*out = new(TcpRoute)
		(*in).DeepCopyInto(*out)
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceDiscovery) DeepCopyInto(out *ServiceDiscovery) {
	*out = *in
	if in.CloudMap != nil {
		in, out := &in.CloudMap, &out.CloudMap
		*out = new(CloudMapServiceDiscovery)
		(*in).DeepCopyInto(*out)
	}
	if in.Dns != nil {
		in, out := &in.Dns, &out.Dns
		*out = new(DnsServiceDiscovery)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceDiscovery.
func (in *ServiceDiscovery) DeepCopy() *ServiceDiscovery {
	if in == nil {
		return nil
	}
	out := new(ServiceDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcpRoute) DeepCopyInto(out *TcpRoute) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcpRoute.
func (in *TcpRoute) DeepCopy() *TcpRoute {
	if in == nil {
		return nil
	}
	out := new(TcpRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcpRouteAction) DeepCopyInto(out *TcpRouteAction) {
	*out = *in
	if in.WeightedTargets != nil {
		in, out := &in.WeightedTargets, &out.WeightedTargets
		*out = make([]WeightedTarget, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcpRouteAction.
func (in *TcpRouteAction) DeepCopy() *TcpRouteAction {
	if in == nil {
		return nil
	}
	out := new(TcpRouteAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNode) DeepCopyInto(out *VirtualNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNode.
func (in *VirtualNode) DeepCopy() *VirtualNode {
	if in == nil {
		return nil
	}
	out := new(VirtualNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeCondition) DeepCopyInto(out *VirtualNodeCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeCondition.
func (in *VirtualNodeCondition) DeepCopy() *VirtualNodeCondition {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeList) DeepCopyInto(out *VirtualNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeList.
func (in *VirtualNodeList) DeepCopy() *VirtualNodeList {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeSpec) DeepCopyInto(out *VirtualNodeSpec) {
	*out = *in
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]Listener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceDiscovery != nil {
		in, out := &in.ServiceDiscovery, &out.ServiceDiscovery
		*out = new(ServiceDiscovery)
		(*in).DeepCopyInto(*out)
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]Backend, len(*in))
		copy(*out, *in)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(Logging)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeSpec.
func (in *VirtualNodeSpec) DeepCopy() *VirtualNodeSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodeStatus) DeepCopyInto(out *VirtualNodeStatus) {
	*out = *in
	if in.MeshArn != nil {
		in, out := &in.MeshArn, &out.MeshArn
		*out = new(string)
		**out = **in
	}
	if in.VirtualNodeArn != nil {
		in, out := &in.VirtualNodeArn, &out.VirtualNodeArn
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VirtualNodeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudMapService != nil {
		in, out := &in.CloudMapService, &out.CloudMapService
		*out = new(CloudMapServiceStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodeStatus.
func (in *VirtualNodeStatus) DeepCopy() *VirtualNodeStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualRouter) DeepCopyInto(out *VirtualRouter) {
	*out = *in
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]VirtualRouterListener, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualRouter.
func (in *VirtualRouter) DeepCopy() *VirtualRouter {
	if in == nil {
		return nil
	}
	out := new(VirtualRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualRouterListener) DeepCopyInto(out *VirtualRouterListener) {
	*out = *in
	out.PortMapping = in.PortMapping
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualRouterListener.
func (in *VirtualRouterListener) DeepCopy() *VirtualRouterListener {
	if in == nil {
		return nil
	}
	out := new(VirtualRouterListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualService) DeepCopyInto(out *VirtualService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService.
func (in *VirtualService) DeepCopy() *VirtualService {
	if in == nil {
		return nil
	}
	out := new(VirtualService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceBackend) DeepCopyInto(out *VirtualServiceBackend) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceBackend.
func (in *VirtualServiceBackend) DeepCopy() *VirtualServiceBackend {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceCondition) DeepCopyInto(out *VirtualServiceCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceCondition.
func (in *VirtualServiceCondition) DeepCopy() *VirtualServiceCondition {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceList) DeepCopyInto(out *VirtualServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceList.
func (in *VirtualServiceList) DeepCopy() *VirtualServiceList {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceSpec) DeepCopyInto(out *VirtualServiceSpec) {
	*out = *in
	if in.VirtualRouter != nil {
		in, out := &in.VirtualRouter, &out.VirtualRouter
		*out = new(VirtualRouter)
		(*in).DeepCopyInto(*out)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceSpec.
func (in *VirtualServiceSpec) DeepCopy() *VirtualServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceStatus) DeepCopyInto(out *VirtualServiceStatus) {
	*out = *in
	if in.VirtualServiceArn != nil {
		in, out := &in.VirtualServiceArn, &out.VirtualServiceArn
		*out = new(string)
		**out = **in
	}
	if in.VirtualRouterArn != nil {
		in, out := &in.VirtualRouterArn, &out.VirtualRouterArn
		*out = new(string)
		**out = **in
	}
	if in.RouteArns != nil {
		in, out := &in.RouteArns, &out.RouteArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VirtualServiceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceStatus.
func (in *VirtualServiceStatus) DeepCopy() *VirtualServiceStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WeightedTarget) DeepCopyInto(out *WeightedTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WeightedTarget.
func (in *WeightedTarget) DeepCopy() *WeightedTarget {
	if in == nil {
		return nil
	}
	out := new(WeightedTarget)
	in.DeepCopyInto(out)
	return out
}
