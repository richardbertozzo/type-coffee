package entity

import "testing"

func TestCreateLink(t *testing.T) {
	linksValids := []string{
		"https://rollingstone.uol.com.br/media/_versions/godzilla-kingking-reprod-twitter-cortada_widelg.jpg",
		"https://hiperideal.vteximg.com.br/arquivos/ids/170582-1000-1000/1993127.jpg",
		"https://loucodocafe.com.br/wp-content/uploads/2019/09/cafe-jacu-02-e1568167519150.jpg",
		"https://loucodocafe.com.br/wp-content/uploads/2019/09/cafe-jacu-02-e1568167519150.jpeg",
		"https://loucodocafe.com.br/wp-content/uploads/2019/09/cafe-jacu-02-e1568167519150.png",
	}
	linkInvalids := []string{
		"https://www.twitch.tv/richardbertozzo",
		"www.sitedodegola.com.br",
		"oitudobem",
		"",
		"256546456456",
		"https://loucodocafe.com.br/wp-content/uploads/2019/09/cafe-jacu-02-e1568167519150",
	}

	t.Run("Must create a valid link", func(t *testing.T) {
		for _, l := range linksValids {
			l, err := NewImageLink(l)
			if err != nil {
				t.Errorf("Must not get an error, but got: %v", err)
			}
			t.Log(l)
		}
	})

	t.Run("Must create an invalid link", func(t *testing.T) {
		for _, l := range linkInvalids {
			_, err := NewImageLink(l)
			if err == nil {
				t.Errorf("Must get an error, but got nil with link: %s", l)
			}
		}
	})
}
