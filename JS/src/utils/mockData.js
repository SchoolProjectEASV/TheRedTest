const { WareHouse, ThreeDRoom, Item } = require('../models');

const warehouseData = [
    new WareHouse(1, new ThreeDRoom(10, 20, 30), [
        new Item(1, 'Item1', 5, 5, 5, new Date('2024-12-20'), new Date('2024-12-25')),
    ]),
    new WareHouse(2, new ThreeDRoom(15, 25, 35), []),
];

module.exports = { warehouseData };