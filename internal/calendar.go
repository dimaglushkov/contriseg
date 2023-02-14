package internal

// Calendar represents contributions calendar by precalculated day quartiles
type Calendar [][]int8

// CalendarColorSchema is a default GitHub color schema for contributions charts
var CalendarColorSchema = [5]string{"#ebedf0", "#9be9a8", "#40c463", "#30a14e", "#216e39"}
