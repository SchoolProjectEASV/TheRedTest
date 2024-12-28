class WarehouseStorageService {
    constructor(warehouses) {
        this.warehouses = warehouses;
    }

    findAvailableWarehouse(startDate, endDate, requiredHeight, requiredWidth, requiredLength) {
        if (startDate <= new Date() || startDate > endDate) {
            throw new Error("The start date cannot be in the past or later than the end date");
        }

        const requiredVolume = requiredHeight * requiredWidth * requiredLength;

        for (const warehouse of this.warehouses) {
            const warehouseVolume = WarehouseExtensions.getWarehouseVolume(warehouse);
            let canAccommodate = true;

            for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
                const occupiedVolume = WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day);
                if (occupiedVolume + requiredVolume > warehouseVolume) {
                    canAccommodate = false;
                    break;
                }
            }

            if (canAccommodate) return warehouse.id;
        }

        return -1;
    }

    getFullyUtilizedDates(startDate, endDate) {
        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        const totalCapacity = this.warehouses.reduce((sum, warehouse) =>
            sum + WarehouseExtensions.getWarehouseVolume(warehouse), 0);

        const allItems = this.warehouses.flatMap(warehouse => warehouse.items.filter(item => item.isActive));
        const fullyUtilizedDates = [];

        for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
            const totalVolumeForDay = allItems
                .filter(item => day >= item.startDate && day <= item.endDate)
                .reduce((sum, item) => sum + WarehouseExtensions.getItemVolume(item), 0);

            if (totalVolumeForDay >= totalCapacity) {
                fullyUtilizedDates.push(new Date(day));
            }
        }

        return fullyUtilizedDates;
    }

    // Calculate available capacity
    calculateAvailableCapacity(startDate, endDate) {
        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        const totalCapacity = this.warehouses.reduce((sum, warehouse) =>
            sum + WarehouseExtensions.getWarehouseVolume(warehouse), 0);

        const allItems = this.warehouses.flatMap(warehouse => warehouse.items.filter(item => item.isActive));
        const capacityMap = {};

        for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
            const totalVolumeForDay = allItems
                .filter(item => day >= item.startDate && day <= item.endDate)
                .reduce((sum, item) => sum + WarehouseExtensions.getItemVolume(item), 0);

            capacityMap[new Date(day)] = totalCapacity - totalVolumeForDay;
        }

        return capacityMap;
    }

    getLeastUsedWarehouse(startDate, endDate) {
        if (startDate > endDate) {
            throw new Error("The start date cannot be later than the end date");
        }

        if (!this.warehouses.length) return -1;

        const usageMap = new Map();

        for (const warehouse of this.warehouses) {
            let totalVolumeDays = 0;

            for (let day = new Date(startDate); day <= endDate; day.setDate(day.getDate() + 1)) {
                const occupiedVolume = WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day);
                totalVolumeDays += occupiedVolume;
            }

            usageMap.set(warehouse.id, totalVolumeDays);
        }

        const leastUsed = [...usageMap.entries()].reduce((least, current) =>
            current[1] < least[1] ? current : least, [null, Infinity]);

        return leastUsed[1] === Infinity ? -1 : leastUsed[0];
    }
}