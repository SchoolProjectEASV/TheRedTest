using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using WarehouseProject.Models;

namespace WarehouseProject.Services
{
    public class WarehouseStorageService
    {
        private readonly IEnumerable<Warehouse> Warehouses;

        public WarehouseStorageService(IEnumerable<Warehouse> Warehouses)
        {
            this.Warehouses = Warehouses;
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
            if (startDate <= DateTime.Today || startDate > endDate)
                throw new ArgumentException("The start date cannot be in the past or later than the end date.");

            double requiredVolume = requiredHeight * requiredWidth * requiredLength;

            foreach (var Warehouse in Warehouses)
            {
                double WarehouseVolume = Warehouse.GetWarehouseVolume();
                bool canAccommodate = true;

                for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
                {
                    double occupiedVolume = Warehouse.GetVolumeOccupiedOnDay(day);
                    if (occupiedVolume + requiredVolume > WarehouseVolume)
                    {
                        canAccommodate = false;
                        break;
                    }
                }

                if (canAccommodate)
                    return Warehouse.Id;
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
            if (startDate > endDate)
                throw new ArgumentException("The start date cannot be later than the end date.");

            List<DateTime> fullyUtilizedDates = new List<DateTime>();

            double totalCapacity = Warehouses.Sum(w => w.GetWarehouseVolume());

            // Get all items from all Warehouses (assume we just flatten)
            var allItems = Warehouses.SelectMany(w => w.Items).Where(i => i.IsActive).ToList();

            for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
            {
                // Sum volume of all items that are active on this day
                double totalVolumeForDay = allItems
                    .Where(i => day >= i.StartDate && day <= i.EndDate)
                    .Sum(i => i.GetItemVolume());

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
        public Dictionary<DateTime, double> CalculateAvailableCapacity(DateTime startDate, DateTime endDate)
        {
            if (startDate > endDate)
                throw new ArgumentException("The start date cannot be later than the end date.");

            double totalCapacity = Warehouses.Sum(w => w.GetWarehouseVolume());
            var allItems = Warehouses.SelectMany(w => w.Items).Where(i => i.IsActive).ToList();
            Dictionary<DateTime, double> capacityMap = new Dictionary<DateTime, double>();

            for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
            {
                double totalVolumeForDay = allItems
                    .Where(i => day >= i.StartDate && day <= i.EndDate)
                    .Sum(i => i.GetItemVolume());

                double available = totalCapacity - totalVolumeForDay;
                capacityMap[day] = available;
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
            if (startDate > endDate)
                throw new ArgumentException("The start date cannot be later than the end date.");

            if (!Warehouses.Any())
                return -1;

            var usageMap = Warehouses.ToDictionary(w => w.Id, w => 0.0);

            // For each Warehouse, sum the volume of items day by day
            foreach (var Warehouse in Warehouses)
            {
                double totalVolumeDays = 0.0;

                for (DateTime day = startDate; day <= endDate; day = day.AddDays(1))
                {
                    double occupiedVolume = Warehouse.GetVolumeOccupiedOnDay(day);
                    totalVolumeDays += occupiedVolume;
                }

                usageMap[Warehouse.Id] = totalVolumeDays;
            }

            // If no usage at all, return -1
            if (usageMap.Values.All(v => v == 0.0))
                return -1;

            // Return the Warehouse with the smallest usage value
            int leastUsedWarehouseId = usageMap.OrderBy(kvp => kvp.Value).First().Key;
            return leastUsedWarehouseId;
        }
    }
}
