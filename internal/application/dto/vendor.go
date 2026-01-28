package dto

// ==================== REQUEST DTOs ====================
// CreateVendorRequest - DTO untuk membuat vendor baru
type CreateVendorRequest struct {
	VendorName  string `json:"vendor_name" validate:"required,min=3,max=255"`
	Address     string `json:"address" validate:"required,min=5,max=255"`
	City        string `json:"city" validate:"required,min=2,max=255"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=20,phone_number"`
	Email       string `json:"email" validate:"required,email,max=255"`
	Description string `json:"description" validate:"required,min=10,max=2000"`
}

// UpdateVendorRequest - DTO untuk update vendor (semua field optional)
type UpdateVendorRequest struct {
	VendorName  *string `json:"vendor_name,omitempty" validate:"omitempty,min=3,max=255"`
	Address     *string `json:"address,omitempty" validate:"omitempty,min=5,max=255"`
	City        *string `json:"city,omitempty" validate:"omitempty,min=2,max=255"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=20,phone_number"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=10,max=2000"`
}

// UpdateVendorStatusRequest - DTO untuk update status vendor
type UpdateVendorStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=active inactive"`
}

// VendorFilterRequest - DTO untuk filter vendors
type VendorFilterRequest struct {
	City          string  `json:"city,omitempty"`
	Status        string  `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
	MinRating     float64 `json:"min_rating,omitempty" validate:"omitempty,gte=0,lte=5"`
	SearchQuery   string  `json:"search_query,omitempty"`
	Page          int     `json:"page,omitempty" validate:"omitempty,min=1"`
	PageSize      int     `json:"page_size,omitempty" validate:"omitempty,min=1,max=100"`
	SortBy        string  `json:"sort_by,omitempty" validate:"omitempty,oneof=name rating created_at"`
	SortDirection string  `json:"sort_direction,omitempty" validate:"omitempty,oneof=asc desc"`
}

// ==================== RESPONSE DTOs ====================
// VendorResponse - DTO untuk response vendor detail
type VendorResponse struct {
	VendorID    string  `json:"vendor_id"`
	OwnerUserID string  `json:"owner_user_id"`
	VendorName  string  `json:"vendor_name"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	RatingStars string  `json:"rating_stars"`
	Status      string  `json:"status"`
	StatusBadge string  `json:"status_badge"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`

	// Optional nested data
	Owner *UserSummary `json:"owner,omitempty"`
}

// VendorDetailResponse - DTO untuk response vendor detail dengan statistics
type VendorDetailResponse struct {
	VendorResponse

	// Statistics
	TotalResources    int     `json:"total_resources"`
	TotalBookings     int     `json:"total_bookings"`
	CompletedBookings int     `json:"completed_bookings"`
	TotalRevenue      float64 `json:"total_revenue"`
	AveragePrice      float64 `json:"average_price"`
	TotalReviews      int     `json:"total_reviews"`

	// Related data
	Resources     []ResourceSummary `json:"resources,omitempty"`
	RecentReviews []ReviewSummary   `json:"recent_reviews,omitempty"`
}

// VendorSummary - DTO untuk nested response (ringkasan vendor)
type VendorSummary struct {
	VendorID   string  `json:"vendor_id"`
	VendorName string  `json:"vendor_name"`
	City       string  `json:"city"`
	Rating     float64 `json:"rating"`
	Status     string  `json:"status"`
}

// VendorListResponse - DTO untuk list vendors dengan pagination
type VendorListResponse struct {
	Vendors   []VendorResponse `json:"vendors"`
	Total     int64            `json:"total"`
	Page      int              `json:"page"`
	PageSize  int              `json:"page_size"`
	TotalPage int              `json:"total_page"`
}

// VendorStatisticsResponse - DTO untuk vendor statistics
type VendorStatisticsResponse struct {
	VendorID          string  `json:"vendor_id"`
	VendorName        string  `json:"vendor_name"`
	TotalResources    int     `json:"total_resources"`
	ActiveResources   int     `json:"active_resources"`
	TotalBookings     int     `json:"total_bookings"`
	PendingBookings   int     `json:"pending_bookings"`
	ConfirmedBookings int     `json:"confirmed_bookings"`
	CompletedBookings int     `json:"completed_bookings"`
	CancelledBookings int     `json:"cancelled_bookings"`
	TotalRevenue      float64 `json:"total_revenue"`
	MonthlyRevenue    float64 `json:"monthly_revenue"`
	TotalReviews      int     `json:"total_reviews"`
	AverageRating     float64 `json:"average_rating"`
	ResponseRate      float64 `json:"response_rate"` // Percentage
}

// VendorDashboardResponse - DTO untuk vendor dashboard
type VendorDashboardResponse struct {
	VendorInfo       VendorResponse           `json:"vendor_info"`
	Statistics       VendorStatisticsResponse `json:"statistics"`
	RecentBookings   []BookingSummary         `json:"recent_bookings"`
	RecentReviews    []ReviewSummary          `json:"recent_reviews"`
	PopularResources []ResourceSummary        `json:"popular_resources"`
	RevenueChart     []RevenueDataPoint       `json:"revenue_chart"`
}

// RevenueDataPoint - DTO untuk revenue chart data
type RevenueDataPoint struct {
	Date     string  `json:"date"`
	Revenue  float64 `json:"revenue"`
	Bookings int     `json:"bookings"`
}

