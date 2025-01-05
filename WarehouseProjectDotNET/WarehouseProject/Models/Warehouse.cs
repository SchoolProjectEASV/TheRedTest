using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace WarehouseProject.Models
{
    public class Warehouse
    {
        public int Id { get; set; }
        public ThreeDRoom MaxCapacity { get; set; }
        public List<Item> Items { get; set; }
    }
}
