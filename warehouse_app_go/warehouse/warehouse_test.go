package warehouse

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var (
	service      WarehouseStorageService
	searchResult int
	searchError  error

	// Use "pretty" formatter + color, search for ../features folder
	opts = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
		Paths:  []string{"../features"},
	}
)

func TestMain(m *testing.M) {
	// Run Godog test suite
	status := godog.TestSuite{
		Name:                 "warehouse",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	// Run any other standard Go tests (if present)
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	// Perform any global setup, if necessary
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the following warehouses exist:$`, theFollowingWarehousesExist)
	ctx.Step(`^the following items exist in warehouse (\d+):$`, theFollowingItemsExistInWarehouse)
	ctx.Step(`^I search for an available warehouse from "([^"]*)" to "([^"]*)" for an item of size (\d+)x(\d+)x(\d+)$`, iSearchForAnAvailableWarehouse)
	ctx.Step(`^the available warehouse ID should be (\d+)$`, theAvailableWarehouseIDShouldBe)
}

func theFollowingWarehousesExist(ctx context.Context, table *godog.Table) (context.Context, error) {
	var warehouses []Warehouse
	for _, row := range table.Rows[1:] {
		var id int
		var height, width, length float64
		if _, err := fmt.Sscanf(row.Cells[0].Value, "%d", &id); err != nil {
			return ctx, err
		}
		if _, err := fmt.Sscanf(row.Cells[1].Value, "%f", &height); err != nil {
			return ctx, err
		}
		if _, err := fmt.Sscanf(row.Cells[2].Value, "%f", &width); err != nil {
			return ctx, err
		}
		if _, err := fmt.Sscanf(row.Cells[3].Value, "%f", &length); err != nil {
			return ctx, err
		}

		warehouses = append(warehouses, Warehouse{
			Id:          id,
			MaxCapacity: ThreeDRoom{Height: height, Width: width, Length: length},
		})
	}
	service = WarehouseStorageService{Warehouses: warehouses}
	return ctx, nil
}

func theFollowingItemsExistInWarehouse(ctx context.Context, warehouseID int, table *godog.Table) (context.Context, error) {
	// We'll parse StartDate and EndDate with a custom layout
	const dateLayout = "2006-01-02T15:04:05"

	for _, row := range table.Rows[1:] {
		var item Item
		if _, err := fmt.Sscanf(row.Cells[0].Value, "%d", &item.ItemId); err != nil {
			return ctx, err
		}
		item.ItemName = row.Cells[1].Value

		if _, err := fmt.Sscanf(row.Cells[2].Value, "%f", &item.ItemHeight); err != nil {
			return ctx, err
		}
		if _, err := fmt.Sscanf(row.Cells[3].Value, "%f", &item.ItemWidth); err != nil {
			return ctx, err
		}
		if _, err := fmt.Sscanf(row.Cells[4].Value, "%f", &item.ItemLength); err != nil {
			return ctx, err
		}

		// Parse the StartDate
		var parseErr error
		if item.StartDate, parseErr = time.Parse(dateLayout, row.Cells[5].Value); parseErr != nil {
			return ctx, fmt.Errorf("failed to parse StartDate %q: %w", row.Cells[5].Value, parseErr)
		}

		// Parse the EndDate
		if item.EndDate, parseErr = time.Parse(dateLayout, row.Cells[6].Value); parseErr != nil {
			return ctx, fmt.Errorf("failed to parse EndDate %q: %w", row.Cells[6].Value, parseErr)
		}

		if _, err := fmt.Sscanf(row.Cells[7].Value, "%t", &item.IsActive); err != nil {
			return ctx, err
		}

		// Attach the item to the specified warehouse
		for i, w := range service.Warehouses {
			if w.Id == warehouseID {
				service.Warehouses[i].Items = append(service.Warehouses[i].Items, item)
				break
			}
		}
	}
	return ctx, nil
}

func iSearchForAnAvailableWarehouse(ctx context.Context, startDate, endDate string, height, width, length int) (context.Context, error) {
	// The Gherkin scenario uses e.g. "2024-12-26" (no time-of-day).
	// So we'll parse with "2006-01-02" for these two fields.
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return ctx, err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return ctx, err
	}
	searchResult, searchError = service.FindAvailableWarehouse(
		start,
		end,
		float64(height),
		float64(width),
		float64(length),
	)
	return ctx, nil
}

func theAvailableWarehouseIDShouldBe(expectedID int) error {
	if searchError != nil {
		return searchError
	}
	if searchResult != expectedID {
		return fmt.Errorf("expected warehouse ID %d, but got %d", expectedID, searchResult)
	}
	return nil
}
