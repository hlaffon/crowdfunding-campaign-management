package model

import "time"

type Project struct {
	Id               int
	Name             string
	Slug             string
	Currency         string
	Type             ProjectType
	Goal             int
	StartDate        time.Time
	EndDate          time.Time
	EndDateExtraTime time.Time
}

type ProjectType string

const (
	TypeProject ProjectType = "project"
	TypePresale ProjectType = "presale"
)

type Contribution struct {
	Id        int
	Amount    int
	UserId    int
	ProjectId int
	CreatedAt time.Time
}

type Visit struct {
	Id        int
	ProjectId int
	Date      time.Time
	Views     int
	Visitors  int
}

// swagger:model AmountPerDay
type AmountPerDay struct {
	Date  time.Time `json:"date"`
	Total float64   `json:"total"`
}

// swagger:model NumberPerDay
type NumberPerDay struct {
	Date  time.Time `json:"date"`
	Total int       `json:"total"`
}

// swagger:model MilestoneContribution
type MilestoneContribution struct {
	CreatedDate    time.Time `json:"date"`
	Milestone      int       `json:"milestone"`
	ContributionId int       `json:"contributionId"`
	UserId         int       `json:"userId"`
	Amount         float64   `json:"amount"`
	Goal           int       `json:"goal"`
	TotalRaised    float64   `json:"totalRaised"`
	ContributionNb int       `json:"contributionNb"`
}

// swagger:model HTTPError
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"an error occurred"`
}
