using Moq;
using System;
using System.Collections.Generic;
using WarehouseProject.Models;
using WarehouseProject.Repository;
using Xunit;

namespace WarehouseProject.Tests
{
    public class WarehouseRepositoryTests
    {
        private readonly Mock<IWarehouseRepository> _mockRepo;
        private readonly List<Warehouse> _testWarehouses;

        public WarehouseRepositoryTests()
        {
            _mockRepo = new Mock<IWarehouseRepository>();
            _testWarehouses = CreateTestWarehouses();
        }

        [Fact]
        public void Test_SameDay_Returns_CorrectCapacity()
        {
            var startDate = new DateTime(2025, 1, 10);
            var endDate = new DateTime(2025, 1, 10);

            var expectedResult = new Dictionary<DateTime, float>
            {
                { startDate, 1000f }  // Example capacity
            };

            _mockRepo.Setup(x => x.CalculateAvailableCapacity(startDate, endDate))
                    .Returns(expectedResult);

            var result = _mockRepo.Object.CalculateAvailableCapacity(startDate, endDate);
            Assert.Equal(expectedResult, result);
        }

        [Fact]
        public void Test_InvalidDateRange_ThrowsException()
        {
            var startDate = new DateTime(2025, 1, 12);
            var endDate = new DateTime(2025, 1, 11);

            _mockRepo.Setup(x => x.CalculateAvailableCapacity(startDate, endDate))
                    .Throws(new ArgumentException("The start date cannot be later than the end date"));

            Assert.Throws<ArgumentException>(() =>
                _mockRepo.Object.CalculateAvailableCapacity(startDate, endDate));
        }

        private List<Warehouse> CreateTestWarehouses()
        {
            return new List<Warehouse>
            {
                new Warehouse
                {
                    Id = 1,
                    MaxCapacity = new ThreeDRoom { Height = 10, Width = 10, Length = 10 },
                    Items = new List<Item>()
                }
            };
        }
    }
}
