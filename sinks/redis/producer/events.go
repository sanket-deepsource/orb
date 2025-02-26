// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Adapted for Orb project, modifications licensed under MPL v. 2.0:
/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package producer

import (
	"time"
)

const (
	SinkPrefix = "sinks."
	SinkCreate = SinkPrefix + "create"
	SinkDelete = SinkPrefix + "delete"
)

type event interface {
	encode() map[string]interface{}
}

var (
	_ event = (*createSinkEvent)(nil)
)

type createSinkEvent struct {
	mfThing   string
	owner     string
	name      string
	content   string
	timestamp time.Time
}

func (cce createSinkEvent) encode() map[string]interface{} {
	return map[string]interface{}{
		"thing_id":  cce.mfThing,
		"owner":     cce.owner,
		"name":      cce.name,
		"content":   cce.content,
		"timestamp": cce.timestamp.Unix(),
		"operation": SinkCreate,
	}
}

type deleteSinkEvent struct {
	id string
}

func (dse deleteSinkEvent) Encode() map[string]interface{} {
	return map[string]interface{}{
		"id":        dse.id,
		"operation": SinkDelete,
	}
}
