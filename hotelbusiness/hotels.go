//go:build !solution

package hotelbusiness

import (
	"maps"
	"slices"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	guestsByDay := make(map[int]int, len(guests))
	for _, guest := range guests {
		for day := guest.CheckInDate; day < guest.CheckOutDate; day++ {
			guestsByDay[day]++
		}

		// Ensure the day after checkout is present to record zero guests if needed
		guestsByDay[guest.CheckOutDate] = guestsByDay[guest.CheckOutDate]
	}

	result := make([]Load, 0, len(guestsByDay))

	for _, day := range slices.Sorted(maps.Keys(guestsByDay)) {
		count := guestsByDay[day]

		if len(result) == 0 || result[len(result)-1].GuestCount != count {
			result = append(result, Load{StartDate: day, GuestCount: count})
		}
	}

	return result
}
