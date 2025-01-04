class Warehouse {
    constructor(id, maxCapacity, items = []) {
        this.id = id;
        this.maxCapacity = maxCapacity;
        this.items = items;
    }
}

module.exports = Warehouse;
