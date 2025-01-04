// services/warehouseStorageService.js
const WarehouseExtensions = require('../utils/warehouseExtensions');

class WarehouseStorageService {
  constructor(warehouses) {
    this.warehouses = warehouses;
  }

  // ---------------------------------------------
  //  findAvailableWarehouse
  // ---------------------------------------------
  findAvailableWarehouse(startDate, endDate, requiredHeight, requiredWidth, requiredLength) {
    if (this.warehouses.length === 0) {
      throw new Error("no warehouses available");
    }

    if (requiredHeight <= 0 || requiredWidth <= 0 || requiredLength <= 0) {
      throw new Error("the 3D model has invalid dimensions (zero or negative)");
    }

    if (startDate > endDate) {
      throw new Error("start date cannot be later than end date");
    }

    if (startDate < new Date()) {
      throw new Error("start date cannot be in the past");
    }

    const requiredVolume = requiredHeight * requiredWidth * requiredLength;

    for (const warehouse of this.warehouses) {
      const warehouseVolume = WarehouseExtensions.getWarehouseVolume(warehouse);
      let canAccommodate = true;

      for (let day = new Date(startDate.getTime()); day <= endDate; day.setDate(day.getDate() + 1)) {
        const occupiedVolume = WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day);
        if (occupiedVolume + requiredVolume > warehouseVolume) {
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

  // ---------------------------------------------
  //  getFullyUtilizedDates
  // ---------------------------------------------
  getFullyUtilizedDates(startDate, endDate) {
    if (this.warehouses.length === 0) {
      throw new Error("no warehouses available");
    }

    if (startDate > endDate) {
      throw new Error("the start date cannot be later than the end date");
    }

    const totalCapacity = this.warehouses.reduce(
      (sum, warehouse) => sum + WarehouseExtensions.getWarehouseVolume(warehouse),
      0
    );

    const fullyUtilizedDates = [];

    for (let day = new Date(startDate.getTime()); day <= endDate; day.setDate(day.getDate() + 1)) {
      const totalVolumeForDay = this.warehouses.reduce(
        (sum, warehouse) => sum + WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day),
        0
      );

      if (totalVolumeForDay >= totalCapacity) {
        fullyUtilizedDates.push(new Date(day));
      }
    }

    return fullyUtilizedDates;
  }

  // ---------------------------------------------
  //  calculateAvailableCapacity
  // ---------------------------------------------
  calculateAvailableCapacity(startDate, endDate) {
    if (this.warehouses.length === 0) {
      throw new Error("no warehouses available");
    }

    if (startDate > endDate) {
      throw new Error("the start date cannot be later than the end date");
    }

    const totalCapacity = this.warehouses.reduce(
      (sum, warehouse) => sum + WarehouseExtensions.getWarehouseVolume(warehouse),
      0
    );

    const capacityMap = {};

    for (let day = new Date(startDate.getTime()); day <= endDate; day.setDate(day.getDate() + 1)) {
      const totalVolumeForDay = this.warehouses.reduce(
        (sum, warehouse) => sum + WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day),
        0
      );
      const dayString = day.toISOString().split('T')[0];
      capacityMap[dayString] = totalCapacity - totalVolumeForDay;
    }

    return capacityMap;
  }

  // ---------------------------------------------
  //  getLeastUsedWarehouse
  // ---------------------------------------------
  getLeastUsedWarehouse(startDate, endDate) {
    if (this.warehouses.length === 0) {
      throw new Error("no warehouses available");
    }

    if (startDate > endDate) {
      throw new Error("the start date cannot be later than the end date");
    }

    const usageMap = new Map();

    for (const warehouse of this.warehouses) {
      let totalVolumeDays = 0;

      for (let day = new Date(startDate.getTime()); day <= endDate; day.setDate(day.getDate() + 1)) {
        totalVolumeDays += WarehouseExtensions.getVolumeOccupiedOnDay(warehouse, day);
      }

      usageMap.set(warehouse.id, totalVolumeDays);
    }

    let leastUsedId = -1;
    let minUsage = Number.MAX_SAFE_INTEGER;

    for (const [id, usage] of usageMap) {
      if (usage < minUsage) {
        minUsage = usage;
        leastUsedId = id;
      }
    }

    if (minUsage === 0) {
      return -1;
    }

    return leastUsedId;
  }
}

module.exports = WarehouseStorageService;
