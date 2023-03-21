# iam-anywhere-ninja

## Introduction
Allows to easily setup AWS iam roles anywhere credential file on a machine, this CLI makes Assumtions that you are using a AWS Private CA, AWS Certificate Maanager and have already set IAM Roles Anywhere Correctly.

## Purpose

## Feautres

- Credential file setup to used iam roles for a given aws profile.
- Generate a x509 certificate on the machine and imports it to the Private Certificate Authority (CA)
- Import a certificate to Amazon certificate Manager (ACM)
- Rotate a certificate by generating and importing a new certificate then marking the old certificate as obsolite.
- Setup a machine to allow for IAM roles anywhere in a CICD pipeline includes, Creating Cred file and importing certificate.

## Prequistis 

- AWS_Signing_helper - Found here https://docs.aws.amazon.com/rolesanywhere/latest/userguide/credential-helper.html
- AWS Private CA provitioned with CLR for Credential Revocation
- AWS IAM Roles Anywhere configured.

## Getting Started

### Installation

### Usage

#### configure-credential

#### setup

#### generate-certificate

#### import-certificate

#### revoke-certificate

#### rotate-certificate
