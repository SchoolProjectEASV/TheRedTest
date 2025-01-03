package warehouse

import (
	"context"
	"strconv"
	"time"

	"github.com/cucumber/godog" // Updated import
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
	ctx.Then(`^the least used warehouse should be (\d+)$`, theLeastUsedWarehouseShouldBe)
	ctx.Then(`^no error is returned$`, noErrorIsReturned)
}

func iHaveWarehousesWithUsage(_ context.Context, table *godog.Table) error {
	// Create warehouses based on table rows
	warehouses := make([]Warehouse, len(table.Rows)-1) // -1 for header

	for i, row := range table.Rows[1:] {
		id, _ := strconv.Atoi(row.Cells[0].Value)
		volume, _ := strconv.ParseFloat(row.Cells[1].Value, 64)
		usage, _ := strconv.ParseFloat(row.Cells[2].Value, 64)

		warehouse := Warehouse{
			Id: id,
			MaxCapacity: ThreeDRoom{
				Height: volume,
				Width:  1,
				Length: 1,
			},
		}

		if usage > 0 {
			warehouse.Items = []Item{{
				ItemId:     1,
				ItemName:   "GeneratedUsage",
				ItemHeight: usage,
				ItemWidth:  1,
				ItemLength: 1,
				StartDate:  makeDate("2025-01-10"),
				EndDate:    makeDate("2025-01-10"),
				IsActive:   true,
			}}
		}

		warehouses[i] = warehouse
	}

	tc.service.Warehouses = warehouses
	return nil
}

func iCallGetLeastUsedWarehouse(_ context.Context, startStr, endStr string) error {
	startDate := makeDate(startStr)
	endDate := makeDate(endStr)

	tc.leastUsedWarehouseResult, tc.leastUsedWarehouseErr = tc.service.GetLeastUsedWarehouse(startDate, endDate)
	return nil
}

func theLeastUsedWarehouseShouldBe(_ context.Context, expectedID int) error {
	assert.NoError(tc.t, tc.leastUsedWarehouseErr)
	assert.Equal(tc.t, expectedID, tc.leastUsedWarehouseResult)
	return nil
}

func noErrorIsReturned(_ context.Context) error {
	assert.NoError(tc.t, tc.leastUsedWarehouseErr)
	return nil
}

func makeDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
