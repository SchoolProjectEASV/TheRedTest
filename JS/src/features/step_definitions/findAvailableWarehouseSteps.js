const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');

let service;
let result;
let error;

Given('today is {string}', function (today) {
    this.today = new Date(today);
});

Given('I have {int} warehouse with total volume {float}', function (count, volume) {
    const warehouses = Array.from({ length: count }, (_, i) => ({
        id: i + 1,
        totalVolume: volume,
        items: [],
    }));
    service = new WarehouseStorageService(warehouses);
});

When('I call FindAvailableWarehouse from {string} to {string} with dimensions {string}', function (startDate, endDate, dimensions) {
    const [height, width, length] = dimensions.split(',').map(Number);
    try {
        result = service.findAvailableWarehouse(new Date(startDate), new Date(endDate), height, width, length);
    } catch (err) {
        error = err.message;
    }
});

Then('I should receive warehouse ID {int}', function (expectedId) {
    assert.strictEqual(result, expectedId);
});
