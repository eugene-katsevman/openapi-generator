/*
 * Swagger Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package petstore

type AdditionalPropertiesClass struct {
	MapProperty map[string]string `json:"map_property,omitempty" xml:"map_property"`

	MapOfMapProperty map[string]map[string]string `json:"map_of_map_property,omitempty" xml:"map_of_map_property"`
}