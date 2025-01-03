package warehouse

import (
	"context"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func initGetFullyUtilizedDatesSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I call GetFullyUtilizedDates from "([^"]*)" to "([^"]*)"$`, iCallGetFullyUtilizedDatesFromTo)
	ctx.Step(`^the fully utilized dates should be:$`, theFullyUtilizedDatesShouldBe)
}

// ----------------------------------------------------------------
// WHEN Step
// ----------------------------------------------------------------

func iCallGetFullyUtilizedDatesFromTo(ctx context.Context, startStr, endStr string) error {
	t := godog.T(ctx)

	startDate := parseDate(t, startStr)
	endDate := parseDate(t, endStr)

	tc.ApplyUsageToWarehouses(tc.usageMap)

	tc.fullyUtilizedDatesResult, tc.fullyUtilizedDatesErr = tc.service.GetFullyUtilizedDates(startDate, endDate)
	return nil
}

// ----------------------------------------------------------------
// THEN Step
// ----------------------------------------------------------------

func theFullyUtilizedDatesShouldBe(ctx context.Context, table *godog.Table) (context.Context, error) {
	t := godog.T(ctx)

	// Check for errors
	assert.NoError(t, tc.fullyUtilizedDatesErr, "unexpected error getting fully utilized dates")

	// Get dates from table
	expectedDates := make([]time.Time, 0)
	for _, row := range table.Rows[1:] {
		if len(row.Cells) > 0 {
			expectedDates = append(expectedDates, parseDate(t, row.Cells[0].Value))
		}
	}

	compareDates(t, expectedDates, tc.fullyUtilizedDatesResult)
	return ctx, nil
}
