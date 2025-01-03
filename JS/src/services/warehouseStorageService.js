class WarehouseStorageService {
    constructor(warehouses) {
        this.warehouses = warehouses;
    }

    // FindAvailableWarehouse
    findAvailableWarehouse(startDate, endDate, requiredHeight, requiredWidth, requiredLength) {
        if (this.warehouses.length === 0) {
            throw new Error("No warehouses available");
        }

        if (requiredHeight <= 0 || requiredWidth <= 0 || requiredLength <= 0) {
            throw new Error("The 3D model has invalid dimensions (zero or negative)");
        }

        if (startDate > endDate) {
            throw new Error("Start date cannot be later than end date");
        }

        if (startDate < new Date()) {
            throw new Error("Start date cannot be in the past");
        }

        const requiredVolume = requiredHeight * requiredWidth * requiredLength;

        for (const warehouse of this.warehouses) {
            const warehouseVolume = warehouse.getWarehouseVolume();
            let canAccommodate = true;

            for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
                const occupiedVolume = warehouse.getVolumeOccupiedOnDay(day);
                if (occupiedVolume + requiredVolume > warehouseVolume) {
                    canAccommodate = false;
                    throw new Error("Required volume cannot be accommodated within the specified dates");
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
        const totalCapacity = this.warehouses.reduce((sum, warehouse) => sum + warehouse.getWarehouseVolume(), 0);

        for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
            const totalVolumeForDay = this.warehouses.reduce((sum, warehouse) =>
                sum + warehouse.getVolumeOccupiedOnDay(day), 0
            );

            if (totalVolumeForDay >= totalCapacity) {
                fullyUtilizedDates.push(new Date(day)); // Ensure dates are added as new instances
            }
        }

        return fullyUtilizedDates;
    }

    // CalculateAvailableCapacity
    calculateAvailableCapacity(startDate, endDate) {
        if (this.warehouses.length === 0) {
            throw new Error("No warehouses available");
        }

        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        const totalCapacity = this.warehouses.reduce((sum, warehouse) => sum + warehouse.getWarehouseVolume(), 0);
        const capacityMap = {};

        for (let day = new Date(startDate); day <= endDate; day = new Date(day.setDate(day.getDate() + 1))) {
            const totalVolumeForDay = this.warehouses.reduce((sum, warehouse) =>
                sum + warehouse.getVolumeOccupiedOnDay(day), 0
            );

            capacityMap[day.toISOString().split('T')[0]] = totalCapacity - totalVolumeForDay;
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