todo:
examples
wrappers



# README.md
1. Replace any `kvendingoldo/terraform-module-template` to <your_org>/<your_repo>
2. Replace `<PROVIDER> <NAME> Terraform module` to your actual name
3. Replace "Terraform module which creates <RESOURCE_NAMES> resources."

All other content will be generated automatically, when you will do the following:
1. Change code
2. git add -A && git commit -m "feat: update code" && pre-commit run -a

Otherwise, you can use a pure terraform-docs command to update docs: terraform-docs -c .terraform-docs.yml .


# .github

1. Edit .github/ISSUE_TEMPLATES, at least change terraform-module-template to <your_repo>.
2. in .github/workflows/linters.yml change tf_version to the actual one, otherwise 1.5.7 will be used. Do the same for tflint_version, by default 0.54.0 will be used.


3. check CONTRIBUTING.md
4.


# functions
delete this folder if you dont' have any functions

# modules
delete this folder if you dont' have any submodules

# tests
delete this folder if you dont' have any tests


# CODEOWNERS
change it or delte


# .tf code
change it or delete

# MAINTAINERS.md
change it or delete

# SECURITY.md

change the:

```
Email: [security@example.com](mailto:security@example.com)
GPG Key (Optional): [Link to GPG Key]
```

# SUPPORT.md
your-org/your-repo
Replace any `kvendingoldo/terraform-module-template` to <your_org>/<your_repo>

## optional

1. check all other configurations
