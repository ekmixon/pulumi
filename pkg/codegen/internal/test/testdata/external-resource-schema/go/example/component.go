// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	storagev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/storage/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Component struct {
	pulumi.CustomResourceState

	Provider       kubernetes.ProviderOutput       `pulumi:"provider"`
	SecurityGroup  ec2.SecurityGroupOutput         `pulumi:"securityGroup"`
	StorageClasses storagev1.StorageClassMapOutput `pulumi:"storageClasses"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		args = &ComponentArgs{}
	}

	var resource Component
	err := ctx.RegisterResource("example::Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("example::Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
	Provider       *kubernetes.Provider               `pulumi:"provider"`
	SecurityGroup  *ec2.SecurityGroup                 `pulumi:"securityGroup"`
	StorageClasses map[string]*storagev1.StorageClass `pulumi:"storageClasses"`
}

type ComponentState struct {
	Provider       kubernetes.ProviderInput
	SecurityGroup  ec2.SecurityGroupInput
	StorageClasses storagev1.StorageClassMapInput
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	Metadata metav1.ObjectMetaPtrInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((*Component)(nil))
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

func (i *Component) ToComponentPtrOutput() ComponentPtrOutput {
	return i.ToComponentPtrOutputWithContext(context.Background())
}

func (i *Component) ToComponentPtrOutputWithContext(ctx context.Context) ComponentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentPtrOutput)
}

type ComponentPtrInput interface {
	pulumi.Input

	ToComponentPtrOutput() ComponentPtrOutput
	ToComponentPtrOutputWithContext(ctx context.Context) ComponentPtrOutput
}

type componentPtrType ComponentArgs

func (*componentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil))
}

func (i *componentPtrType) ToComponentPtrOutput() ComponentPtrOutput {
	return i.ToComponentPtrOutputWithContext(context.Background())
}

func (i *componentPtrType) ToComponentPtrOutputWithContext(ctx context.Context) ComponentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentPtrOutput)
}

// ComponentArrayInput is an input type that accepts ComponentArray and ComponentArrayOutput values.
// You can construct a concrete instance of `ComponentArrayInput` via:
//
//          ComponentArray{ ComponentArgs{...} }
type ComponentArrayInput interface {
	pulumi.Input

	ToComponentArrayOutput() ComponentArrayOutput
	ToComponentArrayOutputWithContext(context.Context) ComponentArrayOutput
}

type ComponentArray []ComponentInput

func (ComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Component)(nil))
}

func (i ComponentArray) ToComponentArrayOutput() ComponentArrayOutput {
	return i.ToComponentArrayOutputWithContext(context.Background())
}

func (i ComponentArray) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentArrayOutput)
}

// ComponentMapInput is an input type that accepts ComponentMap and ComponentMapOutput values.
// You can construct a concrete instance of `ComponentMapInput` via:
//
//          ComponentMap{ "key": ComponentArgs{...} }
type ComponentMapInput interface {
	pulumi.Input

	ToComponentMapOutput() ComponentMapOutput
	ToComponentMapOutputWithContext(context.Context) ComponentMapOutput
}

type ComponentMap map[string]ComponentInput

func (ComponentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Component)(nil))
}

func (i ComponentMap) ToComponentMapOutput() ComponentMapOutput {
	return i.ToComponentMapOutputWithContext(context.Background())
}

func (i ComponentMap) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentMapOutput)
}

type ComponentOutput struct {
	*pulumi.OutputState
}

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Component)(nil))
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentPtrOutput() ComponentPtrOutput {
	return o.ToComponentPtrOutputWithContext(context.Background())
}

func (o ComponentOutput) ToComponentPtrOutputWithContext(ctx context.Context) ComponentPtrOutput {
	return o.ApplyT(func(v Component) *Component {
		return &v
	}).(ComponentPtrOutput)
}

type ComponentPtrOutput struct {
	*pulumi.OutputState
}

func (ComponentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil))
}

func (o ComponentPtrOutput) ToComponentPtrOutput() ComponentPtrOutput {
	return o
}

func (o ComponentPtrOutput) ToComponentPtrOutputWithContext(ctx context.Context) ComponentPtrOutput {
	return o
}

type ComponentArrayOutput struct{ *pulumi.OutputState }

func (ComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Component)(nil))
}

func (o ComponentArrayOutput) ToComponentArrayOutput() ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) Index(i pulumi.IntInput) ComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Component {
		return vs[0].([]Component)[vs[1].(int)]
	}).(ComponentOutput)
}

type ComponentMapOutput struct{ *pulumi.OutputState }

func (ComponentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Component)(nil))
}

func (o ComponentMapOutput) ToComponentMapOutput() ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) MapIndex(k pulumi.StringInput) ComponentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Component {
		return vs[0].(map[string]Component)[vs[1].(string)]
	}).(ComponentOutput)
}

func init() {
	pulumi.RegisterOutputType(ComponentOutput{})
	pulumi.RegisterOutputType(ComponentPtrOutput{})
	pulumi.RegisterOutputType(ComponentArrayOutput{})
	pulumi.RegisterOutputType(ComponentMapOutput{})
}
