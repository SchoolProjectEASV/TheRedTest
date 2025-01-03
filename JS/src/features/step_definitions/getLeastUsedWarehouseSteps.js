const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');

let service;
let result;
let error;

Given('I have {int} warehouses with the following data:', function (dataTable) {
    const warehouses = dataTable.hashes().map(row => ({
        id: parseInt(row['ID']),
        totalVolume: parseFloat(row['totalVolume']),
        items: [
            {
                startDate: new Date('2025-01-10'),
                endDate: new Date('2025-01-10'),
                volume: parseFloat(row['usageOn2025-01-10']),
                isActive: true,
            },
        ],
    }));
    service = new WarehouseStorageService(warehouses);
});

When('I call GetLeastUsedWarehouse from {string} to {string}', function (startDate, endDate) {
    try {
        result = service.getLeastUsedWarehouse(new Date(startDate), new Date(endDate));
    } catch (err) {
        error = err.message;
    }
});

Then('the least used warehouse should be {string}', function (expectedId) {
    assert.strictEqual(result.toString(), expectedId);
});
