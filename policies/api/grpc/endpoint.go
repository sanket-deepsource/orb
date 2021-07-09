// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package grpc

import (
	"context"
	"encoding/json"
	"github.com/ns1labs/orb/policies"

	"github.com/go-kit/kit/endpoint"
)

func retrievePolicyEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(accessByIDReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		policy, err := svc.RetrievePolicyByIDInternal(ctx, req.PolicyID, req.OwnerID)
		if err != nil {
			return policyRes{}, err
		}
		data, err := json.Marshal(policy.Policy)
		if err != nil {
			return policyRes{}, err
		}
		return policyRes{
			name:    policy.Name.String(),
			backend: policy.Backend,
			version: policy.Version,
			data:    data,
		}, nil
	}
}
