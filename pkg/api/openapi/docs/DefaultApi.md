# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafka**](DefaultApi.md#CreateKafka) | **Post** /api/managed-services-api/v1/kafkas | Create a new kafka Request
[**CreateServiceAccount**](DefaultApi.md#CreateServiceAccount) | **Post** /api/managed-services-api/v1/serviceaccounts | Create a service account
[**DeleteKafkaById**](DefaultApi.md#DeleteKafkaById) | **Delete** /api/managed-services-api/v1/kafkas/{id} | Delete a kafka request by id
[**DeleteServiceAccount**](DefaultApi.md#DeleteServiceAccount) | **Delete** /api/managed-services-api/v1/serviceaccounts/{id} | Delete service account
[**GetKafkaById**](DefaultApi.md#GetKafkaById) | **Get** /api/managed-services-api/v1/kafkas/{id} | Get a kafka request by id
[**GetMetricsByInstantQuery**](DefaultApi.md#GetMetricsByInstantQuery) | **Get** /api/managed-services-api/v1/kafkas/{id}/metrics/query | Get metrics with instant query by kafka id.
[**GetMetricsByRangeQuery**](DefaultApi.md#GetMetricsByRangeQuery) | **Get** /api/managed-services-api/v1/kafkas/{id}/metrics/query_range | Get metrics with timeseries range query by kafka id.
[**GetServiceAccountById**](DefaultApi.md#GetServiceAccountById) | **Get** /api/managed-services-api/v1/serviceaccounts/{id} | get service account by id
[**ListCloudProviderRegions**](DefaultApi.md#ListCloudProviderRegions) | **Get** /api/managed-services-api/v1/cloud_providers/{id}/regions | Retrieves the list of supported regions of the supported cloud provider.
[**ListCloudProviders**](DefaultApi.md#ListCloudProviders) | **Get** /api/managed-services-api/v1/cloud_providers | Retrieves the list of supported cloud providers.
[**ListKafkas**](DefaultApi.md#ListKafkas) | **Get** /api/managed-services-api/v1/kafkas | Returns a list of Kafka requests
[**ListServiceAccounts**](DefaultApi.md#ListServiceAccounts) | **Get** /api/managed-services-api/v1/serviceaccounts | List service accounts
[**ResetServiceAccountCreds**](DefaultApi.md#ResetServiceAccountCreds) | **Post** /api/managed-services-api/v1/serviceaccounts/{id}/reset-credentials | reset credentials for the service account



## CreateKafka

> KafkaRequest CreateKafka(ctx, async, kafkaRequestPayload)

Create a new kafka Request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**async** | **bool**| Perform the action in an asynchronous manner | 
**kafkaRequestPayload** | [**KafkaRequestPayload**](KafkaRequestPayload.md)| Kafka data | 

### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccount

> ServiceAccount CreateServiceAccount(ctx, serviceAccountRequest)

Create a service account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountRequest** | [**ServiceAccountRequest**](ServiceAccountRequest.md)| service account request | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaById

> Error DeleteKafkaById(ctx, id, async)

Delete a kafka request by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 
**async** | **bool**| Perform the action in an asynchronous manner | 

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> Error DeleteServiceAccount(ctx, id)

Delete service account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaById

> KafkaRequest GetKafkaById(ctx, id)

Get a kafka request by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 

### Return type

[**KafkaRequest**](KafkaRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByInstantQuery

> MetricsInstantQueryList GetMetricsByInstantQuery(ctx, id, optional)

Get metrics with instant query by kafka id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 
 **optional** | ***GetMetricsByInstantQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMetricsByInstantQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | [**optional.Interface of []string**](string.md)| List of metrics to fetch. Fetch all metrics when empty. List entries are kafka internal metric names. | [default to []]

### Return type

[**MetricsInstantQueryList**](MetricsInstantQueryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByRangeQuery

> MetricsRangeQueryList GetMetricsByRangeQuery(ctx, id, duration, interval, optional)

Get metrics with timeseries range query by kafka id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 
**duration** | **int64**| The length of time in minutes over which to return the metrics. | [default to 5]
**interval** | **int64**| The interval in seconds between data points. | [default to 30]
 **optional** | ***GetMetricsByRangeQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMetricsByRangeQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filters** | [**optional.Interface of []string**](string.md)| List of metrics to fetch. Fetch all metrics when empty. List entries are kafka internal metric names. | [default to []]

### Return type

[**MetricsRangeQueryList**](MetricsRangeQueryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountById

> ServiceAccount GetServiceAccountById(ctx, id)

get service account by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviderRegions

> CloudRegionList ListCloudProviderRegions(ctx, id, optional)

Retrieves the list of supported regions of the supported cloud provider.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 
 **optional** | ***ListCloudProviderRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCloudProviderRegionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| Page index | 
 **size** | **optional.String**| Number of items in each page | 

### Return type

[**CloudRegionList**](CloudRegionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviders

> CloudProviderList ListCloudProviders(ctx, optional)

Retrieves the list of supported cloud providers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCloudProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCloudProvidersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| Page index | 
 **size** | **optional.String**| Number of items in each page | 

### Return type

[**CloudProviderList**](CloudProviderList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKafkas

> KafkaRequestList ListKafkas(ctx, optional)

Returns a list of Kafka requests

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListKafkasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKafkasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| Page index | 
 **size** | **optional.String**| Number of items in each page | 
 **orderBy** | **optional.String**| Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the kafkaRequests fields. For example, in order to retrieve all kafkas ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or in order to retrieve all kafkas ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results will be ordered by name. | 
 **search** | **optional.String**| Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: cloud_provider, name, owner, region and status. Allowed comparators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve kafka request with name equal &#x60;my-kafka&#x60; and region equal &#x60;aws&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-kafka and cloud_provider &#x3D; aws &#x60;&#x60;&#x60;  To retrieve kafka request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the kafkas that the user has permission to see will be returned.  Note. If the query is invalid, an error will be returned  | 

### Return type

[**KafkaRequestList**](KafkaRequestList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccounts

> ServiceAccountList ListServiceAccounts(ctx, )

List service accounts

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceAccountList**](ServiceAccountList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetServiceAccountCreds

> ServiceAccount ResetServiceAccountCreds(ctx, id)

reset credentials for the service account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The id of record | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

