using WarehouseProject.Models;
using WarehouseProject.Repository;
using Xunit.Gherkin.Quick;
using System;
using System.Collections.Generic;
using Xunit;
using Gherkin.Ast;
using System.Globalization;

[FeatureFile("./Features/FindAvailableWarehouse.feature")]
public sealed class FindAvailableWarehouse : Xunit.Gherkin.Quick.Feature
{
    private readonly List<Warehouse> _warehouses = new();
    private readonly WarehouseRepository _repository;
    private int? _resultWarehouseId;
    private Exception _exception;
    private DateTime _today;

    public FindAvailableWarehouse()
    {
        _repository = new WarehouseRepository(_warehouses);
    }

    [Given(@"today is ""(.*)""")]
    public void Given_today_is(string date)
    {
        _today = DateTime.Parse(date);
    }

    [And(@"I have (\d+) warehouses?")]
    public void And_warehouse_count(int count)
    {
        _warehouses.Clear();
        for (int i = 0; i < count; i++)
        {
            _warehouses.Add(new Warehouse
            {
                Id = i + 1,
                MaxCapacity = new ThreeDRoom(),
                Items = new List<Item>()
            });
        }
    }

    [And(@"I have (\d+) warehouse with total volume (\d+\.\d+)")]
    public void And_warehouse_with_volume(int count, float volume)
    {
        _warehouses.Clear();
        for (int i = 0; i < count; i++)
        {
            _warehouses.Add(new Warehouse
            {
                Id = i + 1,
                MaxCapacity = new ThreeDRoom { Height = volume, Width = 1, Length = 1 },
                Items = new List<Item>()
            });
        }
    }

    [And(@"the warehouse usage is empty on all days")]
    public void And_warehouse_usage_is_empty()
    {
        foreach (var warehouse in _warehouses)
        {
            warehouse.Items.Clear();
        }
    }

    [When(@"I call FindAvailableWarehouse from ""(.*)"" to ""(.*)"" with dimensions:")]
    public void When_call_find_available_warehouse(string start, string end, DataTable dataTable)
    {
        var dimensions = new ThreeDRoom
        {
            Height = float.Parse(dataTable.Rows.ElementAt(1).Cells.ElementAt(0).Value, CultureInfo.InvariantCulture),
            Width = float.Parse(dataTable.Rows.ElementAt(1).Cells.ElementAt(1).Value, CultureInfo.InvariantCulture),
            Length = float.Parse(dataTable.Rows.ElementAt(1).Cells.ElementAt(2).Value, CultureInfo.InvariantCulture)
        };

        try
        {
            _resultWarehouseId = _repository.FindAvailableWarehouse(
                DateTime.Parse(start),
                DateTime.Parse(end),
                dimensions
            );
        }
        catch (Exception ex)
        {
            _exception = ex;
        }
    }

    [Then(@"I should receive warehouse ID (\d+)")]
    public void Then_should_receive_warehouse_id(int expectedWarehouseId)
    {
        Assert.NotNull(_resultWarehouseId);
        Assert.Equal(expectedWarehouseId, _resultWarehouseId);
    }

    [Then(@"an error should be returned with message ""(.*)""")]
    public void Then_verify_error(string message)
    {
        Assert.NotNull(_exception);
        Assert.Equal(message, _exception.Message.ToLower());
    }
}
