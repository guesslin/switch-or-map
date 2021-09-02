package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Config struct {
	Command string
	Package string
	Items   []Element
}

type Element struct {
	Key   string
	Value string
}

func main() {
	numbers := flag.Int("n", 4, "")
	input := flag.String("f", "template", "template filepath")
	output := flag.String("o", "", "output folder path")
	lib := flag.String("p", "main", "package name")

	flag.Parse()
	fmt.Println("Start generating", *numbers, *input, *output)

	// load and parse template
	tmpl, err := template.ParseFiles(*input)
	panicOn(err)

	// prepare output writer
	f, err := os.Create(*output)
	panicOn(err)
	defer f.Close()
	writer := io.MultiWriter(f, os.Stdout)

	// create output
	elements := buildElements(*numbers)
	var command string
	if len(os.Args) > 1 {
		command = fmt.Sprintf("generator %s", strings.Join(os.Args[1:], " "))
	}
	tmpl.Execute(writer, Config{Command: command, Package: *lib, Items: elements})
}

func buildElements(numbers int) []Element {
	res := make([]Element, 0, numbers)
	for i := 0; i < numbers; i++ {
		res = append(res, Element{dictionary[2*i], dictionary[2*i+1]})
	}
	return res
}

func panicOn(err error, msg ...string) {
	if err != nil {
		panic(err)
	}
}

func generateWords(i int) (string, string) {
	total := len(dictionary)
	if i < (total / 2) {
		return dictionary[2*i], dictionary[2*i+1]
	}
	n := i / (total / 2)
	return fmt.Sprintf("%s%d", dictionary[2*i], n), fmt.Sprintf("%s%d", dictionary[2*i+1], n)
}

