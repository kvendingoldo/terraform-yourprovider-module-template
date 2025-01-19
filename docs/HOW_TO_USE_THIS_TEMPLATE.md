The `kvendingoldo/terraform-module-template` repository contains a template, that can be used for any Terraform/OpenTofu modules. By default, this repository gathered all best-practices related to the maintaining Terraform modules.

For the fork, use the following naming convention:
`terraform-<provider>-<name>`

# Required Steps

## examples Directory
This folder contains Terraform module examples. The `README.md` file is generated automatically using either pre-commit hooks or terraform-docs.

Do the following steps:
* Use the README.md file as the canonical example for documentation.
* Delete any outdated or redundant Terraform code in this folder.

## wrappers Directory
The configuration in this directory contains an implementation of a single module wrapper pattern, which allows managing several copies of a module in places where using the native Terraform 0.13+ for_each feature is not feasible (e.g., with Terragrunt).

You may want to use a single Terragrunt configuration file to manage multiple resources without duplicating terragrunt.hcl files for each copy of the same module.

This wrapper does not implement any extra functionality.

Do the following steps:
* Delete or re-generate code via https://gist.github.com/antonbabenko/d77f8cf8bf891e589a6b5b0ab0e773ae

## README.md
To customize the README.md file, do the following steps:

* Replace any instances of `kvendingoldo/terraform-module-template` with `<your_org>/<your_repo>`.
* Update <PROVIDER> <NAME> Terraform module with the actual provider and module name.
* Replace the placeholder description "Terraform module which creates <RESOURCE_NAMES> resources" with an accurate and detailed description of the module’s purpose.

All other content will be automatically generated when you do any of that:
* Update the code and run: `git add -A && git commit -m "feat: update code" && pre-commit run -a`
* Run: `terraform-docs -c .terraform-docs.yml .`


## .github Directory
Do the following steps:
* `.github/ISSUE_TEMPLATES`: replace `terraform-module-template` with `<your_repo>`.
* `.github/workflows/linters.yml`: specify your Terraform version (default is `1.5.7`), and update the TFLint version if necessary (default is `0.54.0`).
* Check `.github/CONTRIBUTING.md` to match organizational policies.
* Change or delete `.github/FUNDING.yml`. It configures funding links to platforms like GitHub Sponsors, Patreon, etc.

## functions Directory
Remove this folder if no custom functions are implemented (e.g. it can be AWS lambda or anything like that).

## modules Directory
Delete this folder if there are no Terraform submodules.

## tests Directory
Remove this folder if no test cases are provided (e.g.: it can be `terratest`, `pytest` or any other kind of tests).

## CODEOWNERS
Review and update the CODEOWNERS file to reflect the appropriate ownership structure. Delete the file if it is not required for your GitHub project.

## Terraform Code (.tf files)
Keep the structure, and just delete the content of .tf files. By default, the repository contains a sample Terraform code.

## MAINTAINERS.md
Update the MAINTAINERS.md file to include accurate and current information about repository maintainers. Remove the file if unnecessary.

## SECURITY.md
Correct the following information in the `SECURITY.md` file

```
Email: [security@example.com](mailto:security@example.com)
Optional GPG Key link:
GPG Key (Optional): [Link to GPG Key]
Support Documentation
```

## SUPPORT.md
Update the file by replacing any instances of `kvendingoldo/terraform-module-template` with `<your_org>/<your_repo>`.


# Optional Steps
1. Review all additional configuration files and ensure alignment with enterprise standards.
2. Validate that all scripts, configurations, and documentation adhere to compliance and security guidelines.
