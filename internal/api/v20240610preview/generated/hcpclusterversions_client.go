//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.3, generator: @autorest/go@4.0.0-preview.63)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HcpClusterVersionsClient contains the methods for the HcpClusterVersions group.
// Don't use this type directly, use NewHcpClusterVersionsClient() instead.
type HcpClusterVersionsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewHcpClusterVersionsClient creates a new instance of HcpClusterVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHcpClusterVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HcpClusterVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HcpClusterVersionsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// NewListPager - List HcpOpenShiftVersionResource resources by location
//
// Generated from API version 2024-06-10-preview
//   - location - The name of the Azure region.
//   - options - HcpClusterVersionsClientListOptions contains the optional parameters for the HcpClusterVersionsClient.NewListPager
//     method.
func (client *HcpClusterVersionsClient) NewListPager(location string, options *HcpClusterVersionsClientListOptions) (*runtime.Pager[HcpClusterVersionsClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[HcpClusterVersionsClientListResponse]{
		More: func(page HcpClusterVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HcpClusterVersionsClientListResponse) (HcpClusterVersionsClientListResponse, error) {
		ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HcpClusterVersionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return HcpClusterVersionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
			},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HcpClusterVersionsClient) listCreateRequest(ctx context.Context, location string, options *HcpClusterVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/locations/{location}/providers/Microsoft.RedHatOpenShift/hcpOpenShiftVersions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// listHandleResponse handles the List response.
func (client *HcpClusterVersionsClient) listHandleResponse(resp *http.Response) (HcpClusterVersionsClientListResponse, error) {
	result := HcpClusterVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HcpOpenShiftVersionResourceListResult); err != nil {
		return HcpClusterVersionsClientListResponse{}, err
	}
	return result, nil
}

