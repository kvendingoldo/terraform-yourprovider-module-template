# Wrapper for the root module

The configuration in this directory contains an implementation of a single module wrapper pattern, which allows managing several copies of a module in places where using the native Terraform 0.13+ `for_each` feature is not feasible (e.g., with `Terragrunt`).

You may want to use a single `Terragrunt` configuration file to manage multiple resources without duplicating `terragrunt.hcl` files for each copy of the same module.

This wrapper does not implement any extra functionality.

## Usage with Terragrunt

`terragrunt.hcl`:

```hcl
terraform {
 source = "git::git@github.com:kvendingoldo/terraform-module-template.git//wrappers?ref=main"
}

inputs = {
  defaults = {
    name = "kvendingoldo-example"
  }

  items = {
    my-item = {
      # omitted... can be any argument supported by the module
    }
    my-second-item = {
      # omitted... can be any argument supported by the module
    }
    # omitted...
  }
}
```

## Usage with Terraform

```hcl
module "wrapper" {
  source = "git::git@github.com:kvendingoldo/terraform-module-template.git//wrappers?ref=main"

  defaults = {
    name = "kvendingoldo-example"
  }

  items = {
    my-item = {
      # omitted... can be any argument supported by the module
    }
    my-second-item = {
      # omitted... can be any argument supported by the module
    }
    # omitted...
  }
}
```
