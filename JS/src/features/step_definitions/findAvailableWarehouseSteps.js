const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');


Given('today is {string}', function (today) {
  this.today = new Date(today);
});

Given('the warehouse usage is empty on all days', function () {
  const { warehouses } = this.service;
  warehouses.forEach(wh => {
    wh.items = [];
  });
});


When('I call FindAvailableWarehouse from {string} to {string} with dimensions:', function (startDate, endDate, dataTable) {
  const row = dataTable.hashes()[0];
  const height = parseFloat(row.height);
  const width = parseFloat(row.width);
  const length = parseFloat(row.length);

  try {
    this.result = this.service.findAvailableWarehouse(
      new Date(startDate),
      new Date(endDate),
      height,
      width,
      length
    );
  } catch (err) {
    this.error = err.message;
  }
});

When('I call FindAvailableWarehouse from {string} to {string} with dimensions {string}', function (startDate, endDate, dims) {
  const [height, width, length] = dims.split(',').map(Number);
  try {
    this.result = this.service.findAvailableWarehouse(
      new Date(startDate),
      new Date(endDate),
      height,
      width,
      length
    );
  } catch (err) {
    this.error = err.message;
  }
});

// "Then I should receive warehouse ID N"
Then('I should receive warehouse ID {int}', function (expectedId) {
  assert.strictEqual(this.result, expectedId);
});
