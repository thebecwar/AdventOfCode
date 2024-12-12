package aoc2015

import (
	"fmt"
	"strings"
)

type lookAndSayAtom struct {
	Symbol   string
	Pattern  string
	DecaysTo []string
}

var lookAndSayAtoms = map[string]lookAndSayAtom{
	"H":  {Symbol: "H", Pattern: "22", DecaysTo: []string{"H"}},
	"He": {Symbol: "He", Pattern: "13112221133211322112211213322112", DecaysTo: []string{"Hf.Pa.H.Ca.Li"}},
	"Li": {Symbol: "Li", Pattern: "312211322212221121123222112", DecaysTo: []string{"He"}},
	"Be": {Symbol: "Be", Pattern: "111312211312113221133211322112211213322112", DecaysTo: []string{"Ge.Ca.Li"}},
	"B":  {Symbol: "B", Pattern: "1321132122211322212221121123222112", DecaysTo: []string{"Be"}},
	"C":  {Symbol: "C", Pattern: "3113112211322112211213322112", DecaysTo: []string{"B"}},
	"N":  {Symbol: "N", Pattern: "111312212221121123222112", DecaysTo: []string{"C"}},
	"O":  {Symbol: "O", Pattern: "132112211213322112", DecaysTo: []string{"N"}},
	"F":  {Symbol: "F", Pattern: "31121123222112", DecaysTo: []string{"O"}},
	"Ne": {Symbol: "Ne", Pattern: "111213322112", DecaysTo: []string{"F"}},
	"Na": {Symbol: "Na", Pattern: "123222112", DecaysTo: []string{"Ne"}},
	"Mg": {Symbol: "Mg", Pattern: "3113322112", DecaysTo: []string{"Pm.Na"}},
	"Al": {Symbol: "Al", Pattern: "1113222112", DecaysTo: []string{"Mg"}},
	"Si": {Symbol: "Si", Pattern: "1322112", DecaysTo: []string{"Al"}},
	"P":  {Symbol: "P", Pattern: "311311222112", DecaysTo: []string{"Ho.Si"}},
	"S":  {Symbol: "S", Pattern: "1113122112", DecaysTo: []string{"P"}},
	"Cl": {Symbol: "Cl", Pattern: "132112", DecaysTo: []string{"S"}},
	"Ar": {Symbol: "Ar", Pattern: "3112", DecaysTo: []string{"Cl"}},
	"K":  {Symbol: "K", Pattern: "1112", DecaysTo: []string{"Ar"}},
	"Ca": {Symbol: "Ca", Pattern: "12", DecaysTo: []string{"K"}},
	"Sc": {Symbol: "Sc", Pattern: "3113112221133112", DecaysTo: []string{"Ho.Pa.H.Ca.Co"}},
	"Ti": {Symbol: "Ti", Pattern: "11131221131112", DecaysTo: []string{"Sc"}},
	"V":  {Symbol: "V", Pattern: "13211312", DecaysTo: []string{"Ti"}},
	"Cr": {Symbol: "Cr", Pattern: "31132", DecaysTo: []string{"V"}},
	"Mn": {Symbol: "Mn", Pattern: "111311222112", DecaysTo: []string{"Cr.Si"}},
	"Fe": {Symbol: "Fe", Pattern: "13122112", DecaysTo: []string{"Mn"}},
	"Co": {Symbol: "Co", Pattern: "32112", DecaysTo: []string{"Fe"}},
	"Ni": {Symbol: "Ni", Pattern: "11133112", DecaysTo: []string{"Zn.Co"}},
	"Cu": {Symbol: "Cu", Pattern: "131112", DecaysTo: []string{"Ni"}},
	"Zn": {Symbol: "Zn", Pattern: "312", DecaysTo: []string{"Cu"}},
	"Ga": {Symbol: "Ga", Pattern: "13221133122211332", DecaysTo: []string{"Eu.Ca.Ac.H.Ca.Zn"}},
	"Ge": {Symbol: "Ge", Pattern: "31131122211311122113222", DecaysTo: []string{"Ho.Ga"}},
	"As": {Symbol: "As", Pattern: "11131221131211322113322112", DecaysTo: []string{"Ge.Na"}},
	"Se": {Symbol: "Se", Pattern: "13211321222113222112", DecaysTo: []string{"As"}},
	"Br": {Symbol: "Br", Pattern: "3113112211322112", DecaysTo: []string{"Se"}},
	"Kr": {Symbol: "Kr", Pattern: "11131221222112", DecaysTo: []string{"Br"}},
	"Rb": {Symbol: "Rb", Pattern: "1321122112", DecaysTo: []string{"Kr"}},
	"Sr": {Symbol: "Sr", Pattern: "3112112", DecaysTo: []string{"Rb"}},
	"Y":  {Symbol: "Y", Pattern: "1112133", DecaysTo: []string{"Sr.U"}},
	"Zr": {Symbol: "Zr", Pattern: "12322211331222113112211", DecaysTo: []string{"Y.H.Ca.Tc"}},
	"Nb": {Symbol: "Nb", Pattern: "1113122113322113111221131221", DecaysTo: []string{"Er.Zr"}},
	"Mo": {Symbol: "Mo", Pattern: "13211322211312113211", DecaysTo: []string{"Nb"}},
	"Tc": {Symbol: "Tc", Pattern: "311322113212221", DecaysTo: []string{"Mo"}},
	"Ru": {Symbol: "Ru", Pattern: "132211331222113112211", DecaysTo: []string{"Eu.Ca.Tc"}},
	"Rh": {Symbol: "Rh", Pattern: "311311222113111221131221", DecaysTo: []string{"Ho.Ru"}},
	"Pd": {Symbol: "Pd", Pattern: "111312211312113211", DecaysTo: []string{"Rh"}},
	"Ag": {Symbol: "Ag", Pattern: "132113212221", DecaysTo: []string{"Pd"}},
	"Cd": {Symbol: "Cd", Pattern: "3113112211", DecaysTo: []string{"Ag"}},
	"In": {Symbol: "In", Pattern: "11131221", DecaysTo: []string{"Cd"}},
	"Sn": {Symbol: "Sn", Pattern: "13211", DecaysTo: []string{"In"}},
	"Sb": {Symbol: "Sb", Pattern: "3112221", DecaysTo: []string{"Pm.Sn"}},
	"Te": {Symbol: "Te", Pattern: "1322113312211", DecaysTo: []string{"Eu.Ca.Sb"}},
	"I":  {Symbol: "I", Pattern: "311311222113111221", DecaysTo: []string{"Ho.Te"}},
	"Xe": {Symbol: "Xe", Pattern: "11131221131211", DecaysTo: []string{"I"}},
	"Cs": {Symbol: "Cs", Pattern: "13211321", DecaysTo: []string{"Xe"}},
	"Ba": {Symbol: "Ba", Pattern: "311311", DecaysTo: []string{"Cs"}},
	"La": {Symbol: "La", Pattern: "11131", DecaysTo: []string{"Ba"}},
	"Ce": {Symbol: "Ce", Pattern: "1321133112", DecaysTo: []string{"La.H.Ca.Co"}},
	"Pr": {Symbol: "Pr", Pattern: "31131112", DecaysTo: []string{"Ce"}},
	"Nd": {Symbol: "Nd", Pattern: "111312", DecaysTo: []string{"Pr"}},
	"Pm": {Symbol: "Pm", Pattern: "132", DecaysTo: []string{"Nd"}},
	"Sm": {Symbol: "Sm", Pattern: "311332", DecaysTo: []string{"Pm.Ca.Zn"}},
	"Eu": {Symbol: "Eu", Pattern: "1113222", DecaysTo: []string{"Sm"}},
	"Gd": {Symbol: "Gd", Pattern: "13221133112", DecaysTo: []string{"Eu.Ca.Co"}},
	"Tb": {Symbol: "Tb", Pattern: "3113112221131112", DecaysTo: []string{"Ho.Gd"}},
	"Dy": {Symbol: "Dy", Pattern: "111312211312", DecaysTo: []string{"Tb"}},
	"Ho": {Symbol: "Ho", Pattern: "1321132", DecaysTo: []string{"Dy"}},
	"Er": {Symbol: "Er", Pattern: "311311222", DecaysTo: []string{"Ho.Pm"}},
	"Tm": {Symbol: "Tm", Pattern: "11131221133112", DecaysTo: []string{"Er.Ca.Co"}},
	"Yb": {Symbol: "Yb", Pattern: "1321131112", DecaysTo: []string{"Tm"}},
	"Lu": {Symbol: "Lu", Pattern: "311312", DecaysTo: []string{"Yb"}},
	"Hf": {Symbol: "Hf", Pattern: "11132", DecaysTo: []string{"Lu"}},
	"Ta": {Symbol: "Ta", Pattern: "13112221133211322112211213322113", DecaysTo: []string{"Hf.Pa.H.Ca.W"}},
	"W":  {Symbol: "W", Pattern: "312211322212221121123222113", DecaysTo: []string{"Ta"}},
	"Re": {Symbol: "Re", Pattern: "111312211312113221133211322112211213322113", DecaysTo: []string{"Ge.Ca.W"}},
	"Os": {Symbol: "Os", Pattern: "1321132122211322212221121123222113", DecaysTo: []string{"Re"}},
	"Ir": {Symbol: "Ir", Pattern: "3113112211322112211213322113", DecaysTo: []string{"Os"}},
	"Pt": {Symbol: "Pt", Pattern: "111312212221121123222113", DecaysTo: []string{"Ir"}},
	"Au": {Symbol: "Au", Pattern: "132112211213322113", DecaysTo: []string{"Pt"}},
	"Hg": {Symbol: "Hg", Pattern: "31121123222113", DecaysTo: []string{"Au"}},
	"Tl": {Symbol: "Tl", Pattern: "111213322113", DecaysTo: []string{"Hg"}},
	"Pb": {Symbol: "Pb", Pattern: "123222113", DecaysTo: []string{"Tl"}},
	"Bi": {Symbol: "Bi", Pattern: "3113322113", DecaysTo: []string{"Pm.Pb"}},
	"Po": {Symbol: "Po", Pattern: "1113222113", DecaysTo: []string{"Bi"}},
	"At": {Symbol: "At", Pattern: "1322113", DecaysTo: []string{"Po"}},
	"Rn": {Symbol: "Rn", Pattern: "311311222113", DecaysTo: []string{"Ho.At"}},
	"Fr": {Symbol: "Fr", Pattern: "1113122113", DecaysTo: []string{"Rn"}},
	"Ra": {Symbol: "Ra", Pattern: "132113", DecaysTo: []string{"Fr"}},
	"Ac": {Symbol: "Ac", Pattern: "3113", DecaysTo: []string{"Ra"}},
	"Th": {Symbol: "Th", Pattern: "1113", DecaysTo: []string{"Ac"}},
	"Pa": {Symbol: "Pa", Pattern: "13", DecaysTo: []string{"Th"}},
	"U":  {Symbol: "U", Pattern: "3", DecaysTo: []string{"Pa"}},
}
var atomicPatterns = make(map[string]string)

