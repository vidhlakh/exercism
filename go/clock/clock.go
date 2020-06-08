//Package clock works with time with no date
package clock

import "fmt"

//Clock is the struct type with hours and minutes
type Clock struct {
	hour, min int
}

//New constructor to initialize the hours and minutes
func New(h, m int) Clock {

	c := Clock{h, m}
	return c.Adjust()

}

//String displays the hour and min in string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

//Add adds the minutes to the current clock
func (c Clock) Add(min int) Clock {
	c.min += min
	return c.Adjust()
}

//Subtract subtracts the minutes to the current clock
func (c Clock) Subtract(min int) Clock {
	c.min -= min
	return c.Adjust()
}

//Adjust time in proper way Hour if more than 12 , rollover, minutes if more than 59 ,roll over it by also increasing hour
func (c Clock) Adjust() Clock {
	// if c.min > 59 {
	// 	c.min = c.min - 60
	// 	c.hour = c.hour + 1
	// }
	// if c.hour > 12 {
	// 	c.hour = c.hour - 12
	// }
	// return c
	c.hour += (c.min / 60)
	c.min = (c.min % 60)
	for c.min < 0 {
		c.min += 60
		c.hour--
	}
	c.hour %= 24
	for c.hour < 0 {
		c.hour += 24
	}
	return c
}
