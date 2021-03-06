// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElasticsearchClusterTopologyElement The topology of the Elasticsearch nodes, including the number, capacity, and type of nodes, and where they can be allocated.
// swagger:model ElasticsearchClusterTopologyElement
type ElasticsearchClusterTopologyElement struct {

	// DEPRECATED: Scheduled for removal in a future version of the API.
	//
	// Controls the allocation strategy of this node type using a simplified version of the Elasticsearch filter DSL (together with 'node_configuration')
	AllocatorFilter interface{} `json:"allocator_filter,omitempty"`

	// elasticsearch
	Elasticsearch *ElasticsearchConfiguration `json:"elasticsearch,omitempty"`

	// Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the id of an existing instance configuration.
	InstanceConfigurationID string `json:"instance_configuration_id,omitempty"`

	// The memory capacity in MB for each node of this type built in each zone.
	MemoryPerNode int32 `json:"memory_per_node,omitempty"`

	// DEPRECATED: Scheduled for removal in a future version of the API. Please use `instance_configuration_id`.
	//
	// Controls the allocation strategy of this node type by pointing to the names of pre-registered allocator settings. Unless otherwise specified for this deployment, only 'default' is supported (equivalent to omitting).
	NodeConfiguration string `json:"node_configuration,omitempty"`

	// The number of nodes of this type that are allocated within each zone. (i.e. total capacity per zone = `node_count_per_zone` * `memory_per_node` in MB). Cannot be set for tiebreaker topologies. For dedicated master nodes, must be 1 if an entry exists.
	NodeCountPerZone int32 `json:"node_count_per_zone,omitempty"`

	// node type
	NodeType *ElasticsearchNodeType `json:"node_type,omitempty"`

	// size
	Size *TopologySize `json:"size,omitempty"`

	// The default number of zones in which data nodes will be placed
	ZoneCount int32 `json:"zone_count,omitempty"`
}

// Validate validates this elasticsearch cluster topology element
func (m *ElasticsearchClusterTopologyElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElasticsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterTopologyElement) validateElasticsearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Elasticsearch) { // not required
		return nil
	}

	if m.Elasticsearch != nil {
		if err := m.Elasticsearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterTopologyElement) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeType) { // not required
		return nil
	}

	if m.NodeType != nil {
		if err := m.NodeType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_type")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterTopologyElement) validateSize(formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterTopologyElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterTopologyElement) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterTopologyElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
