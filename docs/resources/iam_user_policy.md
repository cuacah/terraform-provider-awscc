---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_iam_user_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for IAM User Policy
---

# awscc_iam_user_policy (Resource)

Schema for IAM User Policy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy_name` (String) The name of the policy document.
- `user_name` (String) The name of the user to associate the policy with.

### Optional

- `policy_document` (Map of String) The policy document.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_iam_user_policy.example <resource ID>
```
