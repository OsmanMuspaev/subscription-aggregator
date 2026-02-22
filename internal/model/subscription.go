package model

import "time"

type Subscription struct {
	ID          string     `json:"id"`           // UUID
	ServiceName string     `json:"service_name"` // Название сервиса
	Price       int        `json:"price"`        // Стоимость в рублях
	UserID      string     `json:"user_id"`      // UUID пользователя
	StartDate   time.Time  `json:"start_date"`   // Месяц и год начала
	EndDate     *time.Time `json:"end_date"`     // Опционально
}