var (
	dictionary = []string{"aal", "aalii", "aam", "Aani", "aardvark", "aardwolf", "Aaron", "Aaronic", "Aaronical", "Aaronite", "Aaronitic", "Aaru", "Ab", "aba", "Ababdeh", "Ababua", "abac", "abaca", "abacate", "abacay", "abacinate", "abacination", "abaciscus", "abacist", "aback", "abactinal", "abactinally", "abaction", "abactor", "abaculus", "abacus", "Abadite", "abaff", "abaft", "abaisance", "abaiser", "abaissed", "abalienate", "abalienation", "abalone", "Abama", "abampere", "abandon", "abandonable", "abandoned", "abandonedly", "abandonee", "abandoner", "abandonment", "Abanic", "Abantes", "abaptiston", "Abarambo", "Abaris", "abarthrosis", "abarticular", "abarticulation", "abas", "abase", "abased", "abasedly", "abasedness", "abasement", "abaser", "Abasgi", "abash", "abashed", "abashedly", "abashedness", "abashless", "abashlessly", "abashment", "abasia", "abasic", "abask", "Abassin", "abastardize", "abatable", "abate", "abatement", "abater", "abatis", "abatised", "abaton", "abator", "abattoir", "Abatua", "abature", "abave", "abaxial", "abaxile", "abaze", "abb", "Abba", "abbacomes", "abbacy", "Abbadide", "abbas", "abbasi", "abbassi", "Abbasside", "abbatial", "abbatical", "abbess", "abbey", "abbeystede", "Abbie", "abbot", "abbotcy", "abbotnullius", "abbotship", "abbreviate", "abbreviately", "abbreviation", "abbreviator", "abbreviatory", "abbreviature", "Abby", "abcoulomb", "abdal", "abdat", "Abderian", "Abderite", "abdest", "abdicable", "abdicant", "abdicate", "abdication", "abdicative", "abdicator", "Abdiel", "abditive", "abditory", "abdomen", "abdominal", "Abdominales", "abdominalian", "abdominally", "abdominoanterior", "abdominocardiac", "abdominocentesis", "abdominocystic", "abdominogenital", "abdominohysterectomy", "abdominohysterotomy", "abdominoposterior", "abdominoscope", "abdominoscopy", "abdominothoracic", "abdominous", "abdominovaginal", "abdominovesical", "abduce", "abducens", "abducent", "abduct", "abduction", "abductor", "Abe", "abeam", "abear", "abearance", "abecedarian", "abecedarium", "abecedary", "abed", "abeigh", "Abel", "abele", "Abelia", "Abelian", "Abelicea", "Abelite", "abelite", "Abelmoschus", "abelmosk", "Abelonian", "abeltree", "Abencerrages", "abenteric", "abepithymia", "Aberdeen", "aberdevine", "Aberdonian", "Aberia", "aberrance", "aberrancy", "aberrant", "aberrate", "aberration", "aberrational", "aberrator", "aberrometer", "aberroscope", "aberuncator", "abet", "abetment", "abettal", "abettor", "abevacuation", "abey", "abeyance", "abeyancy", "abeyant", "abfarad", "abhenry", "abhiseka", "abhominable", "abhor", "abhorrence", "abhorrency", "abhorrent", "abhorrently", "abhorrer", "abhorrible", "abhorring", "Abhorson", "abidal", "abidance", "abide", "abider", "abidi", "abiding", "abidingly", "abidingness", "Abie", "Abies", "abietate", "abietene", "abietic", "abietin", "Abietineae", "abietineous", "abietinic", "Abiezer", "Abigail", "abigail", "abigailship", "abigeat", "abigeus", "abilao", "ability", "abilla", "abilo", "abintestate", "abiogenesis", "abiogenesist", "abiogenetic", "abiogenetical", "abiogenetically", "abiogenist", "abiogenous", "abiogeny", "abiological", "abiologically", "abiology", "abiosis", "abiotic", "abiotrophic", "abiotrophy", "Abipon", "abir", "abirritant", "abirritate", "abirritation", "abirritative", "abiston", "Abitibi", "abiuret", "abject", "abjectedness", "abjection", "abjective", "abjectly", "abjectness", "abjoint", "abjudge", "abjudicate", "abjudication", "abjunction", "abjunctive", "abjuration", "abjuratory", "abjure", "abjurement", "abjurer", "abkar", "abkari", "Abkhas", "Abkhasian", "ablach", "ablactate", "ablactation", "ablare", "ablastemic", "ablastous", "ablate", "ablation", "ablatitious", "ablatival", "ablative", "ablator", "ablaut", "ablaze", "able", "ableeze", "ablegate", "ableness", "ablepharia", "ablepharon", "ablepharous", "Ablepharus", "ablepsia", "ableptical", "ableptically", "abler", "ablest", "ablewhackets", "ablins", "abloom", "ablow", "ablude", "abluent", "ablush", "ablution", "ablutionary", "abluvion", "ably", "abmho", "Abnaki", "abnegate", "abnegation", "abnegative", "abnegator", "Abner", "abnerval", "abnet", "abneural", "abnormal", "abnormalism", "abnormalist", "abnormality", "abnormalize", "abnormally", "abnormalness", "abnormity", "abnormous", "abnumerable", "Abo", "aboard", "Abobra", "abode", "abodement", "abody", "abohm", "aboil", "abolish", "abolisher", "abolishment", "abolition", "abolitionary", "abolitionism", "abolitionist", "abolitionize", "abolla", "aboma", "abomasum", "abomasus", "abominable", "abominableness", "abominably", "abominate", "abomination", "abominator", "abomine", "Abongo", "aboon", "aborad", "aboral", "aborally", "abord", "aboriginal", "aboriginality", "aboriginally", "aboriginary", "aborigine", "abort", "aborted", "aborticide", "abortient", "abortifacient", "abortin", "abortion", "abortional", "abortionist", "abortive", "abortively", "abortiveness", "abortus", "abouchement", "abound", "abounder", "abounding", "aboundingly", "about", "abouts", "above", "aboveboard", "abovedeck", "aboveground", "aboveproof", "abovestairs", "abox", "abracadabra", "abrachia", "abradant", "abrade", "abrader", "Abraham", "Abrahamic", "Abrahamidae", "Abrahamite", "Abrahamitic", "abraid", "Abram", "Abramis", "abranchial", "abranchialism", "abranchian", "Abranchiata", "abranchiate", "abranchious", "abrasax", "abrase", "abrash", "abrasiometer", "abrasion", "abrasive", "abrastol", "abraum", "abraxas", "abreact", "abreaction", "abreast", "abrenounce", "abret", "abrico", "abridge", "abridgeable", "abridged", "abridgedly", "abridger", "abridgment", "abrim", "abrin", "abristle", "abroach", "abroad", "Abrocoma", "abrocome", "abrogable", "abrogate", "abrogation", "abrogative", "abrogator", "Abroma", "Abronia", "abrook", "abrotanum", "abrotine", "abrupt", "abruptedly", "abruption", "abruptly", "abruptness", "Abrus", "Absalom", "absampere", "Absaroka", "absarokite", "abscess", "abscessed", "abscession", "abscessroot", "abscind", "abscise", "abscision", "absciss", "abscissa", "abscissae", "abscisse", "abscission", "absconce", "abscond", "absconded", "abscondedly", "abscondence"}
)
