package api

import "testing"

func Test_getThing(t *testing.T) {
	api := &bbg{}

	result, err := api.getThing(43)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Names) == 0 {
		t.Fatal("No Games returned should be some")
	}
}

func Test_searchBBG(t *testing.T) {

	api := &bbg{}

	searchMap := make(map[string][]string)

	searchMap["name"] = make([]string, 1)
	searchMap["name"][0] = "verbotene"

	result, err := api.searchBBG(searchMap)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Games) == 0 {
		t.Fatal("No Games returned should be some")
	}

	if result.Games[0].ObjectId == 0 {
		t.Fatal("Object ID not found")
	}

	if len(result.Games[0].Name) == 0 {
		t.Fatal("Name not found")
	}

	if len(result.Games[0].Published) == 0 {
		t.Fatal("Year not found")
	}
}
