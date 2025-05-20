package goxmind

type MarkerId string

// ⭐ 星星
const (
	StarRed    MarkerId = "star-red"
	StarOrange MarkerId = "star-orange"
	StarYellow MarkerId = "star-yellow"
	StarBlue   MarkerId = "star-blue"
	StarGreen  MarkerId = "star-green"
	StarPurple MarkerId = "star-purple"
)

// 🔢 优先级
const (
	Priority1 MarkerId = "priority-1"
	Priority2 MarkerId = "priority-2"
	Priority3 MarkerId = "priority-3"
	Priority4 MarkerId = "priority-4"
	Priority5 MarkerId = "priority-5"
	Priority6 MarkerId = "priority-6"
	Priority7 MarkerId = "priority-7"
	Priority8 MarkerId = "priority-8"
	Priority9 MarkerId = "priority-9"
)

// 😊 表情
const (
	SmileySmile    MarkerId = "smiley-smile"
	SmileyLaugh    MarkerId = "smiley-laugh"
	SmileyAngry    MarkerId = "smiley-angry"
	SmileyCry      MarkerId = "smiley-cry"
	SmileySurprise MarkerId = "smiley-surprise"
	SmileyBoring   MarkerId = "smiley-boring"
)

// ✅ 任务进度
const (
	Task0_8 MarkerId = "task-start"
	Task1_8 MarkerId = "task-oct"
	Task2_8 MarkerId = "task-quarter"
	Task3_8 MarkerId = "task-3oct"
	Task4_8 MarkerId = "task-half"
	Task5_8 MarkerId = "task-5oct"
	Task6_8 MarkerId = "task-3quar"
	Task7_8 MarkerId = "task-7oct"
	Task8_8 MarkerId = "task-done"
)

// 🚩 旗子
const (
	FlagRed    MarkerId = "flag-red"
	FlagOrange MarkerId = "flag-orange"
	FlagYellow MarkerId = "flag-yellow"
	FlagBlue   MarkerId = "flag-blue"
	FlagGreen  MarkerId = "flag-green"
	FlagPurple MarkerId = "flag-purple"
)

// 👤 人像
const (
	PeopleRed    MarkerId = "people-red"
	PeopleOrange MarkerId = "people-orange"
	PeopleYellow MarkerId = "people-yellow"
	PeopleBlue   MarkerId = "people-blue"
	PeopleGreen  MarkerId = "people-green"
	PeoplePurple MarkerId = "people-purple"
)

// 🔁 箭头
const (
	ArrowUp        MarkerId = "arrow-up"
	ArrowUpRight   MarkerId = "arrow-up-right"
	ArrowRight     MarkerId = "arrow-right"
	ArrowDownRight MarkerId = "arrow-down-right"
	ArrowDown      MarkerId = "arrow-down"
	ArrowDownLeft  MarkerId = "arrow-down-left"
	ArrowLeft      MarkerId = "arrow-left"
	ArrowUpLeft    MarkerId = "arrow-up-left"
	ArrowRefresh   MarkerId = "arrow-refresh"
)

// 🧩 符号
const (
	SymbolPlus         MarkerId = "symbol-plus"
	SymbolMinus        MarkerId = "symbol-minus"
	SymbolQuestion     MarkerId = "symbol-question"
	SymbolInfo         MarkerId = "symbol-info"
	SymbolAttention    MarkerId = "symbol-attention"
	SymbolWrong        MarkerId = "symbol-wrong"
	SymbolRight        MarkerId = "symbol-right"
	SymbolPause        MarkerId = "symbol-pause"
	SymbolBarChart     MarkerId = "c_symbol_bar_chart"
	SymbolFlight       MarkerId = "c_symbol_flight"
	SymbolExercise     MarkerId = "c_symbol_exercise"
	SymbolDrink        MarkerId = "c_symbol_drink"
	SymbolDislike      MarkerId = "c_symbol_dislike"
	SymbolContact      MarkerId = "c_symbol_contact"
	SymbolTrophy       MarkerId = "c_symbol_trophy"
	SymbolThermometer  MarkerId = "c_symbol_thermometer"
	SymbolTelephone    MarkerId = "c_symbol_telephone"
	SymbolShoppingCart MarkerId = "c_symbol_shopping_cart"
	SymbolPieChart     MarkerId = "c_symbol_pie_chart"
	SymbolPen          MarkerId = "c_symbol_pen"
	SymbolMusic        MarkerId = "c_symbol_music"
	SymbolMoney        MarkerId = "c_symbol_money"
	SymbolMedals       MarkerId = "c_symbol_medals"
	SymbolLineGraph    MarkerId = "c_symbol_line_graph"
	SymbolLike         MarkerId = "c_symbol_like"
	SymbolHeart        MarkerId = "c_symbol_heart"
)

// 📅 月份
const (
	MonthJan MarkerId = "month-jan"
	MonthFeb MarkerId = "month-feb"
	MonthMar MarkerId = "month-mar"
	MonthApr MarkerId = "month-apr"
	MonthMay MarkerId = "month-may"
	MonthJun MarkerId = "month-jun"
	MonthJul MarkerId = "month-jul"
	MonthAug MarkerId = "month-aug"
	MonthSep MarkerId = "month-sep"
	MonthOct MarkerId = "month-oct"
	MonthNov MarkerId = "month-nov"
	MonthDec MarkerId = "month-dec"
)

// 📆 星期
const (
	WeekSun MarkerId = "week-sun"
	WeekMon MarkerId = "week-mon"
	WeekTue MarkerId = "week-tue"
	WeekWed MarkerId = "week-wed"
	WeekThu MarkerId = "week-thu"
	WeekFri MarkerId = "week-fri"
	WeekSat MarkerId = "week-sat"
)
