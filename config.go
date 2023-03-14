// Copyright © 2023 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cassandra

import "fmt"

//go:generate paramgen -output=paramgen_dest.go DestinationConfig

type DestinationConfig struct {
	// The keyspace (similar to a database in a relational database system) that has the table.
	Keyspace string `json:"keyspace" validate:"required"`
	// The table name.
	Table string `json:"table" validate:"required"`
	// The host to access Cassandra.
	Host string `json:"host" validate:"required"`
	// Cassandra’s TCP port.
	Port int `json:"port" default:"9042"`
	// Authentication mechanism used by Cassandra.
	AuthMechanism string `json:"auth.mechanism" validate:"inclusion=none|basic" default:"none"`
	// Username, only if basic auth is used.
	AuthUsername string `json:"auth.basic.username"`
	// Password, only if basic auth is used.
	AuthPassword string `json:"auth.basic.password"`
}

const (
	AuthMechanismBasic = "basic"
	AuthMechanismNone  = "none"
)

// validateConfig extra validations needed for destination config.
func (d *DestinationConfig) validateConfig() error {
	if d.AuthMechanism == AuthMechanismBasic && (d.AuthUsername == "" || d.AuthPassword == "") {
		return fmt.Errorf("auth.basic.username and auth.basic.password should be provided for basic authentication mechanism")
	}
	return nil
}
