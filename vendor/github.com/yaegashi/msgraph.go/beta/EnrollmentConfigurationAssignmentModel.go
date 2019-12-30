// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EnrollmentConfigurationAssignment Enrollment Configuration Assignment
type EnrollmentConfigurationAssignment struct {
	// Entity is the base model of EnrollmentConfigurationAssignment
	Entity
	// Target Represents an assignment to managed devices in the tenant
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
	// Source Type of resource used for deployment to a group, direct or policySet
	Source *DeviceAndAppManagementAssignmentSource `json:"source,omitempty"`
	// SourceID Identifier for resource used for deployment to a group
	SourceID *string `json:"sourceId,omitempty"`
}