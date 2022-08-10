# {{classname}}

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnowExportAssetsCSV**](ServiceNowApi.md#SnowExportAssetsCSV) | **Get** /export/org/assets.servicenow.csv | Export an asset inventory as CSV for ServiceNow integration
[**SnowExportAssetsJSON**](ServiceNowApi.md#SnowExportAssetsJSON) | **Get** /export/org/assets.servicenow.json | Exports the asset inventory as JSON
[**SnowExportServicesCSV**](ServiceNowApi.md#SnowExportServicesCSV) | **Get** /export/org/services.servicenow.csv | Export a service inventory as CSV for ServiceNow integration

# **SnowExportAssetsCSV**
> *os.File SnowExportAssetsCSV(ctx, )
Export an asset inventory as CSV for ServiceNow integration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnowExportAssetsJSON**
> []AssetServiceNow SnowExportAssetsJSON(ctx, )
Exports the asset inventory as JSON

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AssetServiceNow**](AssetServiceNow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnowExportServicesCSV**
> *os.File SnowExportServicesCSV(ctx, )
Export a service inventory as CSV for ServiceNow integration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

