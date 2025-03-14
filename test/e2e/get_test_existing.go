package e2e

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Get operation", func() {
	var (
		clustersClient *api.HcpOpenShiftClustersClient
	)

	BeforeEach(func() {
		By("Prepare HCP clusters client")
		clustersClient = clients.NewHcpOpenShiftClustersClient()
	})

	It("Get existing cluster", labels.Medium, labels.Positive, func(ctx context.Context) {
		clusterName := "mgahaganhcp"
		By("Send get request for cluster")
		out, err := clustersClient.Get(ctx, customerRGName, clusterName, nil)
		Expect(err).To(BeNil())
		Expect(string(*out.Properties.ProvisioningState)).To(Equal(("Succeeded")))
	})
})
