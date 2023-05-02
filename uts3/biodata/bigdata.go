package biodata

type BigData struct {
	Data1 DataDiri
	Data2 DataDiri
}

func ShowBigData() BigData {

	allData := BigData{
		Data1: ShowDataDiri("gusti", 1234, "Malang"),
		Data2: ShowDataDiri("Pangestu", 321, "Kabupaten Malang"),
	}
	return allData
}
