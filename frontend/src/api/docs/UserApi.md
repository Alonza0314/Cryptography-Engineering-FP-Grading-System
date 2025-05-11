# UserApi

All URIs are relative to *http://ce.alonza.xyz:31413*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**apiUserGroupBigGet**](#apiusergroupbigget) | **GET** /api/user/group/big | Get User Big Groups|
|[**apiUserGroupGet**](#apiusergroupget) | **GET** /api/user/group | Get User Groups|
|[**apiUserGroupGradePost**](#apiusergroupgradepost) | **POST** /api/user/group/grade | Add Group Grade|
|[**apiUserLoginPost**](#apiuserloginpost) | **POST** /api/user/login | User Login|
|[**apiUserTotpInitBeginGet**](#apiusertotpinitbeginget) | **GET** /api/user/totp/init/begin | Begin TOTP Initialization|
|[**apiUserTotpInitFinishPost**](#apiusertotpinitfinishpost) | **POST** /api/user/totp/init/finish | Finish TOTP Initialization|
|[**apiUserTotpPost**](#apiusertotppost) | **POST** /api/user/totp | Verify TOTP Code|

# **apiUserGroupBigGet**
> ApiAdminGroupBigGet200Response apiUserGroupBigGet()


### Example

```typescript
import {
    UserApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

const { status, data } = await apiInstance.apiUserGroupBigGet();
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
|**200** | Successfully retrieved user big groups |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiUserGroupGet**
> ApiUserGroupGet200Response apiUserGroupGet(apiAdminGroupBigPostRequest)


### Example

```typescript
import {
    UserApi,
    Configuration,
    ApiAdminGroupBigPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

let apiAdminGroupBigPostRequest: ApiAdminGroupBigPostRequest; //

const { status, data } = await apiInstance.apiUserGroupGet(
    apiAdminGroupBigPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiAdminGroupBigPostRequest** | **ApiAdminGroupBigPostRequest**|  | |


### Return type

**ApiUserGroupGet200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successfully retrieved user groups |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiUserGroupGradePost**
> ApiAdminGroupBigPost200Response apiUserGroupGradePost(apiUserGroupGradePostRequest)


### Example

```typescript
import {
    UserApi,
    Configuration,
    ApiUserGroupGradePostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

let apiUserGroupGradePostRequest: ApiUserGroupGradePostRequest; //

const { status, data } = await apiInstance.apiUserGroupGradePost(
    apiUserGroupGradePostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiUserGroupGradePostRequest** | **ApiUserGroupGradePostRequest**|  | |


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
|**200** | Grade added successfully |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiUserLoginPost**
> ApiAdminLoginPost200Response apiUserLoginPost(apiUserLoginPostRequest)


### Example

```typescript
import {
    UserApi,
    Configuration,
    ApiUserLoginPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

let apiUserLoginPostRequest: ApiUserLoginPostRequest; //

const { status, data } = await apiInstance.apiUserLoginPost(
    apiUserLoginPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiUserLoginPostRequest** | **ApiUserLoginPostRequest**|  | |


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
|**400** | Bad Request |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiUserTotpInitBeginGet**
> ApiUserTotpInitBeginGet200Response apiUserTotpInitBeginGet()


### Example

```typescript
import {
    UserApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

const { status, data } = await apiInstance.apiUserTotpInitBeginGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**ApiUserTotpInitBeginGet200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | TOTP initialization started successfully |  -  |
|**401** | Unauthorized |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiUserTotpInitFinishPost**
> ApiAdminGroupBigPost200Response apiUserTotpInitFinishPost(apiUserTotpInitFinishPostRequest)


### Example

```typescript
import {
    UserApi,
    Configuration,
    ApiUserTotpInitFinishPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

let apiUserTotpInitFinishPostRequest: ApiUserTotpInitFinishPostRequest; //

const { status, data } = await apiInstance.apiUserTotpInitFinishPost(
    apiUserTotpInitFinishPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiUserTotpInitFinishPostRequest** | **ApiUserTotpInitFinishPostRequest**|  | |


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
|**200** | TOTP initialization completed |  -  |
|**400** | Bad Request |  -  |
|**401** | Unauthorized |  -  |
|**404** | Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **apiUserTotpPost**
> ApiAdminLoginPost200Response apiUserTotpPost(apiUserTotpPostRequest)


### Example

```typescript
import {
    UserApi,
    Configuration,
    ApiUserTotpPostRequest
} from './api';

const configuration = new Configuration();
const apiInstance = new UserApi(configuration);

let apiUserTotpPostRequest: ApiUserTotpPostRequest; //

const { status, data } = await apiInstance.apiUserTotpPost(
    apiUserTotpPostRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **apiUserTotpPostRequest** | **ApiUserTotpPostRequest**|  | |


### Return type

**ApiAdminLoginPost200Response**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | TOTP verification successful |  -  |
|**404** | Not Found |  -  |
|**500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

