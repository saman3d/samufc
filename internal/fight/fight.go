// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package fight

import (
	"time"

	"github.com/saman3d/samufc/internal/fighter"
)

const (
	DefaultRoundDuration = time.Minute * 5
)

type Fight struct {
	ID                   uint64
	FightOrder           int
	Status               Status
	CardSegment          CardSegment
	CardSegmentStartTime time.Time
	Fighters             [2]fighter.Fighter
	Result               Result
	WeightClass          fighter.WeightClass
	Accolades            []Accolade
	Referee              Referee
	RuleSet              Rulset
	Tracking             []Action
}

type Action struct {
	ID      uint64
	Fighter fighter.Fighter
	Type    ActionType
}

type ActionType string

const (
	ActionTypeFightOpen                ActionType = "fight_open"
	ActionTypeWalkOut                  ActionType = "walkout"
	ActionTypeTaleOfTheTape            ActionType = "tale_of_the_tape"
	ActionTypeStareDown                ActionType = "staredown"
	ActionTypeStareRoundStart          ActionType = "round_start"
	ActionTypeStareKnockDown           ActionType = "knockdown"
	ActionTypeTakeDownAttempt          ActionType = "takedown_attempt"
	ActionTypeTakeDown                 ActionType = "takedown"
	ActionTypeRoundEnd                 ActionType = "round_end"
	ActionTypeSubmissionAttempt        ActionType = "submission_attempt"
	ActionTypeSubmission               ActionType = "submission"
	ActionTypeReversal                 ActionType = "reversal"
	ActionTypeResults                  ActionType = "fight_over"
	ActionTypeUnofficialWinnerDecision ActionType = "unofficial_winner_decision"
	ActionTypeFightComplete            ActionType = "fight_complete"
)

type Rulset struct {
	PossibleRounds int
}

type Referee struct {
	ID        uint64
	FirstName string
	LastName  string
}

type Accolade struct {
	Type AccoladeType
	Name string
}

type AccoladeType string

const (
	AccoladeTypeBelt AccoladeType = "Belt"
)

type Result struct {
	Method           ResultMethod
	EndingRound      int
	EndingTime       time.Time
	EndingStrike     ResultEndingStrike
	EndingTarget     ResultEndingTarget
	EndignPosition   ResultEndingPosition
	EndingSubmission ResultEndingSubmission
	EndingNotes      string
	FightOfTheNight  bool
	FightScores      []ResultScore
}

type ResultScore struct {
	By     Judge
	Scores []FighterWithScore
}

type FighterWithScore struct {
	Fighter fighter.Fighter
	Score   int
}

type Judge struct {
	ID        uint64
	FirstName string
	LastName  string
}

type ResultEndingSubmission string

const (
	ResultEndingSubmissionRearNakedChoke ResultEndingSubmission = "Rear Naked Choke"
)

type ResultEndingPosition string

const (
	ResultEndingPostionAtDistance      ResultEndingPosition = "At Distance"
	ResultEndingPostionFromBackControl ResultEndingPosition = "From Back Control"
)

type ResultEndingTarget string

const (
	ResultEndingTargetHead ResultEndingTarget = "Head"
)

type ResultEndingStrike string

const (
	ResultEndingStrikePunch ResultEndingStrike = "Punch"
)

type ResultMethod string

const (
	ResultMethodDecisionUnanimous = "Decision - Unanimous"
)

type Status string

const (
	StatusUpcoming Status = "Upcoming"
	StatusUnderway Status = "Underway"
	StatusFinal    Status = "Final"
)

type CardSegment string

const (
	CardSegmentTypeMain     CardSegment = "Main"
	CardSegmentTypePrelims1 CardSegment = "Prelims1"
	CardSegmentTypePrelims2 CardSegment = "Prelims2"
)
