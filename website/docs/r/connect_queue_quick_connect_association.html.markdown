---
subcategory: "Connect"
layout: "aws"
page_title: "AWS: aws_connect_queue_quick_connect_association"
description: |-
  Provides details about a specific Amazon Queue Quick Connect Association
---

# Resource: aws_connect_queue_quick_connect_association

Provides an Amazon Connect Queue Quick Connect association resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

## Example Usage

```terraform
resource "aws_connect_queue_quick_connect_association" "test" {
  instance_id = "aaaaaaaa-bbbb-cccc-dddd-111111111111"
  queue_id    = "12345678-1234-1234-1234-123456789012"

  quick_connect_ids = [
    "12345678-abcd-1234-abcd-123456789012"
  ]
}
```

## Argument Reference

This resource supports the following arguments:

* `instance_id` - (Required) Specifies the identifier of the hosting Amazon Connect Instance.
* `queue_id` - (Required) Specifies the identifier for the queue.
* `quick_connect_ids` - (Optional) Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the Quick Connect.
* `quick_connect_id` - The identifier for the Quick Connect.
* `id` - The identifier of the hosting Amazon Connect Instance and identifier of the Quick Connect separated by a colon (`:`).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Amazon Connect Quick Connects using the `instance_id` and `quick_connect_id` separated by a colon (`:`). For example:

```terraform
import {
  to = aws_connect_quick_connect.example
  id = "f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5"
}
```

Using `terraform import`, import Amazon Connect Quick Connects using the `instance_id` and `quick_connect_id` separated by a colon (`:`). For example:

```console
% terraform import aws_connect_quick_connect.example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
