package e2e

import (
	"context"
	"os"

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
		nodePoolName    = os.Getenv("NP_NAME")
		clusterName     = os.Getenv("CLUSTER_NAME")
		nodePoolOptions *api.NodePoolsClientListByParentOptions
	)

	It("Get nodepool from cluster", labels.Medium, func(ctx context.Context) {
		By("Send get request for nodepool")
		nodePool, err := NodePoolsClient.Get(ctx, customerRGName, clusterName, nodePoolName, (*api.NodePoolsClientGetOptions)(nodePoolOptions))
		Expect(err).To(BeNil())
		Expect(nodePool).ToNot(BeNil())
		By("Check to see nodepool exists and is successfully provisioned")
		Expect(string(*nodePool.Name)).To(Equal(nodePoolName))
		Expect(string(*nodePool.Properties.ProvisioningState)).To(Equal("Succeeded"))
	})
})
