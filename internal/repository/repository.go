package repository

import (
	"time"

	"github.com/finkord/building_modern_web_app_with_go/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(res models.RoomRestriction) error

	SearchAvailabilityByRoomID(start, end time.Time, roomID int) (bool, error)

	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
