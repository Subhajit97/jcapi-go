/* 
 * JumpCloud APIs
 *
 * V1 & V2 versions of JumpCloud's API. The next version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings. The most recent version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings.
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package v2

type Group struct {

	// ObjectId uniquely identifying a Group.
	Id string `json:"id,omitempty"`

	Type_ GroupType `json:"type,omitempty"`

	// Display name of a Group.
	Name string `json:"name,omitempty"`
}
