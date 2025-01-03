const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');
const WarehouseStorageService = require('../../services/warehouseStorageService');
const Item = require('../../models/item');

let service;
let result;
let error;

Given(/I have (\d+) warehouse with total volume (\d+\.\d+) for available capacity/, function (count, volume) {
    console.log(`Creating ${count} warehouse with total volume ${volume} for available capacity`);
    const warehouses = Array.from({ length: count }, (_, i) => ({
        id: i + 1,
        maxCapacity: volume,
        items: [],
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
    assert.deepStrictEqual(result, expected);
});

Then('an error should be returned with message {string}', function (expectedMessage) {
    assert.strictEqual(error, expectedMessage);
});
