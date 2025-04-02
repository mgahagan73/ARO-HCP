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
	clusterKubeconfig := os.Getenv("CLUSTER_KUBECONFIG")
	clusterApiUrl := ""

	It("Confirm cluster has been created successfully", labels.Medium, labels.Positive, func(ctx context.Context) {
		By("Checking Provisioning state with RP")
		out, err := clustersClient.Get(ctx, customerRGName, clusterName, nil)
		Expect(err).To(BeNil())
		Expect(string(*out.Properties.ProvisioningState)).To(Equal(("Succeeded")))
	})
	It("Log in to existing cluster with kubeadmin password", labels.Medium, labels.Positive, func(ctx context.Context) {
		By("Get cluster api URL from RP")
		Expect(kubeadminPassword).ToNot(BeNil())
		out, err := clustersClient.Get(ctx, customerRGName, clusterName, nil)
		Expect(err).To(BeNil())
		clusterApiUrl = string(*out.Properties.API.URL)
		Expect(clusterApiUrl).To(ContainSubstring("https://"))
		By("Log in with oc command")
		//oc_command := fmt.Sprintf("oc login -u kubeadmin -p " + kubeadminPassword + " " + clusterApiUrl)
		oc_command := "oc --insecure-skip-tls-verify login -u kubeadmin -p " + kubeadminPassword + " " + clusterApiUrl
		stdout, stderr, err := cmdline.RunCMD(oc_command)
		Expect(err).To(BeNil())
		Expect(stderr).ToNot(ContainSubstring("error"))
		Expect(stdout).To(ContainSubstring("Login successful."))
	})
	It("Can access cluster with oc command using kubeconfig file", labels.Medium, labels.Positive, func(ctx context.Context) {
		Expect(clusterKubeconfig).ToNot(BeNil())
		By("Get projects on cluster")
		oc_command := "KUBECONFIG=" + clusterKubeconfig + " oc get projects"
		stdout, stderr, err := cmdline.RunCMD(oc_command)
		Expect(err).To(BeNil())
		Expect(stderr).ToNot(ContainSubstring("error"))
		Expect(stdout).To(ContainSubstring("default"))
	})

})
