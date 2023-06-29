package project

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/tekpriest/poprev/cmd/database"
	"github.com/tekpriest/poprev/internal/constants"
	"github.com/tekpriest/poprev/internal/model"
	"github.com/tekpriest/poprev/internal/query"
)

type ProjectService interface {
	CreateProject(data CreateProjectData) (*model.Project, error)
	ListProjects(query QueryProjectData) (FetchProjectsData, error)
	FetchProject(id string) (*model.Project, error)
	FetchAllProjectTransactions(projectID string) ([]model.Transaction, error)
}

type service struct {
	db *gorm.DB
}

// FetchAllProjectTransactions implements ProjectService.
func (s *service) FetchAllProjectTransactions(projectID string) ([]model.Transaction, error) {
	// var transactions []model.Transaction
	// var deposits []model.Deposit
	// var withdrawal []model.Withdrawal
	panic("")
}

// CreateProject implements ProjectService.
func (s *service) CreateProject(data CreateProjectData) (*model.Project, error) {
	var artist model.Artist

	if err := s.db.Table("artists").Attrs(model.Artist{
		Name: data.ArtistName,
	}).FirstOrCreate(&artist).Error; err != nil {
		return nil, err
	}

	tokens := calculateTokens(data.Amount)

	project := &model.Project{
		Amount:   data.Amount,
		Tokens:   tokens,
		ArtistID: artist.ID,
	}
	if err := s.db.Create(project).Error; err != nil {
		return nil, err
	}

	return s.FetchProject(project.ID)
}

// FetchProject implements ProjectService.
func (s *service) FetchProject(id string) (*model.Project, error) {
	var project model.Project

	if err := s.db.
		Table("projects").
		Preload(clause.Associations).
		First(&project, "id = ?", id).
		Error; err != nil {
		return nil, err
	}

	return &project, nil
}

// ListProjects implements ProjectService.
func (s *service) ListProjects(q QueryProjectData) (FetchProjectsData, error) {
	var projects []model.Project
	var data FetchProjectsData
	var count int64

	if err := s.db.Table("projects").Preload("Rate").
		Scopes(
			query.PaginateRows(int(q.Page), int(q.Limit)),
			query.FilterProjectByStatus(q.Status),
		).Order("created_at DESC").
		Find(&projects).
		Error; err != nil {
		return FetchProjectsData{}, err
	}

	if err := s.db.Table("projects").Count(&count).Scopes(
		query.FilterProjectByStatus(q.Status),
	).
		Error; err != nil {
		return FetchProjectsData{}, err
	}

	data.Projects = projects
	data.Meta = query.Paginate(count, len(projects), int(q.Page), int(q.Limit))

	return data, nil
}

func calculateTokens(amount float64) int {
	// get the rate for creating project
	tokens := int(amount / constants.RATE)

	return tokens
}

func NewProjectService(db database.DatabaseConnection) ProjectService {
	return &service{
		db: db.GetDB(),
	}
}
