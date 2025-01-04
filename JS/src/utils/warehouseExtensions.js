class WarehouseExtensions {
    static getWarehouseVolume(warehouse) {
        if (warehouse.maxCapacity) {
            return warehouse.maxCapacity;
        }
        const { height, width, length } = warehouse;
        if (height && width && length) {
            return height * width * length;
        }
        throw new Error("Warehouse does not have enough information to calculate volume.");
    }

    static getItemVolume(item) {
        return item.volume || 0; 
    }

    static getVolumeOccupiedOnDay(warehouse, day) {
        const normalizedDay = new Date(day).setHours(0, 0, 0, 0);
        return warehouse.items
            .filter(item => {
                const normalizedStartDate = new Date(item.startDate).setHours(0, 0, 0, 0);
                const normalizedEndDate = new Date(item.endDate).setHours(0, 0, 0, 0);
                return (
                    item.isActive !== false &&
                    normalizedDay >= normalizedStartDate &&
                    normalizedDay <= normalizedEndDate
                );
            })
            .reduce((total, item) => total + WarehouseExtensions.getItemVolume(item), 0);
    }
}

module.exports = WarehouseExtensions;
