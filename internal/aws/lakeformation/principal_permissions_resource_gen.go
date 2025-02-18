// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lakeformation_principal_permissions", principalPermissionsResource)
}

// principalPermissionsResource returns the Terraform awscc_lakeformation_principal_permissions resource.
// This Terraform resource corresponds to the CloudFormation AWS::LakeFormation::PrincipalPermissions resource.
func principalPermissionsResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Catalog
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"catalog": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(12, 12),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Permissions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "ALL",
		//	      "SELECT",
		//	      "ALTER",
		//	      "DROP",
		//	      "DELETE",
		//	      "INSERT",
		//	      "DESCRIBE",
		//	      "CREATE_DATABASE",
		//	      "CREATE_TABLE",
		//	      "DATA_LOCATION_ACCESS",
		//	      "CREATE_TAG",
		//	      "ASSOCIATE"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"permissions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"ALL",
						"SELECT",
						"ALTER",
						"DROP",
						"DELETE",
						"INSERT",
						"DESCRIBE",
						"CREATE_DATABASE",
						"CREATE_TABLE",
						"DATA_LOCATION_ACCESS",
						"CREATE_TAG",
						"ASSOCIATE",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PermissionsWithGrantOption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "ALL",
		//	      "SELECT",
		//	      "ALTER",
		//	      "DROP",
		//	      "DELETE",
		//	      "INSERT",
		//	      "DESCRIBE",
		//	      "CREATE_DATABASE",
		//	      "CREATE_TABLE",
		//	      "DATA_LOCATION_ACCESS",
		//	      "CREATE_TAG",
		//	      "ASSOCIATE"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"permissions_with_grant_option": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.OneOf(
						"ALL",
						"SELECT",
						"ALTER",
						"DROP",
						"DELETE",
						"INSERT",
						"DESCRIBE",
						"CREATE_DATABASE",
						"CREATE_TABLE",
						"DATA_LOCATION_ACCESS",
						"CREATE_TAG",
						"ASSOCIATE",
					),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Principal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DataLakePrincipalIdentifier": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"principal": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataLakePrincipalIdentifier
				"data_lake_principal_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrincipalIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"principal_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Resource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Catalog": {
		//	      "additionalProperties": false,
		//	      "type": "object"
		//	    },
		//	    "DataCellsFilter": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TableCatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "TableName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "TableCatalogId",
		//	        "DatabaseName",
		//	        "TableName",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DataLocation": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "ResourceArn": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "ResourceArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Database": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LFTag": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "TagKey": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TagValues": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 256,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          },
		//	          "maxItems": 50,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "TagKey",
		//	        "TagValues"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LFTagPolicy": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "Expression": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "TagKey": {
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "TagValues": {
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "maxLength": 256,
		//	                  "minLength": 0,
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 50,
		//	                "minItems": 1,
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "maxItems": 5,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        },
		//	        "ResourceType": {
		//	          "enum": [
		//	            "DATABASE",
		//	            "TABLE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "ResourceType",
		//	        "Expression"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Table": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TableWildcard": {
		//	          "additionalProperties": false,
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "DatabaseName"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TableWithColumns": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CatalogId": {
		//	          "maxLength": 12,
		//	          "minLength": 12,
		//	          "type": "string"
		//	        },
		//	        "ColumnNames": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 255,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "type": "array"
		//	        },
		//	        "ColumnWildcard": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ExcludedColumnNames": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "maxLength": 255,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "DatabaseName": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Name": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "CatalogId",
		//	        "DatabaseName",
		//	        "Name"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Catalog
				"catalog": schema.MapAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DataCellsFilter
				"data_cells_filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: TableCatalogId
						"table_catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: TableName
						"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DataLocation
				"data_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: ResourceArn
						"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Database
				"database": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LFTag
				"lf_tag": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: TagKey
						"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: TagValues
						"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 50),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(0, 256),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LFTagPolicy
				"lf_tag_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Expression
						"expression": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: TagKey
									"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.LengthBetween(1, 128),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: TagValues
									"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Optional:    true,
										Computed:    true,
										Validators: []validator.List{ /*START VALIDATORS*/
											listvalidator.SizeBetween(1, 50),
											listvalidator.ValueStringsAre(
												stringvalidator.LengthBetween(0, 256),
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
											generic.Multiset(),
											listplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Required: true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 5),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ResourceType
						"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"DATABASE",
									"TABLE",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Table
				"table": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: TableWildcard
						"table_wildcard": schema.MapAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
								mapplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TableWithColumns
				"table_with_columns": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CatalogId
						"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(12, 12),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: ColumnNames
						"column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 255),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: ColumnWildcard
						"column_wildcard": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ExcludedColumnNames
								"excluded_column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.ValueStringsAre(
											stringvalidator.LengthBetween(1, 255),
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										generic.Multiset(),
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"resource_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "A resource schema representing a Lake Formation Permission.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::PrincipalPermissions").WithTerraformTypeName("awscc_lakeformation_principal_permissions")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":                        "Catalog",
		"catalog_id":                     "CatalogId",
		"column_names":                   "ColumnNames",
		"column_wildcard":                "ColumnWildcard",
		"data_cells_filter":              "DataCellsFilter",
		"data_lake_principal_identifier": "DataLakePrincipalIdentifier",
		"data_location":                  "DataLocation",
		"database":                       "Database",
		"database_name":                  "DatabaseName",
		"excluded_column_names":          "ExcludedColumnNames",
		"expression":                     "Expression",
		"lf_tag":                         "LFTag",
		"lf_tag_policy":                  "LFTagPolicy",
		"name":                           "Name",
		"permissions":                    "Permissions",
		"permissions_with_grant_option":  "PermissionsWithGrantOption",
		"principal":                      "Principal",
		"principal_identifier":           "PrincipalIdentifier",
		"resource":                       "Resource",
		"resource_arn":                   "ResourceArn",
		"resource_identifier":            "ResourceIdentifier",
		"resource_type":                  "ResourceType",
		"table":                          "Table",
		"table_catalog_id":               "TableCatalogId",
		"table_name":                     "TableName",
		"table_wildcard":                 "TableWildcard",
		"table_with_columns":             "TableWithColumns",
		"tag_key":                        "TagKey",
		"tag_values":                     "TagValues",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
