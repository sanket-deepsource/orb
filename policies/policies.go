/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package policies

import (
	"context"
	"github.com/ns1labs/orb/pkg/types"
	"time"
)

type Policy struct {
	ID          string
	Name        types.Identifier
	Description string
	MFOwnerID   string
	Backend     string
	Version     int32
	OrbTags     types.Tags
	Policy      types.Metadata
	Created     time.Time
}

type Dataset struct {
	ID           string
	Name         types.Identifier
	MFOwnerID    string
	Valid        bool
	AgentGroupID string
	PolicyID     string
	Metadata     types.Metadata
	Created      time.Time
	Tags         types.Tags
	SinkIDs      []string
}

type PolicyInDataset struct {
	Policy
	DatasetID string
}

type Page struct {
	PageMetadata
	Policies []Policy
}

type PageDataset struct {
	PageMetadata
	Datasets []Dataset
}

type Service interface {
	// AddPolicy creates new agent Policy
	AddPolicy(ctx context.Context, token string, p Policy, format string, policyData string) (Policy, error)

	// ViewPolicyByID retrieving policy by id with token
	ViewPolicyByID(ctx context.Context, token string, policyID string) (Policy, error)

	// ListPolicies
	ListPolicies(ctx context.Context, token string, pm PageMetadata) (Page, error)

	// ViewPolicyByIDInternal gRPC version of retrieving policy by id with no token
	ViewPolicyByIDInternal(ctx context.Context, policyID string, ownerID string) (Policy, error)

	// ListPoliciesByGroupIDInternal gRPC version of retrieving list of policies belonging to specified agent group with no token
	ListPoliciesByGroupIDInternal(ctx context.Context, groupIDs []string, ownerID string) ([]PolicyInDataset, error)

	// EditPolicy edit a existing policy by id with a valid token
	EditPolicy(ctx context.Context, token string, pol Policy, format string, policyData string) (Policy, error)

	// RemovePolicy remove a existing policy owned by the specified user
	RemovePolicy(ctx context.Context, token string, policyID string) error

	// AddDataset creates new Dataset
	AddDataset(ctx context.Context, token string, d Dataset) (Dataset, error)

	// InactivateDatasetByGroupID inactivate a dataset
	InactivateDatasetByGroupID(ctx context.Context, groupID string, token string) error

	// ListDatasetsByPolicyIDInternal retrieves the subset of Datasets by policyID owned by the specified user
	ListDatasetsByPolicyIDInternal(ctx context.Context, policyID string, token string) ([]Dataset, error)

	// ValidatePolicy validates an agent Policy without saving
	ValidatePolicy(ctx context.Context, token string, p Policy, format string, policyData string) (Policy, error)

	// EditDataset edit a existing dataset by id with a valid token
	EditDataset(ctx context.Context, token string, ds Dataset) (Dataset, error)
	// RemoveDataset remove a dataset by id with a valid token
	RemoveDataset(ctx context.Context, token string, dsID string) error

	// ValidateDataset validates a new Dataset without saving it
	ValidateDataset(ctx context.Context, token string, d Dataset) (Dataset, error)

	// ViewDatasetByID retrieving dataset by id with token
	ViewDatasetByID(ctx context.Context, token string, datasetID string) (Dataset, error)

	// ListDatasets retrieve a list of Dataset by owner
	ListDatasets(ctx context.Context, token string, pm PageMetadata) (PageDataset, error)
}

type Repository interface {
	// SavePolicy persists a Policy. Successful operation is indicated by non-nil
	// error response.
	SavePolicy(ctx context.Context, policy Policy) (string, error)

	// RetrievePolicyByID Retrieve policy by id
	RetrievePolicyByID(ctx context.Context, policyID string, ownerID string) (Policy, error)

	// RetrievePoliciesByGroupID Retrieve policy list by group id
	RetrievePoliciesByGroupID(ctx context.Context, groupIDs []string, ownerID string) ([]PolicyInDataset, error)

	// RetrieveAll retrieves the subset of Policies owned by the specified user
	RetrieveAll(ctx context.Context, ownerID string, pm PageMetadata) (Page, error)

	// UpdatePolicy update a existing policy by id with a valid token
	UpdatePolicy(ctx context.Context, ownerID string, pol Policy) error

	// DeletePolicy a existing policy by id owned by the specified user
	DeletePolicy(ctx context.Context, ownerID string, policyID string) error

	// SaveDataset persists a Dataset. Successful operation is indicated by non-nil
	// error response.
	SaveDataset(ctx context.Context, dataset Dataset) (string, error)

	// InactivateDatasetByGroupID inactivate a dataset
	InactivateDatasetByGroupID(ctx context.Context, groupID string, ownerID string) error

	// InactivateDatasetByPolicyID inactivate a dataset by policy id
	InactivateDatasetByPolicyID(ctx context.Context, policyID string, ownerID string) error

	// RetrieveDatasetsByPolicyID retrieves the subset of Datasets by policyID owned by the specified user
	RetrieveDatasetsByPolicyID(ctx context.Context, policyID string, ownerID string) ([]Dataset, error)

	// UpdateDataset update a existing dataset by id with a valid token
	UpdateDataset(ctx context.Context, ownerID string, ds Dataset) error

	//DeleteDataset delete a existing dataset by id by ownerID
	DeleteDataset(ctx context.Context, ownerID string, dsID string) error

	// RetrieveDatasetByID Retrieves dataset by id
	RetrieveDatasetByID(ctx context.Context, datasetID string, ownerID string) (Dataset, error)

	// RetrieveAllDatasetsByOwner retrieves the subset of Datasets owned by the specified user
	RetrieveAllDatasetsByOwner(ctx context.Context, ownerID string, pm PageMetadata) (PageDataset, error)
}