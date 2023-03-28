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
