---
subcategory: "Connect"
layout: "aws"
page_title: "AWS: aws_connect_prompt"
description: |-
  Provides details about a specific Amazon Connect Prompt.
---

# Data Source: aws_connect_prompt

Provides details about a specific Amazon Connect Prompt.

## Example Usage

By `name`

```terraform
data "aws_connect_prompt" "example" {
  instance_id = "aaaaaaaa-bbbb-cccc-dddd-111111111111"
  name        = "Beep.wav"
}
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `instance_id` - (Required) Reference to the hosting Amazon Connect Instance
* `name` - (Required) Returns information on a specific Prompt by name

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Prompt.
* `prompt_id` - Identifier for the prompt.
