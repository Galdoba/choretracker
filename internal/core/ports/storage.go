package ports

import "github.com/Galdoba/choretracker/internal/core/domain"

type Storage interface {
	Create(domain.Chore) error
	Read(int64) (domain.Chore, error)
	Update(domain.Chore) error
	Delete(int64) error
	GetAll() ([]domain.Chore, error)
}
