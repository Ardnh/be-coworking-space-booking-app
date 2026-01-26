package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 1,
		Name:    "create_notification_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS notifications (
			        notification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        booking_id UUID NOT NULL,
			        user_id UUID NOT NULL,
			        notification_type VARCHAR(255) NOT NULL,
			        title VARCHAR(255) NOT NULL,
			        message TEXT NOT NULL,
			        is_read BOOLEAN NOT NULL DEFAULT FALSE,
			        is_sent BOOLEAN NOT NULL DEFAULT FALSE,
			        sent_at TIMESTAMP,
			        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			        -- Foreign Key Constraints
			        CONSTRAINT fk_notification_booking
			            FOREIGN KEY (booking_id)
			            REFERENCES bookings(booking_id)
			            ON DELETE CASCADE,

			        CONSTRAINT fk_notification_user
			            FOREIGN KEY (user_id)
			            REFERENCES users(id)
			            ON DELETE CASCADE,

			        -- Check Constraints
			        CONSTRAINT chk_notification_type
			            CHECK (notification_type IN (
			                'booking_confirmation',
			                'booking_reminder',
			                'booking_cancelled',
			                'payment_success',
			                'payment_failed',
			                'payment_refund',
			                'schedule_change',
			                'promotion',
			                'system_alert'
			            )),

			        CONSTRAINT chk_title_not_empty
			            CHECK (TRIM(title) != ''),

			        CONSTRAINT chk_message_not_empty
			            CHECK (TRIM(message) != ''),

			        CONSTRAINT chk_sent_logic
			            CHECK (
			                (is_sent = FALSE AND sent_at IS NULL) OR
			                (is_sent = TRUE AND sent_at IS NOT NULL)
			            ),

			        CONSTRAINT chk_sent_after_created
			            CHECK (sent_at IS NULL OR sent_at >= created_at),

			        CONSTRAINT chk_dates
			            CHECK (updated_at >= created_at)
			    );

			    -- Indexes untuk performa query
			    CREATE INDEX IF NOT EXISTS idx_notifications_user_id
			        ON notifications(user_id);

			    CREATE INDEX IF NOT EXISTS idx_notifications_booking_id
			        ON notifications(booking_id);

			    CREATE INDEX IF NOT EXISTS idx_notifications_is_read
			        ON notifications(is_read)
			        WHERE is_read = FALSE;

			    CREATE INDEX IF NOT EXISTS idx_notifications_is_sent
			        ON notifications(is_sent)
			        WHERE is_sent = FALSE;

			    CREATE INDEX IF NOT EXISTS idx_notifications_type
			        ON notifications(notification_type);

			    CREATE INDEX IF NOT EXISTS idx_notifications_created_at
			        ON notifications(created_at DESC);

			    -- Composite index untuk query unread notifications per user
			    CREATE INDEX IF NOT EXISTS idx_notifications_user_unread
			        ON notifications(user_id, is_read, created_at DESC)
			        WHERE is_read = FALSE;

			    -- Composite index untuk query unsent notifications
			    CREATE INDEX IF NOT EXISTS idx_notifications_unsent
			        ON notifications(is_sent, created_at)
			        WHERE is_sent = FALSE;
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS notifications").Error
		},
	})
}
