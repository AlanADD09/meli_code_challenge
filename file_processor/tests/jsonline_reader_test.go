package tests

// func TestJSONLinesReader_ParseFile(t *testing.T) {
// 	jsonData := `{"site":"MLA","id":"839084261"}
// {"site":"MLA","id":"813840673"}`
// 	reader := file_processor.NewJSONLinesReader()
// 	result, err := reader.ParseFile(strings.NewReader(jsonData))
// 	if err != nil {
// 		t.Errorf("parseFile failed: %v", err)
// 	}

// 	expected := []file_processor.SiteID{
// 		{Site: "MLA", ID: "839084261"},
// 		{Site: "MLA", ID: "813840673"},
// 	}

// 	if len(result) != len(expected) {
// 		t.Errorf("Expected %d records, got %d", len(expected), len(result))
// 	}

// 	for i, record := range result {
// 		if record != expected[i] {
// 			t.Errorf("Expected record %v, got %v", expected[i], record)
// 		}
// 	}
// }
