# iam-anywhere-ninja

## Introduction
Allows to easily setup AWS iam roles anywhere credential file on a machine, this CLI makes Assumtions that you are using a AWS Private CA, AWS Certificate Maanager and have already set IAM Roles Anywhere Correctly. This was designed to work on Linux, Windows and Mac.

## Purpose
The purpose of this CLI is to make it really easy for development teams to configure, revoke and Rotate Credentials on the machine that needs access to AWS via IAM Roles anywhere. Idealy this would be used in the CICD pipeline. 

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
