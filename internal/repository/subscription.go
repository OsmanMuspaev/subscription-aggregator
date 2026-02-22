package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/yourusername/subscription-service/internal/db"
	"github.com/yourusername/subscription-service/internal/model"
)

type SubscriptionRepo struct{}

func NewSubscriptionRepo() *SubscriptionRepo {
	return &SubscriptionRepo{}
}

// Create добавляет новую подписку
func (r *SubscriptionRepo) Create(sub *model.Subscription) error {
	if sub.ID == "" {
		sub.ID = uuid.New().String()
	}

	_, err := db.Pool.Exec(context.Background(),
		`INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		sub.ID,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate.Format("2006-01-02"),
		func() interface{} {
			if sub.EndDate != nil {
				return sub.EndDate.Format("2006-01-02")
			}
			return nil
		}(),
	)
	return err
}

// Get возвращает подписку по ID
func (r *SubscriptionRepo) Get(id string) (*model.Subscription, error) {
	sub := &model.Subscription{}
	err := db.Pool.QueryRow(context.Background(),
		`SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id=$1`, id,
	).Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return sub, nil
}

// Delete удаляет подписку по ID
func (r *SubscriptionRepo) Delete(id string) error {
	_, err := db.Pool.Exec(context.Background(),
		`DELETE FROM subscriptions WHERE id=$1`, id)
	return err
}

// Update обновляет подписку
func (r *SubscriptionRepo) Update(sub *model.Subscription) error {
	_, err := db.Pool.Exec(context.Background(),
		`UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5 WHERE id=$6`,
		sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate, sub.ID)
	return err
}

// List возвращает список подписок с фильтрацией по пользователю и сервису
func (r *SubscriptionRepo) List(userID, serviceName string) ([]*model.Subscription, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE 1=1`
	args := []interface{}{}
	i := 1

	if userID != "" {
		query += fmt.Sprintf(" AND user_id=$%d", i)
		args = append(args, userID)
		i++
	}
	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name=$%d", i)
		args = append(args, serviceName)
		i++
	}

	rows, err := db.Pool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []*model.Subscription
	for rows.Next() {
		sub := &model.Subscription{}
		err := rows.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
		if err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

// SumCost возвращает сумму подписок за период
func (r *SubscriptionRepo) SumCost(userID, serviceName string, from, to time.Time) (int, error) {
	query := `SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE start_date BETWEEN $1 AND $2`
	args := []interface{}{from, to}

	if userID != "" {
		query += ` AND user_id=$3`
		args = append(args, userID)
	}
	if serviceName != "" {
		query += ` AND service_name=$4`
		args = append(args, serviceName)
	}

	var sum int
	err := db.Pool.QueryRow(context.Background(), query, args...).Scan(&sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}