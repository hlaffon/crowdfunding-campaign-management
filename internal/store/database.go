package store

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"

	"crowdfunding-campaign-management/model"
)

//go:generate mockery --dir=. --name Storage --with-expecter

type Storage interface {
	GetTotalAmountPerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error)
	GetNewContributionsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error)
	GetNewContributorsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error)
	GetPercentageOfGoalPerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error)
	GetAverageContribution(ctx context.Context, projectId int) (*float64, error)
	GetVisitsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error)
	GetConversionRatePerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error)
	GetMilestoneContribution(ctx context.Context, projectId, milestone int) (*model.MilestoneContribution, error)
}

const (
	queryGetTotalAmountPerDayForProject = `WITH cte_contribution_per_day AS(
			SELECT 
				created_date::TIMESTAMP::DATE AS date,
				amount
				FROM contribution
			WHERE project_id = %d
			)
			SELECT 
					date,
					SUM(amount)/100 AS total
					FROM cte_contribution_per_day
					GROUP BY date
					ORDER BY date ASC;`

	queryGetContributionNumberPerDayForProject = `WITH cte_contribution_per_day AS(
			SELECT 
				created_date::TIMESTAMP::DATE AS date
				FROM contribution
			WHERE project_id = %d
			)
			SELECT 
					date,
					COUNT(*) AS total
					FROM cte_contribution_per_day
					GROUP BY date
					ORDER BY date ASC;`

	queryGetContributorsNumberPerDayForProject = `WITH cte_contribution_per_day AS(
			SELECT 
				created_date::TIMESTAMP::DATE AS date,
				user_id,
				COUNT(user_id)
				FROM contribution
			WHERE project_id = %d
			GROUP BY date, user_id
			)
			SELECT 
					date,
					COUNT(*) AS total
					FROM cte_contribution_per_day
							GROUP BY date
					ORDER BY date ASC;`

	queryGetPercentageOfGoalPerDay = `WITH cte_contribution_per_day AS(
			SELECT 
				created_date::TIMESTAMP::DATE AS date,
				amount,
				goal,
				project_type
				FROM contribution
			INNER JOIN project USING (project_id)
			WHERE project_id = %d
			)
			SELECT 
				CASE project_type
					WHEN 'project' THEN
						SUM(amount)*1./goal
					ELSE
						COUNT(*)*100./goal
					END AS total,
					date
			FROM cte_contribution_per_day
			GROUP BY (date, goal, project_type)
			ORDER BY date ASC;`

	queryGetAverageContribution = `SELECT AVG(amount/100) FROM contribution WHERE project_id=%d;`

	queryGetViewsPerDay = `SELECT 
						views AS total, 
						visit_date AS date
						FROM visit
						WHERE project_id = %d 
						GROUP BY (views, visit_date) 
						ORDER BY visit_date ASC;`

	queryGetConversionRatePerDay = `WITH cte_contribution_per_day AS(
									SELECT 
									id,
									created_date::TIMESTAMP::DATE AS contribution_date,
									project_id
									FROM contribution
									WHERE project_id = %d
									),
									cte_views_per_day AS(
									SELECT 
									views, 
									visit_date::TIMESTAMP::DATE AS view_date,
									project_id
									FROM visit
									WHERE project_id = %d
									)
							SELECT 
									contribution_date AS date,
									COUNT(id)*1./views AS total
									FROM cte_contribution_per_day
									INNER JOIN cte_views_per_day USING (project_id)
									WHERE contribution_date=view_date
									GROUP BY (contribution_date, view_date, views)
									ORDER BY contribution_date ASC;`

	queryGetContributionMilestone = `WITH cte_table AS(
									SELECT goal,
									project_type,
									amount/100 AS amount,
									created_date,
									id,
									user_id,
									SUM(amount/100) OVER(
										ORDER BY created_date
									) AS total_raised,
									COUNT(*) OVER(
										ORDER BY created_date
									) AS contribution_nb
									FROM contribution
									INNER JOIN project using(project_id)
									WHERE project_id=%d
									ORDER BY created_date ASC
									)
									SELECT amount,
									created_date,
									id AS contribution_id,
									user_id,
									total_raised,
									contribution_nb,
									goal
									FROM cte_table
									WHERE
										CASE project_type
											WHEN 'project' THEN
											total_raised*100./goal >= %d
											ELSE
											contribution_nb*100./goal >= %d
										END
									LIMIT 1;`
)

func (db *Database) GetTotalAmountPerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error) {
	var amntRaisedDuringCampaign []*model.AmountPerDay
	query := fmt.Sprintf(queryGetTotalAmountPerDayForProject, projectId)
	err := pgxscan.Select(ctx, db.pool, &amntRaisedDuringCampaign, query)
	if err != nil {
		return nil, err
	}
	return amntRaisedDuringCampaign, nil
}

func (db *Database) GetNewContributionsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	var contributionsNumber []*model.NumberPerDay
	query := fmt.Sprintf(queryGetContributionNumberPerDayForProject, projectId)
	err := pgxscan.Select(ctx, db.pool, &contributionsNumber, query)
	if err != nil {
		return nil, err
	}
	return contributionsNumber, nil
}

func (db *Database) GetNewContributorsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	var contributionsNumber []*model.NumberPerDay
	query := fmt.Sprintf(queryGetContributorsNumberPerDayForProject, projectId)
	err := pgxscan.Select(ctx, db.pool, &contributionsNumber, query)
	if err != nil {
		return nil, err
	}
	return contributionsNumber, nil
}

func (db *Database) GetPercentageOfGoalPerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error) {
	var percentagePerDay []*model.AmountPerDay
	query := fmt.Sprintf(queryGetPercentageOfGoalPerDay, projectId)
	err := pgxscan.Select(ctx, db.pool, &percentagePerDay, query)
	if err != nil {
		return nil, err
	}
	return percentagePerDay, nil
}

func (db *Database) GetAverageContribution(ctx context.Context, projectId int) (*float64, error) {
	var averageAmount *float64
	query := fmt.Sprintf(queryGetAverageContribution, projectId)
	err := pgxscan.Get(ctx, db.pool, &averageAmount, query)
	if err != nil {
		return nil, err
	}
	return averageAmount, nil
}

func (db *Database) GetVisitsPerDay(ctx context.Context, projectId int) ([]*model.NumberPerDay, error) {
	var viewsPerDay []*model.NumberPerDay
	query := fmt.Sprintf(queryGetViewsPerDay, projectId)
	err := pgxscan.Select(ctx, db.pool, &viewsPerDay, query)
	if err != nil {
		return nil, err
	}
	return viewsPerDay, nil
}

func (db *Database) GetConversionRatePerDay(ctx context.Context, projectId int) ([]*model.AmountPerDay, error) {
	var conversionRatePerDay []*model.AmountPerDay
	query := fmt.Sprintf(queryGetConversionRatePerDay, projectId, projectId)
	err := pgxscan.Select(ctx, db.pool, &conversionRatePerDay, query)
	if err != nil {
		return nil, err
	}
	return conversionRatePerDay, nil
}

func (db *Database) GetMilestoneContribution(ctx context.Context, projectId, milestone int) (*model.MilestoneContribution, error) {
	var contributions []*model.MilestoneContribution
	query := fmt.Sprintf(queryGetContributionMilestone, projectId, milestone, milestone)
	err := pgxscan.Select(ctx, db.pool, &contributions, query)
	if err != nil {
		return nil, err
	}
	if contributions == nil {
		return nil, nil
	}
	contributions[0].Milestone = milestone
	return contributions[0], nil
}

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(pool *pgxpool.Pool) *Database {
	return &Database{pool: pool}
}
