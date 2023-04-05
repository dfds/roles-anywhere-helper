# roles-anywhere-helper

## Introduction
Allows to easily set up AWS IAM roles anywhere credential file on a machine, this CLI makes assumptions that you are using an AWS Private CA, AWS Certificate Manager and have already set IAM Roles Anywhere Correctly. This was designed to work on Linux, Windows and Mac.

## Purpose
The purpose of this CLI is to make it easy for development teams to configure, revoke and Rotate Credentials on the machine that needs access to AWS via IAM Roles anywhere. Ideally, this would be used in the CICD pipeline. 

## Feautres

```
Available Commands:
  configure-credentials Configures your local AWS credential file for IAM roles anywhere
  generate-certificate  Generates a X509 certificate and issues with AWS PCM
  help                  Help about any command
  import-certificate    Imports a certificate into AWS ACM
  revoke-certificate    Revoke a certificate in AWS PCA
  rotate-certificate    Rotate certificate
  setup-all             Generates a certificate, imports it to AWS ACM and configures roles anywhere credential file
```

## Dependancies

- [AWS_Signing_helper - Found here](https://docs.aws.amazon.com/rolesanywhere/latest/userguide/credential-helper.html)
- AWS Private CA provisioned with CLR for Credential Revocation
- AWS IAM Roles Anywhere configured.

## Getting Started

### Installation
- Download the binary for your target platform from
  the [releases page](https://github.com/dfds/roles-anywhere-helper/releases)

### Usage

#### configure-credentials

```
$ ./roles-anywhere-helper help configure-credentials 

Usage:
  roles-anywhere-helper configure-credentials [flags]

Flags:
  -c, --certificate-path string       Path of the certificate directory on the machine
  -h, --help                          help for configure-credentials
  -k, --private-key-path string       Path of the private key on the machine
  -p, --profile-arn string            ARN of the AWS IAM roles anywhere profile
  -n, --profile-name string           Name of the profile to that the credentials will be created under (default "default")
  -i, --role-arn string               ARN of the role to be assumed with AWS IAM roles Anywhere
  -r, --rolesanywhere-region string   Name of the region to that the credentials will be created under (default "eu-east-1")
  -t, --trust-anchor-arn string       ARN of the AWS IAM roles anywhere trust anchor

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)

```

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
      --access-key-acm string                AWS credentials access key for the Acm
      --access-key-pca string                AWS credentials access keyfor the Acm Pca
      --acm-region string                    Name of the region to be used for access to the ACM (default "eu-east-1")
      --acmpca-arn string                    ARN of the private CA that issues the certificate
      --certificate-directory string         Path of the certificate directory on the machine
      --certificate-expiry-days int          Number of days in which the certificate should expire (default 365)
      --common-name string                   The common name for the X509 certificate
      --country string                       The country name for the X509 certificate
  -h, --help                                 help for setup-all
      --locality string                      The locality name for the X509 certificate
      --organization-name string             The organization name for the X509 certificate
      --organizational-unit string           The organization unit for the X509 certificate
      --pca-region string                    Name of the region to be used for access to the PCA (default "eu-east-1")
      --profile-arn string                   ARN of the AWS IAM roles anywhere profile
      --profile-name-acm string              Name of the profile to be used for access to the ACM (default "default")
      --profile-name-pca string              Name of the profile to be used for access to the PCA (default "default")
      --profile-name-roles-anywhere string   Name of the profile to that the credentials will be created under (default "roles-anywhere")
      --province string                      The state or province name for the X509 certificate
      --role-arn string                      ARN of the role to be assumed with AWS IAM roles Anywhere
      --rolesanywhere-region string          Name of the region to that the credentials will be created under (default "eu-east-1")
      --secret-access-key-acm string         AWS credentials secret access key for the Acm
      --secret-access-key-pca string         AWS credentials secret access keyfor the Acm Pca
      --session-token-acm string             AWS credentials session token for the Acm
      --session-token-pca string             AWS credentials session tokenfor the Acm Pca
      --trust-anchor-arn string              ARN of the AWS IAM roles anywhere trust anchor

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)
```

#### generate-certificate

```
$ ./roles-anywhere-helper help generate-certificate 

Usage:
  roles-anywhere-helper generate-certificate [flags]

Flags:
      --access-key-pca string          AWS credentials access keyfor the Acm Pca
  -a, --acmpca-arn string              ARN of the private CA that issues the certificate
  -d, --certificate-directory string   Path of the certificate directory on the machine
  -e, --certificate-expiry-days int    Number of days in which the certificate should expire (default 365)
  -n, --common-name string             The common name for the X509 certificate
  -k, --country string                 The country name for the X509 certificate
  -h, --help                           help for generate-certificate
  -l, --locality string                The locality name for the X509 certificate
  -o, --organization-name string       The organization name for the X509 certificate
  -u, --organizational-unit string     The organization unit for the X509 certificate
      --pca-region string              Name of the region to be used for access to the PCA (default "eu-east-1")
      --profile-name-acm string        Name of the profile to be used for access to the ACM (default "default")
  -s, --province string                province
      --secret-access-key-pca string   AWS credentials secret access keyfor the Acm Pca
      --session-token-pca string       AWS credentials session tokenfor the Acm Pca

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)
```

#### import-certificate

```
$ ./roles-anywhere-helper help import-certificate 

