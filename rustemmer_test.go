package rustemmer

import (
	"testing"
	"reflect"
)

func TestStemmer(t *testing.T) {
	testWords := map[string]string{
		"результаты"   : "результат",
		"в"            : "в",
		"вавиловка"    : "вавиловк",
		"вагнера"      : "вагнер",
		"вагон"        : "вагон",
		"вагона"       : "вагон",
		"вагоне"       : "вагон",
		"вагонов"      : "вагон",
		"вагоном"      : "вагон",
		"вагоны"       : "вагон",
		"важная"       : "важн",
		"важнее"       : "важн",
		"важнейшие"    : "важн",
		"важнейшими"   : "важн",
		"важничал"     : "важнича",
		"важно"        : "важн",
		"важного"      : "важн",
		"важное"       : "важн",
		"важной"       : "важн",
		"важном"       : "важн",
		"важному"      : "важн",
		"важности"     : "важност",
		"важностию"    : "важност",
		"важность"     : "важност",
		"важностью"    : "важност",
		"важную"       : "важн",
		"важны"        : "важн",
		"важные"       : "важн",
		"важный"       : "важн",
		"важным"       : "важн",
		"важных"       : "важн",
		"вазах"        : "ваз",
		"вазы"         : "ваз",
		"вакса"        : "вакс",
		"вакханка"     : "вакханк",
		"вал"          : "вал",
		"валандался"   : "валанда",
		"валентина"    : "валентин",
		"валериановых" : "валерианов",
		"валерию"      : "валер",
		"валетами"     : "валет",
		"вали"         : "вал",
		"валил"        : "вал",
		"валился"      : "вал",
		"валится"      : "вал",
		"валов"        : "вал",
		"вальдшнепа"   : "вальдшнеп",
		"вальс"        : "вальс",
		"вальса"       : "вальс",
		"вальсе"       : "вальс",
		"вальсишку"    : "вальсишк",
		"вальтера"     : "вальтер",
		"валяется"     : "валя",
		"валялась"     : "валя",
		"валялись"     : "валя",
		"валялось"     : "валя",
		"валялся"      : "валя",
		"валять"       : "валя",
		"валяются"     : "валя",
		"вам"          : "вам",
		"вами"         : "вам",
	}

	for word, base := range testWords {
		testBase := GetWordBase(word)
		if !reflect.DeepEqual(base, testBase) {
			t.Errorf("Not equal: [%s] %s != %s", word, base, testBase)
		}
	}
}