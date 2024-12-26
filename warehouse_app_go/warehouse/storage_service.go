package warehouse

import (
	"errors"
	"fmt"
	"time"
)

type WarehouseStorageService struct {
	Warehouses []Warehouse
}

func (s WarehouseStorageService) FindAvailableWarehouse(startDate, endDate time.Time, requiredHeight, requiredWidth, requiredLength float64) (int, error) {
	if startDate.Before(time.Now()) || startDate.After(endDate) {
		return -1, errors.New("start date cannot be in the past or later than end date")
	}

	requiredVolume := requiredHeight * requiredWidth * requiredLength

	for _, warehouse := range s.Warehouses {
		warehouseVolume := warehouse.GetWarehouseVolume()
		canAccommodate := true

		for day := startDate; !day.After(endDate); day = day.AddDate(0, 0, 1) {
			occupiedVolume := warehouse.GetVolumeOccupiedOnDay(day)
			if occupiedVolume+requiredVolume > warehouseVolume {
				canAccommodate = false
				break
			}
		}

		if canAccommodate {
			return warehouse.Id, nil
		}
	}

	return -1, nil
}

func (service *WarehouseStorageService) GetFullyUtilizedDates(startDate, endDate time.Time) ([]time.Time, error) {
	if startDate.After(endDate) {
		return nil, fmt.Errorf("the start date cannot be later than the end date")
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

func (service *WarehouseStorageService) CalculateAvailableCapacity(startDate, endDate time.Time) (map[time.Time]float64, error) {
	if startDate.After(endDate) {
		return nil, fmt.Errorf("the start date cannot be later than the end date")
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

func (service *WarehouseStorageService) GetLeastUsedWarehouse(startDate, endDate time.Time) (int, error) {
	if startDate.After(endDate) {
		return -1, fmt.Errorf("the start date cannot be later than the end date")
	}

	if len(service.Warehouses) == 0 {
		return -1, nil
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
