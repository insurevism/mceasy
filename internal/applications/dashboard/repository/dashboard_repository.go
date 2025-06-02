package repository

import (
	"context"
	"fmt"
	"time"

	"mceasy/ent"
	"mceasy/ent/attendance"
	"mceasy/ent/employee"
	"mceasy/ent/salarycalculation"
	"mceasy/internal/applications/dashboard/dto"
)

// DashboardRepository defines the interface for dashboard data operations
type DashboardRepository interface {
	GetTodayAttendanceSummary(ctx context.Context) (*dto.TodayAttendanceSummary, error)
	GetDashboardOverview(ctx context.Context) (*dto.DashboardOverview, error)
	GetAttendanceTrends(ctx context.Context, startDate, endDate time.Time, periodType string) (*dto.AttendanceTrendData, error)
	GetSalaryTrends(ctx context.Context, startDate, endDate time.Time) (*dto.SalaryTrendData, error)
	GetAttendanceAlerts(ctx context.Context) (*dto.AttendanceAlerts, error)
	GetWeeklySummary(ctx context.Context, startDate time.Time) (*dto.WeeklySummary, error)
	GetMonthlyStats(ctx context.Context, month time.Time) (*dto.MonthlyStatsSummary, error)
	GetSalaryOverview(ctx context.Context, month time.Time) (*dto.SalaryOverviewData, error)
	GetSystemStats(ctx context.Context) (*dto.SystemStatsData, error)
}

// DashboardRepositoryImpl implements the DashboardRepository interface
type DashboardRepositoryImpl struct {
	client *ent.Client
}

// NewDashboardRepository creates a new dashboard repository instance
func NewDashboardRepository(client *ent.Client) *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{
		client: client,
	}
}

// GetTodayAttendanceSummary retrieves today's attendance summary (core requirement)
func (r *DashboardRepositoryImpl) GetTodayAttendanceSummary(ctx context.Context) (*dto.TodayAttendanceSummary, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	// Get total active employees
	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Where(employee.IsActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count total employees: %w", err)
	}

	// Get today's attendance records
	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch today's attendance: %w", err)
	}

	// Count attendance by status
	var presentCount, absentCount, lateCount, halfDayCount int
	for _, record := range attendanceRecords {
		switch record.Status {
		case attendance.StatusPresent:
			presentCount++
		case attendance.StatusAbsent:
			absentCount++
		case attendance.StatusLate:
			lateCount++
		case attendance.StatusHalfDay:
			halfDayCount++
		}
	}

	// Calculate not marked count (employees without attendance record today)
	notMarkedCount := totalEmployees - len(attendanceRecords)

	// Calculate percentages
	presentPercentage := 0.0
	absentPercentage := 0.0
	if totalEmployees > 0 {
		// Consider present + late + half day as "present" for percentage calculation
		actualPresentCount := presentCount + lateCount + halfDayCount
		presentPercentage = float64(actualPresentCount) / float64(totalEmployees) * 100
		absentPercentage = float64(absentCount+notMarkedCount) / float64(totalEmployees) * 100
	}

	return &dto.TodayAttendanceSummary{
		Date:              startOfDay,
		TotalEmployees:    totalEmployees,
		PresentCount:      presentCount,
		AbsentCount:       absentCount,
		LateCount:         lateCount,
		HalfDayCount:      halfDayCount,
		PresentPercentage: presentPercentage,
		AbsentPercentage:  absentPercentage,
		NotMarkedCount:    notMarkedCount,
	}, nil
}

// GetDashboardOverview retrieves complete dashboard overview
func (r *DashboardRepositoryImpl) GetDashboardOverview(ctx context.Context) (*dto.DashboardOverview, error) {
	// Get today's attendance summary
	todayAttendance, err := r.GetTodayAttendanceSummary(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get today's attendance: %w", err)
	}

	// Get current month stats
	currentMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
	monthlyStats, err := r.GetMonthlyStats(ctx, currentMonth)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly stats: %w", err)
	}

	// Get last 7 days attendance trends
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -7)
	recentTrends, err := r.GetAttendanceTrends(ctx, startDate, endDate, "daily")
	if err != nil {
		return nil, fmt.Errorf("failed to get recent trends: %w", err)
	}

	// Get salary overview for current month
	salaryOverview, err := r.GetSalaryOverview(ctx, currentMonth)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary overview: %w", err)
	}

	// Get system stats
	systemStats, err := r.GetSystemStats(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get system stats: %w", err)
	}

	return &dto.DashboardOverview{
		TodayAttendance: *todayAttendance,
		MonthlyStats:    *monthlyStats,
		RecentTrends:    *recentTrends,
		SalaryOverview:  *salaryOverview,
		SystemStats:     *systemStats,
	}, nil
}

