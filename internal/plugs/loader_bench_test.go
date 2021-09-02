package plugs_test

import (
	"testing"

	target "gxn.tw/performance/switch-or-map/internal/plugs"
	"gxn.tw/performance/switch-or-map/pkg/plugs"
)

func BenchmarkPluginSwitching(b *testing.B) {
	getter := target.LoadGetter("../../plugs/switching/switching.so")
	runBenchmarkPlugin(b, "switching", getter)
}

func BenchmarkPluginMapping(b *testing.B) {
	getter := target.LoadGetter("../../plugs/mapping/mapping.so")
	runBenchmarkPlugin(b, "mapping", getter)
}

func runBenchmarkPlugin(b *testing.B, name string, getter plugs.Getter) {
	b.Helper()
	l := len(dictionary)
	missed := 0
	j := 0
	for i := 0; i < b.N; i++ { // benchmark loop
		_, ok := getter.Get(dictionary[(2*j)%l]) // benchmark target
		if !ok {
			missed++
			j = 0

		}
		j++
	}
	b.Logf("tried: %d, missed: %d\n", b.N, missed)
}

var (
	dictionary = []string{"aal", "aalii", "aam", "Aani", "aardvark", "aardwolf", "Aaron", "Aaronic", "Aaronical", "Aaronite", "Aaronitic", "Aaru", "Ab", "aba", "Ababdeh", "Ababua", "abac", "abaca", "abacate", "abacay", "abacinate", "abacination", "abaciscus", "abacist", "aback", "abactinal", "abactinally", "abaction", "abactor", "abaculus", "abacus", "Abadite", "abaff", "abaft", "abaisance", "abaiser", "abaissed", "abalienate", "abalienation", "abalone", "Abama", "abampere", "abandon", "abandonable", "abandoned", "abandonedly", "abandonee", "abandoner", "abandonment", "Abanic", "Abantes", "abaptiston", "Abarambo", "Abaris", "abarthrosis", "abarticular", "abarticulation", "abas", "abase", "abased", "abasedly", "abasedness", "abasement", "abaser", "Abasgi", "abash", "abashed", "abashedly", "abashedness", "abashless", "abashlessly", "abashment", "abasia", "abasic", "abask", "Abassin", "abastardize", "abatable", "abate", "abatement", "abater", "abatis", "abatised", "abaton", "abator", "abattoir", "Abatua", "abature", "abave", "abaxial", "abaxile", "abaze", "abb", "Abba", "abbacomes", "abbacy", "Abbadide", "abbas", "abbasi", "abbassi", "Abbasside", "abbatial", "abbatical", "abbess", "abbey", "abbeystede", "Abbie", "abbot", "abbotcy", "abbotnullius", "abbotship", "abbreviate", "abbreviately", "abbreviation", "abbreviator", "abbreviatory", "abbreviature", "Abby", "abcoulomb", "abdal", "abdat", "Abderian", "Abderite", "abdest", "abdicable", "abdicant", "abdicate", "abdication", "abdicative", "abdicator", "Abdiel", "abditive", "abditory", "abdomen", "abdominal", "Abdominales", "abdominalian", "abdominally", "abdominoanterior", "abdominocardiac", "abdominocentesis", "abdominocystic", "abdominogenital", "abdominohysterectomy", "abdominohysterotomy", "abdominoposterior", "abdominoscope", "abdominoscopy", "abdominothoracic", "abdominous", "abdominovaginal", "abdominovesical", "abduce", "abducens", "abducent", "abduct", "abduction", "abductor", "Abe", "abeam", "abear", "abearance", "abecedarian", "abecedarium", "abecedary", "abed", "abeigh", "Abel", "abele", "Abelia", "Abelian", "Abelicea", "Abelite", "abelite", "Abelmoschus", "abelmosk", "Abelonian", "abeltree", "Abencerrages", "abenteric", "abepithymia", "Aberdeen", "aberdevine", "Aberdonian", "Aberia", "aberrance", "aberrancy", "aberrant", "aberrate", "aberration", "aberrational", "aberrator", "aberrometer", "aberroscope", "aberuncator", "abet", "abetment", "abettal", "abettor", "abevacuation", "abey", "abeyance", "abeyancy", "abeyant", "abfarad", "abhenry", "abhiseka", "abhominable", "abhor", "abhorrence", "abhorrency", "abhorrent", "abhorrently", "abhorrer", "abhorrible", "abhorring", "Abhorson", "abidal", "abidance", "abide", "abider", "abidi", "abiding", "abidingly", "abidingness", "Abie", "Abies", "abietate", "abietene", "abietic", "abietin", "Abietineae", "abietineous", "abietinic", "Abiezer", "Abigail", "abigail", "abigailship", "abigeat", "abigeus", "abilao", "ability", "abilla", "abilo", "abintestate", "abiogenesis", "abiogenesist", "abiogenetic", "abiogenetical", "abiogenetically", "abiogenist", "abiogenous", "abiogeny", "abiological", "abiologically", "abiology", "abiosis", "abiotic", "abiotrophic", "abiotrophy", "Abipon", "abir", "abirritant", "abirritate", "abirritation", "abirritative", "abiston", "Abitibi", "abiuret", "abject", "abjectedness", "abjection", "abjective", "abjectly", "abjectness", "abjoint", "abjudge", "abjudicate", "abjudication", "abjunction", "abjunctive", "abjuration", "abjuratory", "abjure", "abjurement", "abjurer", "abkar", "abkari", "Abkhas", "Abkhasian", "ablach", "ablactate", "ablactation", "ablare", "ablastemic", "ablastous", "ablate", "ablation", "ablatitious", "ablatival", "ablative", "ablator", "ablaut", "ablaze", "able", "ableeze", "ablegate", "ableness", "ablepharia", "ablepharon", "ablepharous", "Ablepharus", "ablepsia", "ableptical", "ableptically", "abler", "ablest", "ablewhackets", "ablins", "abloom", "ablow", "ablude", "abluent", "ablush", "ablution", "ablutionary", "abluvion", "ably", "abmho", "Abnaki", "abnegate", "abnegation", "abnegative", "abnegator", "Abner", "abnerval", "abnet", "abneural", "abnormal", "abnormalism", "abnormalist", "abnormality", "abnormalize", "abnormally", "abnormalness", "abnormity", "abnormous", "abnumerable", "Abo", "aboard", "Abobra", "abode", "abodement", "abody", "abohm", "aboil", "abolish", "abolisher", "abolishment", "abolition", "abolitionary", "abolitionism", "abolitionist", "abolitionize", "abolla", "aboma", "abomasum", "abomasus", "abominable", "abominableness", "abominably", "abominate", "abomination", "abominator", "abomine", "Abongo", "aboon", "aborad", "aboral", "aborally", "abord", "aboriginal", "aboriginality", "aboriginally", "aboriginary", "aborigine", "abort", "aborted", "aborticide", "abortient", "abortifacient", "abortin", "abortion", "abortional", "abortionist", "abortive", "abortively", "abortiveness", "abortus", "abouchement", "abound", "abounder", "abounding", "aboundingly", "about", "abouts", "above", "aboveboard", "abovedeck", "aboveground", "aboveproof", "abovestairs", "abox", "abracadabra", "abrachia", "abradant", "abrade", "abrader", "Abraham", "Abrahamic", "Abrahamidae", "Abrahamite", "Abrahamitic", "abraid", "Abram", "Abramis", "abranchial", "abranchialism", "abranchian", "Abranchiata", "abranchiate", "abranchious", "abrasax", "abrase", "abrash", "abrasiometer", "abrasion", "abrasive", "abrastol", "abraum", "abraxas", "abreact", "abreaction", "abreast", "abrenounce", "abret", "abrico", "abridge", "abridgeable", "abridged", "abridgedly", "abridger", "abridgment", "abrim", "abrin", "abristle", "abroach", "abroad", "Abrocoma", "abrocome", "abrogable", "abrogate", "abrogation", "abrogative", "abrogator", "Abroma", "Abronia", "abrook", "abrotanum", "abrotine", "abrupt", "abruptedly", "abruption", "abruptly", "abruptness", "Abrus", "Absalom", "absampere", "Absaroka", "absarokite", "abscess", "abscessed", "abscession", "abscessroot", "abscind", "abscise", "abscision", "absciss", "abscissa", "abscissae", "abscisse", "abscission", "absconce", "abscond", "absconded", "abscondedly", "abscondence"}
)
