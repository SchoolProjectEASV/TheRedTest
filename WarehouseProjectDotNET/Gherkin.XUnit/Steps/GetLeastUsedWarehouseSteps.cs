using WarehouseProject.Models;
using Xunit.Gherkin.Quick;
using System;
using System.Collections.Generic;
using Xunit;
using Gherkin.Ast;
using System.Globalization;
using System.Linq;
using WarehouseProject.Services;

[FeatureFile("./Features/GetLeastUsedWarehouse.feature")]
public sealed class GetLeastUsedWarehouseTests : Xunit.Gherkin.Quick.Feature
{
    private readonly List<Warehouse> _warehouses = new();
    private readonly WarehouseStorageService _service;
    private int _resultWarehouseId;
    private Exception _exception;

    public GetLeastUsedWarehouseTests()
    {
        _service = new WarehouseStorageService(_warehouses);
    }

    [Given(@"I have (\d+) warehouses?")]
    public void Given_warehouse_count(int count)
    {
        _warehouses.Clear();
        for (int i = 0; i < count; i++)
        {
            _warehouses.Add(new Warehouse
            {
                Id = i + 1,
                MaxCapacity = new ThreeDRoom { Height = 100, Width = 1, Length = 1 },
                Items = new List<Item>()
            });
        }
    }

    [Given(@"I have warehouses with usage:")]
    public void Given_warehouses_with_usage(DataTable dataTable)
    {
        _warehouses.Clear();

        foreach (var row in dataTable.Rows.Skip(1))
        {
            var id = int.Parse(row.Cells.ElementAt(0).Value);
            var volume = float.Parse(row.Cells.ElementAt(1).Value, CultureInfo.InvariantCulture);
            var usage = float.Parse(row.Cells.ElementAt(2).Value, CultureInfo.InvariantCulture);

            var warehouse = new Warehouse
            {
                Id = id,
                MaxCapacity = new ThreeDRoom { Height = volume, Width = 1, Length = 1 },
                Items = new List<Item>()
            };

            if (usage > 0)
            {
                warehouse.Items.Add(new Item
                {
                    ItemId = 1,
                    ItemName = $"TestItem_{id}",
                    ItemHeight = usage,
                    ItemWidth = 1,
                    ItemLength = 1,
                    StartDate = DateTime.Parse("2025-01-10"),
                    EndDate = DateTime.Parse("2025-01-10"),
                    IsActive = true
                });
            }

            _warehouses.Add(warehouse);
        }
    }

    [When(@"I call GetLeastUsedWarehouse from ""(.*)"" to ""(.*)""")]
    public void When_call_get_least_used_warehouse(string start, string end)
    {
        try
        {
            _resultWarehouseId = _service.GetLeastUsedWarehouse(
                DateTime.Parse(start),
                DateTime.Parse(end)
            );
        }
        catch (Exception ex)
        {
            _exception = ex;
        }
    }

    [Then(@"the least used warehouse should be (\-?\d+)")]
    public void Then_verify_least_used_warehouse(int expectedId)
    {
        Assert.Equal(expectedId, _resultWarehouseId);
    }

    [Then(@"an error should be returned with message ""(.*)""")]
    public void Then_verify_error(string message)
    {
        Assert.NotNull(_exception);
        Assert.Equal(message, _exception.Message.ToLower());
    }

    [And(@"no error is returned")]
    public void And_no_error_is_returned()
    {
        Assert.Null(_exception);
    }
}