package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 9,
		Name:    "create_reviews_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
				CREATE TABLE IF NOT EXISTS reviews (
				    review_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					booking_id UUID NOT NULL,
					user_id UUID NOT NULL,
					resource_id UUID NOT NULL,
					vendor_id UUID NOT NULL,
					rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
					comment TEXT,
					images JSONB,
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

					-- foreign keys
					FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE
				)
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS reviews CASCADE").Error
		},
	})
}
