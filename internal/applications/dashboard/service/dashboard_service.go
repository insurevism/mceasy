package service

import (
	"context"
	"fmt"
	"time"

	"mceasy/internal/applications/dashboard/dto"
	"mceasy/internal/applications/dashboard/repository"
	"mceasy/internal/component/cache"
)

// DashboardService defines the interface for dashboard business logic
type DashboardService interface {
	GetTodayAttendanceSummary(ctx context.Context) (*dto.TodayAttendanceSummary, error)
	GetDashboardOverview(ctx context.Context) (*dto.DashboardOverview, error)
	GetAttendanceTrends(ctx context.Context, params dto.DashboardQueryParams) (*dto.AttendanceTrendData, error)
	GetSalaryTrends(ctx context.Context, params dto.DashboardQueryParams) (*dto.SalaryTrendData, error)
	GetAttendanceAlerts(ctx context.Context) (*dto.AttendanceAlerts, error)
	GetWeeklySummary(ctx context.Context, startDate time.Time) (*dto.WeeklySummary, error)
	GetMonthlyStats(ctx context.Context, month time.Time) (*dto.MonthlyStatsSummary, error)
	GetSalaryOverview(ctx context.Context, month time.Time) (*dto.SalaryOverviewData, error)
	GetSystemStats(ctx context.Context) (*dto.SystemStatsData, error)
}

// DashboardServiceImpl implements the DashboardService interface
type DashboardServiceImpl struct {
	dashboardRepo repository.DashboardRepository
	cache         cache.Cache
}

// NewDashboardService creates a new dashboard service instance
func NewDashboardService(
	dashboardRepo repository.DashboardRepository,
	cache cache.Cache,
) *DashboardServiceImpl {
	return &DashboardServiceImpl{
		dashboardRepo: dashboardRepo,
		cache:         cache,
	}
}

// GetTodayAttendanceSummary retrieves today's attendance summary with caching
func (s *DashboardServiceImpl) GetTodayAttendanceSummary(ctx context.Context) (*dto.TodayAttendanceSummary, error) {
	// Use cache for today's summary (cache for 5 minutes)
	cacheKey := fmt.Sprintf("dashboard:today_summary:%s", time.Now().Format("2006-01-02"))

	var result dto.TodayAttendanceSummary
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	summary, err := s.dashboardRepo.GetTodayAttendanceSummary(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's attendance summary: %w", err)
	}

	// Cache for 5 minutes
	_, _ = s.cache.Create(ctx, cacheKey, summary, 5*time.Minute)

	return summary, nil
}

// GetDashboardOverview retrieves complete dashboard overview with caching
func (s *DashboardServiceImpl) GetDashboardOverview(ctx context.Context) (*dto.DashboardOverview, error) {
	// Use cache for dashboard overview (cache for 10 minutes)
	cacheKey := fmt.Sprintf("dashboard:overview:%s", time.Now().Format("2006-01-02-15"))

	var result dto.DashboardOverview
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	overview, err := s.dashboardRepo.GetDashboardOverview(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard overview: %w", err)
	}

	// Cache for 10 minutes
	_, _ = s.cache.Create(ctx, cacheKey, overview, 10*time.Minute)

	return overview, nil
}

// GetAttendanceTrends retrieves attendance trends with parameters validation
func (s *DashboardServiceImpl) GetAttendanceTrends(ctx context.Context, params dto.DashboardQueryParams) (*dto.AttendanceTrendData, error) {
	// Validate and set default parameters
	startDate := params.StartDate
	endDate := params.EndDate
	periodType := params.PeriodType

	// Set default date range if not provided (last 30 days)
	if startDate.IsZero() {
		startDate = time.Now().AddDate(0, 0, -30)
	}
	if endDate.IsZero() {
		endDate = time.Now()
	}

	// Set default period type
	if periodType == "" {
		periodType = "daily"
	}

	// Validate period type
	if periodType != "daily" && periodType != "weekly" && periodType != "monthly" {
		return nil, fmt.Errorf("invalid period type: %s", periodType)
	}

	// Validate date range
	if startDate.After(endDate) {
		return nil, fmt.Errorf("start date cannot be after end date")
	}

	// Limit date range to prevent excessive data
	maxDays := 365
	if endDate.Sub(startDate).Hours() > float64(maxDays*24) {
		return nil, fmt.Errorf("date range cannot exceed %d days", maxDays)
	}

	// Use cache for trends (cache for 1 hour)
	cacheKey := fmt.Sprintf("dashboard:attendance_trends:%s:%s:%s",
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"),
		periodType)

	var result dto.AttendanceTrendData
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	trends, err := s.dashboardRepo.GetAttendanceTrends(ctx, startDate, endDate, periodType)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance trends: %w", err)
	}

	// Cache for 1 hour
	_, _ = s.cache.Create(ctx, cacheKey, trends, time.Hour)

	return trends, nil
}

