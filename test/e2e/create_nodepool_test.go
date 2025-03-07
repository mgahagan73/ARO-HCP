package e2e

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Nodepool operation", func() {
	var (
		NodePoolsClient *api.NodePoolsClient
	)

	BeforeEach(func() {
		By("Prepare HCP nodepools client")
		NodePoolsClient = clients.NewNodePoolsClient()
	})

	var (
		nodePoolName     = "mynodepool"
		nodePoolResource api.HcpOpenShiftClusterNodePoolResource
		nodePoolOptions  *api.HcpOpenShiftClustersClientBeginCreateOrUpdateOptions
	)
	It("Create invalid nodepool", labels.Medium, labels.Negative, func(ctx context.Context) {
		clusterName := "existing_cluster"
		By("Send get request for nodepool")
		_, err := NodePoolsClient.BeginCreateOrUpdate(ctx, customerRGName, clusterName, nodePoolName, nodePoolResource, (*api.NodePoolsClientBeginCreateOrUpdateOptions)(nodePoolOptions))
		Expect(err).ToNot(BeNil())
		//errMessage := fmt.Sprintf("The resource 'hcpOpenShiftClusters/%s' under resource group '%s' was not found.", clusterName, customerRGName)
		errMessage := fmt.Sprintf("RESPONSE 500: 500 Internal Server Error")
		Expect(err.Error()).To(ContainSubstring(errMessage))
	})
})