// GetAttendanceTrends retrieves attendance trends over a period
func (r *DashboardRepositoryImpl) GetAttendanceTrends(ctx context.Context, startDate, endDate time.Time, periodType string) (*dto.AttendanceTrendData, error) {
	var dataPoints []dto.AttendanceTrendPoint

	// Normalize dates
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, startDate.Location())
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 999999999, endDate.Location())

	// Get total employees for the period (assume constant for simplicity)
	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Where(employee.IsActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count employees: %w", err)
	}

	// Generate data points based on period type
	switch periodType {
	case "daily":
		for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
			// Skip weekends
			if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
				continue
			}

			point, err := r.getAttendancePointForDate(ctx, d, totalEmployees)
			if err != nil {
				continue // Skip on error, don't fail entire request
			}
			dataPoints = append(dataPoints, *point)
		}
	case "weekly":
		// Implement weekly aggregation
		for d := start; !d.After(end); d = d.AddDate(0, 0, 7) {
			weekEnd := d.AddDate(0, 0, 6)
			if weekEnd.After(end) {
				weekEnd = end
			}
			point, err := r.getAttendancePointForWeek(ctx, d, weekEnd, totalEmployees)
			if err != nil {
				continue
			}
			dataPoints = append(dataPoints, *point)
		}
	case "monthly":
		// Implement monthly aggregation
		for d := start; d.Year() < end.Year() || (d.Year() == end.Year() && d.Month() <= end.Month()); d = d.AddDate(0, 1, 0) {
			point, err := r.getAttendancePointForMonth(ctx, d, totalEmployees)
			if err != nil {
				continue
			}
			dataPoints = append(dataPoints, *point)
		}
	}

	return &dto.AttendanceTrendData{
		PeriodType: periodType,
		StartDate:  startDate,
		EndDate:    endDate,
		DataPoints: dataPoints,
	}, nil
}

// getAttendancePointForDate gets attendance data for a specific date
func (r *DashboardRepositoryImpl) getAttendancePointForDate(ctx context.Context, date time.Time, totalEmployees int) (*dto.AttendanceTrendPoint, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var presentCount, absentCount, lateCount int
	for _, record := range attendanceRecords {
		switch record.Status {
		case attendance.StatusPresent, attendance.StatusHalfDay:
			presentCount++
		case attendance.StatusAbsent:
			absentCount++
		case attendance.StatusLate:
			lateCount++
		}
	}

	attendanceRate := 0.0
	if totalEmployees > 0 {
		attendanceRate = float64(presentCount+lateCount) / float64(totalEmployees) * 100
	}

	return &dto.AttendanceTrendPoint{
		Date:           startOfDay,
		PresentCount:   presentCount,
		AbsentCount:    absentCount,
		LateCount:      lateCount,
		AttendanceRate: attendanceRate,
		TotalEmployees: totalEmployees,
	}, nil
}

// getAttendancePointForWeek gets aggregated attendance data for a week
func (r *DashboardRepositoryImpl) getAttendancePointForWeek(ctx context.Context, startDate, endDate time.Time, totalEmployees int) (*dto.AttendanceTrendPoint, error) {
	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDateGTE(startDate)).
		Where(attendance.AttendanceDateLTE(endDate)).
		Where(attendance.DeletedAtIsNil()).
		Where(attendance.IsWeekendEQ(false)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var presentCount, absentCount, lateCount int
	for _, record := range attendanceRecords {
		switch record.Status {
		case attendance.StatusPresent, attendance.StatusHalfDay:
			presentCount++
		case attendance.StatusAbsent:
			absentCount++
		case attendance.StatusLate:
			lateCount++
		}
	}

	attendanceRate := 0.0
	workingDays := r.countWorkingDays(startDate, endDate)
	totalPossibleAttendance := totalEmployees * workingDays
	if totalPossibleAttendance > 0 {
		attendanceRate = float64(presentCount+lateCount) / float64(totalPossibleAttendance) * 100
	}

	return &dto.AttendanceTrendPoint{
		Date:           startDate,
		PresentCount:   presentCount,
		AbsentCount:    absentCount,
		LateCount:      lateCount,
		AttendanceRate: attendanceRate,
		TotalEmployees: totalEmployees,
	}, nil
}

