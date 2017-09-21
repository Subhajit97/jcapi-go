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

type InlineResponse2004Results struct {

	Email string `json:"email,omitempty"`

	Username string `json:"username,omitempty"`

	AllowPublicKey bool `json:"allow_public_key,omitempty"`

	SshKeys []interface{} `json:"ssh_keys,omitempty"`

	Sudo bool `json:"sudo,omitempty"`

	EnableManagedUid bool `json:"enable_managed_uid,omitempty"`

	UnixUid int32 `json:"unix_uid,omitempty"`

	UnixGuid int32 `json:"unix_guid,omitempty"`

	Activated bool `json:"activated,omitempty"`

	PasswordExpired bool `json:"password_expired,omitempty"`

	AccountLocked bool `json:"account_locked,omitempty"`

	PasswordlessSudo bool `json:"passwordless_sudo,omitempty"`

	ExternallyManaged bool `json:"externally_managed,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`

	LdapBindingUser bool `json:"ldap_binding_user,omitempty"`

	EnableUserPortalMultifactor bool `json:"enable_user_portal_multifactor,omitempty"`

	AssociatedTagCount int32 `json:"associatedTagCount,omitempty"`

	TotpEnabled bool `json:"totp_enabled,omitempty"`

	Attributes []interface{} `json:"attributes,omitempty"`

	Id string `json:"_id,omitempty"`
}
