package godogs

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

type WarehouseUsage struct {
	ID     int
	Volume float64
	Usage  float64
}

func initGetLeastUsedWarehouseSteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have warehouses with usage:$`, iHaveWarehousesWithUsage)
	ctx.Given(`^I have (\d+) warehouses?$`, iHaveWarehouses)
	ctx.When(`^I call GetLeastUsedWarehouse from "([^"]*)" to "([^"]*)"$`, iCallGetLeastUsedWarehouse)
	ctx.Then(`^the least used warehouse should be (\-?\d+)$`, theLeastUsedWarehouseShouldBe)
	ctx.Then(`^no error is returned$`, noErrorIsReturned)
}

// ------------------------------------------------------------------
// GIVEN Steps (Arrange)
// ------------------------------------------------------------------

func iHaveWarehousesWithUsage(_ context.Context, table *godog.Table) error {
	return tc.CreateWarehousesFromTable(table)
}

// ------------------------------------------------------------------
// WHEN Step (Act)
// ------------------------------------------------------------------

func iCallGetLeastUsedWarehouse(_ context.Context, startStr, endStr string) error {
	startDate := makeDate(startStr)
	endDate := makeDate(endStr)

	tc.leastUsedWarehouseResult, tc.leastUsedWarehouseErr = tc.service.GetLeastUsedWarehouse(startDate, endDate)
	return nil
}

// ------------------------------------------------------------------
// THEN Steps (Assert)
// ------------------------------------------------------------------

func theLeastUsedWarehouseShouldBe(_ context.Context, expectedID int) error {
	assert.NoError(tc.t, tc.leastUsedWarehouseErr)
	assert.Equal(tc.t, expectedID, tc.leastUsedWarehouseResult)
	return nil
}

func noErrorIsReturned(_ context.Context) error {
	assert.NoError(tc.t, tc.leastUsedWarehouseErr)
	return nil
}
