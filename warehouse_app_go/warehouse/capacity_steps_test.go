package warehouse

import (
	"context"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

type DateCapacity struct {
	Date     time.Time
	Capacity float64
}

func initCalculateCapacitySteps(ctx *godog.ScenarioContext) {
	// GIVEN
	ctx.Given(`^I have (\d+) warehouse with total volume (\d+\.?\d*)$`, iHaveWarehousesWithTotalVolume)
	ctx.Given(`^I have (\d+) warehouses$`, iHaveWarehouses)
	ctx.Given(`^warehouse usage is:$`, warehouseUsageIs)

	// WHEN
	ctx.When(`^I call CalculateAvailableCapacity from "([^"]*)" to "([^"]*)"$`, iCallCalculateAvailableCapacity)

	// THEN
	ctx.Then(`^the available capacities should be:$`, theAvailableCapacitiesShouldBe)
	ctx.Then(`^an error should be returned with message "([^"]*)"$`, anErrorShouldBeReturnedWithMessage)
}

// ------------------------------------------------------------------
// GIVEN Steps (Arrange)
// ------------------------------------------------------------------

func iHaveWarehousesWithTotalVolume(count int, volume float64) error {
	tc.CreateWarehousesWithVolume(count, volume)
	return nil
}

func iHaveWarehouses(count int) error {
	tc.CreateWarehousesWithVolume(count, 100.0)
	return nil
}

func warehouseUsageIs(_ context.Context, table *godog.Table) error {
	tc.usageMap = tableToTimeMap(tc.t, table, 0, 1)
	return nil
}

// ------------------------------------------------------------------
// WHEN Step (Act)
// ------------------------------------------------------------------

func iCallCalculateAvailableCapacity(startStr, endStr string) error {
	startDate := parseDate(tc.t, startStr)
	endDate := parseDate(tc.t, endStr)

	tc.capacityMap, tc.calculateCapacityErr = tc.CalculateAvailableCapacityWithUsage(
		startDate,
		endDate,
		tc.usageMap,
	)
	return nil
}

// ------------------------------------------------------------------
// THEN Steps (Assert)
// ------------------------------------------------------------------

func theAvailableCapacitiesShouldBe(_ context.Context, table *godog.Table) error {
	assert.NoError(tc.t, tc.calculateCapacityErr, "unexpected calculation error")
	assert.NotNil(tc.t, tc.capacityMap, "capacity map should not be nil")

	expectedMap := tableToTimeMap(tc.t, table, 0, 1)
	assert.Equal(tc.t, expectedMap, tc.capacityMap, "capacity map mismatch")
	return nil
}

func anErrorShouldBeReturnedWithMessage(msg string) error {
	allErrors := []error{
		tc.calculateCapacityErr,
		tc.searchError,
		tc.fullyUtilizedDatesErr,
		tc.leastUsedWarehouseErr,
	}

	assert.True(tc.t, findMatchingError(allErrors, msg),
		"expected error containing %q", msg)
	return nil
}
