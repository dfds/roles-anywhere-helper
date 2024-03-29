package flags

const ProfNameAcmPcaDesc string = "Name of the profile to be used for access to the PCA"
const ProfNameAcmDesc string = "Name of the profile to be used for access to the ACM"
const ProfNameRolesAnywhereDesc string = "Name of the profile to that the credentials will be created under"
const RegionNameAcmPcaDesc string = "Name of the region to be used for access to the PCA"
const RegionNameAcmDesc string = "Name of the region to be used for access to the ACM"

// Certificate flags description
const CertDirDesc string = "Path of the certificate directory on the machine"
const CertPathDesc string = "Path of the certificate on the machine"
const PrivKeyPathDesc string = "Path of the private key on the machine"
const CertArnDesc string = "ARN of the certificate to be revoked"

// // IAM Roles Anywhere flags description
const TrustAnchorArnDesc string = "ARN of the AWS IAM roles anywhere trust anchor"
const ProfileArnDesc string = "ARN of the AWS IAM roles anywhere profile"
const RoleArnDesc string = "ARN of the role to be assumed with AWS IAM roles Anywhere"
const AcmPcaArnDesc string = "ARN of the private CA that issues the certificate"
const RegionNameRolesAnywhereDesc string = "Name of the region to that the credentials will be created under"

// X.509 certificate flags attributes description
const CommonNameDesc string = "The common name for the X509 certificate"
const OrgNameDesc string = "The organization name for the X509 certificate"
const OrgUnitDesc string = "The organization unit for the X509 certificate"
const CountryDesc string = "The country name for the X509 certificate"
const LocalityDesc string = "The locality name for the X509 certificate"
const ProvinceDesc string = "The state or province name for the X509 certificate"

// Revocation flags description
const RevocReasonDesc string = "Reason why the certificate is revoked"
const CertificateExpiryDaysDesc string = "Number of days in which the certificate should expire"

// Cred Flags
const accessKeyDesc string = "AWS credentials access key"
const AccessKeyAcmDesc = accessKeyDesc + " for the Acm"
const AccessKeyAcmPcaDesc = accessKeyDesc + "for the Acm Pca"

const secretAccessKeyDesc string = "AWS credentials secret access key"
const SecretAccessKeyAcmDesc = secretAccessKeyDesc + " for the Acm"
const SecretAccessKeyAcmPcaDesc = secretAccessKeyDesc + "for the Acm Pca"

const sessionTokenDesc string = "AWS credentials session token"
const SessionTokenAcmDesc = sessionTokenDesc + " for the Acm"
const SessionTokenAcmPcaDesc = sessionTokenDesc + "for the Acm Pca"
