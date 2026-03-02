package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 13,
		Name:    "create_booking_slot",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
					CREATE TABLE booking_slots (
						booking_slot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
						booking_id UUID NOT NULL,
						date DATE NOT NULL,
						time_from TIMESTAMP NOT NULL,
						time_to TIMESTAMP NOT NULL,
						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE
					);
				`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec(`

				`).Error
		},
	})
}
