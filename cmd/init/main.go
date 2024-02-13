package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"crowdfunding-campaign-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

const (
	dateOnly        = "2006-01-02"
	bothDateAndTime = "2006-01-02 15:04:05+00"

	queryInsertProject = `INSERT INTO project (
							project_id,
							project_name,
							slug,currency,
							project_type,
							goal,
							start_date,
							end_date
							)
							VALUES (%d,'%s','%s','%s','%s', %d,'%s','%s');`

	queryInsertContributions = `INSERT INTO contribution (
									id,
									amount,
									user_id,
									project_id,
									created_date
									) VALUES `
	contributionsValues = `(%d, %d, %d, %d, '%s')`

	queryInsertVisits = `INSERT INTO visit (
							project_id,
							visit_date,
							views,
							visitors
							) VALUES `
	visitsValues = `(%d, '%s', %d, %d)`

	batchSize = 100
)

func main() {
	ctx := context.Background()
	godotenv.Load(".env")
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Getenv("DB_LOGIN"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
	db, err := pgxpool.New(ctx, url)
	if err != nil {
		fmt.Printf("Unable to get new pool: %v\n", err)
	}

	// parse data
	projects, contributions, visits := parseDataSet()

	// write it in db
	initDataset(ctx, db, projects, contributions, visits)

	fmt.Println("data initialization done")
}

func initDataset(ctx context.Context, db *pgxpool.Pool, projects []*model.Project, contributions []*model.Contribution,
	visits []*model.Visit) {

	// insert projects
	for _, p := range projects {
		q := fmt.Sprintf(queryInsertProject, p.Id, p.Name, p.Slug, p.Currency, p.Type, p.Goal,
			p.StartDate.Format(dateOnly), p.EndDate.Format(bothDateAndTime))
		_, err := db.Exec(ctx, q)
		if err != nil {
			fmt.Printf("error inserting project %d: %v\n", p.Id, err)
		}
	}

	// insert contributions by batch
	conts := chunkContributions(contributions, batchSize)
	for i := 0; i < len(conts); i++ {
		query := queryInsertContributions
		for j, c := range conts[i] {
			vals := fmt.Sprintf(contributionsValues, c.Id, c.Amount, c.UserId, c.ProjectId,
				c.CreatedAt.Format(bothDateAndTime))
			query += vals
			if j != len(conts[i])-1 {
				query += ","
			}
		}
		query += ";"
		_, err := db.Exec(ctx, query)
		if err != nil {
			fmt.Printf("error inserting contribution batch %d: %v\n", i, err)
		}
	}

	// insert visits by batch
	vis := chunkVisits(visits, batchSize)
	for i := 0; i < len(vis); i++ {
		query := queryInsertVisits
		for j, visit := range vis[i] {
			vals := fmt.Sprintf(visitsValues, visit.ProjectId, visit.Date.Format(dateOnly), visit.Views, visit.Visitors)
			query += vals
			if j != len(vis[i])-1 {
				query += ","
			}
		}
		query += ";"
		_, err := db.Exec(ctx, query)
		if err != nil {
			fmt.Printf("error inserting contribution batch %d: %v\n", i, err)
		}
	}
}

func parseDataSet() ([]*model.Project, []*model.Contribution, []*model.Visit) {
	projs, err := parseProjects()
	if err != nil {
		fmt.Printf("error parsing projects: %s\n", err)
	}
	fmt.Println("finish parsing projects")

	conts, err := parseContributions()
	if err != nil {
		fmt.Printf("error parsing projects: %s\n", err)
	}
	fmt.Println("finish parsing contributions")

	visits, err := parseVisits()
	if err != nil {
		fmt.Printf("error parsing visits: %s\n", err)
	}
	fmt.Println("finish parsing visits")

	return projs, conts, visits
}

func parseProjects() ([]*model.Project, error) {
	path := "testdata/technical-test/dataset/projects.csv"
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %s", err)
	}

	r := csv.NewReader(f)
	r.Comma = '\t'
	r.LazyQuotes = true
	projs := make([]*model.Project, 0)
	for {
		line, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("error reading file: %s", err)
		}
		params := strings.Split(line[0], ",")
		// skip first line
		if params[0] == "id" {
			continue
		}

		proj, err := parseProject(params)
		if err != nil {
			return nil, err
		}
		projs = append(projs, proj)
	}
	return projs, nil
}

