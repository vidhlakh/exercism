// Package twofer s
package twofer

// ShareWith respond a message depending on the input name. If name is empty , responds with default you
func ShareWith(name string) string {
	// if name is empty. Respond with "you"
	if name == "" {
		name = "you"
	}
	// return the message along with name
	return "One for " + name + ", one for me."

}