// VendorPublicProfileResponse - DTO untuk public profile vendor (untuk customer)
type VendorPublicProfileResponse struct {
	VendorID     string  `json:"vendor_id"`
	VendorName   string  `json:"vendor_name"`
	City         string  `json:"city"`
	Description  string  `json:"description"`
	Rating       float64 `json:"rating"`
	RatingStars  string  `json:"rating_stars"`
	TotalReviews int     `json:"total_reviews"`

	// Public info only
	Resources          []ResourceSummary  `json:"resources"`
	Reviews            []ReviewSummary    `json:"reviews"`
	RatingDistribution RatingDistribution `json:"rating_distribution"`
}

// RatingDistribution - DTO untuk rating distribution
type RatingDistribution struct {
	FiveStar      int     `json:"five_star"`
	FourStar      int     `json:"four_star"`
	ThreeStar     int     `json:"three_star"`
	TwoStar       int     `json:"two_star"`
	OneStar       int     `json:"one_star"`
	AverageRating float64 `json:"average_rating"`
}

// VendorSearchResult - DTO untuk search results
type VendorSearchResult struct {
	VendorID           string   `json:"vendor_id"`
	VendorName         string   `json:"vendor_name"`
	City               string   `json:"city"`
	Rating             float64  `json:"rating"`
	TotalReviews       int      `json:"total_reviews"`
	MinPrice           float64  `json:"min_price"`
	MaxPrice           float64  `json:"max_price"`
	AvailableResources int      `json:"available_resources"`
	Highlights         []string `json:"highlights"` // Search match highlights
}

// ==================== NESTED DTOs ====================

// UserSummary - untuk nested owner info
type UserSummary struct {
	UserID string  `json:"user_id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  *string `json:"phone,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}

// ResourceSummary - untuk nested resource info
type ResourceSummary struct {
	ResourceID   string   `json:"resource_id"`
	ResourceName string   `json:"resource_name"`
	ResourceType string   `json:"resource_type"`
	Capacity     int      `json:"capacity"`
	PricePerUnit float64  `json:"price_per_unit"`
	Status       string   `json:"status"`
	Images       []string `json:"images,omitempty"`
}

// ReviewSummary - untuk nested review info
type ReviewSummary struct {
	ReviewID   string  `json:"review_id"`
	Rating     int     `json:"rating"`
	Comment    *string `json:"comment,omitempty"`
	UserName   string  `json:"user_name"`
	UserAvatar *string `json:"user_avatar,omitempty"`
	CreatedAt  string  `json:"created_at"`
}

// BookingSummary - untuk nested booking info
type BookingSummary struct {
	BookingID     string  `json:"booking_id"`
	BookingCode   string  `json:"booking_code"`
	ResourceName  string  `json:"resource_name"`
	BookingDate   string  `json:"booking_date"`
	BookingStatus string  `json:"booking_status"`
	PaymentStatus string  `json:"payment_status"`
	TotalPrice    float64 `json:"total_price"`
	UserName      string  `json:"user_name"`
}

// ==================== BULK OPERATIONS ====================

// BulkUpdateVendorStatusRequest - untuk bulk update status
type BulkUpdateVendorStatusRequest struct {
	VendorIDs []string `json:"vendor_ids" validate:"required,min=1,dive,uuid"`
	Status    string   `json:"status" validate:"required,oneof=active inactive"`
}

// BulkDeleteVendorsRequest - untuk bulk soft delete
type BulkDeleteVendorsRequest struct {
	VendorIDs []string `json:"vendor_ids" validate:"required,min=1,dive,uuid"`
}

// ==================== ANALYTICS DTOs ====================

// VendorPerformanceReport - untuk performance report
type VendorPerformanceReport struct {
	VendorID         string                `json:"vendor_id"`
	VendorName       string                `json:"vendor_name"`
	Period           string                `json:"period"` // "2024-01", "2024-Q1", "2024"
	TotalBookings    int                   `json:"total_bookings"`
	TotalRevenue     float64               `json:"total_revenue"`
	AverageRating    float64               `json:"average_rating"`
	BookingGrowth    float64               `json:"booking_growth"` // Percentage
	RevenueGrowth    float64               `json:"revenue_growth"` // Percentage
	TopResources     []ResourcePerformance `json:"top_resources"`
	MonthlyBreakdown []MonthlyPerformance  `json:"monthly_breakdown"`
}

// ResourcePerformance - untuk top performing resources
type ResourcePerformance struct {
	ResourceID   string  `json:"resource_id"`
	ResourceName string  `json:"resource_name"`
	Bookings     int     `json:"bookings"`
	Revenue      float64 `json:"revenue"`
	Rating       float64 `json:"rating"`
}

// MonthlyPerformance - untuk monthly breakdown
type MonthlyPerformance struct {
	Month    string  `json:"month"` // "2024-01"
	Bookings int     `json:"bookings"`
	Revenue  float64 `json:"revenue"`
	Rating   float64 `json:"rating"`
}

// VendorComparisonResponse - untuk comparing vendors
type VendorComparisonResponse struct {
	Vendors []VendorComparisonItem `json:"vendors"`
}

// VendorComparisonItem - single vendor in comparison
type VendorComparisonItem struct {
	VendorID          string  `json:"vendor_id"`
	VendorName        string  `json:"vendor_name"`
	Rating            float64 `json:"rating"`
	TotalReviews      int     `json:"total_reviews"`
	TotalResources    int     `json:"total_resources"`
	AveragePrice      float64 `json:"average_price"`
	CompletedBookings int     `json:"completed_bookings"`
	ResponseTime      string  `json:"response_time"` // "< 1 hour", "1-3 hours", etc.
}
