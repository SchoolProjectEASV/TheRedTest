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
        return warehouse.items
            .filter(item =>
                item.isActive &&
                day >= item.startDate &&
                day <= item.endDate
            )
            .reduce((total, item) => total + WarehouseExtensions.getItemVolume(item), 0);
    }
}