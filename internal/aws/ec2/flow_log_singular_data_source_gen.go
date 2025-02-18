// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_flow_log", flowLogDataSource)
}

// flowLogDataSource returns the Terraform awscc_ec2_flow_log data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::FlowLog resource.
func flowLogDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeliverLogsPermissionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
		//	  "type": "string"
		//	}
		"deliver_logs_permission_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "FileFormat": {
		//	      "enum": [
		//	        "plain-text",
		//	        "parquet"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "HiveCompatiblePartitions": {
		//	      "type": "boolean"
		//	    },
		//	    "PerHourPartition": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "FileFormat",
		//	    "HiveCompatiblePartitions",
		//	    "PerHourPartition"
		//	  ],
		//	  "type": "object"
		//	}
		"destination_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FileFormat
				"file_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: HiveCompatiblePartitions
				"hive_compatible_partitions": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: PerHourPartition
				"per_hour_partition": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Flow Log ID",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Flow Log ID",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogDestination
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.",
		//	  "type": "string"
		//	}
		"log_destination": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogDestinationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.",
		//	  "enum": [
		//	    "cloud-watch-logs",
		//	    "s3",
		//	    "kinesis-data-firehose"
		//	  ],
		//	  "type": "string"
		//	}
		"log_destination_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The fields to include in the flow log record, in the order in which they should appear.",
		//	  "type": "string"
		//	}
		"log_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The fields to include in the flow log record, in the order in which they should appear.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
		//	  "type": "string"
		//	}
		"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxAggregationInterval
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).",
		//	  "type": "integer"
		//	}
		"max_aggregation_interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet, network interface, or VPC for which you want to create a flow log.",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet, network interface, or VPC for which you want to create a flow log.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.",
		//	  "enum": [
		//	    "NetworkInterface",
		//	    "Subnet",
		//	    "VPC",
		//	    "TransitGateway",
		//	    "TransitGatewayAttachment"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to apply to the flow logs.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to apply to the flow logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.",
		//	  "enum": [
		//	    "ACCEPT",
		//	    "ALL",
		//	    "REJECT"
		//	  ],
		//	  "type": "string"
		//	}
		"traffic_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::FlowLog",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::FlowLog").WithTerraformTypeName("awscc_ec2_flow_log")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"deliver_logs_permission_arn": "DeliverLogsPermissionArn",
		"destination_options":         "DestinationOptions",
		"file_format":                 "FileFormat",
		"hive_compatible_partitions":  "HiveCompatiblePartitions",
		"id":                          "Id",
		"key":                         "Key",
		"log_destination":             "LogDestination",
		"log_destination_type":        "LogDestinationType",
		"log_format":                  "LogFormat",
		"log_group_name":              "LogGroupName",
		"max_aggregation_interval":    "MaxAggregationInterval",
		"per_hour_partition":          "PerHourPartition",
		"resource_id":                 "ResourceId",
		"resource_type":               "ResourceType",
		"tags":                        "Tags",
		"traffic_type":                "TrafficType",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
