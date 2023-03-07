// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/conduitio/conduit-connector-sdk/cmd/paramgen

package cassandra

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (DestinationConfig) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"auth.basic.password": {
			Default:     "",
			Description: "Password, only if basic auth is used.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"auth.basic.username": {
			Default:     "",
			Description: "Username, only if basic auth is used.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"auth.mechanism": {
			Default:     "none",
			Description: "Authentication mechanism used by Cassandra.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"none", "basic"}},
			},
		},
		"host": {
			Default:     "",
			Description: "The host to access Cassandra.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"keyspace": {
			Default:     "",
			Description: "The keyspace (similar to a database in a relational database system) that has the table.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
		"port": {
			Default:     "9042",
			Description: "Cassandra’s TCP port.",
			Type:        sdk.ParameterTypeInt,
			Validations: []sdk.Validation{},
		},
		"table": {
			Default:     "",
			Description: "The table name.",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
