/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type ServiceAccountState struct {
	CreateDateTime string `json:"createDatetime,omitempty"`

	CreateMethod string `json:"createMethod,omitempty"`

	CreateUser string `json:"createUser,omitempty"`

	HasSecureToken bool `json:"hasSecureToken,omitempty"`

	KeychainExists bool `json:"keychainExists,omitempty"`

	LegacyPasswordAPFSValid bool `json:"legacyPasswordAPFSValid,omitempty"`

	LegacyPasswordODValid bool `json:"legacyPasswordODValid,omitempty"`

	NewPasswordAPFSValid bool `json:"newPasswordAPFSValid,omitempty"`

	NewPasswordODValid bool `json:"newPasswordODValid,omitempty"`

	PasswordAPFSValid bool `json:"passwordAPFSValid,omitempty"`

	PasswordFileExists bool `json:"passwordFileExists,omitempty"`

	PasswordLastRotated string `json:"passwordLastRotated,omitempty"`

	PasswordLastValid string `json:"passwordLastValid,omitempty"`

	PasswordODValid bool `json:"passwordODValid,omitempty"`

	PasswordRecordExists bool `json:"passwordRecordExists,omitempty"`
}
