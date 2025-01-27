// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mocks

type Branchable struct {
	Id        int
	Data      string
	BranchRef *Branchable
}

func (b *Branchable) Branch() (*Branchable, bool) {
	if b.BranchRef != nil {
		return b.BranchRef, true
	}
	return nil, false
}

func (b *Branchable) Equal(other any) bool {
	if o, ok := other.(*Branchable); ok {
		return b.Id == o.Id
	}
	return false
}

func branchableToItself(id int, data string) *Branchable {
	b := &Branchable{Id: id, Data: data}
	b.BranchRef = b
	return b
}

var smallBranchTestDataRoot = &Branchable{Id: 1, Data: "data1", BranchRef: nil}
var small2 = &Branchable{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot}
var small3 = &Branchable{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot}
var small4 = &Branchable{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot}
var small5 = &Branchable{Id: 5, Data: "data5", BranchRef: small3}
var small6 = &Branchable{Id: 6, Data: "data6", BranchRef: small3}
var small7 = &Branchable{Id: 7, Data: "data7", BranchRef: small5}
var small8 = &Branchable{Id: 8, Data: "data8", BranchRef: small6}
var small9 = &Branchable{Id: 9, Data: "data9", BranchRef: small7}
var small10 = &Branchable{Id: 10, Data: "data10", BranchRef: small2}

var SmallBranchableTestData = []*Branchable{
	smallBranchTestDataRoot, small2, small3, small4, small5, small6, small7,
	small8, small9, small10,
}

func smallBranchableSolution() map[string][]*Branchable {
	return map[string][]*Branchable{
		"1": {
			{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot},
			{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot},
			{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot},
		},
		"2": {
			{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot},
			{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot},
			{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot},
			{Id: 10, Data: "data10", BranchRef: &Branchable{Id: 2, Data: "data2"}},
		},
		"3": {
			{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot},
			{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot},
			{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
		},
		"5": {
			{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot},
			{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot},
			{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 7, Data: "data7", BranchRef: &Branchable{Id: 5, Data: "data5"}},
		},
		"6": {
			{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot},
			{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot},
			{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 8, Data: "data8", BranchRef: &Branchable{Id: 6, Data: "data6"}},
		},
		"7": {
			{Id: 2, Data: "data2", BranchRef: smallBranchTestDataRoot},
			{Id: 3, Data: "data3", BranchRef: smallBranchTestDataRoot},
			{Id: 4, Data: "data4", BranchRef: smallBranchTestDataRoot},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 7, Data: "data7", BranchRef: &Branchable{Id: 5, Data: "data5"}},
			{Id: 9, Data: "data9", BranchRef: &Branchable{Id: 7, Data: "data7"}},
		},
	}
}

