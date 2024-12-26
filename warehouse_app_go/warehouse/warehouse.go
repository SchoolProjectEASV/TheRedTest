package warehouse

import "time"

func (w Warehouse) GetWarehouseVolume() float64 {
	return w.MaxCapacity.Height * w.MaxCapacity.Width * w.MaxCapacity.Length
}

func (i Item) GetItemVolume() float64 {
	return i.ItemHeight * i.ItemWidth * i.ItemLength
}

func (w Warehouse) GetVolumeOccupiedOnDay(day time.Time) float64 {
	volume := 0.0
	for _, item := range w.Items {
		if item.IsActive && !day.Before(item.StartDate) && !day.After(item.EndDate) {
			volume += item.GetItemVolume()
		}
	}
	return volume
}
