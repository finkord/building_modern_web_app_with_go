package repository

import (
	"time"

	"github.com/finkord/building_modern_web_app_with_go/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(res models.RoomRestriction) error

	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)

	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)

	GetRoomByID(id int) (models.Room, error)

	GetUserByID(id int) (models.User, error)

	UpdateUser(u models.User) error

	Authenticate(email, password string) (int, string, error)

	AllReservations() ([]models.Reservation, error)

	AllNewReservations() ([]models.Reservation, error)

	GetReservationByID(id int) (models.Reservation, error)

	UpdateReservation(u models.Reservation) error

	DeleteReservation(id int) error

	UpdateProcessedForReservation(id int, processed int) error

	AllRooms() ([]models.Room, error)

	GetRestrictionForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)

	InsertBlockForRoom(id int, startDate time.Time) error

	DeleteBlockByID(id int) error
}