var mediumBranchTestDataRoot1 = &Branchable{Id: 1, Data: "data1", BranchRef: nil}
var mediumBranchTestDataRoot2 = &Branchable{Id: 7, Data: "data7", BranchRef: nil}
var mediumBranchTestDataRoot3 = &Branchable{Id: 9, Data: "data9", BranchRef: nil}
var medium2 = &Branchable{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1}
var medium3 = &Branchable{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1}
var medium4 = &Branchable{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1}
var medium5 = &Branchable{Id: 5, Data: "data5", BranchRef: mediumBranchTestDataRoot3}
var medium6 = &Branchable{Id: 6, Data: "data6", BranchRef: mediumBranchTestDataRoot3}
var medium7 = &Branchable{Id: 7, Data: "data7", BranchRef: mediumBranchTestDataRoot2}
var medium8 = &Branchable{Id: 8, Data: "data8", BranchRef: mediumBranchTestDataRoot2}
var medium9 = &Branchable{Id: 9, Data: "data9", BranchRef: mediumBranchTestDataRoot3}
var medium10 = &Branchable{Id: 10, Data: "data10", BranchRef: mediumBranchTestDataRoot2}
var medium11 = &Branchable{Id: 11, Data: "data11", BranchRef: mediumBranchTestDataRoot1}
var medium12 = &Branchable{Id: 12, Data: "data12", BranchRef: mediumBranchTestDataRoot1}
var medium13 = &Branchable{Id: 13, Data: "data13", BranchRef: mediumBranchTestDataRoot1}
var medium14 = &Branchable{Id: 14, Data: "data14", BranchRef: mediumBranchTestDataRoot1}
var medium15 = &Branchable{Id: 15, Data: "data15", BranchRef: mediumBranchTestDataRoot1}
var medium16 = &Branchable{Id: 16, Data: "data16", BranchRef: mediumBranchTestDataRoot1}
var medium17 = &Branchable{Id: 17, Data: "data17", BranchRef: mediumBranchTestDataRoot1}
var medium18 = &Branchable{Id: 18, Data: "data18", BranchRef: mediumBranchTestDataRoot1}
var medium19 = &Branchable{Id: 19, Data: "data19", BranchRef: mediumBranchTestDataRoot1}
var medium20 = &Branchable{Id: 20, Data: "data20", BranchRef: mediumBranchTestDataRoot1}
var medium21 = &Branchable{Id: 21, Data: "data21", BranchRef: mediumBranchTestDataRoot1}
var medium22 = &Branchable{Id: 22, Data: "data22", BranchRef: mediumBranchTestDataRoot1}
var medium23 = &Branchable{Id: 23, Data: "data23", BranchRef: mediumBranchTestDataRoot1}
var medium24 = &Branchable{Id: 24, Data: "data24", BranchRef: mediumBranchTestDataRoot1}
var medium25 = &Branchable{Id: 25, Data: "data25", BranchRef: mediumBranchTestDataRoot1}
var medium26 = &Branchable{Id: 26, Data: "data26", BranchRef: mediumBranchTestDataRoot1}
var medium27 = &Branchable{Id: 27, Data: "data27", BranchRef: mediumBranchTestDataRoot1}
var medium28 = &Branchable{Id: 28, Data: "data28", BranchRef: mediumBranchTestDataRoot1}
var medium29 = &Branchable{Id: 29, Data: "data29", BranchRef: mediumBranchTestDataRoot1}
var medium30 = &Branchable{Id: 30, Data: "data30", BranchRef: mediumBranchTestDataRoot1}

var mediumBranchableTestData = []*Branchable{
	mediumBranchTestDataRoot1, medium2, medium3, medium4, medium5, medium6, medium7, medium8, medium9, medium10,
	medium11, medium12, medium13, medium14, medium15, medium16, medium17, medium18, medium19, medium20,
	medium21, medium22, medium23, medium24, medium25, medium26, medium27, medium28, medium29, medium30,
}

func mediumBranchableSolution() map[string][]*Branchable {
	return map[string][]*Branchable{
		"1": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
		},
		"2": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
		},
		"3": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
		},
		"6": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 8, Data: "data8", BranchRef: &Branchable{Id: 6, Data: "data6"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
		},
		"7": {
			{Id: 10, Data: "data10", BranchRef: mediumBranchTestDataRoot2},
			{Id: 28, Data: "data28", BranchRef: mediumBranchTestDataRoot2},
			{Id: 28, Data: "data28", BranchRef: mediumBranchTestDataRoot2},
		},
		"9": {
			{Id: 14, Data: "data14", BranchRef: mediumBranchTestDataRoot3},
			{Id: 27, Data: "data27", BranchRef: mediumBranchTestDataRoot3},
		},
		"11": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 12, Data: "data12", BranchRef: &Branchable{Id: 11, Data: "data11"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
		},
		"12": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 12, Data: "data12", BranchRef: &Branchable{Id: 11, Data: "data11"}},
			{Id: 13, Data: "data13", BranchRef: &Branchable{Id: 12, Data: "data12"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
		},
		"15": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
		},
		"17": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 21, Data: "data21", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 22, Data: "data22", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 23, Data: "data23", BranchRef: &Branchable{Id: 17, Data: "data17"}},
		},
		"21": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 21, Data: "data21", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 22, Data: "data22", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 23, Data: "data23", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 24, Data: "data24", BranchRef: &Branchable{Id: 21, Data: "data21"}},
			{Id: 25, Data: "data25", BranchRef: &Branchable{Id: 21, Data: "data21"}},
			{Id: 26, Data: "data26", BranchRef: &Branchable{Id: 21, Data: "data21"}},
		},
		"22": {
			{Id: 2, Data: "data2", BranchRef: mediumBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: mediumBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: mediumBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 21, Data: "data21", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 22, Data: "data22", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 23, Data: "data23", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 29, Data: "data29", BranchRef: &Branchable{Id: 22, Data: "data22"}},
		},
		"27": {
			{Id: 14, Data: "data14", BranchRef: mediumBranchTestDataRoot3},
			{Id: 27, Data: "data27", BranchRef: mediumBranchTestDataRoot3},
			{Id: 30, Data: "data30", BranchRef: &Branchable{Id: 27, Data: "data27"}},
		},
	}
}

