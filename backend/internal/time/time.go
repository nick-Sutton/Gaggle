// The time package defines the Time enum,
// the Day enum, and the TimeSlot struct
// author: Nicholas Sutton
package time

// Enum representing time slots.
// Stored as integers starting at
// Slot_4_5pm = 0 and counting up.
type Time int

const (
	Slot_4_5pm Time = iota // 0
	Slot_5_6pm             // 1
	Slot_6_7pm             // 2
	slot_7_8pm             // 3
	Slot_8_9pm             // 4
)

// Enum representing the days of the week.
// Stored as integers starting at Monday = 0
// and counting up till Friday.
type Day int

const (
	Monday    Day = iota // 0
	Tuesday              // 1
	Wednesday            // 2
	Thursday             // 3
	Friday               // 4
)

// A DailyTimeSlore stores the day, the time, and the number
// of players available to play during that time slot.
type TimeSlot struct {
	Day              Day
	Time             Time
	AvailablePlayers int
}

// The NewTimeSlot function creates a TimeSlot struct
// param: the day of the time slot
// param: the time of the timeslot
// return: a pointer to the new TimeSlot
func NewTimeSlot(day Day, time Time) *TimeSlot {
	t := &TimeSlot{}
	t.Day = day
	t.Time = time
	t.AvailablePlayers = 0

	return t
}
