// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementSettingRegexConstraint undocumented
type DeviceManagementSettingRegexConstraint struct {
	// DeviceManagementConstraint is the base model of DeviceManagementSettingRegexConstraint
	DeviceManagementConstraint
	// Regex The RegEx pattern to match against
	Regex *string `json:"regex,omitempty"`
}