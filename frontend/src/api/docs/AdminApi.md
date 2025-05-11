# AdminApi

All URIs are relative to *http://ce.alonza.xyz:31413*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiAdminGroupBigDelete**](#apiadmingroupbigdelete) | **DELETE** /api/admin/group/big | Delete Big Group|
|[**apiAdminGroupBigGet**](#apiadmingroupbigget) | **GET** /api/admin/group/big | Get Big Groups|
|[**apiAdminGroupBigPost**](#apiadmingroupbigpost) | **POST** /api/admin/group/big | Add Big Group|
|[**apiAdminGroupDelete**](#apiadmingroupdelete) | **DELETE** /api/admin/group | Delete Group|
|[**apiAdminGroupGet**](#apiadmingroupget) | **GET** /api/admin/group | Get Admin Group Information|
|[**apiAdminGroupGradeGet**](#apiadmingroupgradeget) | **GET** /api/admin/group/grade | Get Group Grades|
|[**apiAdminGroupPost**](#apiadmingrouppost) | **POST** /api/admin/group | Add Group|
|[**apiAdminLoginPost**](#apiadminloginpost) | **POST** /api/admin/login | Admin Login|

# **apiAdminGroupBigDelete**
> ApiAdminGroupBigPost200Response apiAdminGroupBigDelete(apiAdminGroupBigPostRequest)


### Example

```typescript
import {
    AdminApi,
    Configuration,
    ApiAdminGroupBigPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let apiAdminGroupBigPostRequest: ApiAdminGroupBigPostRequest; //

const { status, data } = await apiInstance.apiAdminGroupBigDelete(
    apiAdminGroupBigPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminGroupBigPostRequest** | **ApiAdminGroupBigPostRequest**|  | |


### Return type

**ApiAdminGroupBigPost200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Big group deleted successfully |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**404** | Not Found |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminGroupBigGet**
> ApiAdminGroupBigGet200Response apiAdminGroupBigGet()


### Example

```typescript
import {
    AdminApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

const { status, data } = await apiInstance.apiAdminGroupBigGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**ApiAdminGroupBigGet200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successfully retrieved big groups |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminGroupBigPost**
> ApiAdminGroupBigPost200Response apiAdminGroupBigPost(apiAdminGroupBigPostRequest)


### Example

```typescript
import {
    AdminApi,
    Configuration,
    ApiAdminGroupBigPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let apiAdminGroupBigPostRequest: ApiAdminGroupBigPostRequest; //

const { status, data } = await apiInstance.apiAdminGroupBigPost(
    apiAdminGroupBigPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminGroupBigPostRequest** | **ApiAdminGroupBigPostRequest**|  | |


### Return type

**ApiAdminGroupBigPost200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Big group added successfully |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminGroupDelete**
> ApiAdminGroupDelete200Response apiAdminGroupDelete(apiAdminGroupDeleteRequest)


### Example

```typescript
import {
    AdminApi,
    Configuration,
    ApiAdminGroupDeleteRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let apiAdminGroupDeleteRequest: ApiAdminGroupDeleteRequest; //

const { status, data } = await apiInstance.apiAdminGroupDelete(
    apiAdminGroupDeleteRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminGroupDeleteRequest** | **ApiAdminGroupDeleteRequest**|  | |


### Return type

**ApiAdminGroupDelete200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Group deleted successfully |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**404** | Not Found |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminGroupGet**
> ApiAdminGroupGet200Response apiAdminGroupGet()


### Example

```typescript
import {
    AdminApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let bigGroup: string; // (default to undefined)

const { status, data } = await apiInstance.apiAdminGroupGet(
    bigGroup
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **bigGroup** | [**string**] |  | defaults to undefined|


### Return type

**ApiAdminGroupGet200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successfully retrieved group information |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminGroupGradeGet**
> ApiAdminGroupGradeGet200Response apiAdminGroupGradeGet(apiAdminGroupDeleteRequest)


### Example

```typescript
import {
    AdminApi,
    Configuration,
    ApiAdminGroupDeleteRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let apiAdminGroupDeleteRequest: ApiAdminGroupDeleteRequest; //

const { status, data } = await apiInstance.apiAdminGroupGradeGet(
    apiAdminGroupDeleteRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminGroupDeleteRequest** | **ApiAdminGroupDeleteRequest**|  | |


### Return type

**ApiAdminGroupGradeGet200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successfully retrieved group grades |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminGroupPost**
> ApiAdminGroupPost200Response apiAdminGroupPost(apiAdminGroupPostRequest)


### Example

```typescript
import {
    AdminApi,
    Configuration,
    ApiAdminGroupPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let apiAdminGroupPostRequest: ApiAdminGroupPostRequest; //

const { status, data } = await apiInstance.apiAdminGroupPost(
    apiAdminGroupPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminGroupPostRequest** | **ApiAdminGroupPostRequest**|  | |


### Return type

**ApiAdminGroupPost200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Group added successfully |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiAdminLoginPost**
> ApiAdminLoginPost200Response apiAdminLoginPost(apiAdminLoginPostRequest)


### Example

```typescript
import {
    AdminApi,
    Configuration,
    ApiAdminLoginPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new AdminApi(configuration);

let apiAdminLoginPostRequest: ApiAdminLoginPostRequest; //

const { status, data } = await apiInstance.apiAdminLoginPost(
    apiAdminLoginPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminLoginPostRequest** | **ApiAdminLoginPostRequest**|  | |


### Return type

**ApiAdminLoginPost200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Login successful |  -  |
|**400** | Bad request |  -  |
|**401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

