package godogs

import (
	"context"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func initGetFullyUtilizedDatesSteps(ctx *godog.ScenarioContext) {
	// WHEN
	ctx.Step(`^I call GetFullyUtilizedDates from "([^"]*)" to "([^"]*)"$`, iCallGetFullyUtilizedDatesFromTo)

	// THEN
	ctx.Step(`^the fully utilized dates should be:$`, theFullyUtilizedDatesShouldBe)

	// GIVEN
	ctx.Step(`^warehouse usage on "([^"]*)" is (\d+\.?\d*)$`, warehouseUsageOnDateIs)
}

// ----------------------------------------------------------------
// GIVEN Step (Arrange)
// ----------------------------------------------------------------

func warehouseUsageOnDateIs(_ context.Context, dateStr string, usage float64) error {
	if tc.usageMap == nil {
		tc.usageMap = make(map[time.Time]float64)
	}

	date := parseDate(tc.t, dateStr)
	tc.usageMap[date] = usage
	return nil
}

// ----------------------------------------------------------------
// WHEN Step (Act)
// ----------------------------------------------------------------

func iCallGetFullyUtilizedDatesFromTo(_ context.Context, startStr, endStr string) error {
	startDate := parseDate(tc.t, startStr)
	endDate := parseDate(tc.t, endStr)

	tc.ApplyUsageToWarehouses(tc.usageMap)

	tc.fullyUtilizedDatesResult, tc.fullyUtilizedDatesErr = tc.service.GetFullyUtilizedDates(startDate, endDate)
	return nil
}

// ----------------------------------------------------------------
// THEN Step (Assert)
// ----------------------------------------------------------------

func theFullyUtilizedDatesShouldBe(_ context.Context, table *godog.Table) error {
	assert.NoError(tc.t, tc.fullyUtilizedDatesErr, "unexpected error getting fully utilized dates")

	expectedDates := tableToDateSlice(tc.t, table)
	compareDates(tc.t, expectedDates, tc.fullyUtilizedDatesResult)
	return nil
}
