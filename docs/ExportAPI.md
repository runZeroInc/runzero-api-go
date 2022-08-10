# {{classname}}

All URIs are relative to *https://console.runzero.com/api/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAssetTopHWCSV**](ExportApi.md#ExportAssetTopHWCSV) | **Get** /org/assets/top.hw.csv | Top asset hardware products as CSV
[**ExportAssetTopOSCSV**](ExportApi.md#ExportAssetTopOSCSV) | **Get** /org/assets/top.os.csv | Top asset operating systems as CSV
[**ExportAssetTopTagsCSV**](ExportApi.md#ExportAssetTopTagsCSV) | **Get** /org/assets/top.tags.csv | Top asset tags as CSV
[**ExportAssetTopTypesCSV**](ExportApi.md#ExportAssetTopTypesCSV) | **Get** /org/assets/top.types.csv | Top asset types as CSV
[**ExportAssetsCSV**](ExportApi.md#ExportAssetsCSV) | **Get** /export/org/assets.csv | Asset inventory as CSV
[**ExportAssetsJSON**](ExportApi.md#ExportAssetsJSON) | **Get** /export/org/assets.json | Exports the asset inventory
[**ExportAssetsJSONL**](ExportApi.md#ExportAssetsJSONL) | **Get** /export/org/assets.jsonl | Asset inventory as JSON line-delimited
[**ExportAssetsNmapXML**](ExportApi.md#ExportAssetsNmapXML) | **Get** /export/org/assets.nmap.xml | Asset inventory as Nmap-style XML
[**ExportServicesCSV**](ExportApi.md#ExportServicesCSV) | **Get** /export/org/services.csv | Service inventory as CSV
[**ExportServicesJSON**](ExportApi.md#ExportServicesJSON) | **Get** /export/org/services.json | Service inventory as JSON
[**ExportServicesJSONL**](ExportApi.md#ExportServicesJSONL) | **Get** /export/org/services.jsonl | Service inventory as JSON line-delimited
[**ExportServicesTopProductsCSV**](ExportApi.md#ExportServicesTopProductsCSV) | **Get** /org/services/top.products.csv | Top service products as CSV
[**ExportServicesTopProtocolsCSV**](ExportApi.md#ExportServicesTopProtocolsCSV) | **Get** /org/services/top.protocols.csv | Top service protocols as CSV
[**ExportServicesTopTCPCSV**](ExportApi.md#ExportServicesTopTCPCSV) | **Get** /org/services/top.tcp.csv | Top TCP services as CSV
[**ExportServicesTopUDPCSV**](ExportApi.md#ExportServicesTopUDPCSV) | **Get** /org/services/top.udp.csv | Top UDP services as CSV
[**ExportSitesCSV**](ExportApi.md#ExportSitesCSV) | **Get** /export/org/sites.csv | Site list as CSV
[**ExportSitesJSON**](ExportApi.md#ExportSitesJSON) | **Get** /export/org/sites.json | Export all sites
[**ExportSitesJSONL**](ExportApi.md#ExportSitesJSONL) | **Get** /export/org/sites.jsonl | Site list as JSON line-delimited
[**ExportSubnetUtilizationStatsCSV**](ExportApi.md#ExportSubnetUtilizationStatsCSV) | **Get** /org/services/subnet.stats.csv | Subnet utilization statistics as as CSV
[**ExportWirelessCSV**](ExportApi.md#ExportWirelessCSV) | **Get** /export/org/wireless.csv | Wireless inventory as CSV
[**ExportWirelessJSON**](ExportApi.md#ExportWirelessJSON) | **Get** /export/org/wireless.json | Wireless inventory as JSON
[**ExportWirelessJSONL**](ExportApi.md#ExportWirelessJSONL) | **Get** /export/org/wireless.jsonl | Wireless inventory as JSON line-delimited

# **ExportAssetTopHWCSV**
> *os.File ExportAssetTopHWCSV(ctx, )
Top asset hardware products as CSV

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

# **ExportAssetTopOSCSV**
> *os.File ExportAssetTopOSCSV(ctx, )
Top asset operating systems as CSV

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

# **ExportAssetTopTagsCSV**
> *os.File ExportAssetTopTagsCSV(ctx, )
Top asset tags as CSV

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

# **ExportAssetTopTypesCSV**
> *os.File ExportAssetTopTypesCSV(ctx, )
Top asset types as CSV

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

# **ExportAssetsCSV**
> *os.File ExportAssetsCSV(ctx, optional)
Asset inventory as CSV

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportAssetsCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportAssetsCSVOpts struct
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

# **ExportAssetsJSON**
> []Asset ExportAssetsJSON(ctx, optional)
Exports the asset inventory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportAssetsJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportAssetsJSONOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[**[]Asset**](Asset.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportAssetsJSONL**
> *os.File ExportAssetsJSONL(ctx, optional)
Asset inventory as JSON line-delimited

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportAssetsJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportAssetsJSONLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportAssetsNmapXML**
> *os.File ExportAssetsNmapXML(ctx, optional)
Asset inventory as Nmap-style XML

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportAssetsNmapXMLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportAssetsNmapXMLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportServicesCSV**
> *os.File ExportServicesCSV(ctx, optional)
Service inventory as CSV

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportServicesCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportServicesCSVOpts struct
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

# **ExportServicesJSON**
> []Service ExportServicesJSON(ctx, optional)
Service inventory as JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportServicesJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportServicesJSONOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[**[]Service**](Service.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportServicesJSONL**
> *os.File ExportServicesJSONL(ctx, optional)
Service inventory as JSON line-delimited

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportServicesJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportServicesJSONLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportServicesTopProductsCSV**
> *os.File ExportServicesTopProductsCSV(ctx, )
Top service products as CSV

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

# **ExportServicesTopProtocolsCSV**
> *os.File ExportServicesTopProtocolsCSV(ctx, )
Top service protocols as CSV

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

# **ExportServicesTopTCPCSV**
> *os.File ExportServicesTopTCPCSV(ctx, )
Top TCP services as CSV

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

# **ExportServicesTopUDPCSV**
> *os.File ExportServicesTopUDPCSV(ctx, )
Top UDP services as CSV

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

# **ExportSitesCSV**
> *os.File ExportSitesCSV(ctx, )
Site list as CSV

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

# **ExportSitesJSON**
> []Site ExportSitesJSON(ctx, optional)
Export all sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportSitesJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportSitesJSONOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[**[]Site**](Site.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportSitesJSONL**
> *os.File ExportSitesJSONL(ctx, optional)
Site list as JSON line-delimited

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportSitesJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportSitesJSONLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportSubnetUtilizationStatsCSV**
> *os.File ExportSubnetUtilizationStatsCSV(ctx, optional)
Subnet utilization statistics as as CSV

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportSubnetUtilizationStatsCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportSubnetUtilizationStatsCSVOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mask** | **optional.String**| an optional subnet mask size (ex:24) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportWirelessCSV**
> *os.File ExportWirelessCSV(ctx, optional)
Wireless inventory as CSV

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportWirelessCSVOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportWirelessCSVOpts struct
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

# **ExportWirelessJSON**
> []Wireless ExportWirelessJSON(ctx, optional)
Wireless inventory as JSON

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportWirelessJSONOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportWirelessJSONOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[**[]Wireless**](Wireless.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportWirelessJSONL**
> *os.File ExportWirelessJSONL(ctx, optional)
Wireless inventory as JSON line-delimited

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportApiExportWirelessJSONLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportApiExportWirelessJSONLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| an optional search string for filtering results | 
 **fields** | **optional.String**| an optional list of fields to export, comma-separated | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

