package e2e

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "github.com/Azure/ARO-HCP/internal/api/v20240610preview/generated"
	"github.com/Azure/ARO-HCP/test/util/cmdline"
	"github.com/Azure/ARO-HCP/test/util/labels"
)

var _ = Describe("Checking to see if specified cluster is available", func() {
	var (
		clustersClient    *api.HcpOpenShiftClustersClient
		ClusterName       = os.Getenv("CLUSTER_NAME")
		kubeadminPassword = os.Getenv("KUBEADMIN_PASSWORD")
	)

	BeforeEach(func() {
		By("Prepare HCP clusters client")
		clustersClient = clients.NewHcpOpenShiftClustersClient()
	})

	It("Get existing cluster from RP", labels.Medium, labels.Positive, func(ctx context.Context) {
		By("Send get request for cluster to cluster RP to check ProvisioningState")
		out, err := clustersClient.Get(ctx, customerRGName, ClusterName, nil)
		Expect(err).To(BeNil())
		Expect(string(*out.Properties.ProvisioningState)).To(Equal(("Succeeded")))
	})
	It("Try to login to the cluster using the kubeadmin password", labels.Medium, labels.Positive, func(ctx context.Context) {
		By("Get cluster api URL from RP")
		out, err := clustersClient.Get(ctx, customerRGName, ClusterName, nil)
		Expect(err).To(BeNil())
		clusterApiUrl := string(*out.Properties.API.URL)
		Expect(clusterApiUrl).To(ContainSubstring("https://"))
		By("Log in with oc command")
		//oc_command := fmt.Sprintf("oc login -u kubeadmin -p " + kubeadminPassword + " " + clusterApiUrl)
		oc_command := "oc --insecure-skip-tls-verify login -u kubeadmin -p " + kubeadminPassword + " " + clusterApiUrl
		stdout, stderr, err := cmdline.RunCMD(oc_command)
		Expect(stderr).ToNot(ContainSubstring("error"))
		Expect(stdout).To(ContainSubstring("Login successful."))
		Expect(err).To(BeNil())
	})
})
