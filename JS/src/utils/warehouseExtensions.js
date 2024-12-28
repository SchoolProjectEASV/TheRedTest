class WarehouseExtensions {
    // Calculate the volume of a warehouse
    static getWarehouseVolume(warehouse) {
        const { height, width, length } = warehouse.maxCapacity;
        return height * width * length;
    }

    // Calculate the volume of an item
    static getItemVolume(item) {
        return item.itemHeight * item.itemWidth * item.itemLength;
    }

    // Get the total occupied volume in a warehouse on a specific day
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