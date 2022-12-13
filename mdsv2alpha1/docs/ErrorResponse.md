# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | Optional - http status code | [optional] 
**ErrorCode** | **int32** | Optional - Kafka error code (typically 5 digits) | [optional] 
**Type** | **string** | Optional - Type of error | [optional] 
**Message** | **string** | Required - Top level error message | 
**Errors** | [**[]ErrorDetail**](ErrorDetail.md) | Optional - List of errors | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


