// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// MachinePoolStatusApplyConfiguration represents an declarative configuration of the MachinePoolStatus type for use
// with apply.
type MachinePoolStatusApplyConfiguration struct {
	Replicas    *int32                                   `json:"replicas,omitempty"`
	MachineSets []MachineSetStatusApplyConfiguration     `json:"machineSets,omitempty"`
	Conditions  []MachinePoolConditionApplyConfiguration `json:"conditions,omitempty"`
}

// MachinePoolStatusApplyConfiguration constructs an declarative configuration of the MachinePoolStatus type for use with
// apply.
func MachinePoolStatus() *MachinePoolStatusApplyConfiguration {
	return &MachinePoolStatusApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *MachinePoolStatusApplyConfiguration) WithReplicas(value int32) *MachinePoolStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithMachineSets adds the given value to the MachineSets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MachineSets field.
func (b *MachinePoolStatusApplyConfiguration) WithMachineSets(values ...*MachineSetStatusApplyConfiguration) *MachinePoolStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMachineSets")
		}
		b.MachineSets = append(b.MachineSets, *values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *MachinePoolStatusApplyConfiguration) WithConditions(values ...*MachinePoolConditionApplyConfiguration) *MachinePoolStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
