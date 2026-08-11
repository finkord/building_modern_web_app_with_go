package dbrepo

import "github.com/finkord/building_modern_web_app_with_go/internal/models"

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction adds a room restriction to the database
func (m *testDBRepo) InsertRoomRestriction(res models.RoomRestriction) error {
	return nil
}
