package tg

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/ingress/controller/dummy"
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/mocks"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Test_NewTargets(t *testing.T) {
	for _, tc := range []struct {
		name       string
		targetType string
		ingress    *extensions.Ingress
		backend    *extensions.IngressBackend
		output     *Targets
	}{
		{
			name:       "std params",
			targetType: string(corev1.ServiceTypeNodePort),
			ingress:    &extensions.Ingress{},
			backend:    &extensions.IngressBackend{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			output := NewTargets(tc.targetType, tc.ingress, tc.backend)

			assert.Equal(t, output.TargetType, tc.targetType)
			assert.Equal(t, output.Ingress, tc.ingress)
			assert.Equal(t, output.Backend, tc.backend)
		})
	}
}

func newTd(id string, port int64) *elbv2.TargetDescription {
	td := &elbv2.TargetDescription{
		Id:   aws.String(id),
		Port: aws.Int64(port),
	}
	return td
}

func newTh(state string) *elbv2.TargetHealth {
	return &elbv2.TargetHealth{
		State: aws.String(state),
	}
}

type DescribeTargetHealthCall struct {
	TgArn  string
	Output *elbv2.DescribeTargetHealthOutput
	Err    error
}

type DeregisterTargetsCall struct {
	Input *elbv2.DeregisterTargetsInput
	Err   error
}

type RegisterTargetsCall struct {
	Input *elbv2.RegisterTargetsInput
	Err   error
}

type ResolveCall struct {
	InputIngress    *extensions.Ingress
	InputBackend    *extensions.IngressBackend
	InputTargetType string
	Output          []*elbv2.TargetDescription
	Err             error
}

