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

type Body struct {

	DisplayName string `json:"displayName,omitempty"`

	AllowSshPasswordAuthentication bool `json:"allowSshPasswordAuthentication,omitempty"`

	AllowSshRootLogin bool `json:"allowSshRootLogin,omitempty"`

	AllowMultiFactorAuthentication bool `json:"allowMultiFactorAuthentication,omitempty"`

	AllowPublicKeyAuthentication bool `json:"allowPublicKeyAuthentication,omitempty"`

	AgentBoundMessages []SystemsidAgentBoundMessages `json:"agentBoundMessages,omitempty"`

	Id string `json:"_id,omitempty"`
}
