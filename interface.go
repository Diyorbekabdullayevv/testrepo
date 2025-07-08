package main

// type limitChecker interface {
// 	checkIfLimit() bool
// }

// type cars struct {
// 	name    string
// 	country string
// 	cost    float64
// }

// type toys struct {
// 	name    string
// 	country string
// 	cost    float64
// }

// type dishes struct {
// 	name    string
// 	country string
// 	cost    float64
// }

// type courses struct {
// 	name    string
// 	country string
// 	cost    float64
// }

// func (c *cars) checkIfLimit() bool {
// 	var limit float64 = 1000

// 	if c.cost < limit {
// 		return false
// 	}
// 	return true
// }

// func (t *toys) checkIfLimit() bool {
// 	var limit float64 = 1000

// 	if t.cost == limit {
// 		return true
// 	}
// 	return false
// }

// func (d *dishes) checkIfLimit() bool {
// 	var limit float64 = 1000

// 	if d.cost > limit {
// 		return true
// 	}
// 	return false
// }

// func (c *courses) checkIfLimit() bool {
// 	var limit float64 = 1000

// 	if c.cost != limit {
// 		return false
// 	}
// 	return true
// }

// func calculate(l limitChecker) bool {
// 	// fmt.Println("Limit checker struct:", l)
// 	// fmt.Println(path, l.checkIfLimit())
// 	return l.checkIfLimit()
// }
