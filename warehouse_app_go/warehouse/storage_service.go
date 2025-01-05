package warehouse

import (
	"errors"
	"time"
)

type WarehouseStorageService struct {
	Warehouses []Warehouse
}

// -------------------------------------------------
// FindAvailableWarehouse
// -------------------------------------------------
func (s WarehouseStorageService) FindAvailableWarehouse(
	startDate, endDate time.Time,
	requiredHeight, requiredWidth, requiredLength float64,
) (int, error) {

	if len(s.Warehouses) == 0 {
		return -1, errors.New("no warehouses available")
	}

	if requiredHeight <= 0 || requiredWidth <= 0 || requiredLength <= 0 {
		return -1, errors.New("the 3D model has invalid dimensions (zero or negative)")
	}

	if startDate.After(endDate) {
		return -1, errors.New("start date cannot be later than end date")
	}

	if startDate.Before(time.Now()) {
		return -1, errors.New("start date cannot be in the past")
	}

	requiredVolume := requiredHeight * requiredWidth * requiredLength

	for _, warehouse := range s.Warehouses {
		warehouseVolume := warehouse.GetWarehouseVolume()
		canAccommodate := true

		for day := startDate; !day.After(endDate); day = day.AddDate(0, 0, 1) {
			occupiedVolume := warehouse.GetVolumeOccupiedOnDay(day)
			if occupiedVolume+requiredVolume > warehouseVolume {
				canAccommodate = false
				return -1, errors.New("required volume cannot be accommodated within the specified dates")
			}
		}

		if canAccommodate {
			return warehouse.Id, nil
		}
	}

	return -1, nil
}

// -------------------------------------------------
// GetFullyUtilizedDates
// -------------------------------------------------
func (service *WarehouseStorageService) GetFullyUtilizedDates(
	startDate, endDate time.Time,
) ([]time.Time, error) {

	if len(service.Warehouses) == 0 {
		return nil, errors.New("no warehouses available")
	}

	if startDate.After(endDate) {
		return nil, errors.New("the start date cannot be later than the end date")
	}

	var fullyUtilizedDates []time.Time
	totalCapacity := 0.0

	for _, warehouse := range service.Warehouses {
		totalCapacity += warehouse.GetWarehouseVolume()
	}

	for day := startDate; !day.After(endDate); day = day.AddDate(0, 0, 1) {
		totalVolumeForDay := 0.0
		for _, warehouse := range service.Warehouses {
			totalVolumeForDay += warehouse.GetVolumeOccupiedOnDay(day)
		}
		if totalVolumeForDay >= totalCapacity {
			fullyUtilizedDates = append(fullyUtilizedDates, day)
		}
	}

	return fullyUtilizedDates, nil
}

// -------------------------------------------------
// CalculateAvailableCapacity
// -------------------------------------------------
func (service *WarehouseStorageService) CalculateAvailableCapacity(
	startDate, endDate time.Time,
) (map[time.Time]float64, error) {

	if len(service.Warehouses) == 0 {
		return nil, errors.New("no warehouses available")
	}

	if startDate.After(endDate) {
		return nil, errors.New("the start date cannot be later than the end date")
	}

	totalCapacity := 0.0
	for _, warehouse := range service.Warehouses {
		totalCapacity += warehouse.GetWarehouseVolume()
	}

	capacityMap := make(map[time.Time]float64)
	for day := startDate; !day.After(endDate); day = day.AddDate(0, 0, 1) {
		totalVolumeForDay := 0.0
		for _, warehouse := range service.Warehouses {
			totalVolumeForDay += warehouse.GetVolumeOccupiedOnDay(day)
		}
		available := totalCapacity - totalVolumeForDay
		capacityMap[day] = available
	}

	return capacityMap, nil
}

// -------------------------------------------------
// GetLeastUsedWarehouse
// -------------------------------------------------
func (service *WarehouseStorageService) GetLeastUsedWarehouse(
	startDate, endDate time.Time,
) (int, error) {

	if len(service.Warehouses) == 0 {
		return -1, errors.New("no warehouses available")
	}

	if startDate.After(endDate) {
		return -1, errors.New("the start date cannot be later than the end date")
	}

	usageMap := make(map[int]float64)
	for _, warehouse := range service.Warehouses {
		totalVolumeDays := 0.0
		for day := startDate; !day.After(endDate); day = day.AddDate(0, 0, 1) {
			totalVolumeDays += warehouse.GetVolumeOccupiedOnDay(day)
		}
		usageMap[warehouse.Id] = totalVolumeDays
	}

	leastUsedWarehouseId := -1
	minUsage := float64(-1)
	for id, usage := range usageMap {
		if minUsage == -1 || usage < minUsage {
			minUsage = usage
			leastUsedWarehouseId = id
		}
	}

	if minUsage == 0 {
		return -1, nil
	}

	return leastUsedWarehouseId, nil
}