Usage:
  roles-anywhere-helper import-certificate [flags]

Flags:
      --access-key-acm string          AWS credentials access key for the Acm
      --acm-region string              Name of the region to be used for access to the ACM (default "eu-east-1")
  -c, --certificate-arn string         ARN of the certificate to be revoked
  -d, --certificate-directory string   Name of the profile to that the credentials will be created under
  -h, --help                           help for import-certificate
      --profile-name string            Name of the profile to be used for access to the ACM (default "default")
      --secret-access-key-acm string   AWS credentials secret access key for the Acm
      --session-token-acm string       AWS credentials session token for the Acm

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)
```

#### revoke-certificate

```
$ ./roles-anywhere-helper help revoke-certificate

Usage:
  roles-anywhere-helper revoke-certificate [flags]

Flags:
      --access-key-pca string          AWS credentials access keyfor the Acm Pca
  -a, --acmpca-arn string              ARN of the private CA that issues the certificate
  -c, --certificate-arn string         ARN of the certificate to be revoked
  -h, --help                           help for revoke-certificate
      --pca-region string              Name of the region to be used for access to the PCA (default "eu-east-1")
      --profile-name string            Name of the profile to be used for access to the PCA (default "default")
  -r, --revocation-reason string       Reason why the certificate is revoked (default "UNSPECIFIED")
      --secret-access-key-pca string   AWS credentials secret access keyfor the Acm Pca
      --session-token-pca string       AWS credentials session tokenfor the Acm Pca

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)

```
#### rotate-certificate

```
$ ./roles-anywhere-helper help rotate-certificate

Usage:
  roles-anywhere-helper rotate-certificate [flags]

Flags:
      --access-key-acm string          AWS credentials access key for the Acm
      --access-key-pca string          AWS credentials access keyfor the Acm Pca
      --acm-region string              Name of the region to be used for access to the ACM (default "eu-east-1")
  -a, --acmpca-arn string              ARN of the private CA that issues the certificate
  -c, --certificate-arn string         ARN of the certificate to be revoked
  -d, --certificate-directory string   Name of the profile to that the credentials will be created under
  -e, --certificate-expiry-days int    Number of days in which the certificate should expire (default 365)
  -n, --common-name string             The common name for the X509 certificate
  -k, --country string                 The country name for the X509 certificate
  -h, --help                           help for rotate-certificate
  -l, --locality string                The locality name for the X509 certificate
  -o, --organization-name string       The organization name for the X509 certificate
  -u, --organizational-unit string     The organization unit for the X509 certificate
      --pca-region string              Name of the region to be used for access to the PCA (default "eu-east-1")
      --profile-name string            Name of the profile to be used for access to the PCA (default "default")
  -s, --province string                The state or province name for the X509 certificate
      --secret-access-key-acm string   AWS credentials secret access key for the Acm
      --secret-access-key-pca string   AWS credentials secret access keyfor the Acm Pca
      --session-token-acm string       AWS credentials session token for the Acm
      --session-token-pca string       AWS credentials session tokenfor the Acm Pca

Global Flags:
      --config string   config file (default is $HOME/.roles-anywhere-helper.yaml)
```

## Contributions

Contributions are welcome :)

* Feel free to contribute enhancements or bug fixes.
    * Fork this repo, apply your changes and create a PR pointing to this repo and the main branch
* If you have any ideas or suggestions please open an issue and describe your idea or feature request

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