var largeBranchTestDataRoot1 = &Branchable{Id: 1, Data: "data1", BranchRef: nil}
var largeBranchTestDataRoot2 = &Branchable{Id: 31, Data: "data31", BranchRef: nil}
var largeBranchTestDataRoot3 = &Branchable{Id: 61, Data: "data61", BranchRef: nil}

// Edge case: root that points to itself
var largeBranchTestDataSelfRoot = branchableToItself(91, "data91")
var large2 = &Branchable{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1}
var large3 = &Branchable{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1}
var large4 = &Branchable{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1}
var large5 = &Branchable{Id: 5, Data: "data5", BranchRef: large3}
var large6 = &Branchable{Id: 6, Data: "data6", BranchRef: large3}
var large7 = &Branchable{Id: 7, Data: "data7", BranchRef: large5}
var large8 = &Branchable{Id: 8, Data: "data8", BranchRef: large6}
var large9 = &Branchable{Id: 9, Data: "data9", BranchRef: large7}
var large10 = &Branchable{Id: 10, Data: "data10", BranchRef: large2}
var large11 = &Branchable{Id: 11, Data: "data11", BranchRef: large10}
var large12 = &Branchable{Id: 12, Data: "data12", BranchRef: large11}
var large13 = &Branchable{Id: 13, Data: "data13", BranchRef: large12}
var large14 = &Branchable{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1}
var large15 = &Branchable{Id: 15, Data: "data15", BranchRef: large2}
var large16 = &Branchable{Id: 16, Data: "data16", BranchRef: large15}
var large17 = &Branchable{Id: 17, Data: "data17", BranchRef: large3}
var large18 = &Branchable{Id: 18, Data: "data18", BranchRef: large3}
var large19 = &Branchable{Id: 19, Data: "data19", BranchRef: large3}
var large20 = &Branchable{Id: 20, Data: "data20", BranchRef: large3}
var large21 = &Branchable{Id: 21, Data: "data21", BranchRef: large17}
var large22 = &Branchable{Id: 22, Data: "data22", BranchRef: large17}
var large23 = &Branchable{Id: 23, Data: "data23", BranchRef: large17}
var large24 = &Branchable{Id: 24, Data: "data24", BranchRef: large21}
var large25 = &Branchable{Id: 25, Data: "data25", BranchRef: large21}
var large26 = &Branchable{Id: 26, Data: "data26", BranchRef: large21}
var large27 = &Branchable{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1}
var large28 = &Branchable{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2}
var large29 = &Branchable{Id: 29, Data: "data29", BranchRef: large22}
var large30 = &Branchable{Id: 30, Data: "data30", BranchRef: large27}
var large32 = &Branchable{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2}
var large33 = &Branchable{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2}
var large34 = &Branchable{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2}
var large35 = &Branchable{Id: 35, Data: "data35", BranchRef: large33}
var large36 = &Branchable{Id: 36, Data: "data36", BranchRef: large33}
var large37 = &Branchable{Id: 37, Data: "data37", BranchRef: large35}
var large38 = &Branchable{Id: 38, Data: "data38", BranchRef: large36}
var large39 = &Branchable{Id: 39, Data: "data39", BranchRef: large37}
var large40 = &Branchable{Id: 40, Data: "data40", BranchRef: large32}
var large41 = &Branchable{Id: 41, Data: "data41", BranchRef: large40}
var large42 = &Branchable{Id: 42, Data: "data42", BranchRef: large41}
var large43 = &Branchable{Id: 43, Data: "data43", BranchRef: large42}
var large44 = &Branchable{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2}
var large45 = &Branchable{Id: 45, Data: "data45", BranchRef: large33}
var large46 = &Branchable{Id: 46, Data: "data46", BranchRef: large45}
var large47 = &Branchable{Id: 47, Data: "data47", BranchRef: large33}
var large48 = &Branchable{Id: 48, Data: "data48", BranchRef: large33}
var large49 = &Branchable{Id: 49, Data: "data49", BranchRef: large33}
var large50 = &Branchable{Id: 50, Data: "data50", BranchRef: large33}
var large51 = &Branchable{Id: 51, Data: "data51", BranchRef: large33}
var large52 = &Branchable{Id: 52, Data: "data52", BranchRef: large47}
var large53 = &Branchable{Id: 53, Data: "data53", BranchRef: large47}
var large62 = &Branchable{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3}
var large63 = &Branchable{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3}
var large64 = &Branchable{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3}
var large65 = &Branchable{Id: 65, Data: "data65", BranchRef: large63}
var large66 = &Branchable{Id: 66, Data: "data66", BranchRef: large63}
var large67 = &Branchable{Id: 67, Data: "data67", BranchRef: large65}
var large68 = &Branchable{Id: 68, Data: "data68", BranchRef: large66}
var large69 = &Branchable{Id: 69, Data: "data69", BranchRef: large67}
var large70 = &Branchable{Id: 70, Data: "data70", BranchRef: large62}
var large71 = &Branchable{Id: 71, Data: "data71", BranchRef: large70}
var large72 = &Branchable{Id: 72, Data: "data72", BranchRef: large71}
var large73 = &Branchable{Id: 73, Data: "data73", BranchRef: large72}
var large74 = &Branchable{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3}
var large75 = &Branchable{Id: 75, Data: "data75", BranchRef: large63}
var large76 = &Branchable{Id: 76, Data: "data76", BranchRef: large75}
var large77 = &Branchable{Id: 77, Data: "data77", BranchRef: large63}
var large78 = &Branchable{Id: 78, Data: "data78", BranchRef: large63}
var large79 = &Branchable{Id: 79, Data: "data79", BranchRef: large63}
var large80 = &Branchable{Id: 80, Data: "data80", BranchRef: large63}
var large81 = &Branchable{Id: 81, Data: "data81", BranchRef: large63}
var large82 = &Branchable{Id: 82, Data: "data82", BranchRef: large77}
var large83 = &Branchable{Id: 83, Data: "data83", BranchRef: large77}
var large92 = &Branchable{Id: 92, Data: "data92", BranchRef: largeBranchTestDataSelfRoot}
var large93 = &Branchable{Id: 93, Data: "data93", BranchRef: largeBranchTestDataSelfRoot}
var large94 = &Branchable{Id: 94, Data: "data94", BranchRef: largeBranchTestDataSelfRoot}
var large95 = &Branchable{Id: 95, Data: "data95", BranchRef: large93}
var large96 = &Branchable{Id: 96, Data: "data96", BranchRef: large93}
var large97 = &Branchable{Id: 97, Data: "data97", BranchRef: large95}
var large98 = &Branchable{Id: 98, Data: "data98", BranchRef: large96}
var large99 = &Branchable{Id: 99, Data: "data99", BranchRef: large97}