func parseProject(line []string) (*model.Project, error) {
	proj := &model.Project{}
	id, err := strconv.ParseInt(line[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing id : %s", err)
	}
	proj.Id = int(id)
	proj.Name = line[1]
	proj.Slug = line[2]
	proj.Currency = line[3]
	proj.Type = model.ProjectType(line[4])
	amt, err := strconv.ParseInt(line[5], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing goal : %s", err)
	}
	proj.Goal = int(amt)
	start, err := time.Parse(dateOnly, line[6])
	if err != nil {
		return nil, fmt.Errorf("error parsing start date : %s", err)
	}
	proj.StartDate = start
	end, err := time.Parse(bothDateAndTime, line[7])
	if err != nil {
		return nil, fmt.Errorf("error parsing end date : %s", err)
	}
	proj.EndDate = end
	return proj, nil
}

func parseContributions() ([]*model.Contribution, error) {
	path := "testdata/technical-test/dataset/contributions.csv"
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %s", err)
	}

	r := csv.NewReader(f)
	r.Comma = '\t'
	cont := make([]*model.Contribution, 0)
	for {
		line, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("error reading file: %s", err)
		}
		params := strings.Split(line[0], ",")
		// skip first line
		if params[0] == "id" {
			continue
		}
		c, err := parseContribution(params)
		if err != nil {
			return nil, err
		}
		cont = append(cont, c)
	}
	return cont, nil
}

func parseContribution(line []string) (*model.Contribution, error) {
	c := &model.Contribution{}
	id, err := strconv.ParseInt(line[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing id : %s", err)
	}
	c.Id = int(id)
	amt, err := strconv.ParseFloat(line[1], 10)
	if err != nil {
		return nil, fmt.Errorf("error parsing amount : %s", err)
	}
	c.Amount = int(amt * 100)
	userId, err := strconv.ParseInt(line[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing user id : %s", err)
	}
	c.UserId = int(userId)
	projectId, err := strconv.ParseInt(line[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing projectId id : %s", err)
	}
	c.ProjectId = int(projectId)

	creation, err := time.Parse(bothDateAndTime, line[4])
	if err != nil {
		return nil, fmt.Errorf("error parsing creation date : %s", err)
	}
	c.CreatedAt = creation
	return c, nil
}

func parseVisits() ([]*model.Visit, error) {
	path := "testdata/technical-test/dataset/visits.csv"
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %s", err)
	}

	r := csv.NewReader(f)
	r.Comma = '\t'
	visits := make([]*model.Visit, 0)
	for {
		line, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("error reading file: %s", err)
		}
		params := strings.Split(line[0], ",")
		// skip first line
		if params[0] == "project_id" {
			continue
		}
		v, err := parseVisit(params)
		if err != nil {
			return nil, err
		}
		visits = append(visits, v)
	}
	return visits, nil
}

func parseVisit(line []string) (*model.Visit, error) {
	v := &model.Visit{}
	id, err := strconv.ParseInt(line[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing project id : %s", err)
	}
	v.ProjectId = int(id)
	date, err := time.Parse(dateOnly, line[1])
	if err != nil {
		return nil, fmt.Errorf("error parsing date : %s", err)
	}
	v.Date = date
	views, err := strconv.ParseInt(line[2], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing views : %s", err)
	}
	v.Views = int(views)
	visitors, err := strconv.ParseInt(line[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing visitors : %s", err)
	}
	v.Visitors = int(visitors)
	return v, nil
}

func chunkContributions(slice []*model.Contribution, chunkSize int) [][]*model.Contribution {
	var chunks [][]*model.Contribution
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func chunkVisits(slice []*model.Visit, chunkSize int) [][]*model.Visit {
	var chunks [][]*model.Visit
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