// getAttendancePointForMonth gets aggregated attendance data for a month
func (r *DashboardRepositoryImpl) getAttendancePointForMonth(ctx context.Context, month time.Time, totalEmployees int) (*dto.AttendanceTrendPoint, error) {
	startOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	return r.getAttendancePointForWeek(ctx, startOfMonth, endOfMonth, totalEmployees)
}

// countWorkingDays counts working days between two dates (excluding weekends)
func (r *DashboardRepositoryImpl) countWorkingDays(startDate, endDate time.Time) int {
	count := 0
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			count++
		}
	}
	return count
}

// GetSalaryTrends retrieves salary trends over multiple months
func (r *DashboardRepositoryImpl) GetSalaryTrends(ctx context.Context, startDate, endDate time.Time) (*dto.SalaryTrendData, error) {
	var dataPoints []dto.SalaryTrendPoint

	startMonth := time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, startDate.Location())
	endMonth := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, endDate.Location())

	for month := startMonth; !month.After(endMonth); month = month.AddDate(0, 1, 0) {
		calculations, err := r.client.SalaryCalculation.
			Query().
			Where(salarycalculation.CalculationMonth(month)).
			Where(salarycalculation.DeletedAtIsNil()).
			All(ctx)
		if err != nil {
			continue
		}

		var totalBaseSalary, totalFinalSalary, totalDeductions float64
		employeesCount := len(calculations)

		for _, calc := range calculations {
			totalBaseSalary += calc.BaseSalary
			totalFinalSalary += calc.FinalSalary
			totalDeductions += calc.DeductionAmount
		}

		averageDeductionRate := 0.0
		if totalBaseSalary > 0 {
			averageDeductionRate = (totalDeductions / totalBaseSalary) * 100
		}

		dataPoints = append(dataPoints, dto.SalaryTrendPoint{
			Month:                month,
			TotalBaseSalary:      totalBaseSalary,
			TotalFinalSalary:     totalFinalSalary,
			TotalDeductions:      totalDeductions,
			AverageDeductionRate: averageDeductionRate,
			EmployeesCount:       employeesCount,
		})
	}

	return &dto.SalaryTrendData{
		PeriodType: "monthly",
		StartDate:  startDate,
		EndDate:    endDate,
		DataPoints: dataPoints,
	}, nil
}

// GetAttendanceAlerts retrieves attendance-related alerts
func (r *DashboardRepositoryImpl) GetAttendanceAlerts(ctx context.Context) (*dto.AttendanceAlerts, error) {
	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())

	// Get today's attendance records with employee info
	todayAttendance, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		WithEmployee().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch today's attendance: %w", err)
	}

	var lateToday, absentToday []dto.EmployeeAlert

	for _, record := range todayAttendance {
		if record.Edges.Employee == nil {
			continue
		}

		emp := record.Edges.Employee
		alert := dto.EmployeeAlert{
			EmployeeID:   emp.ID,
			EmployeeName: emp.FullName,
			EmployeeCode: emp.EmployeeID,
			Department:   emp.Department,
			Date:         startOfDay,
		}

		switch record.Status {
		case attendance.StatusLate:
			alert.Message = fmt.Sprintf("Arrived late at %s", record.CheckInTime.Format("15:04"))
			alert.Severity = "warning"
			lateToday = append(lateToday, alert)
		case attendance.StatusAbsent:
			alert.Message = "Absent today"
			alert.Severity = "critical"
			absentToday = append(absentToday, alert)
		}
	}

	// Get perfect attendance for current month (bonus feature)
	perfectAttendance, err := r.getPerfectAttendanceEmployees(ctx)
	if err != nil {
		perfectAttendance = []dto.EmployeeAlert{} // Don't fail on this
	}

	// Get frequent late employees (bonus feature)
	frequentLate, err := r.getFrequentLateEmployees(ctx)
	if err != nil {
		frequentLate = []dto.EmployeeAlert{} // Don't fail on this
	}

	return &dto.AttendanceAlerts{
		LateEmployeesToday:     lateToday,
		AbsentEmployeesToday:   absentToday,
		PerfectAttendanceMonth: perfectAttendance,
		FrequentLateEmployees:  frequentLate,
	}, nil
}

// getPerfectAttendanceEmployees gets employees with perfect attendance this month
func (r *DashboardRepositoryImpl) getPerfectAttendanceEmployees(ctx context.Context) ([]dto.EmployeeAlert, error) {
	// currentMonth := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
	// endOfMonth := currentMonth.AddDate(0, 1, -1)

	// This would require more complex logic to calculate perfect attendance
	// For now, return empty array
	return []dto.EmployeeAlert{}, nil
}

