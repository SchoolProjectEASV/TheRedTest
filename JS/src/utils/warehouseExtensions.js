class WarehouseExtensions {
    static getWarehouseVolume(warehouse) {
        if (warehouse.maxCapacity) {
            return warehouse.maxCapacity;
        }
        const { height, width, length } = warehouse;
        return height * width * length;
    }

    static getItemVolume(item) {
        return item.itemHeight * item.itemWidth * item.itemLength;
    }

    static getVolumeOccupiedOnDay(warehouse, day) {
        const normalizedDay = new Date(day.setHours(0, 0, 0, 0));

        return warehouse.items
            .filter(item => {
                const normalizedStartDate = new Date(item.startDate).setHours(0, 0, 0, 0);
                const normalizedEndDate = new Date(item.endDate).setHours(0, 0, 0, 0);
                return (
                    item.isActive &&
                    normalizedDay >= normalizedStartDate &&
                    normalizedDay <= normalizedEndDate
                );
            })
            .reduce((total, item) => total + WarehouseExtensions.getItemVolume(item), 0);
    }
}