package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const size = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(result.Requests) != size {
		t.Errorf("result should have %d"+"requests;but had %d", size, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("Epected url #%d: %s;but"+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}
	if len(result.Items) != size {
		t.Errorf("result should have %d"+"requests;but had %d", size, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("Epected url #%d: %s;but"+
				"was %s",
				i, city, result.Items[i].(string))
		}
	}
	fmt.Printf("%s\n", contents)

}
