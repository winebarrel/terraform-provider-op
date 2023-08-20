# terraform-provider-op

Terraform provider for [1Password CLI](https://developer.1password.com/docs/cli/).

## Preparation

Install 1Password CLI.

see https://developer.1password.com/docs/cli/get-started/

## Usage

```tf
terraform {
  required_providers {
    op = {
      source = "winebarrel/op"
    }
  }
}

provider "op" {
  # command = "/usr/local/bin/op"
}

data "op_item" "yahoo" {
  # id = "xxx..."
  title = "yahoo"
}

output "yahoo_username" {
  value = data.op_item.yahoo.username
}

output "yahoo_password" {
  value     = data.op_item.yahoo.password
  sensitive = true
}

output "yahoo_field_value" {
  value = data.op_item.yahoo.fields["my-label"]
}
```

## Run locally for development

```sh
cp op.tf.sample op.tf
make tf-plan
```
