package models

// BusinessTypeListRes ..
type BusinessTypeListRes struct {
	TipeBisnis []DataModule `json:"tipe_bisnis"`
}

// CategoryTypeListRes ..
type CategoryTypeListRes struct {
	KategoriBisnis []DataModule `json:"kategori_bisnis"`
}

// AcquitisionsDropdownListRes ..
type AcquitisionsDropdownListRes struct {
	LokasiBisnis        []DataModule `json:"lokasi_bisnis"`
	JenisLokasiBisnis   []DataModule `json:"jenis_lokasi_bisnis"`
	JamOperational      []DataModule `json:"jam_operasional"`
	JamKunjunganTerbaik []DataModule `json:"jam_kunjungan_terbaik"`
}

// DataModule ..
type DataModule struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