var largeBranchableTestData = []*Branchable{
	largeBranchTestDataRoot1, largeBranchTestDataRoot2, largeBranchTestDataRoot3,
	largeBranchTestDataSelfRoot, large2, large3, large4, large5, large6, large7,
	large8, large9, large10, large11, large12, large13, large14, large15, large16,
	large17, large18, large19, large20, large21, large22, large23, large24, large25,
	large26, large27, large28, large29, large30, large32, large33, large34,
	large35, large36, large37, large38, large39, large40, large41, large42, large43,
	large44, large45, large46, large47, large48, large49, large50, large51, large52,
	large53, large62, large63, large64, large65, large66, large67, large68, large69,
	large70, large71, large72, large73, large74, large75, large76, large77, large78,
	large79, large80, large81, large82, large83, large92, large93, large94, large95,
	large96, large97, large98, large99,
}

func largeBranchableSolution() map[string][]*Branchable {
	return map[string][]*Branchable{
		"1": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
		},
		"31": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
		},
		"61": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
		},
		"91": {
			{Id: 92, Data: "data92", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 93, Data: "data93", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 94, Data: "data94", BranchRef: largeBranchTestDataSelfRoot},
		},
		"2": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 10, Data: "data10", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
		},
		"3": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
		},
		"5": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 7, Data: "data7", BranchRef: &Branchable{Id: 5, Data: "data5"}},
		},
		"6": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 8, Data: "data8", BranchRef: &Branchable{Id: 6, Data: "data6"}},
		},
		"7": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 7, Data: "data7", BranchRef: &Branchable{Id: 5, Data: "data5"}},
			{Id: 9, Data: "data9", BranchRef: &Branchable{Id: 7, Data: "data7"}},
		},
		"10": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 10, Data: "data10", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 10, Data: "data10"}},
		},
		"11": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 10, Data: "data10", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 10, Data: "data10"}},
			{Id: 12, Data: "data12", BranchRef: &Branchable{Id: 11, Data: "data11"}},
		},
		"12": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 10, Data: "data10", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 11, Data: "data11", BranchRef: &Branchable{Id: 10, Data: "data10"}},
			{Id: 12, Data: "data12", BranchRef: &Branchable{Id: 11, Data: "data11"}},
			{Id: 13, Data: "data13", BranchRef: &Branchable{Id: 12, Data: "data12"}},
		},
		"15": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 10, Data: "data10", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 15, Data: "data15", BranchRef: &Branchable{Id: 2, Data: "data2"}},
			{Id: 16, Data: "data16", BranchRef: &Branchable{Id: 15, Data: "data15"}},
		},
		"17": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 21, Data: "data21", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 22, Data: "data22", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 23, Data: "data23", BranchRef: &Branchable{Id: 17, Data: "data17"}},
		},
		"21": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 21, Data: "data21", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 22, Data: "data22", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 23, Data: "data23", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 24, Data: "data24", BranchRef: &Branchable{Id: 21, Data: "data21"}},
			{Id: 25, Data: "data25", BranchRef: &Branchable{Id: 21, Data: "data21"}},
			{Id: 26, Data: "data26", BranchRef: &Branchable{Id: 21, Data: "data21"}},
		},
		"22": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 5, Data: "data5", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 6, Data: "data6", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 17, Data: "data17", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 18, Data: "data18", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 19, Data: "data19", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 20, Data: "data20", BranchRef: &Branchable{Id: 3, Data: "data3"}},
			{Id: 21, Data: "data21", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 22, Data: "data22", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 23, Data: "data23", BranchRef: &Branchable{Id: 17, Data: "data17"}},
			{Id: 29, Data: "data29", BranchRef: &Branchable{Id: 22, Data: "data22"}},
		},
		"27": {
			{Id: 2, Data: "data2", BranchRef: largeBranchTestDataRoot1},
			{Id: 3, Data: "data3", BranchRef: largeBranchTestDataRoot1},
			{Id: 4, Data: "data4", BranchRef: largeBranchTestDataRoot1},
			{Id: 14, Data: "data14", BranchRef: largeBranchTestDataRoot1},
			{Id: 27, Data: "data27", BranchRef: largeBranchTestDataRoot1},
			{Id: 30, Data: "data30", BranchRef: &Branchable{Id: 27, Data: "data27"}},
		},
		"32": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 40, Data: "data40", BranchRef: &Branchable{Id: 32, Data: "data32"}},
		},
		"33": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 35, Data: "data35", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 36, Data: "data36", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 45, Data: "data45", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 47, Data: "data47", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 48, Data: "data48", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 49, Data: "data49", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 50, Data: "data50", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 51, Data: "data51", BranchRef: &Branchable{Id: 33, Data: "data33"}},
		},
		"35": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 35, Data: "data35", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 36, Data: "data36", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 45, Data: "data45", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 47, Data: "data47", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 48, Data: "data48", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 49, Data: "data49", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 50, Data: "data50", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 51, Data: "data51", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 37, Data: "data37", BranchRef: &Branchable{Id: 35, Data: "data35"}},
		},
		"36": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 35, Data: "data35", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 36, Data: "data36", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 45, Data: "data45", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 47, Data: "data47", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 48, Data: "data48", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 49, Data: "data49", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 50, Data: "data50", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 51, Data: "data51", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 38, Data: "data38", BranchRef: &Branchable{Id: 36, Data: "data36"}},
		},
		"37": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 35, Data: "data35", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 36, Data: "data36", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 45, Data: "data45", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 47, Data: "data47", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 48, Data: "data48", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 49, Data: "data49", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 50, Data: "data50", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 51, Data: "data51", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 37, Data: "data37", BranchRef: &Branchable{Id: 35, Data: "data35"}},
			{Id: 39, Data: "data39", BranchRef: &Branchable{Id: 37, Data: "data37"}},
		},
		"40": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 40, Data: "data40", BranchRef: &Branchable{Id: 32, Data: "data32"}},
			{Id: 41, Data: "data41", BranchRef: &Branchable{Id: 40, Data: "data40"}},
		},
		"41": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 40, Data: "data40", BranchRef: &Branchable{Id: 32, Data: "data32"}},
			{Id: 41, Data: "data41", BranchRef: &Branchable{Id: 40, Data: "data40"}},
			{Id: 42, Data: "data42", BranchRef: &Branchable{Id: 41, Data: "data41"}},
		},
		"42": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 40, Data: "data40", BranchRef: &Branchable{Id: 32, Data: "data32"}},
			{Id: 41, Data: "data41", BranchRef: &Branchable{Id: 40, Data: "data40"}},
			{Id: 42, Data: "data42", BranchRef: &Branchable{Id: 41, Data: "data41"}},
			{Id: 43, Data: "data43", BranchRef: &Branchable{Id: 42, Data: "data42"}},
		},
		"45": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 35, Data: "data35", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 36, Data: "data36", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 45, Data: "data45", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 47, Data: "data47", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 48, Data: "data48", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 49, Data: "data49", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 50, Data: "data50", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 51, Data: "data51", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 46, Data: "data46", BranchRef: &Branchable{Id: 45, Data: "data45"}},
		},
		"47": {
			{Id: 28, Data: "data28", BranchRef: largeBranchTestDataRoot2},
			{Id: 32, Data: "data32", BranchRef: largeBranchTestDataRoot2},
			{Id: 33, Data: "data33", BranchRef: largeBranchTestDataRoot2},
			{Id: 34, Data: "data34", BranchRef: largeBranchTestDataRoot2},
			{Id: 44, Data: "data44", BranchRef: largeBranchTestDataRoot2},
			{Id: 35, Data: "data35", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 36, Data: "data36", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 45, Data: "data45", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 47, Data: "data47", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 48, Data: "data48", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 49, Data: "data49", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 50, Data: "data50", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 51, Data: "data51", BranchRef: &Branchable{Id: 33, Data: "data33"}},
			{Id: 52, Data: "data52", BranchRef: &Branchable{Id: 47, Data: "data47"}},
			{Id: 53, Data: "data53", BranchRef: &Branchable{Id: 47, Data: "data47"}},
		},
		"62": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 70, Data: "data70", BranchRef: &Branchable{Id: 62, Data: "data62"}},
		},
		"63": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 65, Data: "data65", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 66, Data: "data66", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 77, Data: "data77", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 78, Data: "data78", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 79, Data: "data79", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 80, Data: "data80", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 81, Data: "data81", BranchRef: &Branchable{Id: 63, Data: "data63"}},
		},
		"65": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 65, Data: "data65", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 66, Data: "data66", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 77, Data: "data77", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 78, Data: "data78", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 79, Data: "data79", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 80, Data: "data80", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 81, Data: "data81", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 67, Data: "data67", BranchRef: &Branchable{Id: 65, Data: "data65"}},
		},
		"66": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 65, Data: "data65", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 66, Data: "data66", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 77, Data: "data77", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 78, Data: "data78", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 79, Data: "data79", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 80, Data: "data80", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 81, Data: "data81", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 68, Data: "data68", BranchRef: &Branchable{Id: 66, Data: "data66"}},
		},
		"67": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 65, Data: "data65", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 66, Data: "data66", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 77, Data: "data77", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 78, Data: "data78", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 79, Data: "data79", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 80, Data: "data80", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 81, Data: "data81", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 67, Data: "data67", BranchRef: &Branchable{Id: 65, Data: "data65"}},
			{Id: 69, Data: "data69", BranchRef: &Branchable{Id: 67, Data: "data67"}},
		},
		"70": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 70, Data: "data70", BranchRef: &Branchable{Id: 62, Data: "data62"}},
			{Id: 71, Data: "data71", BranchRef: &Branchable{Id: 70, Data: "data70"}},
		},
		"71": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 70, Data: "data70", BranchRef: &Branchable{Id: 62, Data: "data62"}},
			{Id: 71, Data: "data71", BranchRef: &Branchable{Id: 70, Data: "data70"}},
			{Id: 72, Data: "data72", BranchRef: &Branchable{Id: 71, Data: "data71"}},
		},
		"72": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 70, Data: "data70", BranchRef: &Branchable{Id: 62, Data: "data62"}},
			{Id: 71, Data: "data71", BranchRef: &Branchable{Id: 70, Data: "data70"}},
			{Id: 72, Data: "data72", BranchRef: &Branchable{Id: 71, Data: "data71"}},
			{Id: 73, Data: "data73", BranchRef: &Branchable{Id: 72, Data: "data72"}},
		},
		"75": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 65, Data: "data65", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 66, Data: "data66", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 77, Data: "data77", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 78, Data: "data78", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 79, Data: "data79", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 80, Data: "data80", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 81, Data: "data81", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 76, Data: "data76", BranchRef: &Branchable{Id: 75, Data: "data75"}},
		},
		"77": {
			{Id: 62, Data: "data62", BranchRef: largeBranchTestDataRoot3},
			{Id: 63, Data: "data63", BranchRef: largeBranchTestDataRoot3},
			{Id: 64, Data: "data64", BranchRef: largeBranchTestDataRoot3},
			{Id: 74, Data: "data74", BranchRef: largeBranchTestDataRoot3},
			{Id: 65, Data: "data65", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 66, Data: "data66", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 77, Data: "data77", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 78, Data: "data78", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 79, Data: "data79", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 80, Data: "data80", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 81, Data: "data81", BranchRef: &Branchable{Id: 63, Data: "data63"}},
			{Id: 82, Data: "data82", BranchRef: &Branchable{Id: 77, Data: "data77"}},
			{Id: 83, Data: "data83", BranchRef: &Branchable{Id: 77, Data: "data77"}},
		},
		"93": {
			{Id: 92, Data: "data92", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 93, Data: "data93", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 94, Data: "data94", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 95, Data: "data95", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 96, Data: "data96", BranchRef: &Branchable{Id: 93, Data: "data93"}},
		},
		"95": {
			{Id: 92, Data: "data92", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 93, Data: "data93", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 94, Data: "data94", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 95, Data: "data95", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 96, Data: "data96", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 97, Data: "data97", BranchRef: &Branchable{Id: 95, Data: "data95"}},
		},
		"96": {
			{Id: 92, Data: "data92", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 93, Data: "data93", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 94, Data: "data94", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 95, Data: "data95", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 96, Data: "data96", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 98, Data: "data98", BranchRef: &Branchable{Id: 96, Data: "data96"}},
		},
		"97": {
			{Id: 92, Data: "data92", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 93, Data: "data93", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 94, Data: "data94", BranchRef: largeBranchTestDataSelfRoot},
			{Id: 95, Data: "data95", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 96, Data: "data96", BranchRef: &Branchable{Id: 93, Data: "data93"}},
			{Id: 97, Data: "data97", BranchRef: &Branchable{Id: 95, Data: "data95"}},
			{Id: 99, Data: "data99", BranchRef: &Branchable{Id: 97, Data: "data97"}},
		},
	}
}
