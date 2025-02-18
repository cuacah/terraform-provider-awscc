// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigatewayv2_deployment", deploymentDataSource)
}

// deploymentDataSource returns the Terraform awscc_apigatewayv2_deployment data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGatewayV2::Deployment resource.
func deploymentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"deployment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StageName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"stage_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGatewayV2::Deployment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::Deployment").WithTerraformTypeName("awscc_apigatewayv2_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":        "ApiId",
		"deployment_id": "DeploymentId",
		"description":   "Description",
		"stage_name":    "StageName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
