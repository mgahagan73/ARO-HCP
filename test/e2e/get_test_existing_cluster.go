package e2e

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/cmdline"
	"github.com/Azure/ARO-HCP/test/util/integration"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Confirm HCPCluster is operational", func() {

	var (
		clustersClient    *api.HcpOpenShiftClustersClient
		clusterKubeconfig string
		customerEnv       *integration.CustomerEnv
		clusterInfo       *integration.Cluster
	)

	BeforeEach(func() {
		By("Prepare cluster client")
		clustersClient = clients.NewHcpOpenShiftClustersClient()
		By("Preparing customer environment values")
		customerEnv = &e2eSetup.CustomerEnv
	})

	It("Confirm cluster has been created successfully", labels.Medium, func(ctx context.Context) {
		By("Checking Provisioning state with RP")
		out, err := clustersClient.Get(ctx, customerEnv.CustomerRGName, clusterInfo.Name, nil)
		Expect(err).To(BeNil())
		Expect(string(*out.Properties.ProvisioningState)).To(Equal(("Succeeded")))
	})
	It("Check access to cluster with oc command using kubeconfig file", labels.Medium, func(ctx context.Context) {
		By("Getting Kubeconfig from RP")
		adminPoller, err := clustersClient.BeginRequestAdminCredential(ctx, customerEnv.CustomerRGName, clusterInfo.Name, nil)
		Expect(err).To(BeNil())
		for !adminPoller.Done() {
			status, err := adminPoller.PollUntilDone(ctx, nil)
			Expect(err).To(BeNil())
			Expect(status).ToNot(BeNil())
		}
		adminCredentialResponse, err := adminPoller.Result(ctx)
		Expect(err).To(BeNil())
		Expect(adminCredentialResponse).ToNot(BeNil())
		clusterKubeconfig = *adminCredentialResponse.Kubeconfig
		By("Attempting to list projects on cluster and confirm default project is present")
		oc_command := "KUBECONFIG=" + clusterKubeconfig + " oc get projects"
		stdout, stderr, err := cmdline.RunCMD(oc_command)
		Expect(err).To(BeNil())
		Expect(stderr).ToNot(ContainSubstring("error"))
		Expect(stdout).To(ContainSubstring("default"))
	})

})
