const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');
const WarehouseStorageService = require('../../services/warehouseStorageService');
const Item = require('../../models/item');

let service;
let result;
let error;

Given(/I have (\d+) warehouse/, function (count) {
    // Create the warehouse with a default volume if none is specified
    const warehouses = Array.from({ length: parseInt(count, 10) }, (_, i) => ({
        id: i + 1,
        maxCapacity: 100.0,  // Default volume for the warehouse
        items: [],
        getWarehouseVolume() {
            return this.maxCapacity;
        },
        getVolumeOccupiedOnDay(date) {
            return this.items
                .filter(item => item.startDate <= date && item.endDate >= date)
                .reduce((sum, item) => sum + item.volume, 0);
        },
    }));

    service = new WarehouseStorageService(warehouses);
});


Given('warehouse usage on {string} is {float}', function (date, usage) {
    console.log(`Setting warehouse usage on ${date} to ${usage}`);
    const warehouse = service.warehouses[0];
    warehouse.items.push(new Item(
        warehouse.items.length + 1,
        `Item ${warehouse.items.length + 1}`,
        1, 1, 1,
        new Date(date).setHours(0, 0, 0, 0),
        new Date(date).setHours(0, 0, 0, 0),
        usage
    ));
});

When('I call CalculateAvailableCapacity from {string} to {string}', function (startDate, endDate) {
    const start = new Date(startDate);
    const end = new Date(endDate);

    start.setHours(0, 0, 0, 0);
    end.setHours(0, 0, 0, 0);

    try {
        if (start > end) {
            throw new Error("the start date cannot be later than the end date");
        }

        result = service.calculateAvailableCapacity(start, end);
    } catch (err) {
        error = err.message;
    }
});

Then('the available capacities should be:', function (dataTable) {
    const expected = dataTable.hashes().map(row => ({
        date: row['Date'],
        capacity: parseFloat(row['Capacity']),
    }));

    const transformedResult = Object.entries(result || {}).map(([date, capacity]) => ({
        date,
        capacity,
    }));

    console.log('Expected:', expected);
    console.log('Actual:', transformedResult);

    assert.deepStrictEqual(transformedResult, expected);
});

Then('an error should be returned with message {string}', function (expectedMessage) {
    assert.strictEqual(error, expectedMessage);
});
