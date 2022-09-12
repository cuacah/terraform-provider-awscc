// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iot_role_alias", roleAliasResourceType)
}

// roleAliasResourceType returns the Terraform awscc_iot_role_alias resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::RoleAlias resource type.
func roleAliasResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"credential_duration_seconds": {
			// Property: CredentialDurationSeconds
			// CloudFormation resource type schema:
			// {
			//   "default": 3600,
			//   "maximum": 43200,
			//   "minimum": 900,
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntBetween(900, 43200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Int64{Value: 3600}),
				resource.UseStateForUnknown(),
			},
		},
		"role_alias": {
			// Property: RoleAlias
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[\\w=,@-]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[\\w=,@-]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"role_alias_arn": {
			// Property: RoleAliasArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[\\w=,@-]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "arn:(aws[a-zA-Z-]*)?:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("arn:(aws[a-zA-Z-]*)?:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+"), ""),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
							validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
							validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Use the AWS::IoT::RoleAlias resource to declare an AWS IoT RoleAlias.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::RoleAlias").WithTerraformTypeName("awscc_iot_role_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"credential_duration_seconds": "CredentialDurationSeconds",
		"key":                         "Key",
		"role_alias":                  "RoleAlias",
		"role_alias_arn":              "RoleAliasArn",
		"role_arn":                    "RoleArn",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
