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
						slot_date DATE NOT NULL,
						slot_hour INT  NOT NULL,
						created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
						FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE,
						FOREIGN KEY (resource_id) REFERENCES resources(resource_id) ON DELETE RESTRICT,
						UNIQUE (resource_id, slot_date, slot_hour, booking_id)
					);
				`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec(`

				`).Error
		},
	})
}
