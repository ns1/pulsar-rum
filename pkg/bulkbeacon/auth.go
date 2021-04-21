// Copyright 2021 NSONE, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bulkbeacon

import "context"

// Auth it's a simple structure to manage the authentication. It implements
// grpc.PerRPCCredentials interface.
type Auth struct {
	key string
}

// GetRequestMetadata sets the authentication key into the metadata map.
func (a Auth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	m := map[string]string{
		"X-NSONE-Key": a.key,
	}
	return m, nil
}

// RequireTransportSecurity must return true if we are using TLS.
func (a Auth) RequireTransportSecurity() bool {
	return true
}

// NewAuth is a simple constructor for Auth. It takes the NS1 api key as
// argument.
func NewAuth(key string) Auth {
	return Auth{key: key}
}
