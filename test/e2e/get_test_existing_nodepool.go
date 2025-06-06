package e2e

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/integration"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Nodepool operation", func() {
	var (
		NodePoolsClient *api.NodePoolsClient
		//nodePoolName    = os.Getenv("NP_NAME")
		clusterName     = os.Getenv("CLUSTER_NAME")
		nodePoolOptions *api.NodePoolsClientGetOptions
		customerEnv     *integration.CustomerEnv
		nodePools       *[]integration.Nodepool
	)

	BeforeEach(func() {
		By("Prepare HCP nodepools client")
		NodePoolsClient = clients.NewNodePoolsClient()
		By("Preparing customer environment values")
		customerEnv = &e2eSetup.CustomerEnv
		nodePools = &e2eSetup.Nodepools
	})

	It("Get each nodepool from cluster", labels.Medium, func(ctx context.Context) {
		if nodePools != nil {
			nps := *nodePools
			for np := range nps {
				By("Send get request for nodepool")
				clusterNodePool, err := NodePoolsClient.Get(ctx, customerEnv.CustomerRGName, clusterName, nps[np].Name, nodePoolOptions)
				Expect(err).To(BeNil())
				Expect(clusterNodePool).ToNot(BeNil())
				By("Check to see nodepool exists and is successfully provisioned")
				Expect(string(*clusterNodePool.Name)).To(Equal(nps[np].Name))
				Expect(string(*clusterNodePool.Properties.ProvisioningState)).To(Equal("Succeeded"))
			}
		}
	})
})
