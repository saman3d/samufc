// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fighter

import (
	"github.com/saman3d/samufc/internal/common"
	"net/url"
	"time"
)

type Fighter struct {
	Profile               Profile
	Corner                CornerType
	WeighIn               int
	PerformanceOfTheNight bool
	SubmissionOfTheNight  bool
	KOOfTheNight          bool
}

type CornerType string

const (
	CornerTypeRed  = "Red"
	CornerTypeBlue = "Blue"
)

type Profile struct {
	ID            uint64
	MMAID         uint64
	Name          Name
	Born          *common.Location
	FightingOutOf *common.Location
	Record        Record
	DateOfBirth   time.Time
	Weight        int
	Height        int
	ArmReach      int
	LegReach      int
	UFCLink       *url.URL
	WeightClasses []WeightClass
}

type StanceType string

const (
	StanceTypeSouthpaw = "Southpaw"
	StanceTypeOrthodox = "Orthodox"
)

type BirthInfo struct {
	Location common.Location
}

type Record struct {
	Wins       int
	Losses     int
	Draws      int
	NoContests int
}

type Name struct {
	FirstName string
	LastName  string
	NickName  string
}

type WeightClass struct {
	ID           uint64
	Order        int
	Description  string
	Abbreviation string
}
