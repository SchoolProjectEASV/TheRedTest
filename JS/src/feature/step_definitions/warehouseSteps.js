const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');
const WarehouseStorageService = require('../../services/WarehouseStorageService');
const { warehouseData } = require('../mockData'); // Import mock data

let service;
let result;

Given('the following warehouses:', function () {
    service = new WarehouseStorageService(warehouseData);
});

Given('the following items:', function (dataTable) {
    dataTable.hashes().forEach((row) => {
        const item = new Item(
            parseInt(row.ItemId),
            '',
            parseFloat(row.Height),
            parseFloat(row.Width),
            parseFloat(row.Length),
            new Date(row.StartDate),
            new Date(row.EndDate)
        );

        // Add item to the correct warehouse
        const warehouse = warehouseData.find(w => w.id === parseInt(row.WarehouseId));
        if (warehouse) {
            warehouse.items.push(item);
        }
    });
});

When('I try to store an item with dimensions {float}x{float}x{float} from {string} to {string}', function (height, width, length, startDate, endDate) {
    result = service.findAvailableWarehouse(new Date(startDate), new Date(endDate), height, width, length);
});

Then('I should be told that Warehouse {int} can accommodate the item', function (expectedId) {
    assert.strictEqual(result, expectedId);
});