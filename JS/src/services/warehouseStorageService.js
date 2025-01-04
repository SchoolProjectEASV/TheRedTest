const WarehouseExtensions = require('../utils/warehouseExtensions');

class WarehouseStorageService {
    constructor(warehouses) {
        this.warehouses = warehouses;
    }

    // FindAvailableWarehouse
    findAvailableWarehouse(startDate, endDate, requiredHeight, requiredWidth, requiredLength) {
        const requiredVolume = requiredHeight * requiredWidth * requiredLength;

        for (const warehouse of this.warehouses) {
            let canAccommodate = true;
            for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
                const totalVolumeForDay = WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day);
                const warehouseVolume = WarehouseExtensions.getWarehouseVolume(warehouse);
                if (totalVolumeForDay + requiredVolume > warehouseVolume) {
                    canAccommodate = false;
                    break;
                }
            }
            if (canAccommodate) {
                return warehouse.id;
            }
        }
        return -1;
    }


    // GetFullyUtilizedDates
    getFullyUtilizedDates(startDate, endDate) {
        if (this.warehouses.length === 0) {
            throw new Error("No warehouses available");
        }

        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        const fullyUtilizedDates = [];
        const totalCapacity = this.warehouses.reduce((sum, warehouse) =>
            sum + WarehouseExtensions.getWarehouseVolume(warehouse), 0);

        for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
            const totalVolumeForDay = this.warehouses.reduce((sum, warehouse) =>
                sum + WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day), 0);

            if (totalVolumeForDay >= totalCapacity) {
                fullyUtilizedDates.push(new Date(day));  // Add fully utilized date to array
            }
        }

        return fullyUtilizedDates;
    }

    // CalculateAvailableCapacity
    calculateAvailableCapacity(startDate, endDate) {
        if (this.warehouses.length === 0) {
            throw new Error("no warehouses available");
        }

        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        const totalCapacity = this.warehouses.reduce((sum, warehouse) =>
            sum + WarehouseExtensions.getWarehouseVolume(warehouse), 0
        );

        const capacityMap = {};

        for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
            const currentDate = new Date(day).toISOString().split('T')[0];
            const totalVolumeForDay = this.warehouses.reduce((sum, warehouse) =>
                sum + WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day), 0
            );

            capacityMap[currentDate] = totalCapacity - totalVolumeForDay;
        }

        return capacityMap;
    }


    // GetLeastUsedWarehouse
    getLeastUsedWarehouse(startDate, endDate) {
        if (this.warehouses.length === 0) {
            throw new Error("No warehouses available");
        }

        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        const usageMap = new Map();

        for (const warehouse of this.warehouses) {
            let totalVolumeDays = 0;

            for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
                totalVolumeDays += warehouse.getVolumeOccupiedOnDay(day);
            }

            usageMap.set(warehouse.id, totalVolumeDays);
        }

        let leastUsedWarehouseId = -1;
        let minUsage = Number.MAX_SAFE_INTEGER;

        for (const [id, usage] of usageMap) {
            if (usage < minUsage) {
                minUsage = usage;
                leastUsedWarehouseId = id;
            }
        }

        if (minUsage === 0) {
            return -1;
        }

        return leastUsedWarehouseId;
    }
}

module.exports = WarehouseStorageService;
