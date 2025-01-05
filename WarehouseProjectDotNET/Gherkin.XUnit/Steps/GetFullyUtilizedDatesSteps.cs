using WarehouseProject.Models;
using Xunit.Gherkin.Quick;
using System;
using System.Collections.Generic;
using Xunit;
using Gherkin.Ast;
using System.Globalization;
using System.Linq;
using WarehouseProject.Services;

[FeatureFile("./Features/GetFullyUtilizedDates.feature")]
public sealed class GetFullyUtilizedDatesTests : Xunit.Gherkin.Quick.Feature
{
    private readonly List<Warehouse> _warehouses = new();
    private List<DateTime> _resultDates;
    private Exception _exception;
    private readonly WarehouseStorageService _service;

    public GetFullyUtilizedDatesTests()
    {
        _service = new WarehouseStorageService(_warehouses);
    }

    [Given(@"I have (\d+) warehouses?")]
    public void Given_warehouse_count(int count)
    {
        _warehouses.Clear();
        _warehouses.AddRange(new Warehouse[count].Select((_, i) => new Warehouse
        {
            Id = i + 1,
            MaxCapacity = new ThreeDRoom { Height = 100, Width = 1, Length = 1 },
            Items = new List<Item>()
        }));
    }


    [Given(@"I have (\d+) warehouse with total volume (\d+\.\d+)")]
    public void Given_warehouse_with_volume(int count, float volume)
    {
        _warehouses.Clear();
        _warehouses.AddRange(new Warehouse[count].Select((_, i) => new Warehouse
        {
            Id = i + 1,
            MaxCapacity = new ThreeDRoom { Height = volume, Width = 1, Length = 1 },
            Items = new List<Item>()
        }));
    }


    [And(@"warehouse usage on ""(.*)"" is (\d+\.\d+)")]
    public void And_warehouse_usage_on_date(string date, float volume)
    {
        var targetDate = DateTime.Parse(date);
        var warehouse = _warehouses.First();

        warehouse.Items.Add(new Item
        {
            ItemId = warehouse.Items.Count + 1,
            ItemName = $"TestItem_{warehouse.Items.Count + 1}",
            ItemHeight = volume,
            ItemWidth = 1,
            ItemLength = 1,
            StartDate = targetDate,
            EndDate = targetDate,
            IsActive = true
        });
    }

    [When(@"I call GetFullyUtilizedDates from ""(.*)"" to ""(.*)""")]
    public void When_call_get_fully_utilized_dates(string start, string end)
    {
        try
        {
            _resultDates = _service.GetFullyUtilizedDates(
                DateTime.Parse(start),
                DateTime.Parse(end)
            ).ToList();
        }
        catch (Exception ex)
        {
            _exception = ex;
        }
    }

    [Then(@"the fully utilized dates should be:")]
    public void Then_verify_fully_utilized_dates(DataTable dataTable)
    {
        Assert.NotNull(_resultDates);

        var expectedDates = dataTable.Rows
            .Skip(1)
            .Select(row => DateTime.Parse(row.Cells.First().Value))
            .ToList();

        Assert.Equal(expectedDates.Count, _resultDates.Count);

        Assert.True(expectedDates.All(ed => _resultDates.Any(rd => rd.Date == ed.Date)), "The dates do not match.");
    }


    [Then(@"an error should be returned with message ""(.*)""")]
    public void Then_verify_error(string message)
    {
        Assert.NotNull(_exception);
        Assert.Equal(message, _exception.Message.ToLower());
    }
}