<!-- BADGES -->
[![Github release](https://img.shields.io/github/v/release/kvendingoldo/terraform-module-template?style=for-the-badge)](https://github.com/kvendingoldo/terraform-module-template/releases) [![Contributors](https://img.shields.io/github/contributors/kvendingoldo/terraform-module-template?style=for-the-badge)](https://github.com/kvendingoldo/terraform-module-template/graphs/contributors) ![maintenance status](https://img.shields.io/maintenance/yes/2025.svg?style=for-the-badge)


# <PROVIDER> <NAME> Terraform module

Terraform module which creates <RESOURCE_NAMES> resources.

## Examples

Examples codified under
the [`examples`](https://github.com/kvendingoldo/terraform-module-template/tree/main/examples) are intended
to give users references for how to use the module(s) as well as testing/validating changes to the source code of the
module. If contributing to the project, please be sure to make any appropriate updates to the relevant examples to allow
maintainers to test your changes and to keep the examples up to date for users. Thank you!

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_instance.example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ami_id"></a> [ami\_id](#input\_ami\_id) | The default AWS EC2 AMI | `string` | `"ami-0c55b159cbfafe1f0"` | no |
| <a name="input_name"></a> [name](#input\_name) | terraform-module-template | `string` | `"The name of AWS EC2 VM"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_instance_id"></a> [instance\_id](#output\_instance\_id) | AWS EC2 instance id |
<!-- END_TF_DOCS -->


## Authors

<!-- markdownlint-disable no-inline-html -->
<a href="https://github.com/kvendingoldo/terraform-module-template/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=kvendingoldo/terraform-module-template" />
</a>

<a href="https://star-history.com/#kvendingoldo/terraform-module-template&Date">
  <picture>
    <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=kvendingoldo/terraform-module-template&type=Date" />
  </picture>
</a>
<!-- markdownlint-enable no-inline-html -->

## License

Apache-2.0 Licensed.
See [LICENSE](https://github.com/kvendingoldo/terraform-module-template/blob/main/LICENSE).
