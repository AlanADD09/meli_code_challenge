package tests

// func TestCSVReader_ParseFile(t *testing.T) {
// 	csvData := `site,id
// MLA,839084261
// MLA,813840673`
// 	reader := file_processor.NewCSVReader(',')
// 	result, err := reader.ParseFile(strings.NewReader(csvData))
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
