/* 
 * JumpCloud APIs
 *
 * V1 and V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package v1

type Body5 struct {

	// A unique name for the Tag.
	Name string `json:"name"`

	// An array of system ids that are associated to the Tag.
	Systems []ErrorUnknown `json:"systems,omitempty"`

	// An array of system user ids that are associated to the Tag.
	Systemusers []ErrorUnknown `json:"systemusers,omitempty"`

	// A date timestamp indicating when this Tag will expire itself. When a Tag expires it will revoke any system user to system associations.
	ExpirationTime int32 `json:"expirationTime,omitempty"`
}
