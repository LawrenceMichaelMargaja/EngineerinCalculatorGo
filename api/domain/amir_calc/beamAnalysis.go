package amir_calc

import (
	"fmt"
	"math"
)

func BeamColumnAnalysis(Analysis string, Method string, Units string,

	//user input FACTORS(tab)
	tensileFactor float64, compressFactor float64, bendingFactor float64, shearFactor float64,

	//user input FORCES(tab) loading
	Vx float64, Vy float64,
	Mx float64, My float64,
	P float64,
	//Tr float64,

	//user input properties for MEMBER(tab) properties
	Lb float64,
	Lx float64, Kx float64,
	Ly float64, Ky float64,
	LLT float64,
	cbx float64, cby float64,
	SRc float64, SRt float64,

	//Steel Properties, (under MEMBERS tab)
	ModE float64, YieldStr float64, UltStr float64,

	//Section shape, (under MEMBERS tab)
	Shape string,

	//Section name database values,  (under MEMBERS tab)
	T_F string,
	W float64,
	A float64,
	d float64,
	ddet float64,
	Ht float64,
	h float64,
	OD float64,
	bf float64,
	bfdet float64,
	B float64,
	b float64,
	ID float64,
	tw float64,
	twdet float64,
	twdet_2 float64,
	tf float64,
	tfdet float64,
	t float64,
	tnom float64,
	tdes float64,
	kdes float64,
	kdet float64,
	k1 float64,
	x float64,
	y float64,
	eo float64,
	xp float64,
	yp float64,
	bf_2tf float64,
	b_t float64,
	b_tdes float64,
	h_tw float64,
	h_tdes float64,
	D_t float64,
	Ix float64,
	Zx float64,
	Sx float64,
	rx float64,
	Iy float64,
	Zy float64,
	Sy float64,
	ry float64,
	Iz float64,
	rz float64,
	Sz float64,
	J float64,
	Cw float64,
	C float64,
	Wno float64,
	Sw1 float64,
	Sw2 float64,
	Sw3 float64,
	Qf float64,
	Qw float64,
	ro float64,
	H float64,
	tan float64,
	Iw float64,
	zA float64,
	zB float64,
	zC float64,
	wA float64,
	wB float64,
	wC float64,
	SwA float64,
	SwB float64,
	SwC float64,
	SzA float64,
	SzB float64,
	SzC float64,
	rts float64,
	ho float64,
	PA float64,
	PA2 float64,
	PB float64,
	PC float64,
	PD float64,
	T float64,
	WGi float64,
	WGo float64) (float64,
	float64,
	float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {

	fmt.Println("hello friends")
	//user input for if axial FORCES(tab) is in Tension bolted
	var boltsAlongTension float64 = 0
	var boltsOnWeb bool = false
	var numOfHoles float64 = 0
	var diameterOfHoles float64 = 0

	//user input for if axial FORCES(tab) is in Tension welded
	var lengthOfConnection float64 = 0
	var lengthOfConnection1 float64 = 0
	var lengthOfConnection2 float64 = 0

	//user input for if axial FORCES(tab) is in Tension gusset dimensions
	var twoSideGusset bool = false
	var gussetThickness float64 = 0

	if YieldStr == 0 {
		YieldStr = 250.0
	}
	if ModE == 0 {
		ModE = 200000.0
	}

	var alpha float64 = 1.0
	var factor float64 = 0.9
	if Method == "ASD" {
		alpha = 1.6
		factor = 1 / 1.67
	}

	if tensileFactor <= 0 {
		tensileFactor = factor
	}
	if compressFactor <= 0 {
		compressFactor = factor
	}
	if bendingFactor <= 0 {
		bendingFactor = factor
	}
	if shearFactor <= 0 {
		shearFactor = factor
	}

	var Tr float64 = 0.0            //omitted from calc
	var holeIncrement float64 = 2.0 //for nominal diameter of holes if SI

	//Unit conversion to SI
	if Units == "English" {
		ModE = uConverToSI(ModE, "ksi")
		YieldStr = uConverToSI(YieldStr, "ksi")
		UltStr = uConverToSI(UltStr, "ksi")
		P = uConverToSI(P, "kips")
		Vx = uConverToSI(Vx, "kips")
		Vy = uConverToSI(Vy, "kips")
		Mx = uConverToSI(Mx, "kip-ft")
		My = uConverToSI(My, "kip-ft")
		Tr = uConverToSI(Tr, "kip-ft")
		Lb = uConverToSI(Lb, "ft")
		Lx = uConverToSI(Lx, "ft")
		Ly = uConverToSI(Ly, "ft")
		LLT = uConverToSI(LLT, "ft")
		holeIncrement = 1.5875
		diameterOfHoles = uConverToSI(diameterOfHoles, "in")
		lengthOfConnection1 = uConverToSI(lengthOfConnection1, "in")
		lengthOfConnection2 = uConverToSI(lengthOfConnection2, "in")
		gussetThickness = uConverToSI(gussetThickness, "in")
	}

	//Conversion in terms of N, mm, and, MPa for in-program calculation.
	P *= 1000
	Vx *= 1000
	Vy *= 1000
	Mx *= 1000000
	My *= 1000000
	Tr *= 1000000
	Lb *= 1000
	Lx *= 1000
	Ly *= 1000
	LLT *= 1000

	//Pdelta effects
	var Peb, pfact float64
	Peb = 3.141592 * 3.141592 * ModE * mMin(Ix, Iy) / Lb / Lb
	pfact = math.Sqrt(1 + alpha*math.Abs(P)/Peb)
	cbx = cbx * pfact
	cby = cby * pfact
	var thickRatio float64 = h_tw + h_tdes + D_t
	var thickRatio2 float64 = bf_2tf + b_t + b_tdes
	var Lcx float64 = Kx * Lx
	var Lcy float64 = Ky * Ly
	var LvHSS float64 = Lb * 0.5
	var shearArea float64
	var Shape1 string = "Weak"

	//substitute missing database values with alternate values in database
	switch Shape {
	case "I":
	case "C":
	case "T":
	case "2L":
		thickRatio = d / t
		thickRatio2 = b / t
	case "L":
	case "recHSS":
		tw = tdes
		tf = tdes
		d = Ht
		Shape1 = Shape
	case "roundHSS":
		Shape = "Pipe"
		Shape1 = Shape
	default:
		Shape1 = Shape
	}

	if t == 0.0 {
		t = tdes
	}

	//Setting P state (compression, tension)
	var PInTension bool = (P <= 0)
	P = math.Abs(P)

	//Axial Tensile Capacity
	var Pt float64 = tensileFactor * cTension(Shape, YieldStr, UltStr, A, d, OD, bf, B, b, tw, tf, t, tdes, x, y, Zx, Zy, H, holeIncrement, boltsAlongTension, boltsOnWeb, numOfHoles, diameterOfHoles, lengthOfConnection, lengthOfConnection1, lengthOfConnection2, twoSideGusset, gussetThickness)

	//Axial Compressive Capacity
	var Pc float64 = compressFactor * cCompression(Shape, YieldStr, ModE, Lcx, Lcy, SRc, thickRatio, thickRatio2, A, d, Ht, h, bf, B, b, tw, tf, t, tdes, x, y, Ix, rx, Iy, ry, rz, J, Cw, ro, H, alpha)

	var Ix1 float64 = Ix
	var Zx1 float64 = Zx
	var Sx1 float64 = Sx
	var rx1 float64 = rx
	var Iy1 float64 = Iy
	var Zy1 float64 = Zy
	var Sy1 float64 = Sy
	var ry1 float64 = ry

	//Flexural Capacity about the X-axis
	var Mcx float64 = bendingFactor * cFlexure(Shape, YieldStr, ModE, P, cbx, thickRatio, thickRatio2, LLT, Lx, A, d, bf, b, tw, tf, t, tdes, x, y, Ix1, Zx1, Sx1, rx1, Iy1, Zy1, Sy1, ry1, rz, J, rts)

	if Shape == "recHSS" {
		Ix1 = Iy
		Zx1 = Zy
		Sx1 = Sy
		rx1 = ry
		Iy1 = Ix
		Zy1 = Zx
		Sy1 = Sx
		ry1 = rx
	}

	//Flexural Capacity about the Y-axis
	var Mcy float64 = bendingFactor * cFlexure(Shape1, YieldStr, ModE, P, cby, thickRatio, thickRatio2, LLT, Ly, A, d, bf, b, tw, tf, t, tdes, x, y, Ix1, Zx1, Sx1, rx1, Iy1, Zy1, Sy1, ry1, rz, J, rts)

	//Torsional Capacity
	var Tc float64 = factor * cTorBuckling(Shape, ModE, YieldStr, thickRatio, C, LvHSS, OD, Zx, Sy*1.97)

	//Shear Capacity about the X-axis
	//Resists shear with Y direction
	var shearFactorHolder float64 = shearFactor
	switch Shape {
	case "I":
		shearArea = d * tw
		if math.Round(uConverToEN(YieldStr, "MPa")) >= 50.0 && (h_tw <= 54.0) {
			shearFactorHolder = shearFactor / 0.9
		}
	case "C":
		shearArea = d * tw
	case "T":
		shearArea = d * tw
	case "2L":
		shearArea = 2 * d * t
	case "L":
		shearArea = d * t
	case "recHSS":
		shearArea = 2 * h * t
	default:
		shearArea = A * 0.5
	}

	var Vcy float64 = shearFactorHolder * cShear(Shape, ModE, YieldStr, thickRatio, shearArea, LvHSS, OD)

	//Shear Capacity about the Y-axis
	//Resists shear with X direction
	shearFactorHolder = shearFactor
	switch Shape {
	case "I":
		shearArea = 2 * bf * tf
		/*
			if math.Round(uConverToEN(YieldStr,"MPa"))==50.0&&(h_tw<=54.0) {
				shearFactorHolder=shearFactor/0.9
			}
		*/
	case "C":
		shearArea = 2 * bf * tf
	case "T":
		shearArea = bf * tf
	case "2L":
		shearArea = 2 * b * t
	case "L":
		shearArea = b * t
	case "recHSS":
		shearArea = 2 * b * t
	default:
		shearArea = A * 0.5
	}

	var Vcx float64 = shearFactorHolder * cShear(Shape, ModE, YieldStr, thickRatio2, shearArea, LvHSS, OD)

	//Converting N and mm to kN and m
	P = P / 1000
	Vx = Vx / 1000
	Vy = Vy / 1000
	Mx = Mx / 1000000
	My = My / 1000000
	Tr = Tr * 12 / 1000000
	Pt = Pt / 1000
	Pc = Pc / 1000
	Vcx = Vcx / 1000
	Vcy = Vcy / 1000
	Mcx = Mcx / 1000000
	Mcy = Mcy / 1000000
	Tc = Tc * 12 / 1000000

	//Convert back if necessary
	if Units == "English" {
		P = uConverToEN(P, "kN")
		Vx = uConverToEN(Vx, "kN")
		Vy = uConverToEN(Vy, "kN")
		Mx = uConverToEN(Mx, "kN-m")
		My = uConverToEN(My, "kN-m")
		Tr = uConverToEN(Tr, "kN-m")
		Pt = uConverToEN(Pt, "kN")
		Pc = uConverToEN(Pc, "kN")
		Vcx = uConverToEN(Vcx, "kN")
		Vcy = uConverToEN(Vcy, "kN")
		Mcx = uConverToEN(Mcx, "kN-m")
		Mcy = uConverToEN(Mcy, "kN-m")
		Tc = uConverToEN(Tc, "kN-m")
	}

	var PRatio = P / Pc
	var slenderness = mMax(Lcy/ry, Lcx/rx) / SRc
	if PInTension {
		PRatio = P / Pt
		slenderness = mMax(Ly/ry, Lx/rx) / SRt
		if rz != 0 {
			slenderness = mMax(slenderness, Ly/rz/SRt, Lx/rz/SRt)
		}
	}

	return Pt, Pc, Mcx, Mcy, Vcx, Vcy, PRatio, Mx / Mcx, My / Mcy, Vx / Vcx, Vy / Vcy, cInteraction(Shape, PInTension, P, Mx, My, Vx, Vy, Tr, Pt, Pc, Mcx, Mcy, Vcx, Vcy, Tc), slenderness
}

//Misc. Functions
func mMax(flots ...float64) float64 {
	mNow := flots[0]
	for _, i := range flots {
		if mNow < i {
			mNow = i
		}
	}
	return mNow
}
func mMin(flots ...float64) float64 {
	mNow := flots[0]
	for _, i := range flots {
		if mNow > i {
			mNow = i
		}
	}
	return mNow
}

//Functions for Nominal Capacity
func cTension(Shape string, YieldStr float64, UltStr float64, A float64, d float64, OD float64, bf float64, B float64, b float64, tw float64, tf float64, t float64, tdes float64, x float64, y float64, Zx float64, Zy float64, H float64, holeIncrement float64, boltsAlongTension float64, boltsOnWeb bool, numOfHoles float64, diameterOfHoles float64, lengthOfConnection float64, lengthOfConnection1 float64, lengthOfConnection2 float64, twoSideGusset bool, gussetThickness float64) float64 {

	var weldDist float64 = bf + b
	var U float64
	if (lengthOfConnection + lengthOfConnection1 + lengthOfConnection2 + boltsAlongTension) == 0 {
		lengthOfConnection = weldDist
		lengthOfConnection1 = weldDist
		lengthOfConnection2 = weldDist
		U = 1.0
		return mMin(A*YieldStr, A*UltStr)
	}
	var thn = t + tf + tdes //Relevant thickness
	var xbar float64 = 0.0
	switch Shape {
	case "I":
		U = 2 * bf * tf / A
		xbar = d/2 - Zx/A
		if boltsOnWeb {
			U = 1 - U
			xbar = Zy / A
			thn = tw
		}
	case "C":
		U = 2 * bf * tf / A
		xbar = d/2 - Zx/A
		if boltsOnWeb {
			U = 1 - U
			xbar = x
			thn = tw
		}
	case "T":
		U = bf * tf / A
		xbar = y
		if boltsOnWeb {
			U = 1 - U
			xbar = Zy / A
			thn = tw
		}
	case "2L":
		U = 0.5
		xbar = y
		if boltsOnWeb {
			xbar = Zy / A
			thn = thn * 2
		}
	case "L":
		U = 0.5
		xbar = y
		if boltsOnWeb {
			xbar = x
		}
	case "recHSS":
		xbar = 0.25 * (B*B + 2*B*H) / (B + H)
		numOfHoles = 2
		diameterOfHoles = gussetThickness
		if twoSideGusset {
			numOfHoles = 0
			xbar = 0.25 * (B * B) / (B + H)
		}
	case "Pipe":
		numOfHoles = 2
		diameterOfHoles = gussetThickness
		if lengthOfConnection < 1.3*OD {
			xbar = OD / 3.141592
		}
	default:
	}
	if lengthOfConnection > 0 {
		U = mMax(U, 1.0-xbar/lengthOfConnection)
	}
	if boltsAlongTension <= 0 {
		var weldThickness float64 = thn
		if thn > 0.25*25.4 {
			weldThickness = weldThickness - holeIncrement
		}
		lengthOfConnection = mMax(0.5*(lengthOfConnection1+lengthOfConnection2), 4*weldThickness)
	}

	if boltsAlongTension > 0 {
		switch Shape {
		case "I":
			if boltsOnWeb && (boltsAlongTension >= 4) {
				U = mMax(U, 0.70)
			}
			if !boltsOnWeb && (boltsAlongTension >= 3) {
				U = mMax(U, 0.85)
				if bf >= (2 / 3 * d) {
					U = mMax(U, 0.90)
				}
			}
		case "L":
			U = mMax(U, 0.60)
			if boltsAlongTension >= 4 {
				U = mMax(U, 0.80)
			}
		case "2L":
			U = mMax(U, 0.60)
			if boltsAlongTension >= 4 {
				U = mMax(U, 0.80)
			}
		default:
			U = 1.0
		}
	} else {
		U = U * (3 * lengthOfConnection * lengthOfConnection) / (3*lengthOfConnection*lengthOfConnection + weldDist*weldDist)
		//if condition for null pass.
		if (lengthOfConnection1 + lengthOfConnection2) == 0 {
			U = 1.0
		}
	}

	var Anet float64 = A - numOfHoles*(diameterOfHoles+holeIncrement)*thn
	//fail-safe if can't compute Anet
	if Anet <= 0 {
		Anet = A //Error Catch only
	}
	Anet = mMin(Anet, 0.75*A)
	var Aeff = Anet * U
	return mMin(A*YieldStr, Aeff*UltStr)
}
func cCompression(Shape string, YieldStr float64, ModE float64, Lcx float64, Lcy float64, SRc float64, thickRatio float64, thickRatio2 float64, A float64, d float64, Ht float64, h float64, bf float64, B float64, b float64, tw float64, tf float64, t float64, tdes float64, x float64, y float64, Ix float64, rx float64, Iy float64, ry float64, rz float64, J float64, Cw float64, ro float64, H float64, alpha float64) float64 {
	var Lcz float64 = mMax(Lcx, Lcy)
	var KLrx float64 = Lcx / rx
	var KLry float64 = Lcy / ry
	var FB = 1073741823.0
	var FTB = 1073741823.0
	var Ae float64
	var thn = t + tf + tdes
	var xo = mMax(0.0, x-thn*0.5)
	if Shape == "T" || Shape == "2L" {
		xo = 0.0
		Cw = 1.0
	}
	var yo = mMax(0.0, y-thn*0.5)
	var ro2 float64 = ro * ro
	if ro == 0 {
		ro2 = (xo*xo*A + yo*yo*A + Ix + Iy) / A
	}
	if H == 0 {
		H = 1 - (xo*xo+yo*yo)/ro2
	}
	var Fex = 3.141592 * 3.141592 * ModE / KLrx / KLrx
	var Fey = 3.141592 * 3.141592 * ModE / KLry / KLry
	var Fez = (3.141592*3.141592*ModE*Cw/Lcz/Lcz + 77200*J) / (A * ro2)
	var Fe = 0.5 * (Fey + Fez) * (1 - math.Sqrt(1-4*Fey*Fez*H/(Fey+Fez)/(Fey+Fez))) / H
	Fe = mMin(Fex, Fey, Fez, Fe)
	//Note: If calculating Pnx and Pny, use Fe=Fex and Fe=Fey respectively. Otherwise, use the least Fe value since it is the weakest state it can have.
	var Fcr = 0.877 * Fe
	if YieldStr/Fe < 2.25 {
		Fcr = YieldStr * math.Pow(0.658, YieldStr/Fe)
	}
	switch {
	case Shape == "I" || Shape == "C":
		var de = d * miniAe(0.18, 1.31, 1.49, thickRatio, ModE, YieldStr, Fcr)
		if thickRatio <= 1.49*math.Sqrt(ModE/Fcr) {
			de = d
		}
		var bfe = bf * miniAe(0.22, 1.49, 0.56, thickRatio2, ModE, YieldStr, Fcr)
		if thickRatio2 <= 0.56*math.Sqrt(ModE/Fcr) {
			bfe = bf
		}
		Ae = A - (bf-bfe)*tf*2 - (d-de)*tw
		if Shape == "C" {
			Fe = 0.5 * (Fex + Fez) * (1 - math.Sqrt(1-4*Fex*Fez*H/(Fex+Fez)/(Fex+Fez))) / H
		}
		Fe = mMin(Fex, Fey, Fe)
		Fcr = 0.877 * Fe
		if YieldStr/Fe < 2.25 {
			Fcr = YieldStr * math.Pow(0.658, YieldStr/Fe)
		}
		if KLrx > SRc {
			Fcr = 0.6 * YieldStr
		}
	case Shape == "recHSS":
		var Hte = h * miniAe(0.2, 1.38, 1.40, thickRatio, ModE, YieldStr, Fcr)
		if thickRatio <= 1.40*math.Sqrt(ModE/Fcr) {
			Hte = h
		}
		var Be = b * miniAe(0.2, 1.38, 1.40, thickRatio2, ModE, YieldStr, Fcr)
		if thickRatio2 <= 1.40*math.Sqrt(ModE/Fcr) {
			Be = b
		}
		Ae = A - 2*(h-Hte)*tdes - 2*(b-Be)*tdes*0
	case Shape == "Pipe":
		Ae = A * mMin(1.0, 2/3+0.038*ModE/YieldStr/thickRatio)
	case Shape == "T":
		var de = d * miniAe(0.22, 1.49, 0.75, thickRatio, ModE, YieldStr, Fcr)
		if thickRatio <= 0.75*math.Sqrt(ModE/Fcr) {
			de = d
		}
		var bfe = bf * miniAe(0.22, 1.49, 0.56, thickRatio2, ModE, YieldStr, Fcr)
		if thickRatio2 <= 0.56*math.Sqrt(ModE/Fcr) {
			bfe = bf
		}
		Ae = A - (d-de)*tw - (bf-bfe)*tf
	case Shape == "2L":
		var Ki = 0.50
		var ri = rz
		if Lcy/ri/3 <= 40 {
			Ki = 0
		}
		var KLri2 = KLry*KLry + Ki*Ki*Lcy/3*Lcy/3/ri/ri
		Fey = 3.141592 * 3.141592 * ModE / KLri2
		Fez = (77200 * J * 2) / (A * ro2)
		Fe = 0.5 * (Fey + Fez) * (1 - math.Sqrt(1-4*Fey*Fez*H/(Fey+Fez)/(Fey+Fez))) / H
		Fe = mMin(Fex, Fey, Fez, Fe)
		Fcr = 0.877 * Fe
		if YieldStr/Fe < 2.25 {
			Fcr = YieldStr * math.Pow(0.658, YieldStr/Fe)
		}
		if KLrx > SRc {
			Fcr = 0.6 * YieldStr
		}
		var de = d * miniAe(0.22, 1.49, 0.45, thickRatio, ModE, YieldStr, Fcr)
		if thickRatio <= 0.45*math.Sqrt(ModE/Fcr) {
			de = d
		}
		var be = b * miniAe(0.22, 1.49, 0.45, thickRatio2, ModE, YieldStr, Fcr)
		if thickRatio2 <= 0.45*math.Sqrt(ModE/Fcr) {
			be = b
		}
		Ae = A - (b-be)*t*2 - (d-de)*t*2
	default:
		var de = d * miniAe(0.22, 1.49, 0.45, thickRatio, ModE, YieldStr, Fcr)
		if thickRatio <= 0.45*math.Sqrt(ModE/Fcr) {
			de = d
		}
		var be = b * miniAe(0.22, 1.49, 0.45, thickRatio2, ModE, YieldStr, Fcr)
		if thickRatio2 <= 0.45*math.Sqrt(ModE/Fcr) {
			be = b
		}
		Ae = A - (b-be)*t - (d-de)*t
	}
	FB = Ae * Fcr
	FTB = Ae * Fcr
	return mMin(FB, FTB)
}
func cFlexure(Shape string, YieldStr float64, ModE float64, P float64, cb float64, thickRatio float64, thickRatio2 float64, LLT float64, Lb float64, A float64, d float64, bf float64, b float64, tw float64, tf float64, t float64, tdes float64, x float64, y float64, Ix float64, Zx float64, Sx float64, rx float64, Iy float64, Zy float64, Sy float64, ry float64, rz float64, J float64, rts float64) float64 {
	var Mp float64 = mMin(YieldStr*Zx, 1.6*YieldStr*Sx)
	var LB = 1073741823000.0
	var LLB = 1073741823000.0
	var LTB float64 = 1073741823000.0
	var FLB float64 = 1073741823000.0
	var WLB float64 = 1073741823000.0
	var CFLB float64 = 1073741823000.0
	var CFY float64 = 1073741823000.0
	switch {
	case Shape == "I" || Shape == "C":
		var c float64 = 1.0
		if Shape == "C" {
			c = (d - tf) * 0.5 * Iy / (bf * bf / (12 + 2*(d-tf)*tw/bf/tf)) / Sx
		}
		var aw float64 = thickRatio * tw * tw / bf / tf
		var Lp float64 = 1.1 * rts * math.Sqrt(ModE/YieldStr)
		var Lr float64 = 1.95 * rts * ModE / YieldStr / 0.7 * math.Sqrt(J*c/Sx/(d-tf)+math.Sqrt(J*c/Sx/(d-tf)*J*c/Sx/(d-tf)+6.76*(0.7*YieldStr/ModE)*(0.7*YieldStr/ModE)))
		var Rpg = mMin(1-aw/(1200+300*aw)*(thickRatio-5.7*math.Sqrt(ModE/YieldStr)), 1.0)
		var Me = YieldStr * Sx
		var Rpc = mMax(mMin(Mp/Me-(Mp/Me-1.0)*(thickRatio-3.76*math.Sqrt(ModE/YieldStr))/1.94/math.Sqrt(ModE/YieldStr), Mp/Me), 1.0)
		if (thickRatio > 5.7*math.Sqrt(ModE/YieldStr)) && Shape == "I" {
			J = 0
			Lr = 3.141592 * rts * math.Sqrt(ModE/0.7/YieldStr)
		} else if (thickRatio > 3.76*math.Sqrt(ModE/YieldStr)) && Shape == "I" {
		} else {
			Lp = Lp * 1.6 * ry / rts
		}
		Lb = mMin(LLT, Lb)
		LTB = mMin(cb*(Rpg*(Rpc*Me-(Rpc*Me-0.7*YieldStr*Sx)*(Lb-Lp)/(Lr-Lp))), cb*Sx*(Rpg*3.141592*3.141592*ModE*rts*rts/Lb/Lb*math.Sqrt(1+0.078*J*Lb*Lb/Sx/(d-tf)/rts/rts)))
		CFLB = mMin(Rpg*(Rpc*Me-(Rpc-0.7)*Me*(thickRatio2-0.38*math.Sqrt(ModE/YieldStr))/0.62/math.Sqrt(ModE/YieldStr)), 3.6*ModE*Sx/(thickRatio2)/(thickRatio2)/math.Sqrt(thickRatio))
		CFY = Rpg * Rpc * Me
		return mMin(Mp, LTB, CFLB, CFY)
	case Shape == "recHSS":
		var Lconst = 2 * ModE * ry * math.Sqrt(J*A)
		var Lp = 0.065 * Lconst / Mp
		var Lr = Lconst / 0.7 / YieldStr / Sx
		var Rpg = 1 - ((d-tf)*tw/bf/tf)/(1200+300*((d-tf)*tw/bf/tf))*(thickRatio-5.7*math.Sqrt(ModE/YieldStr))
		var be = mMin(1.92*tf*math.Sqrt(ModE/YieldStr)*(1-0.38*tf/b*math.Sqrt(ModE/YieldStr)), b)
		var Ieff = Ix - (b-be)*tf*tf*tf/6 - (b-be)*tf*0.5*(d-tf)*(d-tf)
		var Se = 2 * Ieff / d
		Lb = mMin(LLT, Lb)
		LTB = mMin(cb*(Mp-(Mp-0.7*YieldStr*Sx)*(Lb-Lp)/(Lr-Lp)), cb*Lconst/Lb)
		FLB = mMax(Mp-(Mp-YieldStr*Se)*(3.57*thickRatio2*math.Sqrt(YieldStr/ModE)-4.0), YieldStr*Se)
		WLB = mMin(Mp-(Mp-YieldStr*Sx)*(0.305*thickRatio*math.Sqrt(YieldStr/ModE)-0.738), Rpg*YieldStr*Sx, Rpg*3.6*ModE/thickRatio2/thickRatio2)
		return mMin(Mp, FLB, WLB, LTB)
	case Shape == "Pipe":
		LB = Sx * mMin((0.021*ModE/thickRatio+YieldStr), 0.33*ModE/thickRatio)
		return mMin(Mp, LB)
	case Shape == "2L" || Shape == "T":
		var Lp = 1.76 * ry * math.Sqrt(ModE/YieldStr)
		var Lr = 1.95 * ModE * math.Sqrt(Iy*J) / (YieldStr * Sx) * math.Sqrt(2.36*d*YieldStr*Sx/ModE/J+1)
		if LLT < Lp {
			Lp = Lr
		}
		var Sxc float64 = Ix / y
		var B = 2.3 * (d - y) / Lb * math.Sqrt(Iy/J)
		if P < 0 {
			if Shape == "2L" {
				Mp = YieldStr * Sx
			} else {
				Mp = 1.5 * YieldStr * Sx
			}
		}
		Lb = mMin(LLT, Lb)
		LTB = mMin(Mp-(Mp-YieldStr*Sx)*(Lb-Lp)/(Lr-Lp), 1.95*ModE*math.Sqrt(Iy*J)/Lb*(B+math.Sqrt(1+B*B)))
		FLB = mMin(1.6*YieldStr*Sx, Mp-(Mp-0.7*YieldStr*Sxc)*(thickRatio2-0.38*math.Sqrt(ModE/YieldStr))/0.62/math.Sqrt(ModE/YieldStr), 0.7*Sxc*ModE/thickRatio2/thickRatio2)
		LB = Sxc * mMin(YieldStr, YieldStr*(1.43-0.515*thickRatio*math.Sqrt(YieldStr/ModE)), 1.52*ModE/thickRatio/thickRatio)
		return mMin(Mp, LTB, FLB, LB)
	case Shape == "L":
		var Bw = (b*b - d*d) / b * .9354
		var Mcr = 1.125 * ModE * A * rz * t * cb / Lb * (math.Sqrt(1+(4.4*Bw*rz/Lb/t)*(4.4*Bw*rz/Lb/t)) + (4.4 * Bw * rz / Lb / t))
		if P <= 0 && Bw == 0 {
			Mcr = 0.58 * ModE * b * b * b * b * t * cb / Lb / Lb * (math.Sqrt(1+0.88*Lb*Lb*t*t/b/b/b/b) - 1)
		}
		Lb = mMin(LLT, Lb)
		LTB = mMin((1.92-1.17*math.Sqrt(YieldStr*Sx/Mcr))*YieldStr*Sx, (0.92-0.17*Mcr/YieldStr/Sx)*Mcr)
		LLB = Sx * mMin(YieldStr*(2.43-1.72*thickRatio*math.Sqrt(YieldStr/ModE)), 0.71*ModE/thickRatio/thickRatio)
		return mMin(1.5*YieldStr*Sx, LTB, LLB)
	default:
		Mp = mMin(YieldStr*Zy, 1.6*YieldStr*Sy)
		var FLB = mMin(Mp-(Mp-0.7*YieldStr*Sy)*((thickRatio2-0.38*math.Sqrt(ModE/YieldStr))/0.62/math.Sqrt(ModE/YieldStr)), 0.69*ModE/thickRatio2/thickRatio2*Sy)
		return mMin(Mp, FLB)
	}
}
func cShear(Shape string, ModE float64, YieldStr float64, thickRatio float64, shearArea float64, LvHSS float64, OD float64) float64 {
	var adm = math.Sqrt(ModE/YieldStr) / thickRatio
	switch {
	case Shape == "I" || Shape == "C":
		var k = 5.34
		var cv = mMin(1.0, 1.10*math.Sqrt(k)*adm)
		return 0.6 * YieldStr * shearArea * cv
	case Shape == "T" || Shape == "L":
		var k = 1.20
		var cv = mMin(1.0, 1.10*math.Sqrt(k)*adm, 1.5125*k*adm*adm)
		return 0.6 * YieldStr * shearArea * cv
	case Shape == "2L" || Shape == "recHSS":
		var k = 5.00
		var cv = mMin(1.0, 1.10*math.Sqrt(k)*adm, 1.5125*k*adm*adm)
		return 0.6 * YieldStr * shearArea * cv
	case Shape == "Pipe":
		var Fcr = mMin(uConverToSI(mMax(1.60/math.Sqrt(LvHSS/OD)/math.Pow(thickRatio, 1.25), 0.78/math.Pow(thickRatio, 1.5))*uConverToEN(ModE, "MPa"), "ksi"), 0.6*YieldStr)
		return Fcr * shearArea
	default:
		var k = 1.20
		var cv = mMin(1.0, 1.10*math.Sqrt(k)*adm, 1.5125*k*adm*adm)
		if YieldStr < 482.633 {
			cv = 1.0
		}
		return 0.6 * YieldStr * shearArea * cv
	}
}
func cTorBuckling(Shape string, ModE float64, YieldStr float64, thickRatio float64, C float64, LvHSS float64, OD float64, plasticMod float64, Xfact float64) float64 {
	switch {
	case Shape == "recHSS":
		var Fcr = mMin(0.6*YieldStr, 0.6*YieldStr*2.45*math.Sqrt(ModE/YieldStr)/thickRatio, 0.458*3.141592*3.141592*ModE/thickRatio/thickRatio)
		return Fcr * C
	case Shape == "Pipe":
		var Fcr = mMin(mMax(1.60/math.Sqrt(LvHSS/OD)/math.Pow(thickRatio, 1.25), 0.78/math.Pow(thickRatio, 1.5))*ModE/1.3, 0.6*YieldStr)
		if C == 0 {
			C = 3.141592 * 0.5 * (OD - OD/thickRatio) * (OD - OD/thickRatio) * OD / thickRatio
		}
		return Fcr * C
	default:
		var Fcr = mMin(0.6*YieldStr, 0.6*YieldStr*2.45*math.Sqrt(ModE/YieldStr)/thickRatio, 0.458*3.141592*3.141592*ModE/thickRatio/thickRatio)
		return Fcr * Xfact
	}
}

//Functions for Convenience
func miniAe(c1 float64, c2 float64, lambda float64, thickRatio float64, ModE float64, YieldStr float64, Fcr float64) float64 {
	lambda = lambda * math.Sqrt(ModE/YieldStr)
	var rFact = (c2 * lambda / thickRatio) * math.Sqrt(YieldStr/Fcr)
	return (1 - c1*rFact) * rFact
}

//Essential Functions
func uConverToEN(conVal float64, unit string) float64 {
	switch unit {
	//to kips
	case "kN":
		return 1 / 4.4482216 * conVal

	//to kip-ft
	case "kN-m":
		return 1 / 1.35581795 * conVal

	//to kip-ft
	case "kN-mm":
		return 1 / 1.35581795 / 1000 * conVal

	//to kip-ft
	case "N-m":
		return 1 / 1.35581795 / 1000 * conVal

	//to kip-ft
	case "N-mm":
		return 1 / 1.35581795 / 1000000 * conVal

	//to kips/ft
	case "kN/m":
		return 1 / 14.5939029 * conVal

	//to ft
	case "m":
		return 1 / 0.3048 * conVal

	//to in
	case "mm":
		return 1 / 25.4 * conVal

	//to ksi
	case "MPa":
		return 1 / 6.8947573 * conVal

	default:
		return conVal
	}
}
func uConverToSI(conVal float64, unit string) float64 {
	switch unit {
	//to kN
	case "kips":
		return 4.4482216 * conVal

	//to kN-m
	case "kip-ft":
		return 1.35581795 * conVal

	//to kN-m
	case "kip-in":
		return 1.35581795 / 12 * conVal

	//to kN/m
	case "kips/ft":
		return 14.5939029 * conVal

	//to m
	case "ft":
		return 0.3048 * conVal

	//to mm
	case "in":
		return 25.4 * conVal

	//to MPa
	case "ksi":
		return 6.8947573 * conVal

	default:
		return conVal
	}
}
func cInteraction(Shape string, PInTension bool, P float64, Mx float64, My float64, Vx float64, Vy float64, Tr float64, Pt float64, Pc float64, Mcx float64, Mcy float64, Vcx float64, Vcy float64, Tc float64) float64 {
	if PInTension {
		Pc = Pt
	}
	var Mcheck = 1.0
	if (math.Abs(Mx) + math.Abs(My)) == 0 {
		Mcheck = 0
	}
	//Check if individual failure exists...
	if mMax(P/Pc, Mx/Mcx, My/Mcy, Vx/Vcx, Vy/Vcy, Tr/Tc) > 1.0 {
		return 0.0
		//...
		// return mMax(P/Pc,Mx/Mcx,My/Mcy,Vx/Vcx,Vy/Vcy,Tr/Tc)
	}
	if Vx != 0 || Vy != 0 || Tr != 0 {
		var TrTc = 0.0
		if Tc != 0.0 {
			TrTc = Tr / Tc
		}
		if (Shape != "recHSS" || Shape != "Pipe") && Tr != 0 {
			return TrTc
		} else {
			return (P/Pc + Mx/Mcx + My/Mcy)
			// ...
			// return (P/Pc+Mx/Mcx+My/Mcy)+(Vx/Vcx+Vy/Vcy+Tr/Tc)*(Vx/Vcx+Vy/Vcy+TrTc)
		}
	} else if P*Mcheck != 0 {
		if Shape == "L" {
			return P/Pc + Mx/Mcx + My/Mcy
		} else {
			return mMax(P/Pc+(Mx/Mcx+My/Mcy)*8/9, 0.5*P/Pc+(Mx/Mcx+My/Mcy))
		}
	} else {
		return 0.0
		//...
		// return mMax(P/Pc,Mx/Mcx,My/Mcy,Vx/Vcx,Vy/Vcy,Tr/Tc)
	}
}

func main() {

	//setting the units to SI

	//start of test code lines

	//design method
	var Analysis string = "Beam-Column"
	var design string = "LRFD"
	var units string = "English"

	//user input factors
	var tensileFactor float64 = 0
	var compressFactor float64 = 0
	var bendingFactor float64 = 0
	var shearFactor float64 = 0

	//user input loading (already factored)
	var Mx float64 = 0
	var My float64 = 0
	var Vx float64 = 0
	var Vy float64 = 289.6
	var P float64 = 0

	//user input properties for member
	var Lb float64 = 10
	var Lx float64 = 10
	var Kx float64 = 1.0
	var Ly float64 = 10
	var Ky float64 = 1.0
	var LLT float64 = 10
	var cbx float64 = 1.0
	var cby float64 = 1.0
	var LSC float64 = 200.0
	var LST float64 = 300.0

	//steel properties
	var ModE float64 = 29000
	var YieldStr float64 = 50
	var UltStr float64 = 65

	var Shape string = "I"

	//end of test code lines AND/OR UI input

	//User Analysis Type ("Beam-Column" for all cases involving beam and columns except for struts and significantly slanted columns)
	//(design) variable refers to either "LRFD" or "ASD"
	//Steel Member Shape (i.e. "I","C","recHSS","Pipe","L","T","2L")
	var Pt, Pc, Mcx, Mcy, Vcx, Vcy, Pratio, MxRatio, MyRatio, VxRatio, VyRatio, Combined, KLr = BeamColumnAnalysis(Analysis, design, units,

		//Input FACTORS
		tensileFactor, compressFactor, bendingFactor, shearFactor,

		//Input FORCES Loading (factored)
		Vx, Vy, Mx, My, P,

		//Input MEMBER dimensional properties
		Lb, Lx, Kx, Ly, Ky, LLT, cbx, cby, LSC, LST,

		//Steel Properties
		ModE, YieldStr, UltStr,

		//Section shape
		Shape,

		//Section database values
		"F", 92, 11700, 602, 603, 0, 0, 0, 179, 178, 0, 0, 0, 10.9, 11.1, 6.35, 15, 14.3, 0, 0, 0, 27.7, 38.1, 27, 0, 0, 0, 0, 0, 5.97, 0, 0, 50.1, 0, 0, 645000000, 2510000, 2150000, 234, 14400000, 257000, 161000, 35.1, 0, 0, 0, 712000, 1240000000000, 0, 26300, 17600000, 0, 0, 369000, 1230000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 44.4, 587, 1700, 0, 1870, 1380, 1560, 527, 88.9, 0)

	fmt.Println("\nDesign Capacities = ", Pt, Pc, Mcx, Mcy, Vcx, Vcy)
	fmt.Println("\nDesign Ratio = ", Pratio, MxRatio, MyRatio, VxRatio, VyRatio, Combined, KLr)
}
