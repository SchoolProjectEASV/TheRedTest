using System;
using System.Collections.Generic;
using System.Linq;
using WarehouseProject.Models;

namespace WarehouseProject.Services
{
    public class WarehouseStorageService
    {
        private readonly IEnumerable<Warehouse> Warehouses;

        public WarehouseStorageService(IEnumerable<Warehouse> warehouses)
        {
            this.Warehouses = warehouses ?? throw new ArgumentNullException(nameof(warehouses));
        }

        /// <summary>
        /// Finds an available Warehouse that can accommodate an item of given dimensions (requiredHeight, requiredWidth, requiredLength)
        /// for the entire date range (startDate to endDate).
        /// A Warehouse is considered available if, for every day in the date range, 
        /// sum of currently occupied volume + required item volume <= Warehouse total volume.
        /// Returns the ID of the found Warehouse or -1 if none is found.
        /// </summary>
        public int FindAvailableWarehouse(DateTime startDate, DateTime endDate, double requiredHeight, double requiredWidth, double requiredLength)
        {
            if (!Warehouses.Any())
                throw new InvalidOperationException("no warehouses available");

            if (requiredHeight <= 0 || requiredWidth <= 0 || requiredLength <= 0)
                throw new ArgumentException("the 3d model has invalid dimensions (zero or negative)");

            if (startDate > endDate)
                throw new ArgumentException("start date cannot be later than end date");

            if (startDate <= DateTime.Today)
                throw new ArgumentException("start date cannot be later than end date");

            double requiredVolume = requiredHeight * requiredWidth * requiredLength;

            foreach (var warehouse in Warehouses)
            {
                double warehouseVolume = warehouse.GetWarehouseVolume();
                bool canAccommodate = true;

                for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
                {
                    double occupiedVolume = warehouse.GetVolumeOccupiedOnDay(day);
                    if (occupiedVolume + requiredVolume > warehouseVolume)
                    {
                        canAccommodate = false;
                        break;
                    }
                }

                if (canAccommodate)
                    return warehouse.Id;
            }

            return -1;
        }

        /// <summary>
        /// Retrieves a list of dates where the total volume of all Warehouses is fully utilized or exceeded.
        /// A date is considered fully utilized if the sum of all active items' volumes across all Warehouses
        /// equals or exceeds the sum of all Warehouses' volumes.
        /// </summary>
        public List<DateTime> GetFullyUtilizedDates(DateTime startDate, DateTime endDate)
        {
            if (!Warehouses.Any())
                throw new InvalidOperationException("no warehouses available");

            if (startDate > endDate)
                throw new ArgumentException("start date cannot be later than end date");

            var fullyUtilizedDates = new List<DateTime>();
            double totalCapacity = Warehouses.Sum(w => w.GetWarehouseVolume());

            for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
            {
                double totalVolumeForDay = Warehouses.Sum(w => w.GetVolumeOccupiedOnDay(day));
                if (totalVolumeForDay >= totalCapacity)
                {
                    fullyUtilizedDates.Add(day);
                }
            }

            return fullyUtilizedDates;
        }

        /// <summary>
        /// Calculates the total available volume across all Warehouses for the specified date range.
        /// Returns a dictionary keyed by date, indicating how much volume is free on that date.
        /// If a day is overbooked (total item volume > total capacity), it will show a negative value.
        /// </summary>
        /// 
        public Dictionary<DateTime, double> CalculateAvailableCapacity(DateTime startDate, DateTime endDate)
        {
            if (!Warehouses.Any())
                throw new InvalidOperationException("no warehouses available");

            if (startDate > endDate)
                throw new ArgumentException("start date cannot be later than end date");

            double totalCapacity = Warehouses.Sum(w => w.GetWarehouseVolume());
            var capacityMap = new Dictionary<DateTime, double>();

            for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
            {
                double totalVolumeForDay = Warehouses.Sum(w => w.GetVolumeOccupiedOnDay(day));
                capacityMap[day] = totalCapacity - totalVolumeForDay;
            }

            return capacityMap;
        }

        /// <summary>
        /// Gets the least used Warehouse within the given date range, based on the total occupied volume-days.
        /// For example:
        ///   If a Warehouse has 100 units of volume used each day for 5 days => 500 volume-days total.
        /// This method returns the Id of the Warehouse with the least usage (lowest sum of volume-days).
        /// If there's a tie, returns one arbitrarily.
        /// If no items exist in the given range, returns -1.
        /// </summary>
        public int GetLeastUsedWarehouse(DateTime startDate, DateTime endDate)
        {
            if (!Warehouses.Any())
                throw new InvalidOperationException("no warehouses available");

            if (startDate > endDate)
                throw new ArgumentException("start date cannot be later than end date");

            var usageMap = Warehouses.ToDictionary(w => w.Id, w => 0.0);

            foreach (var warehouse in Warehouses)
            {
                double totalVolumeDays = 0.0;

                for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
                {
                    totalVolumeDays += warehouse.GetVolumeOccupiedOnDay(day);
                }

                usageMap[warehouse.Id] = totalVolumeDays;
            }

            if (usageMap.Values.All(v => v == 0.0))
                return -1;

            int leastUsedWarehouseId = usageMap.OrderBy(kvp => kvp.Value).First().Key;
            return leastUsedWarehouseId;
        }
    }
}
