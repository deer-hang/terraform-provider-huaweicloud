---
subcategory: "Identity and Access Management (IAM)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_identity_users"
description: ""
---

# huaweicloud_identity_users

Use this data source to query the IAM user list within HuaweiCloud.

-> **NOTE:** You *must* have IAM read privileges to use this data source.

## Example Usage

```hcl
data "huaweicloud_identity_users" "all" {}

data "huaweicloud_identity_users" "one" {
  name = "user_name"
}
```

## Argument Reference

* `name` - (Optional, String) Specifies the IAM user name.

* `enabled` - (Optional, Bool) Specifies the status of the IAM user, the default value is **true**.

## Attribute Reference

* `id` - The data source ID.

* `users` - The details of the queried IAM users. The structure is documented below.

The `users` block contains:

* `id` - Indicates the ID of the User.

* `name` - Indicates the IAM user name.

* `description` - Indicates the description of the IAM user.

* `enabled` - Indicates the whether the IAM user is enabled.

* `groups` - Indicates the user groups to which an IAM user belongs.

* `password_expires_at` - Indicates the time when the password will expire.
  If this value is not set, the password will not expire.

* `password_status` - Indicates the password status. True means that the password needs to be changed,
  and false means that the password is normal.
