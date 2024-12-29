using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using WarehouseProject.Models;

namespace WarehouseProject.Services
{
    public static class WarehouseExtensions
    {
        public static double GetWarehouseVolume(this Warehouse warehouse)
        {
            return warehouse.MaxCapacity.Height * warehouse.MaxCapacity.Width * warehouse.MaxCapacity.Length;
        }

        public static double GetItemVolume(this Item item)
        {
            return item.ItemHeight * item.ItemWidth * item.ItemLength;
        }

        /// <summary>
        /// Returns the sum of the volumes of all items that are active and overlap the given day.
        /// </summary>
        public static double GetVolumeOccupiedOnDay(this Warehouse warehouse, DateTime day)
        {
            return warehouse.Items
                .Where(i => i.IsActive && day >= i.StartDate && day <= i.EndDate)
                .Sum(i => i.GetItemVolume());
        }
    }
}
