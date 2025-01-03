using System;
using System.Collections.Generic;
using System.Linq;
using WarehouseProject.Models;

namespace WarehouseProject.Repository
{
    public interface IWarehouseRepository
    {
        Dictionary<DateTime, float> CalculateAvailableCapacity(DateTime startDate, DateTime endDate);
        int FindAvailableWarehouse(DateTime startDate, DateTime endDate, ThreeDRoom dimensions);
    }

    public class WarehouseRepository : IWarehouseRepository
    {
        private readonly List<Warehouse> _warehouses;

        public WarehouseRepository(List<Warehouse> warehouses)
        {
            _warehouses = warehouses;
        }

        public Dictionary<DateTime, float> CalculateAvailableCapacity(DateTime startDate, DateTime endDate)
        {
            ValidateDates(startDate, endDate);
            ValidateWarehouses();

            var results = new Dictionary<DateTime, float>();
            var currentDate = startDate;

            while (currentDate <= endDate)
            {
                float totalCapacity = 0;
                float usedCapacity = 0;

                foreach (var warehouse in _warehouses)
                {
                    totalCapacity += CalculateTotalCapacity(warehouse.MaxCapacity);
                    usedCapacity += CalculateUsedCapacity(warehouse.Items, currentDate);
                }

                results[currentDate] = totalCapacity - usedCapacity;
                currentDate = currentDate.AddDays(1);
            }

            return results;
        }

        public int FindAvailableWarehouse(DateTime startDate, DateTime endDate, ThreeDRoom dimensions)
        {
            ValidateDates(startDate, endDate);
            ValidateDimensions(dimensions);
            ValidateWarehouses();

            var requiredVolume = CalculateTotalCapacity(dimensions);

            foreach (var warehouse in _warehouses)
            {
                var availableCapacity = CalculateAvailableCapacityForWarehouse(warehouse, startDate, endDate);
                if (availableCapacity >= requiredVolume)
                {
                    return warehouse.Id;
                }
            }

            throw new InvalidOperationException("Required volume cannot be accommodated within the specified dates");
        }

        private void ValidateDates(DateTime startDate, DateTime endDate)
        {
            if (startDate > endDate)
            {
                throw new ArgumentException("start date cannot be later than end date");
            }
        }

        private void ValidateDimensions(ThreeDRoom dimensions)
        {
            if (dimensions.Height <= 0 || dimensions.Width <= 0 || dimensions.Length <= 0)
            {
                throw new ArgumentException("The 3d model has invalid dimensions (zero or negative)");
            }
        }

        private void ValidateWarehouses()
        {
            if (_warehouses == null || !_warehouses.Any())
            {
                throw new InvalidOperationException("No warehouses available");
            }
        }

        private float CalculateTotalCapacity(ThreeDRoom room)
        {
            return room.Height * room.Width * room.Length;
        }

        private float CalculateUsedCapacity(List<Item> items, DateTime date)
        {
            return items
                .Where(item => item.IsActive &&
                       item.StartDate <= date &&
                       item.EndDate >= date)
                .Sum(item => item.ItemHeight * item.ItemWidth * item.ItemLength);
        }

        private float CalculateAvailableCapacityForWarehouse(Warehouse warehouse, DateTime startDate, DateTime endDate)
        {
            var totalCapacity = CalculateTotalCapacity(warehouse.MaxCapacity);
            var usedCapacity = 0f;

            var currentDate = startDate;
            while (currentDate <= endDate)
            {
                usedCapacity += CalculateUsedCapacity(warehouse.Items, currentDate);
                currentDate = currentDate.AddDays(1);
            }

            return totalCapacity - usedCapacity;
        }
    }
}