// getFrequentLateEmployees gets employees who are frequently late
func (r *DashboardRepositoryImpl) getFrequentLateEmployees(ctx context.Context) ([]dto.EmployeeAlert, error) {
	// This would require more complex logic to calculate frequent lateness
	// For now, return empty array
	return []dto.EmployeeAlert{}, nil
}

// GetWeeklySummary retrieves weekly attendance summary
func (r *DashboardRepositoryImpl) GetWeeklySummary(ctx context.Context, startDate time.Time) (*dto.WeeklySummary, error) {
	endDate := startDate.AddDate(0, 0, 6)

	var dailySummaries []dto.TodayAttendanceSummary
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		// Skip weekends
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			continue
		}

		summary, err := r.getDailySummaryForDate(ctx, d)
		if err != nil {
			continue
		}
		dailySummaries = append(dailySummaries, *summary)
	}

	// Calculate weekly averages
	var totalAttendanceRate, totalLateRate, totalAbsentRate float64
	for _, summary := range dailySummaries {
		totalAttendanceRate += summary.PresentPercentage
		lateRate := 0.0
		if summary.TotalEmployees > 0 {
			lateRate = float64(summary.LateCount) / float64(summary.TotalEmployees) * 100
		}
		totalLateRate += lateRate
		totalAbsentRate += summary.AbsentPercentage
	}

	weeklyAverages := dto.WeeklyAverageStats{
		AverageAttendanceRate: totalAttendanceRate / float64(len(dailySummaries)),
		AverageLateRate:       totalLateRate / float64(len(dailySummaries)),
		AverageAbsentRate:     totalAbsentRate / float64(len(dailySummaries)),
		TotalWorkingDays:      len(dailySummaries),
	}

	return &dto.WeeklySummary{
		WeekStartDate:  startDate,
		WeekEndDate:    endDate,
		DailySummaries: dailySummaries,
		WeeklyAverages: weeklyAverages,
		TopPerformers:  []dto.EmployeeAlert{}, // Implement if needed
		NeedsAttention: []dto.EmployeeAlert{}, // Implement if needed
	}, nil
}

// getDailySummaryForDate gets attendance summary for a specific date
func (r *DashboardRepositoryImpl) getDailySummaryForDate(ctx context.Context, date time.Time) (*dto.TodayAttendanceSummary, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Where(employee.IsActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return nil, err
	}

	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDate(startOfDay)).
		Where(attendance.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var presentCount, absentCount, lateCount, halfDayCount int
	for _, record := range attendanceRecords {
		switch record.Status {
		case attendance.StatusPresent:
			presentCount++
		case attendance.StatusAbsent:
			absentCount++
		case attendance.StatusLate:
			lateCount++
		case attendance.StatusHalfDay:
			halfDayCount++
		}
	}

	notMarkedCount := totalEmployees - len(attendanceRecords)

	presentPercentage := 0.0
	absentPercentage := 0.0
	if totalEmployees > 0 {
		actualPresentCount := presentCount + lateCount + halfDayCount
		presentPercentage = float64(actualPresentCount) / float64(totalEmployees) * 100
		absentPercentage = float64(absentCount+notMarkedCount) / float64(totalEmployees) * 100
	}

	return &dto.TodayAttendanceSummary{
		Date:              startOfDay,
		TotalEmployees:    totalEmployees,
		PresentCount:      presentCount,
		AbsentCount:       absentCount,
		LateCount:         lateCount,
		HalfDayCount:      halfDayCount,
		PresentPercentage: presentPercentage,
		AbsentPercentage:  absentPercentage,
		NotMarkedCount:    notMarkedCount,
	}, nil
}

