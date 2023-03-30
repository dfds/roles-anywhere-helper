[![Go Report Card](https://goreportcard.com/badge/github.com/dfds/roles-anywhere-helper)](https://goreportcard.com/report/github.com/dfds/roles-anywhere-helper)


# iam-anywhere-ninja

## Introduction
Allows to easily set up AWS IAM roles anywhere credential file on a machine, this CLI makes assumptions that you are using an AWS Private CA, AWS Certificate Manager and have already set IAM Roles Anywhere Correctly. This was designed to work on Linux, Windows and Mac.

## Purpose
The purpose of this CLI is to make it easy for development teams to configure, revoke and Rotate Credentials on the machine that needs access to AWS via IAM Roles anywhere. Ideally, this would be used in the CICD pipeline. 

## Feautres

- Credential file setup to use IAM roles for a given AWS profile.
- Generate an x509 certificate on the machine and imports it to the Private Certificate Authority (CA)
- Import a certificate to the AWS certificate manager (ACM)
- Rotate a certificate by generating and importing a new certificate and then marking the old certificate as obsolete.
- Set up a machine to allow for IAM roles anywhere in a CICD pipeline includes, including creating the credential file and importing the certificate.

## Prequistis 

- [AWS_Signing_helper - Found here](https://docs.aws.amazon.com/rolesanywhere/latest/userguide/credential-helper.html)
- AWS Private CA provisioned with CLR for Credential Revocation
- AWS IAM Roles Anywhere configured.

## Getting Started

### Installation
- Download the binary for your target platform from
  the [releases page](https://github.com/dfds/iam-anywhere-helper/releases)

### Usage

#### configure-credential

#### setup-all

```
$ ./roles-anywhere-helper help setup-all 

Setup the whole process of configuring AWS IAM Roles Anywhere.
                List of operations:
                        - Generates certificate and issues that in AWS ACM PCA.
                        - Imports certificate into AWS ACM.
                        - Configures AWS profile in credentials file to use AWS Signing Helper with created certificate.

Usage:
  roles-anywhere-helper setup-all [flags]

Flags:
      --acm-region string                    Name of the region to be used for access to the ACM (default "eu-east-1")
      --acmpca-arn string                    ARN of the private CA that issues the certificate
      --certificate-directory string         Path of the certificate directory on the machine
      --common-name string                   The common name for the X509 certificate
      --country string                       The country name for the X509 certificate
  -h, --help                                 help for setup-all
      --locality string                      The locality name for the X509 certificate
      --organization-name string             The organization name for the X509 certificate
      --organizational-unit string           The organization unit for the X509 certificate
      --pca-region string                    Name of the region to be used for access to the PCA (default "eu-east-1")
      --profile-arn string                   The Arn of the AWS IAM roles Anywhere profile
      --profile-name-acm string              Name of the profile to be used for access to the ACM (default "default")
      --profile-name-pca string              Name of the profile to be used for access to the PCA (default "default")
      --profile-name-roles-anywhere string   Name of the profile to that the credentials will be created under (default "roles-anywhere")
      --province string                      The state or province name for the X509 certificate
      --role-arn string                      The Arn of the role to be assumed with AWS IAM roles Anywhere
      --rolesanywhere-region string          Name of the region to that the credentials will be created under (default "eu-east-1")
      --trust-anchor-arn string              The Arn of the AWS IAM roles anywhere trust anchor

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)
```

#### generate-certificate

#### import-certificate

#### revoke-certificate

#### rotate-certificate

## Contributions

Contributions are welcome :)

* Feel free to contribute enhancements or bug fixes.
    * Fork this repo, apply your changes and create a PR pointing to this repo and the main branch
* If you have any ideas or suggestions please open an issue and describe your idea or feature request

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
