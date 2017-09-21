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

type InlineResponse200ConfigConstantAttributes struct {

	Label string `json:"label,omitempty"`

	ReadOnly bool `json:"readOnly,omitempty"`

	Tooltip InlineResponse200ConfigIdpEntityIdTooltip `json:"tooltip,omitempty"`

	Type_ string `json:"type,omitempty"`

	Value []InlineResponse200ConfigConstantAttributesValue `json:"value,omitempty"`

	Visible bool `json:"visible,omitempty"`

	Mutable bool `json:"mutable,omitempty"`

	Required bool `json:"required,omitempty"`

	Position int32 `json:"position,omitempty"`
}
