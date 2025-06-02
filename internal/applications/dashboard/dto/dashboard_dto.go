package dto

import (
	"time"
)

// TodayAttendanceSummary represents today's attendance summary (core requirement)
type TodayAttendanceSummary struct {
	Date              time.Time `json:"date"`
	TotalEmployees    int       `json:"total_employees"`
	PresentCount      int       `json:"present_count"`
	AbsentCount       int       `json:"absent_count"`
	LateCount         int       `json:"late_count"`
	HalfDayCount      int       `json:"half_day_count"`
	PresentPercentage float64   `json:"present_percentage"`
	AbsentPercentage  float64   `json:"absent_percentage"`
	NotMarkedCount    int       `json:"not_marked_count"`
}

// DashboardOverview represents overall dashboard statistics
type DashboardOverview struct {
	TodayAttendance TodayAttendanceSummary `json:"today_attendance"`
	MonthlyStats    MonthlyStatsSummary    `json:"monthly_stats"`
	RecentTrends    AttendanceTrendData    `json:"recent_trends"`
	SalaryOverview  SalaryOverviewData     `json:"salary_overview"`
	SystemStats     SystemStatsData        `json:"system_stats"`
}

// MonthlyStatsSummary represents current month statistics
type MonthlyStatsSummary struct {
	Month                          time.Time `json:"month"`
	TotalWorkingDays               int       `json:"total_working_days"`
	WorkingDaysElapsed             int       `json:"working_days_elapsed"`
	AverageAttendanceRate          float64   `json:"average_attendance_rate"`
	TotalSalaryCalculated          float64   `json:"total_salary_calculated"`
	TotalSalaryPending             float64   `json:"total_salary_pending"`
	EmployeesWithPerfectAttendance int       `json:"employees_with_perfect_attendance"`
}

// AttendanceTrendData represents attendance trends over time
type AttendanceTrendData struct {
	PeriodType string                 `json:"period_type"` // "daily", "weekly", "monthly"
	StartDate  time.Time              `json:"start_date"`
	EndDate    time.Time              `json:"end_date"`
	DataPoints []AttendanceTrendPoint `json:"data_points"`
}

// AttendanceTrendPoint represents a single point in attendance trend
type AttendanceTrendPoint struct {
	Date           time.Time `json:"date"`
	PresentCount   int       `json:"present_count"`
	AbsentCount    int       `json:"absent_count"`
	LateCount      int       `json:"late_count"`
	AttendanceRate float64   `json:"attendance_rate"`
	TotalEmployees int       `json:"total_employees"`
}

// SalaryOverviewData represents salary-related statistics
type SalaryOverviewData struct {
	CurrentMonth          time.Time `json:"current_month"`
	TotalBaseSalary       float64   `json:"total_base_salary"`
	TotalCalculatedSalary float64   `json:"total_calculated_salary"`
	TotalDeductions       float64   `json:"total_deductions"`
	EmployeesCalculated   int       `json:"employees_calculated"`
	EmployeesPending      int       `json:"employees_pending"`
	AverageDeductionRate  float64   `json:"average_deduction_rate"`
	HighestSalary         float64   `json:"highest_salary"`
	LowestSalary          float64   `json:"lowest_salary"`
}

// SystemStatsData represents overall system statistics
type SystemStatsData struct {
	TotalEmployees          int       `json:"total_employees"`
	ActiveEmployees         int       `json:"active_employees"`
	TotalAttendanceRecords  int       `json:"total_attendance_records"`
	TotalSalaryCalculations int       `json:"total_salary_calculations"`
	SystemUptime            string    `json:"system_uptime"`
	LastUpdated             time.Time `json:"last_updated"`
}

// DashboardQueryParams represents query parameters for dashboard data
type DashboardQueryParams struct {
	StartDate  time.Time `query:"start_date" validate:"omitempty"`
	EndDate    time.Time `query:"end_date" validate:"omitempty"`
	PeriodType string    `query:"period_type" validate:"omitempty,oneof=daily weekly monthly"`
	EmployeeID uint64    `query:"employee_id" validate:"omitempty"`
}

// AttendanceAlerts represents attendance-related alerts and notifications
type AttendanceAlerts struct {
	LateEmployeesToday     []EmployeeAlert `json:"late_employees_today"`
	AbsentEmployeesToday   []EmployeeAlert `json:"absent_employees_today"`
	PerfectAttendanceMonth []EmployeeAlert `json:"perfect_attendance_month"`
	FrequentLateEmployees  []EmployeeAlert `json:"frequent_late_employees"`
}

// EmployeeAlert represents an individual employee alert
type EmployeeAlert struct {
	EmployeeID   uint64    `json:"employee_id"`
	EmployeeName string    `json:"employee_name"`
	EmployeeCode string    `json:"employee_code"`
	Department   string    `json:"department"`
	Message      string    `json:"message"`
	Severity     string    `json:"severity"` // "info", "warning", "critical"
	Date         time.Time `json:"date"`
}

// WeeklySummary represents weekly attendance summary
type WeeklySummary struct {
	WeekStartDate  time.Time                `json:"week_start_date"`
	WeekEndDate    time.Time                `json:"week_end_date"`
	DailySummaries []TodayAttendanceSummary `json:"daily_summaries"`
	WeeklyAverages WeeklyAverageStats       `json:"weekly_averages"`
	TopPerformers  []EmployeeAlert          `json:"top_performers"`
	NeedsAttention []EmployeeAlert          `json:"needs_attention"`
}

// WeeklyAverageStats represents average statistics for a week
type WeeklyAverageStats struct {
	AverageAttendanceRate float64 `json:"average_attendance_rate"`
	AverageLateRate       float64 `json:"average_late_rate"`
	AverageAbsentRate     float64 `json:"average_absent_rate"`
	TotalWorkingDays      int     `json:"total_working_days"`
}

// SalaryTrendData represents salary trends over multiple months
type SalaryTrendData struct {
	PeriodType string             `json:"period_type"` // "monthly", "quarterly"
	StartDate  time.Time          `json:"start_date"`
	EndDate    time.Time          `json:"end_date"`
	DataPoints []SalaryTrendPoint `json:"data_points"`
}

// SalaryTrendPoint represents a single point in salary trend
type SalaryTrendPoint struct {
	Month                time.Time `json:"month"`
	TotalBaseSalary      float64   `json:"total_base_salary"`
	TotalFinalSalary     float64   `json:"total_final_salary"`
	TotalDeductions      float64   `json:"total_deductions"`
	AverageDeductionRate float64   `json:"average_deduction_rate"`
	EmployeesCount       int       `json:"employees_count"`
}