func init() {
	for k, atom := range lookAndSayAtoms {
		atom.DecaysTo = strings.Split(atom.DecaysTo[0], ".")
		lookAndSayAtoms[k] = atom
		atomicPatterns[atom.Pattern] = atom.Symbol
	}
}

func lookAndSayLength(symbol string, iterations int) int {
	atom, ok := lookAndSayAtoms[symbol]
	if !ok {
		panic("ugh")
	}

	if iterations == 0 {
		return len(atom.Pattern)
	}

	totalLength := 0
	for _, decay := range atom.DecaysTo {
		totalLength += lookAndSayLength(decay, iterations-1)
	}

	return totalLength
}

func lookAndSay(s string) string {
	output := ""
	count := 1
	current := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != current {
			output += fmt.Sprintf("%d%c", count, current)
			count = 1
			current = s[i]
		} else {
			count++
		}
	}
	output += fmt.Sprintf("%d%c", count, current)
	return output
}

func Day10Part1() {
	input := "1113222113" // Polonium
	length := lookAndSayLength(atomicPatterns[input], 40)
	fmt.Printf("Day 10 Part 1: %d\n", length)
}

func Day10Part2() {
	input := "1113222113"
	length := lookAndSayLength(atomicPatterns[input], 50)
	fmt.Printf("Day 10 Part 2: %d\n", length)
}
