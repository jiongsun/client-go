/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Basis describes the base image portion (Note) of the DockerImage relationship.  Linked occurrences are derived from this or an equivalent image via:   FROM <Basis.resource_url> Or an equivalent reference, e.g. a tag of the resource_url.
type DockerImageBasis struct {
	// The resource_url for the resource representing the basis of associated occurrence images.
	ResourceUrl string `json:"resource_url,omitempty"`

	Fingerprint *DockerImageFingerprint `json:"fingerprint,omitempty"`
}