// GetSalaryTrends retrieves salary trends with parameters validation
func (s *DashboardServiceImpl) GetSalaryTrends(ctx context.Context, params dto.DashboardQueryParams) (*dto.SalaryTrendData, error) {
	// Validate and set default parameters
	startDate := params.StartDate
	endDate := params.EndDate

	// Set default date range if not provided (last 12 months)
	if startDate.IsZero() {
		startDate = time.Now().AddDate(-1, 0, 0)
	}
	if endDate.IsZero() {
		endDate = time.Now()
	}

	// Validate date range
	if startDate.After(endDate) {
		return nil, fmt.Errorf("start date cannot be after end date")
	}

	// Limit date range to prevent excessive data (max 24 months)
	maxMonths := 24
	monthsDiff := int(endDate.Sub(startDate).Hours() / (24 * 30))
	if monthsDiff > maxMonths {
		return nil, fmt.Errorf("date range cannot exceed %d months", maxMonths)
	}

	// Use cache for salary trends (cache for 2 hours)
	cacheKey := fmt.Sprintf("dashboard:salary_trends:%s:%s",
		startDate.Format("2006-01-02"),
		endDate.Format("2006-01-02"))

	var result dto.SalaryTrendData
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	trends, err := s.dashboardRepo.GetSalaryTrends(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary trends: %w", err)
	}

	// Cache for 2 hours
	_, _ = s.cache.Create(ctx, cacheKey, trends, 2*time.Hour)

	return trends, nil
}

// GetAttendanceAlerts retrieves attendance alerts with caching
func (s *DashboardServiceImpl) GetAttendanceAlerts(ctx context.Context) (*dto.AttendanceAlerts, error) {
	// Use cache for alerts (cache for 15 minutes)
	cacheKey := fmt.Sprintf("dashboard:alerts:%s", time.Now().Format("2006-01-02"))

	var result dto.AttendanceAlerts
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	alerts, err := s.dashboardRepo.GetAttendanceAlerts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance alerts: %w", err)
	}

	// Cache for 15 minutes
	_, _ = s.cache.Create(ctx, cacheKey, alerts, 15*time.Minute)

	return alerts, nil
}

// GetWeeklySummary retrieves weekly attendance summary with validation
func (s *DashboardServiceImpl) GetWeeklySummary(ctx context.Context, startDate time.Time) (*dto.WeeklySummary, error) {
	// Validate that startDate is a Monday
	if startDate.Weekday() != time.Monday {
		// Adjust to the Monday of the week containing startDate
		daysToSubtract := int(startDate.Weekday()) - int(time.Monday)
		if daysToSubtract < 0 {
			daysToSubtract += 7
		}
		startDate = startDate.AddDate(0, 0, -daysToSubtract)
	}

	// Normalize to start of day
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())

	// Use cache for weekly summary (cache for 1 hour)
	cacheKey := fmt.Sprintf("dashboard:weekly_summary:%s", startDate.Format("2006-01-02"))

	var result dto.WeeklySummary
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	summary, err := s.dashboardRepo.GetWeeklySummary(ctx, startDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get weekly summary: %w", err)
	}

	// Cache for 1 hour
	_, _ = s.cache.Create(ctx, cacheKey, summary, time.Hour)

	return summary, nil
}

// GetMonthlyStats retrieves monthly statistics with validation
func (s *DashboardServiceImpl) GetMonthlyStats(ctx context.Context, month time.Time) (*dto.MonthlyStatsSummary, error) {
	// Normalize to start of month
	month = time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())

	// Use cache for monthly stats (cache for 1 hour)
	cacheKey := fmt.Sprintf("dashboard:monthly_stats:%s", month.Format("2006-01"))

	var result dto.MonthlyStatsSummary
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	stats, err := s.dashboardRepo.GetMonthlyStats(ctx, month)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly stats: %w", err)
	}

	// Cache for 1 hour
	_, _ = s.cache.Create(ctx, cacheKey, stats, time.Hour)

	return stats, nil
}

// GetSalaryOverview retrieves salary overview with validation
func (s *DashboardServiceImpl) GetSalaryOverview(ctx context.Context, month time.Time) (*dto.SalaryOverviewData, error) {
	// Normalize to start of month
	month = time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())

	// Use cache for salary overview (cache for 2 hours)
	cacheKey := fmt.Sprintf("dashboard:salary_overview:%s", month.Format("2006-01"))

	var result dto.SalaryOverviewData
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	overview, err := s.dashboardRepo.GetSalaryOverview(ctx, month)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary overview: %w", err)
	}

	// Cache for 2 hours
	_, _ = s.cache.Create(ctx, cacheKey, overview, 2*time.Hour)

	return overview, nil
}

// GetSystemStats retrieves system statistics with caching
func (s *DashboardServiceImpl) GetSystemStats(ctx context.Context) (*dto.SystemStatsData, error) {
	// Use cache for system stats (cache for 30 minutes)
	cacheKey := "dashboard:system_stats"

	var result dto.SystemStatsData
	cachedData, err := s.cache.Get(ctx, cacheKey, &result)
	if err == nil && cachedData != nil {
		return &result, nil
	}

	// Get from repository if not in cache
	stats, err := s.dashboardRepo.GetSystemStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get system stats: %w", err)
	}

	// Cache for 30 minutes
	_, _ = s.cache.Create(ctx, cacheKey, stats, 30*time.Minute)

	return stats, nil
}
