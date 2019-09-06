package main

type coord string

type square struct {
	piece     piece
	isRosette bool
	isWarzone bool
	isFinish  bool
}
type board map[coord]square

//define board with 20 squares plus 2 winning squares (invisible)
func instantiateBoard() board {
	b := make(map[coord]square)

	 //Populate square settings
		b["A1"] = square{
			isRosette: true,
		}
		b["A2"] = square{}
		b["A3"] = square{}
		b["A4"] = square{}
		b["A6"] = square{
			isFinish: true,
		}
		b["A7"] = square{
			isRosette: true,
		}
		b["A8"] = square{}

		b["B1"] = square{
			isWarzone: true,
		}
		b["B2"] = square{
			isWarzone: true,
		}
		b["B3"] = square{
			isWarzone: true,
		}
		b["B4"] = square{
			isRosette: true,
			isWarzone: true,
		}
		b["B5"] = square{
			isWarzone: true,
		}
		b["B6"] = square{
			isWarzone: true,
		}
		b["B7"] = square{
			isWarzone: true,
		}
		b["B8"] = square{
			isWarzone: true,
		}

		b["C1"] = square{
			isRosette: true,
		}
		b["C2"] = square{}
		b["C3"] = square{}
		b["C4"] = square{}
		b["C6"] = square{
			isFinish: true,
		}
		b["C7"] = square{
			isRosette: true,
		}
		b["C8"] = square{}
	

	return b
}
