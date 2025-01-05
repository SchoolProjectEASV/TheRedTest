const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');

Given('I have warehouses with usage:', function (dataTable) {
  const rows = dataTable.hashes();
  const warehouses = rows.map(row => ({
    id: parseInt(row.id, 10),
    maxCapacity: parseFloat(row.volume),
    items: [
      {
        startDate: new Date('2025-01-10'),
        endDate: new Date('2025-01-10'),
        volume: parseFloat(row.usage),
      }
    ]
  }));

  this.service = new (require('../../services/warehouseStorageService'))(warehouses);
});

When('I call GetLeastUsedWarehouse from {string} to {string}', function (startDate, endDate) {
  try {
    this.result = this.service.getLeastUsedWarehouse(new Date(startDate), new Date(endDate));
  } catch (err) {
    this.error = err.message;
  }
});

Then('the least used warehouse should be {int}', function (expectedId) {
  assert.strictEqual(this.result, expectedId);
});

Then('no error is returned', function () {
  assert.strictEqual(this.error, undefined);
});
