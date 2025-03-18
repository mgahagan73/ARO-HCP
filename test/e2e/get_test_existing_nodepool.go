package e2e

import (
	"context"
	//"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/labels"
	"github.com/Azure/ARO-HCP/test/util/log"
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
		//nodePoolName = "mgnodes1"
		//nodePoolResource api.HcpOpenShiftClusterNodePoolResource
		nodePoolOptions *api.NodePoolsClientListByParentOptions
	)
	It("List nodepools", labels.Medium, labels.Negative, func(ctx context.Context) {
		clusterName := "mgahagannp"
		By("Send get request for nodepool")
		nodePoolList := NodePoolsClient.NewListByParentPager(customerRGName, clusterName, nodePoolOptions)
		Expect(nodePoolList).ToNot(BeNil())
		//errMessage := fmt.Sprintf("The resource 'hcpOpenShiftClusters/%s' under resource group '%s' was not found.", clusterName, customerRGName)
		By("List all nodepools")
		for nodePoolList.More() {
			nodePools, err := nodePoolList.NextPage(ctx)
			Expect(err).To(BeNil())
			log.Logger.Infoln("Number of nodePools:", len(nodePools.Value))
			Expect(nodePools.HcpOpenShiftClusterNodePoolResourceListResult.MarshalJSON()).Should(MatchJSON(`{"name":"mgnodes1"}`))
		}
	})
})
