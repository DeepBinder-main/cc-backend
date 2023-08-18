# \SlurmAPI

All URIs are relative to *http://localhost/slurm/v0.0.37*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmctldCancelJob**](SlurmAPI.md#SlurmctldCancelJob) | **Delete** /job/{job_id} | cancel or signal job
[**SlurmctldDiag**](SlurmAPI.md#SlurmctldDiag) | **Get** /diag/ | get diagnostics
[**SlurmctldGetJob**](SlurmAPI.md#SlurmctldGetJob) | **Get** /job/{job_id} | get job info
[**SlurmctldGetJobs**](SlurmAPI.md#SlurmctldGetJobs) | **Get** /jobs/ | get list of jobs
[**SlurmctldGetNode**](SlurmAPI.md#SlurmctldGetNode) | **Get** /node/{node_name} | get node info
[**SlurmctldGetNodes**](SlurmAPI.md#SlurmctldGetNodes) | **Get** /nodes/ | get all node info
[**SlurmctldGetPartition**](SlurmAPI.md#SlurmctldGetPartition) | **Get** /partition/{partition_name} | get partition info
[**SlurmctldGetPartitions**](SlurmAPI.md#SlurmctldGetPartitions) | **Get** /partitions/ | get all partition info
[**SlurmctldGetReservation**](SlurmAPI.md#SlurmctldGetReservation) | **Get** /reservation/{reservation_name} | get reservation info
[**SlurmctldGetReservations**](SlurmAPI.md#SlurmctldGetReservations) | **Get** /reservations/ | get all reservation info
[**SlurmctldPing**](SlurmAPI.md#SlurmctldPing) | **Get** /ping/ | ping test
[**SlurmctldSubmitJob**](SlurmAPI.md#SlurmctldSubmitJob) | **Post** /job/submit | submit new job
[**SlurmctldUpdateJob**](SlurmAPI.md#SlurmctldUpdateJob) | **Post** /job/{job_id} | update job



## SlurmctldCancelJob

> SlurmctldCancelJob(ctx, jobId).Signal(signal).Execute()

cancel or signal job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    jobId := "jobId_example" // string | Slurm Job ID
    signal := openapiclient.v0.0.37_signal("HUP") // V0037Signal | signal to send to job (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmAPI.SlurmctldCancelJob(context.Background(), jobId).Signal(signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldCancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | [**V0037Signal**](V0037Signal.md) | signal to send to job | 

### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldDiag

> V0037Diag SlurmctldDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldDiag`: V0037Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldDiagRequest struct via the builder pattern


### Return type

[**V0037Diag**](V0037Diag.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetJob

> V0037JobsResponse SlurmctldGetJob(ctx, jobId).Execute()

get job info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    jobId := "jobId_example" // string | Slurm JobID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJob`: V0037JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm JobID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0037JobsResponse**](V0037JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetJobs

> V0037JobsResponse SlurmctldGetJobs(ctx).UpdateTime(updateTime).Execute()

get list of jobs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetJobs(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJobs`: V0037JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037JobsResponse**](V0037JobsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetNode

> V0037NodesResponse SlurmctldGetNode(ctx, nodeName).Execute()

get node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    nodeName := "nodeName_example" // string | Slurm Node Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNode`: V0037NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0037NodesResponse**](V0037NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetNodes

> V0037NodesResponse SlurmctldGetNodes(ctx).UpdateTime(updateTime).Execute()

get all node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetNodes(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNodes`: V0037NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037NodesResponse**](V0037NodesResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetPartition

> V0037PartitionsResponse SlurmctldGetPartition(ctx, partitionName).UpdateTime(updateTime).Execute()

get partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    partitionName := "partitionName_example" // string | Slurm Partition Name
    updateTime := int64(789) // int64 | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetPartition(context.Background(), partitionName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartition`: V0037PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if there were no partition changes (not limited to partition in URL endpoint) since update_time. | 

### Return type

[**V0037PartitionsResponse**](V0037PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetPartitions

> V0037PartitionsResponse SlurmctldGetPartitions(ctx).UpdateTime(updateTime).Execute()

get all partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetPartitions(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartitions`: V0037PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037PartitionsResponse**](V0037PartitionsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetReservation

> V0037ReservationsResponse SlurmctldGetReservation(ctx, reservationName).UpdateTime(updateTime).Execute()

get reservation info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    reservationName := "reservationName_example" // string | Slurm Reservation Name
    updateTime := int64(789) // int64 | Filter if no reservation (not limited to reservation in URL) changed since update_time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetReservation(context.Background(), reservationName).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetReservation`: V0037ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservationName** | **string** | Slurm Reservation Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTime** | **int64** | Filter if no reservation (not limited to reservation in URL) changed since update_time. | 

### Return type

[**V0037ReservationsResponse**](V0037ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetReservations

> V0037ReservationsResponse SlurmctldGetReservations(ctx).UpdateTime(updateTime).Execute()

get all reservation info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateTime := int64(789) // int64 | Filter if changed since update_time. Use of this parameter can result in faster replies. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldGetReservations(context.Background()).UpdateTime(updateTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldGetReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetReservations`: V0037ReservationsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldGetReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **int64** | Filter if changed since update_time. Use of this parameter can result in faster replies. | 

### Return type

[**V0037ReservationsResponse**](V0037ReservationsResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldPing

> V0037Pings SlurmctldPing(ctx).Execute()

ping test

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldPing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldPing`: V0037Pings
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldPingRequest struct via the builder pattern


### Return type

[**V0037Pings**](V0037Pings.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldSubmitJob

> V0037JobSubmissionResponse SlurmctldSubmitJob(ctx).V0037JobSubmission(v0037JobSubmission).Execute()

submit new job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    v0037JobSubmission := *openapiclient.NewV0037JobSubmission("Script_example") // V0037JobSubmission | submit new job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlurmAPI.SlurmctldSubmitJob(context.Background()).V0037JobSubmission(v0037JobSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldSubmitJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldSubmitJob`: V0037JobSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmAPI.SlurmctldSubmitJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldSubmitJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0037JobSubmission** | [**V0037JobSubmission**](V0037JobSubmission.md) | submit new job | 

### Return type

[**V0037JobSubmissionResponse**](V0037JobSubmissionResponse.md)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldUpdateJob

> SlurmctldUpdateJob(ctx, jobId).V0037JobProperties(v0037JobProperties).Execute()

update job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    jobId := "jobId_example" // string | Slurm Job ID
    v0037JobProperties := *openapiclient.NewV0037JobProperties(map[string]interface{}(123)) // V0037JobProperties | update job

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlurmAPI.SlurmctldUpdateJob(context.Background(), jobId).V0037JobProperties(v0037JobProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmAPI.SlurmctldUpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldUpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0037JobProperties** | [**V0037JobProperties**](V0037JobProperties.md) | update job | 

### Return type

 (empty response body)

### Authorization

[user](../README.md#user), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

