const { Given, When, Then } = require('@cucumber/cucumber');
const assert = require('assert');

Given('warehouse usage on {string} is {float}', function (dateStr, usage) {
  const usageDate = new Date(dateStr);
  usageDate.setHours(0, 0, 0, 0);

  const wh = this.service.warehouses[0];
  wh.items.push({
    startDate: usageDate,
    endDate: usageDate,
    volume: usage,
  });
});

When('I call GetFullyUtilizedDates from {string} to {string}', function (startDate, endDate) {
  try {
    this.result = this.service.getFullyUtilizedDates(new Date(startDate), new Date(endDate));
  } catch (err) {
    this.error = err.message;
  }
});

Then('the fully utilized dates should be:', function (dataTable) {
  const expectedDates = dataTable
    .raw()
    .slice(1)
    .map(row => row[0]);

  const actualDates = (this.result || []).map(d => d.toISOString().split('T')[0]);
  assert.deepStrictEqual(actualDates, expectedDates);
});
