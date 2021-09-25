// Code generated by generators/resource/main.go; DO NOT EDIT.

package sso

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_sso_instance_access_control_attribute_configuration", instanceAccessControlAttributeConfigurationResourceType)
}

// instanceAccessControlAttributeConfigurationResourceType returns the Terraform awscc_sso_instance_access_control_attribute_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSO::InstanceAccessControlAttributeConfiguration resource type.
func instanceAccessControlAttributeConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_control_attributes": {
			// Property: AccessControlAttributes
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Source": {
			//             "insertionOrder": true,
			//             "items": {
			//               "maxLength": 256,
			//               "minLength": 0,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "maxItems": 1,
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Source"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"source": {
									// Property: Source
									Type:     types.ListType{ElemType: types.StringType},
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtMost(1),
										validate.ArrayForEach(validate.StringLenBetween(0, 256)),
									},
								},
							},
						),
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"instance_access_control_attribute_configuration": {
			// Property: InstanceAccessControlAttributeConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.",
			//   "properties": {
			//     "AccessControlAttributes": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Key": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Value": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Source": {
			//                 "insertionOrder": true,
			//                 "items": {
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "maxItems": 1,
			//                 "type": "array"
			//               }
			//             },
			//             "required": [
			//               "Source"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "Key",
			//           "Value"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 50,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "AccessControlAttributes"
			//   ],
			//   "type": "object"
			// }
			Description: "The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"access_control_attributes": {
						// Property: AccessControlAttributes
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 128),
									},
								},
								"value": {
									// Property: Value
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"source": {
												// Property: Source
												Type:     types.ListType{ElemType: types.StringType},
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenAtMost(1),
													validate.ArrayForEach(validate.StringLenBetween(0, 256)),
												},
											},
										},
									),
									Required: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Optional: true,
		},
		"instance_arn": {
			// Property: InstanceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the AWS SSO instance under which the operation will be executed.",
			//   "maxLength": 1224,
			//   "minLength": 10,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ARN of the AWS SSO instance under which the operation will be executed.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(10, 1224),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for SSO InstanceAccessControlAttributeConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSO::InstanceAccessControlAttributeConfiguration").WithTerraformTypeName("awscc_sso_instance_access_control_attribute_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_control_attributes":                       "AccessControlAttributes",
		"instance_access_control_attribute_configuration": "InstanceAccessControlAttributeConfiguration",
		"instance_arn": "InstanceArn",
		"key":          "Key",
		"source":       "Source",
		"value":        "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
