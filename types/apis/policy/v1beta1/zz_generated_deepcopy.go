package v1beta1

import (
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetList) DeepCopyInto(out *PodDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]policyv1beta1.PodDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetList.
func (in *PodDisruptionBudgetList) DeepCopy() *PodDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}