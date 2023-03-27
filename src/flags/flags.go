package flags

// AWS profile name flags
const ProfileName string = "profile-name"
const ProfileNameAcm = ProfileName + "-acm"
const ProfileNameAcmPca = ProfileName + "-pca"
const ProfileNameRolesAnywhere = ProfileName + "-roles-anywhere"

// Certificate flags
const CertificateDirectory string = "certificate-directory"
const CertificatePath string = "certificate-path"
const PrivateKeyPath string = "private-key-path"
const CertificateArn string = "certificate-arn"

// IAM Roles Anywhere flags
const TrustAnchor string = "trust-anchor-arn"
const ProfileArn string = "profile-arn"
const RoleArn string = "role-arn"
const Region string = "region"

const AcmpcaArn string = "acmpca-arn"

// X.509 certificate flags attributes flags 
const CommonName string = "common-name"
const OrganizationName string = "organization-name"
const OrganizationalUnit string = "organizational-unit"
const Country string = "country"
const Locality string = "locality"
const Province string = "province"

const RevocationReason string = "revocation-reason"
