package main

import (
	"fmt"
	"time"
	"warehouse_app_go/warehouse"
)

func main() {
	maxCapacity := warehouse.ThreeDRoom{Height: 10, Width: 10, Length: 10}
	items := []warehouse.Item{
		{ItemId: 1, ItemName: "Item1", ItemHeight: 2, ItemWidth: 2, ItemLength: 2, StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 10), IsActive: true},
	}
	warehouses := []warehouse.Warehouse{
		{Id: 1, MaxCapacity: maxCapacity, Items: items},
	}

	service := warehouse.WarehouseStorageService{Warehouses: warehouses}

	startDate := time.Now().AddDate(0, 0, 1)
	endDate := time.Now().AddDate(0, 0, 5)

	id, err := service.FindAvailableWarehouse(startDate, endDate, 2, 2, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Available Warehouse ID:", id)
	}

	fullyUtilizedDates, err := service.GetFullyUtilizedDates(startDate, endDate)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Fully Utilized Dates:", fullyUtilizedDates)
	}

	availableCapacity, err := service.CalculateAvailableCapacity(startDate, endDate)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Available Capacity:", availableCapacity)
	}

	leastUsedWarehouseId, err := service.GetLeastUsedWarehouse(startDate, endDate)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Least Used Warehouse ID:", leastUsedWarehouseId)
	}
}
