package e2e

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Create Cluster operation", func() {
	var (
		clustersClient *api.HcpOpenShiftClustersClient
	)

	BeforeEach(func() {
		By("Prepare HCP clusters client")
		clustersClient = clients.NewHcpOpenShiftClustersClient()
	})

	It("Attempt to create cluster with non-existant Resource Group", labels.Medium, labels.Negative, func(ctx context.Context) {
		clusterName := "non-existing-cluster"
		customerRGName := "non-existing-group"
		var (
			clusterResource api.HcpOpenShiftClusterResource
			clusterOptions  *api.HcpOpenShiftClustersClientBeginCreateOrUpdateOptions
		)
		By("Send request to create cluster")
		_, err := clustersClient.BeginCreateOrUpdate(ctx, customerRGName, clusterName, clusterResource, clusterOptions)
		// _, err := clustersClient.Get(ctx, customerRGName, clusterName, nil)
		Expect(err).ToNot(BeNil())
		errMessage := fmt.Sprintf("The resource 'hcpOpenShiftClusters/%s' under resource group '%s' was not found.", clusterName, customerRGName)
		Expect(err.Error()).To(ContainSubstring(errMessage))
	})
})
