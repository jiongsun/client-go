/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A set of properties that uniquely identify a given Docker image.
type DockerImageFingerprint struct {
	// The layer-id of the final layer in the Docker image's v1 representation. This field can be used as a filter in list requests.
	V1Name string `json:"v1_name,omitempty"`

	// The ordered list of v2 blobs that represent a given image.
	V2Blob []string `json:"v2_blob,omitempty"`

	// Output only. The name of the image's v2 blobs computed via:   [bottom] := v2_blob[bottom]   [N] := sha256(v2_blob[N] + \" \" + v2_name[N+1]) Only the name of the final blob is kept. This field can be used as a filter in list requests.
	V2Name string `json:"v2_name,omitempty"`
}
