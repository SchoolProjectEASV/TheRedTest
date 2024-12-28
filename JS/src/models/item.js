class Item {
    constructor(itemId, itemName, itemHeight, itemWidth, itemLength, startDate, endDate, isActive = true) {
        this.itemId = itemId;
        this.itemName = itemName;
        this.itemHeight = itemHeight;
        this.itemWidth = itemWidth;
        this.itemLength = itemLength;
        this.startDate = new Date(startDate);
        this.endDate = new Date(endDate);
        this.isActive = isActive;
    }
}