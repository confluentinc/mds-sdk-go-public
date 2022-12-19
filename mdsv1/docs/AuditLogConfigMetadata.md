# AuditLogConfigMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVersion** | **string** | An opaque, server-generated marker, used to detect concurrent modification collisions. When submitting an update, the request will be rejected unless this value agrees with the version expected by the server. GET the current configuration first, to find out the expected &#x60;&#x60;resource_version&#x60;&#x60;.  | 
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | When the audit log configuration was last modified. This value is set by the server, and ignored when submitted.  | [optional] 
**ModifiedSince** | Pointer to [**time.Time**](time.Time.md) | When the server detects that the audit log configuration has been modified through another mechanism but is not sure when the modification occurred, the &#x60;&#x60;updated_at&#x60;&#x60; property is renamed to &#x60;&#x60;modified_since&#x60;&#x60;. This value is set by the server, and ignored when submitted.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


