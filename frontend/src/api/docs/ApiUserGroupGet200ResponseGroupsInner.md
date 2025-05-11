# ApiUserGroupGet200ResponseGroupsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**big_group** | **string** |  | [optional] [default to undefined]
**group_id** | **number** |  | [optional] [default to undefined]
**group_name** | **string** |  | [optional] [default to undefined]
**leader_name** | **string** |  | [optional] [default to undefined]
**leader_student_id** | **string** |  | [optional] [default to undefined]
**members** | [**Array&lt;ApiAdminGroupGet200ResponseGroupsInnerMembersInner&gt;**](ApiAdminGroupGet200ResponseGroupsInnerMembersInner.md) |  | [optional] [default to undefined]
**graded_student_id** | **{ [key: string]: boolean; }** | Map of student IDs who have graded | [optional] [default to undefined]

## Example

```typescript
import { ApiUserGroupGet200ResponseGroupsInner } from './api';

const instance: ApiUserGroupGet200ResponseGroupsInner = {
    big_group,
    group_id,
    group_name,
    leader_name,
    leader_student_id,
    members,
    graded_student_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
