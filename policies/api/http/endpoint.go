// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/ns1labs/orb/pkg/types"
	"github.com/ns1labs/orb/policies"
)

func addPolicyEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPolicyReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		nID, err := types.NewIdentifier(req.Name)
		if err != nil {
			return nil, err
		}

		policy := policies.Policy{
			Name:        nID,
			Backend:     req.Backend,
			Policy:      req.Policy,
			Description: req.Description,
			OrbTags:     req.Tags,
		}

		saved, err := svc.AddPolicy(ctx, req.token, policy, req.Format, req.PolicyData)
		if err != nil {
			return nil, err
		}

		res := policyRes{
			ID:          saved.ID,
			Name:        saved.Name.String(),
			Description: saved.Description,
			Tags:        saved.OrbTags,
			Backend:     saved.Backend,
			Policy:      saved.Policy,
			Version:     saved.Version,
			created:     true,
		}

		return res, nil
	}
}

func viewPolicyEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(viewResourceReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		policy, err := svc.ViewPolicyByID(ctx, req.token, req.id)
		if err != nil {
			return nil, err
		}

		res := policyRes{
			ID:          policy.ID,
			Name:        policy.Name.String(),
			Description: policy.Description,
			Tags:        policy.OrbTags,
			Backend:     policy.Backend,
			Policy:      policy.Policy,
			Version:     policy.Version,
		}
		return res, nil
	}
}

func listPoliciesEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listResourcesReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		page, err := svc.ListPolicies(ctx, req.token, req.pageMetadata)
		if err != nil {
			return nil, err
		}

		res := policiesPageRes{
			pageRes: pageRes{
				Total:  page.Total,
				Offset: page.Offset,
				Limit:  page.Limit,
				Order:  page.Order,
				Dir:    page.Dir,
			},
			Policies: []policyRes{},
		}
		for _, ag := range page.Policies {
			view := policyRes{
				ID:          ag.ID,
				Name:        ag.Name.String(),
				Description: ag.Description,
				Version:     ag.Version,
				Backend:     ag.Backend,
			}
			res.Policies = append(res.Policies, view)
		}
		return res, nil
	}
}

func editPoliciyEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updatePolicyReq)
		if err := req.validate(); err != nil {
			return policyUpdateRes{}, err
		}

		nameID, err := types.NewIdentifier(req.Name)
		if err != nil {
			return policyUpdateRes{}, err
		}
		plcy := policies.Policy{
			ID:          req.id,
			Name:        nameID,
			Description: req.Description,
			OrbTags:     req.Tags,
			Policy:      req.Policy,
		}

		res, err := svc.EditPolicy(ctx, req.token, plcy, req.Format, req.PolicyData)
		if err != nil {
			return policyUpdateRes{}, err
		}

		plcyRes := policyUpdateRes{
			ID:          res.ID,
			Name:        res.Name.String(),
			Description: res.Description,
			Tags:        res.OrbTags,
			Policy:      res.Policy,
			Version:     res.Version,
		}

		return plcyRes, nil
	}
}

func removePolicyEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(viewResourceReq)
		if err := req.validate(); err != nil {
			return removeRes{}, err
		}
		err = svc.RemovePolicy(ctx, req.token, req.id)
		if err != nil {
			return removeRes{}, err
		}
		return removeRes{}, nil
	}
}

func addDatasetEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addDatasetReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		nID, err := types.NewIdentifier(req.Name)
		if err != nil {
			return nil, err
		}

		d := policies.Dataset{
			Name:         nID,
			AgentGroupID: req.AgentGroupID,
			PolicyID:     req.PolicyID,
			SinkIDs:      req.SinkIDs,
		}

		saved, err := svc.AddDataset(ctx, req.token, d)
		if err != nil {
			return nil, err
		}

		res := datasetRes{
			ID:           saved.ID,
			Name:         saved.Name.String(),
			Valid:        saved.Valid,
			AgentGroupID: saved.AgentGroupID,
			PolicyID:     saved.PolicyID,
			SinkIDs:      saved.SinkIDs,
			Metadata:     saved.Metadata,
			TsCreated:    saved.Created,
			Tags:         saved.Tags,
			created:      true,
		}

		return res, nil
	}
}

func editDatasetEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(updateDatasetReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		dataset := policies.Dataset{
			ID:      req.id,
			Tags:    req.Tags,
			SinkIDs: req.SinkIDs,
		}

		ds, err := svc.EditDataset(ctx, req.token, dataset)
		if err != nil {
			return nil, err
		}
		return ds, nil
	}
}

func validatePolicyEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addPolicyReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		nID, err := types.NewIdentifier(req.Name)
		if err != nil {
			return nil, err
		}

		policy := policies.Policy{
			Name:        nID,
			Backend:     req.Backend,
			Policy:      req.Policy,
			OrbTags:     req.Tags,
			Description: req.Description,
		}

		validated, err := svc.ValidatePolicy(ctx, req.token, policy, req.Format, req.PolicyData)
		if err != nil {
			return nil, err
		}

		res := policyValidateRes{
			Name:        validated.Name.String(),
			Backend:     validated.Backend,
			Tags:        validated.OrbTags,
			Policy:      validated.Policy,
			Description: validated.Description,
		}

		return res, nil
	}
}

func removeDatasetEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(viewResourceReq)
		if err := req.validate(); err != nil {
			return removeRes{}, err
		}
		if err := svc.RemoveDataset(ctx, req.token, req.id); err != nil {
			return removeRes{}, err
		}
		return removeRes{}, nil
	}
}

func validateDatasetEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addDatasetReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		nID, err := types.NewIdentifier(req.Name)
		if err != nil {
			return nil, err
		}

		d := policies.Dataset{
			Name:         nID,
			AgentGroupID: req.AgentGroupID,
			PolicyID:     req.PolicyID,
			SinkIDs:      req.SinkIDs,
			Tags:         req.Tags,
		}

		validated, err := svc.ValidateDataset(ctx, req.token, d)
		if err != nil {
			return nil, err
		}

		res := validateDatasetRes{
			Name:         validated.Name.String(),
			Valid:        true,
			Tags:         validated.Tags,
			AgentGroupID: validated.AgentGroupID,
			PolicyID:     validated.PolicyID,
			SinkIDs:      validated.SinkIDs,
		}

		return res, nil
	}
}

func viewDatasetEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(viewResourceReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		dataset, err := svc.ViewDatasetByID(ctx, req.token, req.id)
		if err != nil {
			return nil, err
		}

		res := datasetRes{
			ID:           dataset.ID,
			Name:         dataset.Name.String(),
			PolicyID:     dataset.PolicyID,
			SinkIDs:      dataset.SinkIDs,
			AgentGroupID: dataset.AgentGroupID,
		}
		return res, nil
	}
}

func listDatasetEndpoint(svc policies.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listResourcesReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		page, err := svc.ListDatasets(ctx, req.token, req.pageMetadata)
		if err != nil {
			return nil, err
		}

		res := datasetPageRes{
			pageRes: pageRes{
				Total:  page.Total,
				Offset: page.Offset,
				Limit:  page.Limit,
				Order:  page.Order,
				Dir:    page.Dir,
			},
			Datasets: []datasetRes{},
		}
		for _, dataset := range page.Datasets {
			view := datasetRes{
				ID:           dataset.ID,
				Name:         dataset.Name.String(),
				PolicyID:     dataset.PolicyID,
				SinkIDs:      dataset.SinkIDs,
				AgentGroupID: dataset.AgentGroupID,
			}
			res.Datasets = append(res.Datasets, view)
		}
		return res, nil
	}
}
