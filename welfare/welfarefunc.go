package welfare

import "regexp"

type Welfarepoint struct {
	Company          string `json:"company"`
	Three            bool   `json:"three"`
	Yearend          bool   `json:"yearend"`
	Bitrh            bool   `json:"bitrh"`
	Marry            bool   `json:"marry"`
	Maternity        bool   `json:"maternity"`
	Patent           bool   `json:"patent"`
	Longterm         bool   `json:"longterm"`
	Insurance        bool   `json:"insurance"`
	Stock            bool   `json:"stock"`
	Annual           bool   `json:"annual"`
	Attendance       bool   `json:"attendance"`
	Performance      bool   `json:"performance"`
	Travel           bool   `json:"travel"`
	Consolation      bool   `json:"consolation"`
	Health           bool   `json:"health"`
	Flexible         bool   `json:"flexible"`
	Paternityleave   bool   `json:"paternityleave"`
	Travelleave      bool   `json:"travelleave"`
	Physiologyleave  bool   `json:"Physiologyleave"`
	Fullpaysickleave bool   `json:"fullpaysickleave"`
	Dorm             bool   `json:"dorm"`
	Restaurant       bool   `json:"restaurant"`
	Childcare        bool   `json:"childcare"`
	Transport        bool   `json:"transport"`
	Servemeals       bool   `json:"servemeals"`
	Snack            bool   `json:"snack"`
	Afternoon        bool   `json:"afternoon"`
	Gym              bool   `json:"gym"`
	Education        bool   `json:"education"`
	Tail             bool   `json:"tail"`
	Employeetravel   bool   `json:"employeetravel"`
	Society          bool   `json:"society"`
	Overtime         bool   `json:"overtime"`
	Shift            bool   `json:"shift"`
	Permanent        bool   `json:"permanent"`
}

