const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');

let service;
let result;
let error;

Given('I have {int} warehouse with total volume {float}', function (count, volume) {
    const warehouses = Array.from({ length: count }, (_, i) => ({
        id: i + 1,
        totalVolume: volume,
        items: [],
    }));
    service = new WarehouseStorageService(warehouses);
});

When('I call GetFullyUtilizedDates from {string} to {string}', function (startDate, endDate) {
    try {
        result = service.getFullyUtilizedDates(new Date(startDate), new Date(endDate));
    } catch (err) {
        error = err.message;
    }
});

Then('the fully utilized dates should be {string}', function (expectedDates) {
    const expected = JSON.parse(expectedDates);
    assert.deepStrictEqual(result.map(date => date.toISOString().split('T')[0]), expected);
});
