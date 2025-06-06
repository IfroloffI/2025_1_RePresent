package csat

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	csatEntity "retarget/internal/csat-service/entity/csat"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

type CsatRepositoryInterface interface {
	AddReview(review csatEntity.Review) (*csatEntity.Review, error)
	GetAllReviews() ([]csatEntity.Review, error)
	GetReviewsByUser(userID int) ([]csatEntity.Review, error)
	GetQuestionsByPage(page string) ([]string, error)

	scanReview(rows *sql.Rows) (*csatEntity.Review, error)
	CloseConnection() error
}

type CsatRepository struct {
	db *sql.DB
}

func NewCsatRepository(dsn string) *CsatRepository {
	csatRepo := &CsatRepository{}
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("failed to ping DB: %w", err)
		return nil
	}
	csatRepo.db = db
	return csatRepo
}

func (r *CsatRepository) CloseConnection() error {
	return r.db.Close()
}
func (r *CsatRepository) scanReview(rows *sql.Rows) (*csatEntity.Review, error) {
	var review csatEntity.Review
	var createdAt time.Time
	var id string

	err := rows.Scan(
		&id,
		&review.UserID,
		&review.Question,
		&review.Page,
		&review.Comment,
		&review.Rating,
		&createdAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan review: %w", err)
	}

	review.ID = id
	review.CreatedAt = createdAt
	return &review, nil
}

func (r *CsatRepository) AddReview(review csatEntity.Review) error {
	const addQuery = `
		INSERT INTO reviews (
			user_id, 
			question, 
			page, 
			comment,
			rating
		) VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(addQuery,
		review.UserID,
		review.Question,
		review.Page,
		review.Comment,
		review.Rating,
	)

	if err != nil {
		return fmt.Errorf("failed to insert review: %w", err)
	}

	return nil
}

func (r *CsatRepository) GetAllReviews() ([]csatEntity.Review, error) {
	const query = `
		SELECT 
			id,
			user_id,
			question,
			page,
			comment,
			rating,
			created_at
		FROM reviews
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query reviews: %w", err)
	}
	defer rows.Close()

	var reviews []csatEntity.Review
	for rows.Next() {
		review, err := r.scanReview(rows)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, *review)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return reviews, nil
}

func (r *CsatRepository) GetReviewsByUser(userID int) ([]csatEntity.Review, error) {
	const query = `
		SELECT 
			id,
			user_id,
			question,
			page,
			comment,
			rating,
			created_at
		FROM reviews
		WHERE user_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user reviews: %w", err)
	}
	defer rows.Close()

	var reviews []csatEntity.Review
	for rows.Next() {
		review, err := r.scanReview(rows)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, *review)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return reviews, nil
}

func (r *CsatRepository) GetQuestionsByPage(page string) ([]string, error) {
	pageQuestions := map[string][]string{
		"Profile": {
			"Насколько вам удобно пользоваться профилем?",
			"Достаточно ли информации отображается в вашем профиле?",
			"Хотели бы вы видеть дополнительные функции в профиле?",
		},
		"BannerEditor": {
			"Насколько удобен интерфейс редактора баннеров?",
			"Достаточно ли функций предоставляет редактор?",
			"Как часто вы используете редактор баннеров?",
		},
		"Auth": {
			"Насколько удобен процесс авторизации?",
			"Как вы оцениваете безопасность системы авторизации?",
			"Хотели бы вы видеть дополнительные способы входа?",
		},
	}

	questions, exists := pageQuestions[page]
	if !exists {
		return nil, fmt.Errorf("unknown page: %s", page)
	}

	return questions, nil
}
