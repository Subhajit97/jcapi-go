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

type SystemGroupMembersReq struct {

	// How to modify the membership connection.
	Op string `json:"op"`

	// The member type.
	Type_ string `json:"type"`

	// The ObjectID of member being added or removed.
	Id string `json:"id"`
}
