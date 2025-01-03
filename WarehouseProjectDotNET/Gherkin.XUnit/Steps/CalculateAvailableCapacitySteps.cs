using WarehouseProject.Models;
using WarehouseProject.Repository;
using Xunit.Gherkin.Quick;
using System;
using System.Collections.Generic;
using Xunit;
using Gherkin.Ast;

[FeatureFile("./Features/CalculateAvailableCapacity.feature")]
public sealed class CalculateAvailableCapacity : Xunit.Gherkin.Quick.Feature
{
    private readonly List<Warehouse> _warehouses = new();
    private readonly WarehouseRepository _repository;
    private Dictionary<DateTime, float> _result;
    private Exception _exception;

    public CalculateAvailableCapacity()
    {
        _repository = new WarehouseRepository(_warehouses);
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
                MaxCapacity = new ThreeDRoom(),
                Items = new List<Item>()
            });
        }
    }

    [Given(@"I have (\d+) warehouse with total volume (\d+\.\d+)")]
    public void Given_warehouse_with_volume(int count, float volume)
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

    [And(@"warehouse usage is:")]
    public void Given_warehouse_usage(DataTable dataTable)
    {
        foreach (var row in dataTable.Rows.Skip(1))
        {
            var date = DateTime.Parse(row.Cells.ElementAt(0).Value);
            var usage = float.Parse(row.Cells.ElementAt(1).Value);

            foreach (var warehouse in _warehouses)
            {
                warehouse.Items.Add(new Item
                {
                    ItemHeight = usage,
                    ItemWidth = 1,
                    ItemLength = 1,
                    StartDate = date,
                    EndDate = date,
                    IsActive = true
                });
            }
        }
    }

    [When(@"I call CalculateAvailableCapacity from ""(.*)"" to ""(.*)""")]
    public void When_calculate_capacity(string start, string end)
    {
        try
        {
            _result = _repository.CalculateAvailableCapacity(
                DateTime.Parse(start),
                DateTime.Parse(end)
            );
        }
        catch (Exception ex)
        {
            _exception = ex;
        }
    }

    [Then(@"the available capacities should be:")]
    public void Then_verify_capacities(DataTable dataTable)
    {
        foreach (var row in dataTable.Rows.Skip(1))
        {
            var date = DateTime.Parse(row.Cells.ElementAt(0).Value);
            var expectedCapacity = float.Parse(row.Cells.ElementAt(1).Value);
            Assert.True(_result.ContainsKey(date));
            Assert.Equal(expectedCapacity, _result[date]);
        }
    }

    [Then(@"an error should be returned with message ""(.*)""")]
    public void Then_verify_error(string message)
    {
        Assert.NotNull(_exception);
        Assert.Equal(message, _exception.Message.ToLower());
    }
}
