package e2e

import (
	"context"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/cmdline"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Cluster login operation", func() {

	var (
		clustersClient *api.HcpOpenShiftClustersClient
	)

	BeforeEach(func() {
		By("Prepare cluster client")
		clustersClient = clients.NewHcpOpenShiftClustersClient()
	})

	clusterName := os.Getenv("CLUSTER_NAME")
	kubeadminPassword := os.Getenv("KUBEADMIN_PASSWORD")
	//kubeconfig := os.Getenv("KUBECONFIG")

	It("Log in to existing cluster with kubeadmin password", labels.Medium, labels.Positive, func(ctx context.Context) {
		By("Get cluster api URL from RP")
		out, err := clustersClient.Get(ctx, customerRGName, clusterName, nil)
		Expect(err).To(BeNil())
		Expect(string(*out.Properties.ProvisioningState)).To(Equal(("Succeeded")))
		clusterApiUrl := string(*out.Properties.api.URL)
		Expect(clusterApiUrl).To(ContainSubstring("https//"))
		By("Log in with oc command")
		oc_command := fmt.Sprintf("oc login -u kubeadmin -p " + kubeadminPassword + " " + clusterApiUrl)
		stdout, _, err := cmdline.RunCMD(oc_command)
		Expect(err).To(BeNil())
		Expect(stdout).To(ContainSubstring("blah blah"))
	})

})
