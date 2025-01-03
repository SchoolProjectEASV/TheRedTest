package warehouse

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/require"
)

// TestContext is a struct that holds scenario data for tests.
// It's not part of the domain; it's purely for test setup.
type TestState struct {
	t        *testing.T // Add this field for assertions
	service  WarehouseStorageService
	usageMap map[time.Time]float64

	calculateCapacityErr  error
	fullyUtilizedDatesErr error
	leastUsedWarehouseErr error

	capacityMap              map[time.Time]float64
	searchResult             int
	searchError              error
	currentDate              time.Time
	fullyUtilizedDatesResult []time.Time
	leastUsedWarehouseResult int
}

func NewTestContext(t *testing.T) *TestState {
	return &TestState{
		t:        t,
		usageMap: make(map[time.Time]float64),
		service:  WarehouseStorageService{},
	}
}

func (tc *TestState) CreateWarehouses(count int, volume float64) {
	for i := 1; i <= count; i++ {
		w := Warehouse{
			Id: i,
			MaxCapacity: ThreeDRoom{
				Height: volume,
				Width:  1,
				Length: 1,
			},
		}
		tc.service.Warehouses = append(tc.service.Warehouses, w)
	}
}

func (tc *TestState) ClearAllWarehousesUsage() {
	for i := range tc.service.Warehouses {
		tc.service.Warehouses[i].Items = nil
	}
}

func (tc *TestState) SetUsage(day time.Time, usageVal float64) {
	tc.usageMap[day] = usageVal
}

func (tc *TestState) CreateWarehousesWithVolume(count int, volume float64) {
	for i := 1; i <= count; i++ {
		wh := Warehouse{
			Id: i,
			MaxCapacity: ThreeDRoom{
				Height: volume,
				Width:  1,
				Length: 1,
			},
		}
		tc.service.Warehouses = append(tc.service.Warehouses, wh)
	}
}

func (tc *TestState) ApplyUsageToWarehouses(usage map[time.Time]float64) {
	for i := range tc.service.Warehouses {
		wh := &tc.service.Warehouses[i]
		wh.Items = nil
		for day, usageVal := range usage {
			wh.Items = append(wh.Items, Item{
				ItemId:     1,
				ItemName:   "GeneratedUsage",
				ItemHeight: usageVal,
				ItemWidth:  1,
				ItemLength: 1,
				StartDate:  day,
				EndDate:    day,
				IsActive:   true,
			})
		}
	}
}

func (tc *TestState) ApplyUsageMap(usageMap map[time.Time]float64) {
	for i := range tc.service.Warehouses {
		for day, usageVal := range usageMap {
			tc.service.Warehouses[i].Items = append(tc.service.Warehouses[i].Items, Item{
				ItemId:     1,
				ItemName:   "GeneratedUsage",
				ItemHeight: usageVal,
				ItemWidth:  1,
				ItemLength: 1,
				StartDate:  day,
				EndDate:    day,
				IsActive:   true,
			})
		}
	}
}

func (tc *TestState) CalculateAvailableCapacityWithUsage(
	start, end time.Time,
	usage map[time.Time]float64,
) (map[time.Time]float64, error) {
	// Internally loop to apply usage to each warehouse
	for i := range tc.service.Warehouses {
		w := &tc.service.Warehouses[i]
		w.Items = nil
		for day, usageVal := range usage {
			w.Items = append(w.Items, Item{
				ItemId:     1,
				ItemName:   "GeneratedUsage",
				ItemHeight: usageVal,
				ItemWidth:  1,
				ItemLength: 1,
				StartDate:  day,
				EndDate:    day,
				IsActive:   true,
			})
		}
	}
	return tc.service.CalculateAvailableCapacity(start, end)
}

// Utility function to check if error exists in slice
func findMatchingError(errors []error, msg string) bool {
	for _, err := range errors {
		if err != nil && strings.Contains(err.Error(), msg) {
			return true
		}
	}
	return false
}

// Parse date string in format "2006-01-02"
func parseDate(t require.TestingT, dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	require.NoError(t, err, "invalid date format %q", dateStr)
	return date
}

func compareDates(t require.TestingT, expected, actual []time.Time) {
	require.Equal(t, len(expected), len(actual), "date array length mismatch")

	for i := range expected {
		require.True(t, expected[i].Equal(actual[i]),
			"date mismatch at index %d: expected %s, got %s",
			i, expected[i].Format("2006-01-02"), actual[i].Format("2006-01-02"))
	}
}

func (tc *TestState) Service() WarehouseStorageService {
	return tc.service
}

func tableToTimeMap(t require.TestingT, table *godog.Table, dateCol, valueCol int) map[time.Time]float64 {
	result := make(map[time.Time]float64)
	for _, row := range table.Rows[1:] {
		date := parseDate(t, row.Cells[dateCol].Value)
		value := parseFloat(row.Cells[valueCol].Value)
		result[date] = value
	}
	return result
}

func parseFloat(value string) float64 {
	val, _ := strconv.ParseFloat(value, 64)
	return val
}

func parseDimensionsTable(table *godog.Table) (*ThreeDRoom, error) {
	if len(table.Rows) != 2 {
		return nil, fmt.Errorf("dimensions table should have exactly 2 rows (header and values), got %d", len(table.Rows))
	}

	headerRow := table.Rows[0]
	valueRow := table.Rows[1]

	if len(headerRow.Cells) != 3 || len(valueRow.Cells) != 3 {
		return nil, fmt.Errorf("dimensions table should have exactly 3 columns, got %d", len(headerRow.Cells))
	}

	// Create a map of header names to values
	values := make(map[string]string)
	for i, header := range headerRow.Cells {
		values[header.Value] = valueRow.Cells[i].Value
	}

	// Parse each dimension
	height, err := strconv.ParseFloat(values["height"], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid height value: %s", values["height"])
	}

	width, err := strconv.ParseFloat(values["width"], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid width value: %s", values["width"])
	}

	length, err := strconv.ParseFloat(values["length"], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid length value: %s", values["length"])
	}

	return &ThreeDRoom{
		Height: height,
		Width:  width,
		Length: length,
	}, nil
}
