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

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type SambaDomainsApi struct {
	Configuration *Configuration
}

func NewSambaDomainsApi() *SambaDomainsApi {
	configuration := NewConfiguration()
	return &SambaDomainsApi{
		Configuration: configuration,
	}
}

func NewSambaDomainsApiWithBasePath(basePath string) *SambaDomainsApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &SambaDomainsApi{
		Configuration: configuration,
	}
}

/**
 * Delete Samba Domain
 * This endpoint allows you to delete a samba domain from an LDAP server.
 *
 * @param ldapserverId Unique identifier o f the LDAP server.
 * @param id Unique identifier of the samba domain.
 * @param contentType 
 * @param accept 
 * @return *string
 */
func (a SambaDomainsApi) LdapserversSambaDomainsDelete(ldapserverId string, id string, contentType string, accept string) (*string, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ldapservers/{ldapserver_id}/sambadomains/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	var successPayload = new(string)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "LdapserversSambaDomainsDelete", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Get Samba Domain
 * This endpoint returns a specific samba domain for an LDAP server.
 *
 * @param ldapserverId Unique identifier o f the LDAP server.
 * @param id Unique identifier of the samba domain.
 * @param contentType 
 * @param accept 
 * @return *SambaDomainOutput
 */
func (a SambaDomainsApi) LdapserversSambaDomainsGet(ldapserverId string, id string, contentType string, accept string) (*SambaDomainOutput, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ldapservers/{ldapserver_id}/sambadomains/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	var successPayload = new(SambaDomainOutput)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "LdapserversSambaDomainsGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * List Samba Domains
 * This endpoint returns all samba domains for an LDAP server.
 *
 * @param ldapserverId Unique identifier of the LDAP server.
 * @param contentType 
 * @param accept 
 * @param fields The comma separated fields included in the returned records. If omitted the default list of fields will be returned. 
 * @param filter Supported operators are: eq, ne, gt, ge, lt, le, between, search
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @param sort The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending. 
 * @return []SambaDomainOutput
 */
func (a SambaDomainsApi) LdapserversSambaDomainsList(ldapserverId string, contentType string, accept string, fields string, filter string, limit int32, skip int32, sort string) ([]SambaDomainOutput, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ldapservers/{ldapserver_id}/sambadomains"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))
	localVarQueryParams.Add("filter", a.Configuration.APIClient.ParameterToString(filter, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))
	localVarQueryParams.Add("sort", a.Configuration.APIClient.ParameterToString(sort, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	var successPayload = new([]SambaDomainOutput)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "LdapserversSambaDomainsList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return *successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return *successPayload, localVarAPIResponse, err
}

/**
 * Create Samba Domain
 * This endpoint allows you to create a samba domain for an LDAP server.
 *
 * @param ldapserverId Unique identifier of the LDAP server.
 * @param body 
 * @param contentType 
 * @param accept 
 * @return *SambaDomainOutput
 */
func (a SambaDomainsApi) LdapserversSambaDomainsPost(ldapserverId string, body SambaDomainInput, contentType string, accept string) (*SambaDomainOutput, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ldapservers/{ldapserver_id}/sambadomains"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	// body params
	localVarPostBody = &body
	var successPayload = new(SambaDomainOutput)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "LdapserversSambaDomainsPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * Update Samba Domain
 * This endpoint allows you to update the samba domain information for an LDAP server.
 *
 * @param ldapserverId Unique identifier o f the LDAP server.
 * @param id Unique identifier of the samba domain.
 * @param body 
 * @param contentType 
 * @param accept 
 * @return *SambaDomainOutput
 */
func (a SambaDomainsApi) LdapserversSambaDomainsPut(ldapserverId string, id string, body SambaDomainInput, contentType string, accept string) (*SambaDomainOutput, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/ldapservers/{ldapserver_id}/sambadomains/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ldapserver_id"+"}", fmt.Sprintf("%v", ldapserverId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(x-api-key)' required
	// set key with prefix in header
	localVarHeaderParams["x-api-key"] = a.Configuration.GetAPIKeyWithPrefix("x-api-key")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// header params "Content-Type"
	localVarHeaderParams["Content-Type"] = a.Configuration.APIClient.ParameterToString(contentType, "")
	// header params "Accept"
	localVarHeaderParams["Accept"] = a.Configuration.APIClient.ParameterToString(accept, "")
	// body params
	localVarPostBody = &body
	var successPayload = new(SambaDomainOutput)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "LdapserversSambaDomainsPut", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

