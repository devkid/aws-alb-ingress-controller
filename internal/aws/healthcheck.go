package aws

import (
	"net/http"

	"github.com/golang/glog"
	"k8s.io/apiserver/pkg/server/healthz"
)

type AWSHealthChecker struct{}

var _ healthz.HealthzChecker = (*AWSHealthChecker)(nil)

func (c *AWSHealthChecker) Name() string {
	// TODO, can we rename it to, e.g. AWS? is there any dependencies on the name?
	return "aws-alb-ingress-controller"
}

// TODO, validate the call health check frequency
func (c *AWSHealthChecker) Check(_ *http.Request) error {
	for _, fn := range []func() error{
		Cloudsvc.StatusACM(),
		Cloudsvc.StatusEC2(),
		Cloudsvc.StatusIAM(),
	} {
		err := fn()
		if err != nil {
			glog.Errorf("Controller health check failed: %v", err.Error())
			return err
		}
	}
	return nil
}
