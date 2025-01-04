const { Given, Then, setWorldConstructor } = require('@cucumber/cucumber');
const assert = require('assert');
const WarehouseStorageService = require('../../services/warehouseStorageService');

class CustomWorld {
  constructor() {
    this.service = undefined;
    this.result = undefined;
    this.error = undefined;
  }
}

setWorldConstructor(CustomWorld);


Then('an error should be returned with message {string}', function (expectedMessage) {
  assert.strictEqual(this.error, expectedMessage);
});

Given('I have {int} warehouse with total volume {float}', function (count, volume) {
  const warehouses = Array.from({ length: count }, (_, i) => ({
    id: i + 1,
    maxCapacity: volume,
    items: []
  }));
  this.service = new WarehouseStorageService(warehouses);
  this.result = null;
  this.error = null;
});


Given('I have {int} warehouses', function (count) {
  const warehouses = Array.from({ length: count }, (_, i) => ({
    id: i + 1,
    maxCapacity: 100.0, 
    items: []
  }));
  this.service = new WarehouseStorageService(warehouses);
  this.result = null;
  this.error = null;
});


Given('I have {int} default warehouse(s)', function (count) {
  const DEFAULT_VOLUME = 100.0;
  const warehouses = Array.from({ length: count }, (_, i) => ({
    id: i + 1,
    maxCapacity: DEFAULT_VOLUME,
    items: []
  }));
  this.service = new WarehouseStorageService(warehouses);
  this.result = null;
  this.error = null;
});