func Test_TargetsReconcile(t *testing.T) {
	tgArn := "arn:"
	serviceName := "name"
	servicePort := intstr.FromInt(123)
	backend := &extensions.IngressBackend{ServiceName: serviceName, ServicePort: servicePort}

	for _, tc := range []struct {
		Name                     string
		Targets                  *Targets
		DescribeTargetHealthCall *DescribeTargetHealthCall
		RegisterTargetsCall      *RegisterTargetsCall
		DeregisterTargetsCall    *DeregisterTargetsCall
		ResolveCall              *ResolveCall
		ExpectedError            error
	}{
		{
			Name:          "Resolve endpoint throws error",
			Targets:       &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			ExpectedError: errors.New("ERROR STRING"),
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
				Err:             errors.New("ERROR STRING"),
			},
		},
		{
			Name:    "DescribeTargetHealth throws error",
			Targets: &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			DescribeTargetHealthCall: &DescribeTargetHealthCall{
				TgArn: tgArn,
				Err:   fmt.Errorf("ERROR STRING"),
			},
			ExpectedError: errors.New("ERROR STRING"),
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
			},
		},
		{
			Name:    "deregister a target",
			Targets: &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			DescribeTargetHealthCall: &DescribeTargetHealthCall{
				TgArn: tgArn,
				Output: &elbv2.DescribeTargetHealthOutput{TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
					{Target: newTd("id", 123), TargetHealth: newTh(elbv2.TargetHealthStateEnumHealthy)},
				}},
			},
			DeregisterTargetsCall: &DeregisterTargetsCall{
				Input: &elbv2.DeregisterTargetsInput{TargetGroupArn: aws.String(tgArn), Targets: []*elbv2.TargetDescription{newTd("id", 123)}},
			},
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
			},
		},
		{
			Name:    "deregister a target with error",
			Targets: &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			DescribeTargetHealthCall: &DescribeTargetHealthCall{
				TgArn: tgArn,
				Output: &elbv2.DescribeTargetHealthOutput{TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
					{Target: newTd("id", 123), TargetHealth: newTh(elbv2.TargetHealthStateEnumHealthy)},
				}},
			},
			DeregisterTargetsCall: &DeregisterTargetsCall{
				Err:   errors.New("ERROR STRING"),
				Input: &elbv2.DeregisterTargetsInput{TargetGroupArn: aws.String(tgArn), Targets: []*elbv2.TargetDescription{newTd("id", 123)}},
			},
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
			},
			ExpectedError: errors.New("ERROR STRING"),
		},
		{
			Name:    "add a target",
			Targets: &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			DescribeTargetHealthCall: &DescribeTargetHealthCall{
				TgArn: tgArn,
				Output: &elbv2.DescribeTargetHealthOutput{TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
					{Target: newTd("id", 123), TargetHealth: newTh(elbv2.TargetHealthStateEnumHealthy)},
				}},
			},
			RegisterTargetsCall: &RegisterTargetsCall{
				Input: &elbv2.RegisterTargetsInput{TargetGroupArn: aws.String(tgArn), Targets: []*elbv2.TargetDescription{newTd("id2", 1234)}},
			},
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
				Output:          []*elbv2.TargetDescription{newTd("id", 123), newTd("id2", 1234)},
			},
		},
		{
			Name:    "add targets when there the it's been drained",
			Targets: &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			DescribeTargetHealthCall: &DescribeTargetHealthCall{
				TgArn: tgArn,
				Output: &elbv2.DescribeTargetHealthOutput{TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
					{Target: newTd("id", 123), TargetHealth: newTh(elbv2.TargetHealthStateEnumDraining)},
				}},
			},
			RegisterTargetsCall: &RegisterTargetsCall{
				Input: &elbv2.RegisterTargetsInput{TargetGroupArn: aws.String(tgArn), Targets: []*elbv2.TargetDescription{newTd("id", 123)}},
			},
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
				Output:          []*elbv2.TargetDescription{newTd("id", 123)},
			},
		},
		{
			Name:    "add a target with error",
			Targets: &Targets{TgArn: tgArn, Ingress: dummy.NewIngress(), Backend: backend, TargetType: elbv2.TargetTypeEnumInstance},
			DescribeTargetHealthCall: &DescribeTargetHealthCall{
				TgArn: tgArn,
				Output: &elbv2.DescribeTargetHealthOutput{TargetHealthDescriptions: []*elbv2.TargetHealthDescription{
					{Target: newTd("id", 123), TargetHealth: newTh(elbv2.TargetHealthStateEnumHealthy)},
				}},
			},
			RegisterTargetsCall: &RegisterTargetsCall{
				Input: &elbv2.RegisterTargetsInput{TargetGroupArn: aws.String(tgArn), Targets: []*elbv2.TargetDescription{newTd("id2", 1234)}},
				Err:   errors.New("ERROR STRING"),
			},
			ResolveCall: &ResolveCall{
				InputIngress:    dummy.NewIngress(),
				InputBackend:    backend,
				InputTargetType: elbv2.TargetTypeEnumInstance,
				Output:          []*elbv2.TargetDescription{newTd("id", 123), newTd("id2", 1234)},
			},
			ExpectedError: errors.New("ERROR STRING"),
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			endpointResolver := &mocks.EndpointResolver{}
			if tc.ResolveCall != nil {
				endpointResolver.On("Resolve", tc.ResolveCall.InputIngress, tc.ResolveCall.InputBackend, tc.ResolveCall.InputTargetType).Return(tc.ResolveCall.Output, tc.ResolveCall.Err)
			}

			cloud := &mocks.CloudAPI{}
			if tc.DescribeTargetHealthCall != nil {
				cloud.On("DescribeTargetHealth", &elbv2.DescribeTargetHealthInput{TargetGroupArn: aws.String(tc.DescribeTargetHealthCall.TgArn)}).Return(tc.DescribeTargetHealthCall.Output, tc.DescribeTargetHealthCall.Err)
			}
			if tc.RegisterTargetsCall != nil {
				cloud.On("RegisterTargets", tc.RegisterTargetsCall.Input).Return(nil, tc.RegisterTargetsCall.Err)
			}
			if tc.DeregisterTargetsCall != nil {
				cloud.On("DeregisterTargets", tc.DeregisterTargetsCall.Input).Return(nil, tc.DeregisterTargetsCall.Err)
			}

			controller := NewTargetsController(cloud, endpointResolver)
			err := controller.Reconcile(context.Background(), tc.Targets)

			if tc.ExpectedError != nil {
				assert.Equal(t, tc.ExpectedError, err)
			} else {
				assert.NoError(t, err)
			}
			cloud.AssertExpectations(t)
			endpointResolver.AssertExpectations(t)
		})

	}
}
func Test_targetChangeSets(t *testing.T) {
	for _, tc := range []struct {
		name      string
		a         []*elbv2.TargetDescription
		b         []*elbv2.TargetDescription
		addSet    []*elbv2.TargetDescription
		removeSet []*elbv2.TargetDescription
	}{
		{
			name:   "a empty, b adds one td",
			a:      []*elbv2.TargetDescription{},
			b:      []*elbv2.TargetDescription{newTd("id", 1)},
			addSet: []*elbv2.TargetDescription{newTd("id", 1)},
		},
		{
			name:      "b empty, b removes one td",
			a:         []*elbv2.TargetDescription{newTd("id", 1)},
			b:         []*elbv2.TargetDescription{},
			removeSet: []*elbv2.TargetDescription{newTd("id", 1)},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			addSet, removeSet := targetChangeSets(tc.a, tc.b)
			assert.Equal(t, tc.addSet, addSet, "add set not as expected")
			assert.Equal(t, tc.removeSet, removeSet, "remove set not as expected")
		})
	}
}

func Test_tdsString(t *testing.T) {
	for _, tc := range []struct {
		name     string
		a        []*elbv2.TargetDescription
		expected string
	}{
		{
			name:     "with port test",
			a:        []*elbv2.TargetDescription{newTd("id", 1), newTd("id2", 2)},
			expected: "id:1, id2:2",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := tdsString(tc.a)
			assert.Equal(t, tc.expected, s)

		})

	}
}