// GetMonthlyStats retrieves monthly statistics
func (r *DashboardRepositoryImpl) GetMonthlyStats(ctx context.Context, month time.Time) (*dto.MonthlyStatsSummary, error) {
	startOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	now := time.Now()

	totalWorkingDays := r.countWorkingDays(startOfMonth, endOfMonth)
	workingDaysElapsed := r.countWorkingDays(startOfMonth, now)
	if workingDaysElapsed > totalWorkingDays {
		workingDaysElapsed = totalWorkingDays
	}

	// Get salary calculations for this month
	salaryCalculations, err := r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.CalculationMonth(startOfMonth)).
		Where(salarycalculation.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch salary calculations: %w", err)
	}

	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Where(employee.IsActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count employees: %w", err)
	}

	var totalSalaryCalculated float64
	for _, calc := range salaryCalculations {
		totalSalaryCalculated += calc.FinalSalary
	}

	calculatedEmployees := len(salaryCalculations)
	pendingEmployees := totalEmployees - calculatedEmployees
	totalSalaryPending := float64(pendingEmployees) * 10000000 // Base salary

	// Calculate average attendance rate for the month
	attendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.AttendanceDateGTE(startOfMonth)).
		Where(attendance.AttendanceDateLTE(endOfMonth)).
		Where(attendance.DeletedAtIsNil()).
		Where(attendance.IsWeekendEQ(false)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch attendance records: %w", err)
	}

	var presentCount int
	for _, record := range attendanceRecords {
		if record.Status == attendance.StatusPresent || record.Status == attendance.StatusLate || record.Status == attendance.StatusHalfDay {
			presentCount++
		}
	}

	averageAttendanceRate := 0.0
	totalPossibleAttendance := totalEmployees * workingDaysElapsed
	if totalPossibleAttendance > 0 {
		averageAttendanceRate = float64(presentCount) / float64(totalPossibleAttendance) * 100
	}

	return &dto.MonthlyStatsSummary{
		Month:                          startOfMonth,
		TotalWorkingDays:               totalWorkingDays,
		WorkingDaysElapsed:             workingDaysElapsed,
		AverageAttendanceRate:          averageAttendanceRate,
		TotalSalaryCalculated:          totalSalaryCalculated,
		TotalSalaryPending:             totalSalaryPending,
		EmployeesWithPerfectAttendance: 0, // Implement complex logic if needed
	}, nil
}

// GetSalaryOverview retrieves salary overview for a month
func (r *DashboardRepositoryImpl) GetSalaryOverview(ctx context.Context, month time.Time) (*dto.SalaryOverviewData, error) {
	startOfMonth := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())

	calculations, err := r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.CalculationMonth(startOfMonth)).
		Where(salarycalculation.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch salary calculations: %w", err)
	}

	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Where(employee.IsActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count employees: %w", err)
	}

	var totalBaseSalary, totalCalculatedSalary, totalDeductions float64
	var highestSalary, lowestSalary float64
	employeesCalculated := len(calculations)
	employeesPending := totalEmployees - employeesCalculated

	if employeesCalculated > 0 {
		highestSalary = calculations[0].FinalSalary
		lowestSalary = calculations[0].FinalSalary

		for _, calc := range calculations {
			totalBaseSalary += calc.BaseSalary
			totalCalculatedSalary += calc.FinalSalary
			totalDeductions += calc.DeductionAmount

			if calc.FinalSalary > highestSalary {
				highestSalary = calc.FinalSalary
			}
			if calc.FinalSalary < lowestSalary {
				lowestSalary = calc.FinalSalary
			}
		}
	}

	averageDeductionRate := 0.0
	if totalBaseSalary > 0 {
		averageDeductionRate = (totalDeductions / totalBaseSalary) * 100
	}

	return &dto.SalaryOverviewData{
		CurrentMonth:          startOfMonth,
		TotalBaseSalary:       totalBaseSalary,
		TotalCalculatedSalary: totalCalculatedSalary,
		TotalDeductions:       totalDeductions,
		EmployeesCalculated:   employeesCalculated,
		EmployeesPending:      employeesPending,
		AverageDeductionRate:  averageDeductionRate,
		HighestSalary:         highestSalary,
		LowestSalary:          lowestSalary,
	}, nil
}

// GetSystemStats retrieves overall system statistics
func (r *DashboardRepositoryImpl) GetSystemStats(ctx context.Context) (*dto.SystemStatsData, error) {
	totalEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count total employees: %w", err)
	}

	activeEmployees, err := r.client.Employee.
		Query().
		Where(employee.DeletedAtIsNil()).
		Where(employee.IsActiveEQ(true)).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count active employees: %w", err)
	}

	totalAttendanceRecords, err := r.client.Attendance.
		Query().
		Where(attendance.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count attendance records: %w", err)
	}

	totalSalaryCalculations, err := r.client.SalaryCalculation.
		Query().
		Where(salarycalculation.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count salary calculations: %w", err)
	}

	return &dto.SystemStatsData{
		TotalEmployees:          totalEmployees,
		ActiveEmployees:         activeEmployees,
		TotalAttendanceRecords:  totalAttendanceRecords,
		TotalSalaryCalculations: totalSalaryCalculations,
		SystemUptime:            "N/A", // Could implement if needed
		LastUpdated:             time.Now(),
	}, nil
}
