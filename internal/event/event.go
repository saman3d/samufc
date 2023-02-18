// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package event

import (
	"github.com/saman3d/samufc/internal/common"
	"time"
)

type Event struct {
	ID           uint64
	Name         string
	UFCID        uint64
	StartTime    time.Time
	TimeZone     time.Location
	Status       Status
	Live         LiveEvent
	Organization Organization
	Venue        Venue
}

type Venue struct {
	ID       uint64
	Name     string
	Locaiton common.Location
}

type Organization struct {
	ID   uint64
	Name string
}

type LiveEvent struct {
	ID               uint64
	FightID          uint64
	RoundNumber      uint64
	RoundElapsedTime uint64
}

type Status string

const (
	StatusFinal = "Final"
)
