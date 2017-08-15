package resources

// AWS::EMR::Cluster.Application AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html
type AWSEMRClusterApplication struct {

	// AdditionalInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-additionalinfo
	AdditionalInfo map[string]AWSEMRClusterApplicationstring `json:"AdditionalInfo"`

	// Args AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-args
	Args []AWSEMRClusterApplicationstring `json:"Args"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-name
	Name string `json:"Name"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-application.html#cfn-emr-cluster-application-version
	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterApplication) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterApplication) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}