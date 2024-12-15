package v1

import (
	"bytes"
	"context"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("NamespaceValidator Webhook", func() {
	var (
		testNamespace string
		logBuffer     *bytes.Buffer
	)

	BeforeEach(func() {
		testNamespace = "test-namespace"
		logBuffer = &bytes.Buffer{}
		log.SetOutput(logBuffer) // Redirect logs to buffer
	})

	AfterEach(func() {
		log.SetOutput(GinkgoWriter) // Reset log output
	})

	Context("when performing CRUD operations on namespaces", func() {
		It("should log namespace actions", func() {
			By("creating a namespace")
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: testNamespace,
				},
			}
			Expect(k8sClient.Create(context.TODO(), ns)).To(Succeed())

			createdNS := &corev1.Namespace{}
			Expect(k8sClient.Get(context.TODO(), client.ObjectKey{Name: testNamespace}, createdNS)).To(Succeed())

			By("updating the namespace")
			createdNS.Labels = map[string]string{"env": "test"}
			Expect(k8sClient.Update(context.TODO(), createdNS)).To(Succeed())

			By("deleting the namespace")
			Expect(k8sClient.Delete(context.TODO(), createdNS)).To(Succeed())

			Eventually(func() error {
				return k8sClient.Get(context.TODO(), client.ObjectKey{Name: testNamespace}, &corev1.Namespace{})
			}).ShouldNot(Succeed())

			By("verifying logs")
			Expect(logBuffer.String()).To(ContainSubstring("namespace action"))
			Expect(logBuffer.String()).To(ContainSubstring(testNamespace))
		})
	})
})
