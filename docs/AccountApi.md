# {{classname}}

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountCredential**](AccountApi.md#CreateAccountCredential) | **Put** /account/credentials | Create a new credential
[**CreateAccountGroup**](AccountApi.md#CreateAccountGroup) | **Post** /account/groups | Create a new group
[**CreateAccountGroupMapping**](AccountApi.md#CreateAccountGroupMapping) | **Post** /account/sso/groups | Create a new SSO group mapping
[**CreateAccountKey**](AccountApi.md#CreateAccountKey) | **Put** /account/keys | Create a new key
[**CreateAccountOrganization**](AccountApi.md#CreateAccountOrganization) | **Put** /account/orgs | Create a new organization
[**CreateAccountScanTemplate**](AccountApi.md#CreateAccountScanTemplate) | **Post** /account/tasks/templates | Create a new scan template
[**CreateAccountUser**](AccountApi.md#CreateAccountUser) | **Put** /account/users | Create a new user account
[**CreateAccountUserInvite**](AccountApi.md#CreateAccountUserInvite) | **Put** /account/users/invite | Create a new user account and send an email invite
[**DeleteAccountOrganizationExportToken**](AccountApi.md#DeleteAccountOrganizationExportToken) | **Delete** /account/orgs/{org_id}/exportToken | Removes the export token from the specified organization
[**ExportEventsJSON**](AccountApi.md#ExportEventsJSON) | **Get** /account/events.json | System event log as JSON
[**ExportEventsJSONL**](AccountApi.md#ExportEventsJSONL) | **Get** /account/events.jsonl | System event log as JSON line-delimited
[**GetAccountAgents**](AccountApi.md#GetAccountAgents) | **Get** /account/agents | Get all agents across all organizations
[**GetAccountCredential**](AccountApi.md#GetAccountCredential) | **Get** /account/credentials/{credential_id} | Get credential details
[**GetAccountCredentials**](AccountApi.md#GetAccountCredentials) | **Get** /account/credentials | Get all account credentials
[**GetAccountGroup**](AccountApi.md#GetAccountGroup) | **Get** /account/groups/{group_id} | Get group details
[**GetAccountGroupMapping**](AccountApi.md#GetAccountGroupMapping) | **Get** /account/sso/groups/{group_mapping_id} | Get SSO group mapping details
[**GetAccountGroupMappings**](AccountApi.md#GetAccountGroupMappings) | **Get** /account/sso/groups | Get all SSO group mappings
[**GetAccountGroups**](AccountApi.md#GetAccountGroups) | **Get** /account/groups | Get all groups
[**GetAccountKey**](AccountApi.md#GetAccountKey) | **Get** /account/keys/{key_id} | Get key details
[**GetAccountKeys**](AccountApi.md#GetAccountKeys) | **Get** /account/keys | Get all active API keys
[**GetAccountLicense**](AccountApi.md#GetAccountLicense) | **Get** /account/license | Get license details
[**GetAccountOrganization**](AccountApi.md#GetAccountOrganization) | **Get** /account/orgs/{org_id} | Get organization details
[**GetAccountOrganizations**](AccountApi.md#GetAccountOrganizations) | **Get** /account/orgs | Get all organization details
[**GetAccountScanTemplate**](AccountApi.md#GetAccountScanTemplate) | **Get** /account/tasks/templates/{scan_template_id} | Get scan template details
[**GetAccountScanTemplates**](AccountApi.md#GetAccountScanTemplates) | **Get** /account/tasks/templates | Get all scan templates across all organizations (up to 1000)
[**GetAccountSites**](AccountApi.md#GetAccountSites) | **Get** /account/sites | Get all sites details across all organizations
[**GetAccountTasks**](AccountApi.md#GetAccountTasks) | **Get** /account/tasks | Get all task details across all organizations (up to 1000)
[**GetAccountUser**](AccountApi.md#GetAccountUser) | **Get** /account/users/{user_id} | Get user details
[**GetAccountUsers**](AccountApi.md#GetAccountUsers) | **Get** /account/users | Get all users
[**RemoveAccountCredential**](AccountApi.md#RemoveAccountCredential) | **Delete** /account/credentials/{credential_id} | Remove this credential
[**RemoveAccountGroup**](AccountApi.md#RemoveAccountGroup) | **Delete** /account/groups/{group_id} | Remove this group
[**RemoveAccountGroupMapping**](AccountApi.md#RemoveAccountGroupMapping) | **Delete** /account/sso/groups/{group_mapping_id} | Remove this SSO group mapping
[**RemoveAccountKey**](AccountApi.md#RemoveAccountKey) | **Delete** /account/keys/{key_id} | Remove this key
[**RemoveAccountOrganization**](AccountApi.md#RemoveAccountOrganization) | **Delete** /account/orgs/{org_id} | Remove this organization
[**RemoveAccountScanTemplate**](AccountApi.md#RemoveAccountScanTemplate) | **Delete** /account/tasks/templates/{scan_template_id} | Remove scan template
[**RemoveAccountUser**](AccountApi.md#RemoveAccountUser) | **Delete** /account/users/{user_id} | Remove this user
[**ResetAccountUserLockout**](AccountApi.md#ResetAccountUserLockout) | **Patch** /account/users/{user_id}/resetLockout | Resets the user&#x27;s lockout status
[**ResetAccountUserMFA**](AccountApi.md#ResetAccountUserMFA) | **Patch** /account/users/{user_id}/resetMFA | Resets the user&#x27;s MFA tokens
[**ResetAccountUserPassword**](AccountApi.md#ResetAccountUserPassword) | **Patch** /account/users/{user_id}/resetPassword | Sends the user a password reset email
[**RotateAccountKey**](AccountApi.md#RotateAccountKey) | **Patch** /account/keys/{key_id}/rotate | Rotates the key secret
[**RotateAccountOrganizationExportToken**](AccountApi.md#RotateAccountOrganizationExportToken) | **Patch** /account/orgs/{org_id}/exportToken/rotate | Rotates the organization export token and returns the updated organization
[**UpdateAccountGroup**](AccountApi.md#UpdateAccountGroup) | **Put** /account/groups | Update an existing group
[**UpdateAccountGroupMapping**](AccountApi.md#UpdateAccountGroupMapping) | **Put** /account/sso/groups | Update an existing SSO group mapping
[**UpdateAccountOrganization**](AccountApi.md#UpdateAccountOrganization) | **Patch** /account/orgs/{org_id} | Update organization details
[**UpdateAccountScanTemplate**](AccountApi.md#UpdateAccountScanTemplate) | **Put** /account/tasks/templates | Update scan template
[**UpdateAccountUser**](AccountApi.md#UpdateAccountUser) | **Patch** /account/users/{user_id} | Update a user&#x27;s details

# **CreateAccountCredential**
> Credential CreateAccountCredential(ctx, body)
Create a new credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CredentialOptions**](CredentialOptions.md)| credential parameters | 

### Return type

[**Credential**](Credential.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountGroup**
> CreateAccountGroup(ctx, body)
Create a new group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupPost**](GroupPost.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountGroupMapping**
> CreateAccountGroupMapping(ctx, body)
Create a new SSO group mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMapping**](GroupMapping.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountKey**
> ApiKey CreateAccountKey(ctx, body)
Create a new key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiKeyOptions**](ApiKeyOptions.md)| key parameters | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountOrganization**
> Organization CreateAccountOrganization(ctx, body)
Create a new organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrgOptions**](OrgOptions.md)| organization definition | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountScanTemplate**
> ScanTemplate CreateAccountScanTemplate(ctx, body)
Create a new scan template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScanTemplateOptions**](ScanTemplateOptions.md)|  | 

### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountUser**
> User CreateAccountUser(ctx, body)
Create a new user account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserOptions**](UserOptions.md)| user parameters | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccountUserInvite**
> User CreateAccountUserInvite(ctx, body)
Create a new user account and send an email invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserInviteOptions**](UserInviteOptions.md)| user invite parameters | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountOrganizationExportToken**
> DeleteAccountOrganizationExportToken(ctx, orgId)
Removes the export token from the specified organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)| UUID of the organization to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportEventsJSON**
> []Event ExportEventsJSON(ctx, optional)
System event log as JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiExportEventsJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiExportEventsJSONOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[**[]Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportEventsJSONL**
> []Event ExportEventsJSONL(ctx, optional)
System event log as JSON line-delimited

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiExportEventsJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiExportEventsJSONLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[**[]Event**](Event.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountAgents**
> []Agent GetAccountAgents(ctx, optional)
Get all agents across all organizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetAccountAgentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountAgentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Agent**](Agent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountCredential**
> Credential GetAccountCredential(ctx, credentialId)
Get credential details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialId** | [**string**](.md)| UUID of the credential to retrieve | 

### Return type

[**Credential**](Credential.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountCredentials**
> []Credential GetAccountCredentials(ctx, optional)
Get all account credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetAccountCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountCredentialsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Credential**](Credential.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroup**
> GetAccountGroup(ctx, groupId)
Get group details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | [**string**](.md)| UUID of the group | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroupMapping**
> GetAccountGroupMapping(ctx, groupMappingId)
Get SSO group mapping details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupMappingId** | [**string**](.md)| UUID of the SSO group mapping | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroupMappings**
> GetAccountGroupMappings(ctx, )
Get all SSO group mappings

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountGroups**
> GetAccountGroups(ctx, )
Get all groups

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountKey**
> GetAccountKey(ctx, keyId)
Get key details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | [**string**](.md)| UUID of the key to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountKeys**
> []ApiKey GetAccountKeys(ctx, )
Get all active API keys

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ApiKey**](APIKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountLicense**
> GetAccountLicense(ctx, )
Get license details

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountOrganization**
> GetAccountOrganization(ctx, orgId)
Get organization details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)| UUID of the organization to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountOrganizations**
> []Organization GetAccountOrganizations(ctx, optional)
Get all organization details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetAccountOrganizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountOrganizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScanTemplate**
> ScanTemplate GetAccountScanTemplate(ctx, scanTemplateId)
Get scan template details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scanTemplateId** | [**string**](.md)| UUID of the scan template to retrieve | 

### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScanTemplates**
> []ScanTemplate GetAccountScanTemplates(ctx, optional)
Get all scan templates across all organizations (up to 1000)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetAccountScanTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountScanTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]ScanTemplate**](ScanTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountSites**
> []Site GetAccountSites(ctx, optional)
Get all sites details across all organizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetAccountSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Site**](Site.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountTasks**
> []Task GetAccountTasks(ctx, optional)
Get all task details across all organizations (up to 1000)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetAccountTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[**[]Task**](Task.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountUser**
> GetAccountUser(ctx, userId)
Get user details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | [**string**](.md)| UUID of the user to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountUsers**
> []User GetAccountUsers(ctx, )
Get all users

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountCredential**
> RemoveAccountCredential(ctx, credentialId)
Remove this credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentialId** | [**string**](.md)| UUID of the credential to delete | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountGroup**
> RemoveAccountGroup(ctx, groupId)
Remove this group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | [**string**](.md)| UUID of the group | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountGroupMapping**
> RemoveAccountGroupMapping(ctx, groupMappingId)
Remove this SSO group mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupMappingId** | [**string**](.md)| UUID of the SSO group mapping | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountKey**
> RemoveAccountKey(ctx, keyId)
Remove this key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | [**string**](.md)| UUID of the key to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountOrganization**
> RemoveAccountOrganization(ctx, orgId)
Remove this organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)| UUID of the organization to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountScanTemplate**
> ScanTemplate RemoveAccountScanTemplate(ctx, scanTemplateId)
Remove scan template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scanTemplateId** | [**string**](.md)| UUID of the scan template to remove | 

### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountUser**
> RemoveAccountUser(ctx, userId)
Remove this user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | [**string**](.md)| UUID of the user to delete | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetAccountUserLockout**
> ResetAccountUserLockout(ctx, userId)
Resets the user's lockout status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | [**string**](.md)| UUID of the user to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetAccountUserMFA**
> ResetAccountUserMFA(ctx, userId)
Resets the user's MFA tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | [**string**](.md)| UUID of the user to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetAccountUserPassword**
> ResetAccountUserPassword(ctx, userId)
Sends the user a password reset email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | [**string**](.md)| UUID of the user to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateAccountKey**
> ApiKey RotateAccountKey(ctx, keyId)
Rotates the key secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyId** | [**string**](.md)| UUID of the key to retrieve | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateAccountOrganizationExportToken**
> Organization RotateAccountOrganizationExportToken(ctx, orgId)
Rotates the organization export token and returns the updated organization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)| UUID of the organization to retrieve | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountGroup**
> UpdateAccountGroup(ctx, body)
Update an existing group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupPut**](GroupPut.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountGroupMapping**
> UpdateAccountGroupMapping(ctx, body)
Update an existing SSO group mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GroupMapping**](GroupMapping.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountOrganization**
> Organization UpdateAccountOrganization(ctx, body, orgId)
Update organization details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrgOptions**](OrgOptions.md)| organization options | 
  **orgId** | [**string**](.md)| UUID of the organization to retrieve | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountScanTemplate**
> ScanTemplate UpdateAccountScanTemplate(ctx, body)
Update scan template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScanTemplate**](ScanTemplate.md)|  | 

### Return type

[**ScanTemplate**](ScanTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountUser**
> UpdateAccountUser(ctx, body, userId)
Update a user's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserOptions**](UserOptions.md)| user parameters | 
  **userId** | [**string**](.md)| UUID of the user to retrieve | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

