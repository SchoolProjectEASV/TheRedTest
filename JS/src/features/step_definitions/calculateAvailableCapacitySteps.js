const { When, Then } = require('@cucumber/cucumber');
const assert = require('assert');

Then('warehouse usage is:', function (dataTable) {
  const usageData = dataTable.hashes();
  const warehouse = this.service.warehouses[0]; 

  usageData.forEach(row => {
    const usageDate = new Date(row.date);
    usageDate.setHours(0, 0, 0, 0);

    warehouse.items.push({
      startDate: usageDate,
      endDate: usageDate,
      volume: parseFloat(row.usage),
    });
  });
});

When('I call CalculateAvailableCapacity from {string} to {string}', function (startStr, endStr) {
  try {
    const start = new Date(startStr);
    const end = new Date(endStr);
    this.result = this.service.calculateAvailableCapacity(start, end);
  } catch (err) {
    this.error = err.message;
  }
});

Then('the available capacities should be:', function (dataTable) {
  const expected = dataTable.hashes().map(r => ({
    date: r.date,
    capacity: parseFloat(r.capacity),
  }));

  const capacityMap = this.result || {};
  const actualArray = Object.entries(capacityMap).map(([date, capacity]) => ({
    date,
    capacity,
  }));

  assert.deepStrictEqual(actualArray, expected);
});
