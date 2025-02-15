/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type System struct {

	Id string `json:"_id,omitempty"`

	Active bool `json:"active,omitempty"`

	AgentVersion string `json:"agentVersion,omitempty"`

	AllowMultiFactorAuthentication bool `json:"allowMultiFactorAuthentication,omitempty"`

	AllowPublicKeyAuthentication bool `json:"allowPublicKeyAuthentication,omitempty"`

	AllowSshPasswordAuthentication bool `json:"allowSshPasswordAuthentication,omitempty"`

	AllowSshRootLogin bool `json:"allowSshRootLogin,omitempty"`

	AmazonInstanceID string `json:"amazonInstanceID,omitempty"`

	Arch string `json:"arch,omitempty"`

	ConnectionHistory []interface{} `json:"connectionHistory,omitempty"`

	Created string `json:"created,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	Fde *Fde `json:"fde,omitempty"`

	FileSystem string `json:"fileSystem,omitempty"`

	HasServiceAccount bool `json:"hasServiceAccount,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	LastContact string `json:"lastContact,omitempty"`

	Mdm *Mdm `json:"mdm,omitempty"`

	ModifySSHDConfig bool `json:"modifySSHDConfig,omitempty"`

	NetworkInterfaces []SystemNetworkInterfaces `json:"networkInterfaces,omitempty"`

	Organization string `json:"organization,omitempty"`

	Os string `json:"os,omitempty"`

	OsFamily string `json:"osFamily,omitempty"`

	RemoteIP string `json:"remoteIP,omitempty"`

	SerialNumber string `json:"serialNumber,omitempty"`

	ServiceAccountState ServiceAccountState `json:"serviceAccountState,omitempty"`

	SshRootEnabled bool `json:"sshRootEnabled,omitempty"`

	SshdParams []SystemSshdParams `json:"sshdParams,omitempty"`

	SystemInsights *SystemSystemInsights `json:"systemInsights,omitempty"`

	SystemTimezone int32 `json:"systemTimezone,omitempty"`

	SystemToken string `json:"systemToken,omitempty"`

	Tags []string `json:"tags,omitempty"`

	TemplateName string `json:"templateName,omitempty"`

	UsernameHashes map[string]string `json:"usernameHashes,omitempty"`

	UserMetrics []UserMetric `json:"userMetrics,omitempty"`

	Version string `json:"version,omitempty"`
}
