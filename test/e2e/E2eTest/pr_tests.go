package E2eTest

import (
	"context"
	opv1 "github.com/hwameistor/hwameistor-operator/api/v1alpha1"
	"github.com/hwameistor/hwameistor-operator/test/e2e/framework"
	"github.com/hwameistor/hwameistor-operator/test/e2e/utils"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

var _ = ginkgo.Describe("pr test ", ginkgo.Ordered, ginkgo.Label("pr-e2e"), func() {
	var f *framework.Framework
	var client ctrlclient.Client
	ctx := context.TODO()
	f = framework.NewDefaultFramework(opv1.AddToScheme)
	client = f.GetClient()
	ginkgo.It("install hwameistor-operator", func() {
		err := wait.PollImmediate(10*time.Second, 20*time.Minute, func() (done bool, err error) {
			output := utils.RunInLinux("kubectl get pod -A  |grep -v Running |wc -l")
			if output != "1\n" {
				return false, nil
			} else {
				logrus.Info("k8s ready")
				return true, nil
			}

		})
		if err != nil {
			logrus.Error(err)
		}
		logrus.Infof("helm install hwameistor-operator")
		_ = utils.RunInLinux("helm install hwameistor-operator -n hwameistor-operator ../../helm/operator --create-namespace ")

		Operator := &appsv1.Deployment{}
		OperatorKey := k8sclient.ObjectKey{
			Name:      "hwameistor-operator",
			Namespace: "default",
		}
		err = wait.PollImmediate(3*time.Second, 20*time.Minute, func() (done bool, err error) {
			err = client.Get(ctx, OperatorKey, Operator)
			if err != nil {
				logrus.Error(err)
			}
			if Operator.Status.AvailableReplicas == int32(1) {
				logrus.Infof("hwameistor-operator ready")
				return true, nil

			}
			return false, nil
		})
		if err != nil {
			logrus.Error(err)
		}
		gomega.Expect(err).To(gomega.BeNil())

	})
	ginkgo.Context("create a hmcluster", func() {
		ginkgo.It("create a hmcluster", func() {
			//create sc
			exampleCluster := &opv1.Cluster{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Cluster",
					APIVersion: "hwameistor.io/v1alpha1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "cluster-sample",
				},
				Spec:   opv1.ClusterSpec{},
				Status: opv1.ClusterStatus{},
			}

			err := client.Create(ctx, exampleCluster)
			if err != nil {
				logrus.Printf("Create hmcluster failed ：%+v ", err)
				f.ExpectNoError(err)
			}
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

})