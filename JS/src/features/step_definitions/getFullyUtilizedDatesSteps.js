const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');
const WarehouseStorageService = require("../../services/warehouseStorageService");

let service;
let result;
let error;

Given('I have {int} warehouse with the given volume {float}', function (count, volume) {
    const warehouses = Array.from({ length: count }, (_, i) => ({
        id: i + 1,
        maxCapacity: volume,
        items: []
    }));

    console.log("Warehouses created:", warehouses);  // Debugging log

    service = new WarehouseStorageService(warehouses);
    console.log("Service initialized:", service);  // Debugging log
});



When('I call GetFullyUtilizedDates from {string} to {string}', function (startDate, endDate) {
    console.log("In When step");
    try {
        result = service.getFullyUtilizedDates(new Date(startDate), new Date(endDate));
        console.log("Result from getFullyUtilizedDates:", result); // Log the result
    } catch (err) {
        error = err.message;
    }
});


Then('the fully utilized dates should be {string}', function (expectedDates) {
    const expected = JSON.parse(expectedDates); // Parse the expected date string
    console.log("Expected dates:", expected); // Log the expected dates
    console.log("Result:", result); // Log the actual result

    assert.deepStrictEqual(result.map(date => date.toISOString().split('T')[0]), expected);
});

