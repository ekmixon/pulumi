// Copyright 2017 Pulumi, Inc. All rights reserved.

package ec2

import (
	"github.com/pulumi/coconut/pkg/resource/idl"
)

// A SecurityGroup is an Amazon EC2 Security Group.  For more information, see
// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html.
type SecurityGroup struct {
	idl.NamedResource
	// A required description about the security group.
	GroupDescription string `coco:"groupDescription,replaces"`
	// The VPC in which this security group resides (or blank if the default VPC).
	VPC *VPC `coco:"vpc,optional,replaces"`
	// A list of Amazon EC2 security group egress rules.
	SecurityGroupEgress *[]SecurityGroupRule `coco:"securityGroupEgress,optional"`
	// A list of Amazon EC2 security group ingress rules.
	SecurityGroupIngress *[]SecurityGroupRule `coco:"securityGroupIngress,optional"`
	// The group ID of the specified security group, such as `sg-94b3a1f6`.
	GroupID string `coco:"groupID,out"`
}

// A SecurityGroupRule describes an EC2 security group rule embedded within a SecurityGroup.
type SecurityGroupRule struct {
	// The IP name or number.
	IPProtocol string `coco:"ipProtocol"`
	// Specifies a CIDR range.
	CIDRIP *string `coco:"cidrIp,optional"`
	// The start of port range for the TCP and UDP protocols, or an ICMP type number.  An ICMP type number of `-1`
	// indicates a wildcard (i.e., any ICMP type number).
	FromPort *float64 `coco:"fromPort,optional"`
	// The end of port range for the TCP and UDP protocols, or an ICMP code.  An ICMP code of `-1` indicates a
	// wildcard (i.e., any ICMP code).
	ToPort *float64 `coco:"toPort,optional"`
}
