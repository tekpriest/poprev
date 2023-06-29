package token

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/tekpriest/poprev/cmd/api/project"
	"github.com/tekpriest/poprev/cmd/database"
	"github.com/tekpriest/poprev/internal/constants"
	"github.com/tekpriest/poprev/internal/model"
	"github.com/tekpriest/poprev/internal/query"
)

type TokenService interface {
	BuyToken(userID string, data BuyTokenData) (map[string]interface{}, error)
	FetchAllTokens(userID string, query QueryUserTokensData) (FetchTokensData, error)
	CreateSale(userID string, data CreateSaleData) (*model.Sale, error)
	FetchSales(userID string, query QueryUserSalesData) (FetchSalesData, error)
	FetchSale(ID string) (*model.Sale, error)
}

type service struct {
	db *gorm.DB
	ps project.ProjectService
}

// FetchSale implements TokenService.
func (s *service) FetchSale(ID string) (*model.Sale, error) {
	var sale model.Sale

	if err := s.db.First(&sale, "id = ?", ID).Error; err != nil {
		return nil, err
	}

	return &sale, nil
}

// FetchSales implements TokenService.
func (s *service) FetchSales(userID string, q QueryUserSalesData) (FetchSalesData, error) {
	var count int64
	var sales []model.Sale
	var data FetchSalesData

	if err := s.db.Table("sales").
		Where("seller_id = ?", userID).
		Scopes(
			query.PaginateRows(int(q.Page), int(q.Limit)),
		).
		Order("created_at DESC").
		Find(&sales).
		Error; err != nil {
		return data, err
	}

	if err := s.db.Table("sales").
		Where("seller_id = ?", userID).
		Count(&count).
		Error; err != nil {
		return data, err
	}

	data.Sales = sales
	data.Meta = query.Paginate(count, len(sales), int(q.Page), int(q.Limit))

	return data, nil
}

// CreateSale implements TokenService.
func (s *service) CreateSale(userID string, data CreateSaleData) (*model.Sale, error) {
	var token model.Token

	if err := s.db.First(&token, "id = ?", data.TokenID).Error; err != nil {
		return nil, err
	}

	sale := &model.Sale{
		Quantity: data.Quantity,
		MinOrder: data.MinOrder,
		MaxOrder: data.MaxOrder,
		TokenID:  data.TokenID,
		SellerID: userID,
	}

	if err := s.db.Create(sale).Error; err != nil {
		return nil, err
	}

	return sale, nil
}

// BuyToken implements TokenService.
func (s *service) BuyToken(userID string, data BuyTokenData) (map[string]interface{}, error) {
	totalClaimed := 0
	project, err := s.ps.FetchProject(data.ProjectID)
	if err != nil {
		return nil, err
	}

	// simple ledger system to check bought token count
	if len(project.Claimed) > 0 {
		for _, token := range project.Claimed {
			totalClaimed += token.Quantity
		}
	}

	tokenBalance := project.Tokens - totalClaimed
	if tokenBalance-data.Quantity < 0 {
		return nil, errors.New("Insufficient tokens to buy")
	}

	// calculating amount based on the rate
	finalAmount := float64(data.Quantity) * float64(project.Rate.Buy) * constants.TRANSACTION_FEE

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		deposit := &model.Deposit{
			Amount:         finalAmount,
			UserID:         userID,
			TokenRequested: data.Quantity,
			ProjectID:      data.ProjectID,
		}
		if err := tx.Create(deposit).Error; err != nil {
			return err
		}
		transaction := &model.Transaction{
			Fee:       constants.TRANSACTION_FEE,
			Amount:    finalAmount,
			EventID:   deposit.ID,
			ProjectID: data.ProjectID,
			Type:      model.TransactionEventType(model.DepositEvent),
		}
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// creating a token now but ideally, I often use webhooks to
		// make final operations when the transaction is confirmed
		token := &model.Token{
			Amount:    finalAmount,
			Quantity:  data.Quantity,
			ProjectID: data.ProjectID,
			UserID:    userID,
		}
		if err := tx.Create(token).Error; err != nil {
			return err
		}

		// do a fake transaction update assuming deposit went through
		if err := tx.
			Table("transactions").
			Where("id = ?", transaction.ID).
			Update("status",
				model.TransactionStatus(model.TransactionSuccessful)).
			Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"note":          fmt.Sprintf("Please make payment of %f", finalAmount),
		"accountNumber": "12345678990",
		"bankName":      "TEST BANK",
		"amount":        finalAmount,
	}, nil
}

// FetchAllTokens implements TokenService.
func (s *service) FetchAllTokens(userID string, q QueryUserTokensData) (FetchTokensData, error) {
	var count int64
	var tokens []model.Token
	var data FetchTokensData

	if err := s.db.
		Table("tokens").
		Where("user_id = ?", userID).
		Scopes(
			query.PaginateRows(int(q.Page), int(q.Limit)),
		).
		Order("created_at DESC").
		Find(&tokens).
		Error; err != nil {
		return data, err
	}

	if err := s.db.Table("tokens").
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return data, err
	}

	data.Tokens = tokens
	data.Meta = query.Paginate(count, len(tokens), int(q.Page), int(q.Limit))

	return data, nil
}

func NewTokenService(db database.DatabaseConnection) TokenService {
	ps := project.NewProjectService(db)
	return &service{
		db: db.GetDB(),
		ps: ps,
	}
}
