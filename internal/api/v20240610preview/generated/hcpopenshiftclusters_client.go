//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.3, generator: @autorest/go@4.0.0-preview.63)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// HcpOpenShiftClustersClient contains the methods for the HcpOpenShiftClusters group.
// Don't use this type directly, use NewHcpOpenShiftClustersClient() instead.
type HcpOpenShiftClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHcpOpenShiftClustersClient creates a new instance of HcpOpenShiftClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHcpOpenShiftClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HcpOpenShiftClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HcpOpenShiftClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hcpOpenShiftClusterName - Name of HCP cluster
//   - resource - Resource create parameters.
//   - options - HcpOpenShiftClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the HcpOpenShiftClustersClient.BeginCreateOrUpdate
//     method.
func (client *HcpOpenShiftClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, resource HcpOpenShiftClusterResource, options *HcpOpenShiftClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[HcpOpenShiftClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hcpOpenShiftClusterName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HcpOpenShiftClustersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HcpOpenShiftClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
func (client *HcpOpenShiftClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, resource HcpOpenShiftClusterResource, options *HcpOpenShiftClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "HcpOpenShiftClustersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hcpOpenShiftClusterName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HcpOpenShiftClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, resource HcpOpenShiftClusterResource, options *HcpOpenShiftClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftClusters/{hcpOpenShiftClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hcpOpenShiftClusterName == "" {
		return nil, errors.New("parameter hcpOpenShiftClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcpOpenShiftClusterName}", url.PathEscape(hcpOpenShiftClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hcpOpenShiftClusterName - Name of HCP cluster
//   - options - HcpOpenShiftClustersClientBeginDeleteOptions contains the optional parameters for the HcpOpenShiftClustersClient.BeginDelete
//     method.
func (client *HcpOpenShiftClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, options *HcpOpenShiftClustersClientBeginDeleteOptions) (*runtime.Poller[HcpOpenShiftClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hcpOpenShiftClusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HcpOpenShiftClustersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HcpOpenShiftClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
func (client *HcpOpenShiftClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, options *HcpOpenShiftClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "HcpOpenShiftClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hcpOpenShiftClusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HcpOpenShiftClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, options *HcpOpenShiftClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftClusters/{hcpOpenShiftClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hcpOpenShiftClusterName == "" {
		return nil, errors.New("parameter hcpOpenShiftClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcpOpenShiftClusterName}", url.PathEscape(hcpOpenShiftClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hcpOpenShiftClusterName - Name of HCP cluster
//   - options - HcpOpenShiftClustersClientGetOptions contains the optional parameters for the HcpOpenShiftClustersClient.Get
//     method.
func (client *HcpOpenShiftClustersClient) Get(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, options *HcpOpenShiftClustersClientGetOptions) (HcpOpenShiftClustersClientGetResponse, error) {
	var err error
	const operationName = "HcpOpenShiftClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hcpOpenShiftClusterName, options)
	if err != nil {
		return HcpOpenShiftClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HcpOpenShiftClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HcpOpenShiftClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HcpOpenShiftClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, options *HcpOpenShiftClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftClusters/{hcpOpenShiftClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hcpOpenShiftClusterName == "" {
		return nil, errors.New("parameter hcpOpenShiftClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcpOpenShiftClusterName}", url.PathEscape(hcpOpenShiftClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HcpOpenShiftClustersClient) getHandleResponse(resp *http.Response) (HcpOpenShiftClustersClientGetResponse, error) {
	result := HcpOpenShiftClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcpOpenShiftClusterResource); err != nil {
		return HcpOpenShiftClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List HcpOpenShiftClusterResource resources by resource group
//
// Generated from API version 2024-06-10-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - HcpOpenShiftClustersClientListByResourceGroupOptions contains the optional parameters for the HcpOpenShiftClustersClient.NewListByResourceGroupPager
//     method.
func (client *HcpOpenShiftClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *HcpOpenShiftClustersClientListByResourceGroupOptions) *runtime.Pager[HcpOpenShiftClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[HcpOpenShiftClustersClientListByResourceGroupResponse]{
		More: func(page HcpOpenShiftClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HcpOpenShiftClustersClientListByResourceGroupResponse) (HcpOpenShiftClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HcpOpenShiftClustersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return HcpOpenShiftClustersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *HcpOpenShiftClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *HcpOpenShiftClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *HcpOpenShiftClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (HcpOpenShiftClustersClientListByResourceGroupResponse, error) {
	result := HcpOpenShiftClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcpOpenShiftClusterResourceListResult); err != nil {
		return HcpOpenShiftClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List HcpOpenShiftClusterResource resources by subscription ID
//
// Generated from API version 2024-06-10-preview
//   - options - HcpOpenShiftClustersClientListBySubscriptionOptions contains the optional parameters for the HcpOpenShiftClustersClient.NewListBySubscriptionPager
//     method.
func (client *HcpOpenShiftClustersClient) NewListBySubscriptionPager(options *HcpOpenShiftClustersClientListBySubscriptionOptions) *runtime.Pager[HcpOpenShiftClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[HcpOpenShiftClustersClientListBySubscriptionResponse]{
		More: func(page HcpOpenShiftClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HcpOpenShiftClustersClientListBySubscriptionResponse) (HcpOpenShiftClustersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HcpOpenShiftClustersClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return HcpOpenShiftClustersClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *HcpOpenShiftClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *HcpOpenShiftClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *HcpOpenShiftClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (HcpOpenShiftClustersClientListBySubscriptionResponse, error) {
	result := HcpOpenShiftClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcpOpenShiftClusterResourceListResult); err != nil {
		return HcpOpenShiftClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hcpOpenShiftClusterName - Name of HCP cluster
//   - properties - The resource properties to be updated.
//   - options - HcpOpenShiftClustersClientBeginUpdateOptions contains the optional parameters for the HcpOpenShiftClustersClient.BeginUpdate
//     method.
func (client *HcpOpenShiftClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, properties HcpOpenShiftClusterPatch, options *HcpOpenShiftClustersClientBeginUpdateOptions) (*runtime.Poller[HcpOpenShiftClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, hcpOpenShiftClusterName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HcpOpenShiftClustersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HcpOpenShiftClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a HcpOpenShiftClusterResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-10-preview
func (client *HcpOpenShiftClustersClient) update(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, properties HcpOpenShiftClusterPatch, options *HcpOpenShiftClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "HcpOpenShiftClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hcpOpenShiftClusterName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *HcpOpenShiftClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hcpOpenShiftClusterName string, properties HcpOpenShiftClusterPatch, options *HcpOpenShiftClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftClusters/{hcpOpenShiftClusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hcpOpenShiftClusterName == "" {
		return nil, errors.New("parameter hcpOpenShiftClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcpOpenShiftClusterName}", url.PathEscape(hcpOpenShiftClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