func (w *Welfarepoint) Match(welfare string) {
	//key word regex for money
	three := regexp.MustCompile("三節")
	yearend := regexp.MustCompile("年終")
	yearend1 := regexp.MustCompile("年節")
	birth := regexp.MustCompile("生日")
	marry := regexp.MustCompile("結婚")
	maternity := regexp.MustCompile("生育")
	patent := regexp.MustCompile("專利")
	longterm := regexp.MustCompile("久任")
	insurance := regexp.MustCompile("團保")
	stock := regexp.MustCompile("股票")
	stock1 := regexp.MustCompile("入股")
	annual := regexp.MustCompile("分紅")
	attendance := regexp.MustCompile("全勤")
	performance := regexp.MustCompile("績效")
	travel := regexp.MustCompile("旅遊補助")
	travel1 := regexp.MustCompile("旅遊津貼")
	consolation := regexp.MustCompile("慰問")
	health := regexp.MustCompile("健康檢查")
	health1 := regexp.MustCompile("體檢")
	health2 := regexp.MustCompile("健檢")

	//key word regex for working time

	flexible := regexp.MustCompile("彈性上下班")
	paternityleave := regexp.MustCompile("陪產假")
	travelleave := regexp.MustCompile("旅遊假")
	physiologyleave := regexp.MustCompile("生理假")
	fullpaysickleave := regexp.MustCompile("全薪病假")
	fullpaysickleave1 := regexp.MustCompile("不扣薪病假")

	//key word regex for infrastructure

	dorm := regexp.MustCompile("宿舍")
	restaurant := regexp.MustCompile("餐廳")
	childcare := regexp.MustCompile("托兒")
	childcare1 := regexp.MustCompile("育兒")
	transport := regexp.MustCompile("交通")
	servemeals := regexp.MustCompile("供餐")
	servemeals1 := regexp.MustCompile("餐點")
	servemeals2 := regexp.MustCompile("供應午餐")
	servemeals3 := regexp.MustCompile("供應晚餐")
	afternoon := regexp.MustCompile("下午茶")
	snack := regexp.MustCompile("點心")
	gym := regexp.MustCompile("健身房")

	//key word regex for entertainment

	education := regexp.MustCompile("教育訓練")
	education1 := regexp.MustCompile("培訓")
	tail := regexp.MustCompile("尾牙")
	tail1 := regexp.MustCompile("旺年會")
	employeetravel := regexp.MustCompile("員工旅遊")
	society := regexp.MustCompile("社團")

	//key word regex for unusually

	overtime := regexp.MustCompile("加班")
	overtime1 := regexp.MustCompile("無加班")
	overtime2 := regexp.MustCompile("不加班")
	overtime3 := regexp.MustCompile("不鼓勵加班")
	overtime4 := regexp.MustCompile("免加班")
	shift := regexp.MustCompile("輪班")
	shift1 := regexp.MustCompile("無需輪班")
	shift2 := regexp.MustCompile("不輪班")
	shift3 := regexp.MustCompile("不須輪班")
	shift4 := regexp.MustCompile("免加班/輪班")
	permanent := regexp.MustCompile("外派")
	permanent1 := regexp.MustCompile("長駐")

	w.Three = three.MatchString(welfare)
	w.Yearend = yearend.MatchString(welfare) || yearend1.MatchString(welfare)
	w.Bitrh = birth.MatchString(welfare)
	w.Marry = marry.MatchString(welfare)
	w.Maternity = maternity.MatchString(welfare)
	w.Patent = patent.MatchString(welfare)
	w.Longterm = longterm.MatchString(welfare)
	w.Insurance = insurance.MatchString(welfare)
	w.Stock = stock.MatchString(welfare) || stock1.MatchString(welfare)
	w.Annual = annual.MatchString(welfare)
	w.Attendance = attendance.MatchString(welfare)
	w.Performance = performance.MatchString(welfare)
	w.Travel = travel.MatchString(welfare) || travel1.MatchString(welfare)
	w.Consolation = consolation.MatchString(welfare)
	w.Health = health.MatchString(welfare) || health1.MatchString(welfare) || health2.MatchString(welfare)
	w.Flexible = flexible.MatchString(welfare)
	w.Paternityleave = paternityleave.MatchString(welfare)
	w.Travelleave = travelleave.MatchString(welfare)
	w.Physiologyleave = physiologyleave.MatchString(welfare)
	w.Fullpaysickleave = fullpaysickleave.MatchString(welfare) || fullpaysickleave1.MatchString(welfare)
	w.Dorm = dorm.MatchString(welfare)
	w.Restaurant = restaurant.MatchString(welfare)
	w.Childcare = childcare.MatchString(welfare) || childcare1.MatchString(welfare)
	w.Transport = transport.MatchString(welfare)
	w.Servemeals = servemeals.MatchString(welfare) || servemeals1.MatchString(welfare) || servemeals2.MatchString(welfare) || servemeals3.MatchString(welfare)
	w.Snack = snack.MatchString(welfare)
	w.Afternoon = afternoon.MatchString(welfare)
	w.Gym = gym.MatchString(welfare)
	w.Education = education.MatchString(welfare) || education1.MatchString(welfare)
	w.Tail = tail.MatchString(welfare) || tail1.MatchString(welfare)
	w.Employeetravel = employeetravel.MatchString(welfare)
	w.Society = society.MatchString(welfare)
	w.Overtime = overtime.MatchString(welfare) && !(overtime1.MatchString(welfare) || overtime2.MatchString(welfare) || overtime3.MatchString(welfare) || overtime4.MatchString(welfare))
	w.Shift = shift.MatchString(welfare) && !(shift1.MatchString(welfare) || shift2.MatchString(welfare) || shift3.MatchString(welfare) || shift4.MatchString(welfare))
	w.Permanent = permanent.MatchString(welfare) || permanent1.MatchString(welfare)
}

func (w Welfarepoint) Wtoi() int {
	result :=
		btou(w.Three)*2 + btou(w.Yearend)*2 + btou(w.Bitrh)*2 + btou(w.Marry)*2 + btou(w.Maternity)*2 +
			btou(w.Patent)*2 + btou(w.Longterm)*2 + btou(w.Insurance)*3 + btou(w.Stock)*3 + btou(w.Annual)*2 +
			btou(w.Attendance) + btou(w.Performance)*2 + btou(w.Travel)*2 + btou(w.Consolation)*2 + btou(w.Health)*3 +
			btou(w.Flexible)*3 + btou(w.Paternityleave)*3 + btou(w.Travelleave)*3 + btou(w.Physiologyleave)*3 + btou(w.Fullpaysickleave)*3 +
			btou(w.Dorm)*3 + btou(w.Restaurant)*2 + btou(w.Childcare)*3 + btou(w.Transport)*2 + btou(w.Servemeals)*2 +
			btou(w.Snack) + btou(w.Afternoon) + btou(w.Gym)*3 + btou(w.Education)*3 + btou(w.Tail)*3 +
			btou(w.Employeetravel)*3 + btou(w.Society)*3 + btou(w.Overtime)*(-1) + btou(w.Shift)*(-1) + btou(w.Permanent)*(-1)
	return result
}

func btou(b bool) int {
	if b {
		return 1
	}
	return 0
}
