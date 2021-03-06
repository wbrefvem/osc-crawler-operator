// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlJob) DeepCopyInto(out *CrawlJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlJob.
func (in *CrawlJob) DeepCopy() *CrawlJob {
	if in == nil {
		return nil
	}
	out := new(CrawlJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CrawlJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlJobList) DeepCopyInto(out *CrawlJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CrawlJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlJobList.
func (in *CrawlJobList) DeepCopy() *CrawlJobList {
	if in == nil {
		return nil
	}
	out := new(CrawlJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CrawlJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlJobSpec) DeepCopyInto(out *CrawlJobSpec) {
	*out = *in
	if in.StartURLs != nil {
		in, out := &in.StartURLs, &out.StartURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedDomains != nil {
		in, out := &in.AllowedDomains, &out.AllowedDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlJobSpec.
func (in *CrawlJobSpec) DeepCopy() *CrawlJobSpec {
	if in == nil {
		return nil
	}
	out := new(CrawlJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlJobStatus) DeepCopyInto(out *CrawlJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlJobStatus.
func (in *CrawlJobStatus) DeepCopy() *CrawlJobStatus {
	if in == nil {
		return nil
	}
	out := new(CrawlJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Crawlable) DeepCopyInto(out *Crawlable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Crawlable.
func (in *Crawlable) DeepCopy() *Crawlable {
	if in == nil {
		return nil
	}
	out := new(Crawlable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Crawlable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlableList) DeepCopyInto(out *CrawlableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Crawlable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlableList.
func (in *CrawlableList) DeepCopy() *CrawlableList {
	if in == nil {
		return nil
	}
	out := new(CrawlableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CrawlableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlableSpec) DeepCopyInto(out *CrawlableSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlableSpec.
func (in *CrawlableSpec) DeepCopy() *CrawlableSpec {
	if in == nil {
		return nil
	}
	out := new(CrawlableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrawlableStatus) DeepCopyInto(out *CrawlableStatus) {
	*out = *in
	if in.CrawlJobs != nil {
		in, out := &in.CrawlJobs, &out.CrawlJobs
		*out = make([]CrawlJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrawlableStatus.
func (in *CrawlableStatus) DeepCopy() *CrawlableStatus {
	if in == nil {
		return nil
	}
	out := new(CrawlableStatus)
	in.DeepCopyInto(out)
	return out
}
