package warehouse

import (
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
	t        *testing.T
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

func (tc *TestState) GetAllWarehouseItems() []Item {
	var allItems []Item
	for _, wh := range tc.service.Warehouses {
		allItems = append(allItems, wh.Items...)
	}
	return allItems
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

func findMatchingError(errors []error, msg string) bool {
	for _, err := range errors {
		if err != nil && strings.Contains(err.Error(), msg) {
			return true
		}
	}
	return false
}

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

func parseDimensionsTable(table *godog.Table) *ThreeDRoom {

	headerRow := table.Rows[0]
	valueRow := table.Rows[1]

	values := make(map[string]string)
	for i, header := range headerRow.Cells {
		values[header.Value] = valueRow.Cells[i].Value
	}

	height := parseFloat(values["height"])

	width := parseFloat(values["width"])

	length := parseFloat(values["length"])

	return &ThreeDRoom{
		Height: height,
		Width:  width,
		Length: length,
	}
}

func tableToDateSlice(t require.TestingT, table *godog.Table) []time.Time {
	var dates []time.Time
	for _, row := range table.Rows[1:] {
		if len(row.Cells) > 0 {
			dates = append(dates, parseDate(t, row.Cells[0].Value))
		}
	}
	return dates
}

func (tc *TestState) CreateWarehousesFromTable(table *godog.Table) error {
	warehouses := make([]Warehouse, len(table.Rows)-1)

	for i, row := range table.Rows[1:] {
		id, _ := strconv.Atoi(row.Cells[0].Value)
		volume, _ := strconv.ParseFloat(row.Cells[1].Value, 64)
		usage, _ := strconv.ParseFloat(row.Cells[2].Value, 64)

		warehouses[i] = Warehouse{
			Id: id,
			MaxCapacity: ThreeDRoom{
				Height: volume,
				Width:  1,
				Length: 1,
			},
		}

		if usage > 0 {
			warehouses[i].Items = []Item{{
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
	}

	tc.service.Warehouses = warehouses
	return nil
}

func makeDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
