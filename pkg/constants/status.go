package constants

const (
	UserStatusActive    = "active"
	UserStatusInactive  = "inactive"
	UserStatusDeleted   = "deleted"
	UserStatusSuspended = "suspended"
	UserStatusBanned    = "banned"
	UserStatusVerified  = "verified"
)

const (
	UserTypeAdmin    = "admin"
	UserTypeCustomer = "customer"
	UserTypeVendor   = "vendor"
	UserTypeStaff    = "staff"
)

const (
	BookingStatusPending   = "pending"
	BookingStatusConfirmed = "confirmed"
	BookingStatusCancelled = "cancelled"
	BookingStatusCompleted = "completed"
	BookingStatusRejected  = "rejected"
)

const (
	Sunday    int16 = 0
	Monday    int16 = 1
	Tuesday   int16 = 2
	Wednesday int16 = 3
	Thursday  int16 = 4
	Friday    int16 = 5
	Saturday  int16 = 6
)

const (
	PaymentMethodCreditCard   = "credit_card"
	PaymentMethodDebitCard    = "debit_card"
	PaymentMethodBankTransfer = "bank_transfer"
	PaymentMethodEWallet      = "ewallet"
	PaymentMethodCash         = "cash"
)

const (
	HistoryStatusPending     = "pending"
	HistoryStatusConfirmed   = "confirmed"
	HistoryStatusCancelled   = "cancelled"
	HistoryStatusCompleted   = "completed"
	HistoryStatusNoShow      = "no_show"
	HistoryStatusRescheduled = "rescheduled"
)

const (
	NotificationBookingConfirmation = "booking_confirmation"
	NotificationBookingReminder     = "booking_reminder"
	NotificationBookingCancelled    = "booking_cancelled"
	NotificationPaymentSuccess      = "payment_success"
	NotificationPaymentFailed       = "payment_failed"
	NotificationPaymentRefund       = "payment_refund"
	NotificationScheduleChange      = "schedule_change"
	NotificationPromotion           = "promotion"
	NotificationSystemAlert         = "system_alert"
)

const (
	ResourceStatusActive   = "active"
	ResourceStatusInactive = "inactive"
)
