package service

// ByEffort sorts routes by duration and by distance (if duration is equal)
type ByEffort []Route

func (routes ByEffort) Len() int {
	return len(routes)
}

func (routes ByEffort) Swap(i, j int) {
	routes[i], routes[j] = routes[j], routes[i]
}

func (routes ByEffort) Less(i, j int) bool {
	if routes[i].Duration != routes[j].Duration {
		return routes[i].Duration < routes[j].Duration
	}
	return routes[i].Distance < routes[j].Distance
}
