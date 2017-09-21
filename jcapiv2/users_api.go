/* 
 * JumpCloud APIs
 *
 * V1 & V2 versions of JumpCloud's API. The next version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings. The most recent version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings.
 *
 * OpenAPI spec version: 2.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package jcapiv2

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type UsersApi struct {
	Configuration *Configuration
}

func NewUsersApi() *UsersApi {
	configuration := NewConfiguration()
	return &UsersApi{
		Configuration: configuration,
	}
}

func NewUsersApiWithBasePath(basePath string) *UsersApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &UsersApi{
		Configuration: configuration,
	}
}

/**
 * List the associations of a User
 * This endpoint returns the _direct_ associations of a User.  A direct association can be a non-homogenous relationship between 2 different objects. for example Users and Systems.   #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/associations?targets&#x3D;user_group &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param targets 
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphConnection
 */
func (a UsersApi) GraphUserAssociationsList(userId string, targets []string, contentType string, accept string, limit int32, skip int32) ([]GraphConnection, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	var targetsCollectionFormat = "csv"
	localVarQueryParams.Add("targets", a.Configuration.APIClient.ParameterToString(targets, targetsCollectionFormat))

	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphConnection)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserAssociationsList", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Manage the associations of a User
 * This endpoint allows you to manage the _direct_ associations of a User.  A direct association can be a non-homogenous relationship between 2 different objects. for example Users and Systems.   #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/associations &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param body 
 * @return void
 */
func (a UsersApi) GraphUserAssociationsPost(userId string, contentType string, accept string, body GraphManagementReq) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/associations"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserAssociationsPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 * List the parent Groups of a User
 * This endpoint returns all the User Groups a User is a member of.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/memberof &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserMemberOf(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/memberof"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserMemberOf", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the Applications associated with a User
 * This endpoint will return Applications associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Application; this array represents all grouping and/or associations that would have to be removed to deprovision the Application from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/applications &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseApplication(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseApplication", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the Directories associated with a User
 * This endpoint will return Directories associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Directory; this array represents all grouping and/or associations that would have to be removed to deprovision the Directory from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/directories &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseDirectory(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/directories"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseDirectory", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the G Suite instances associated with a User
 * This endpoint will return G Suite instances associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding G Suite instance; this array represents all grouping and/or associations that would have to be removed to deprovision the G Suite instance from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/gsuites &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseGSuite(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/gsuites"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseGSuite", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the LDAP servers associated with a User
 * This endpoint will return LDAP Servers associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding LDAP Server; this array represents all grouping and/or associations that would have to be removed to deprovision the LDAP Server from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/ldapservers &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseLdapServer(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/ldapservers"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseLdapServer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the Office 365 instances associated with User
 * This endpoint will return Office 365 instances associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding Office 365 instance; this array represents all grouping and/or associations that would have to be removed to deprovision the Office 365 instance from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/office365s &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseOffice365(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/office365s"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseOffice365", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the RADIUS Servers associated with a User
 * This endpoint will return RADIUS Servers associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding RADIUS Server; this array represents all grouping and/or associations that would have to be removed to deprovision the RADIUS Server from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/radiusservers &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseRadiusServer(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/radiusservers"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseRadiusServer", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * List the Systems associated with a User
 * This endpoint will return Systems associated with a User. Each element will contain the type, id, attributes and paths.  The &#x60;attributes&#x60; object is a key/value hash of attributes specifically set for this group.  The &#x60;paths&#x60; array enumerates each path from this User to the corresponding System; this array represents all grouping and/or associations that would have to be removed to deprovision the System from this User.  See &#x60;/members&#x60; and &#x60;/associations&#x60; endpoints to manage those collections.  #### Sample Request &#x60;&#x60;&#x60; https://console.jumpcloud.com/api/v2/users/{user_id}/systems &#x60;&#x60;&#x60;
 *
 * @param userId ObjectID of the User.
 * @param contentType 
 * @param accept 
 * @param limit The number of records to return at once.
 * @param skip The offset into the records to return.
 * @return []GraphObjectWithPaths
 */
func (a UsersApi) GraphUserTraverseSystem(userId string, contentType string, accept string, limit int32, skip int32) ([]GraphObjectWithPaths, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/users/{user_id}/systems"
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", fmt.Sprintf("%v", userId), -1)

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
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("skip", a.Configuration.APIClient.ParameterToString(skip, ""))

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
	var successPayload = new([]GraphObjectWithPaths)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GraphUserTraverseSystem", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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

