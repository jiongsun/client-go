/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Message encapsulating the signature of the verified build.
type ApiBuildSignature struct {
	// Public key of the builder which can be used to verify that the related findings are valid and unchanged. If `key_type` is empty, this defaults to PEM encoded public keys.  This field may be empty if `key_id` references an external key.  For Cloud Container Builder based signatures, this is a PEM encoded public key. To verify the Cloud Container Builder signature, place the contents of this field into a file (public.pem). The signature field is base64-decoded into its binary representation in signature.bin, and the provenance bytes from `BuildDetails` are base64-decoded into a binary representation in signed.bin. OpenSSL can then verify the signature: `openssl sha256 -verify public.pem -signature signature.bin signed.bin`
	PublicKey string `json:"public_key,omitempty"`

	// Signature of the related `BuildProvenance`, encoded in a base64 string.
	Signature string `json:"signature,omitempty"`

	// An Id for the key used to sign. This could be either an Id for the key stored in `public_key` (such as the Id or fingerprint for a PGP key, or the CN for a cert), or a reference to an external key (such as a reference to a key in Cloud Key Management Service).
	KeyId string `json:"key_id,omitempty"`

	KeyType *BuildSignatureKeyType `json:"key_type,omitempty"`
}
