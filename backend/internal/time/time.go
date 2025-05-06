package time

/*
 * Enum representing time slots.
 * Stored as integers starting at
 * Slot_4_5pm = 0 and counting up.
 */
type Time int

const (
	Slot_4_5pm Time = iota
	Slot_5_6pm
	Slot_6_7pm
	slot_7_8pm
	Slot_8_9pm
)

/*
 * Enum representing the days of the week.
 * Stored as integers starting at Monday = 0
 * and counting up till Friday.
 */
type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
)

/*
 * A DailyTimeSlore stores the day, a map
 * of available times, and a map of prefered times
 */
type DailyTimeSlot struct {
	day            Day
	AvailableTimes map[Time]bool
	PreferredTimes map[Time]bool
}
