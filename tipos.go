package main

type Parametros struct {
	Number_slots_left  int
	Number_slots_right int
	Strike_gap         int
	Interval_frame     int
	Asserts            map[string]int //tipo de assert y rango

}

type Coin struct {
	TypeCoin string
	Value    int
}
