package presenter



type UserPerformance struct {
	NbEvents   int			  `json:"nb_events,omitempty"`
	Stats     []UserStats `json:"stats,omitempty"`
}
