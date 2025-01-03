package warehouse

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

// Define a struct to match the table structure
type Dimensions struct {
	Height float64
	Width  float64
	Length float64
}

func initFindAvailableWarehouseSteps(ctx *godog.ScenarioContext) {
	// GIVEN
	ctx.Given(`^today is "([^"]*)"$`, todayIs)
	ctx.Given(`^I have (\d+) warehouses?$`, iHaveWarehouses)
	ctx.Given(`^I have (\d+) warehouse with total volume (\d+\.?\d*)$`, iHaveWarehouseWithVolume)
	ctx.Given(`^the warehouse usage is empty on all days$`, theWarehouseUsageIsEmptyOnAllDays)

	// WHEN
	ctx.When(`^I call FindAvailableWarehouse from "([^"]*)" to "([^"]*)" with dimensions:$`,
		iCallFindAvailableWarehouseFromToWithDimensions)

	// THEN
	ctx.Then(`^I should receive warehouse ID (\d+)$`, iShouldReceiveWarehouseID)
	ctx.Then(`^an error should be returned with message "([^"]*)"$`, iShouldReceiveErrorMessage)
}

// -------------------
// GIVEN Steps
// -------------------

func todayIs(ctx context.Context, dateStr string) {
	t := godog.T(ctx)
	tc.currentDate = parseDate(t, dateStr)
}

func iHaveWarehouseWithVolume(ctx context.Context, count int, volume float64) {
	tc.CreateWarehouses(count, volume)
}

func theWarehouseUsageIsEmptyOnAllDays(ctx context.Context) {
	t := godog.T(ctx)
	tc.ClearAllWarehousesUsage()

	for _, wh := range tc.service.Warehouses {
		assert.Empty(t, wh.Items, "warehouse %d should have no items", wh.Id)
	}
}

// -------------------
// WHEN Step
// -------------------

func iCallFindAvailableWarehouseFromToWithDimensions(ctx context.Context, startStr, endStr string, table *godog.Table) error {
	t := godog.T(ctx)

	startDate := parseDate(t, startStr)
	endDate := parseDate(t, endStr)

	dims, err := parseDimensionsTable(table)
	if err != nil {
		return err
	}

	tc.searchResult, tc.searchError = tc.service.FindAvailableWarehouse(
		startDate,
		endDate,
		dims.Height,
		dims.Width,
		dims.Length,
	)
	return nil
}

// -------------------
// THEN Steps
// -------------------

func iShouldReceiveWarehouseID(ctx context.Context, expectedID int) {
	t := godog.T(ctx)

	assert.NoError(t, tc.searchError, "unexpected error")
	assert.Equal(t, expectedID, tc.searchResult, "warehouse ID mismatch")
}

func iShouldReceiveErrorMessage(ctx context.Context, expectedMsg string) {
	t := godog.T(ctx)

	assert.Error(t, tc.searchError, "expected an error")
	assert.Equal(t, expectedMsg, tc.searchError.Error(), "error message mismatch")
}
