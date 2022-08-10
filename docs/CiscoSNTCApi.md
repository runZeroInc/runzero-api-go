# {{classname}}

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAssetsCiscoCSV**](CiscoSNTCApi.md#ExportAssetsCiscoCSV) | **Get** /export/org/assets.cisco.csv | Cisco serial number and model name export for Cisco Smart Net Total Care Service.

# **ExportAssetsCiscoCSV**
> *os.File ExportAssetsCiscoCSV(ctx, optional)
Cisco serial number and model name export for Cisco Smart Net Total Care Service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CiscoSNTCApiExportAssetsCiscoCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CiscoSNTCApiExportAssetsCiscoCSVOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

