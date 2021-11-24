package models

import "time"

// AcquisitionsMerchantDetailReq ..
type AcquisitionsMerchantDetailReq struct {
	MerchantPhone 	string `json:"merchant_phone"`
}

// AcquisitionsMerchantDetailRes ..
type AcquisitionsMerchantDetailRes struct {
	PhotoKTP            	string		`json:"photo_ktp"`
	PhotoLocation           string		`json:"photo_location"`
	PhotoLocation2         	string		`json:"photo_location_2"`
	PhotoSelfie			 	string		`json:"photo_selfie"`
	PhotoSign			 	string		`json:"photo_sign"`
	MerchantGroupName       string		`json:"merchant_group_name"`
	//MerchantGroup		    string		`json:"merchantgroup"`
	PriorityLevel		    string		`json:"priority_level"`

	StoreName            	string		`json:"store_name"`
	StoreJenisUsaha         string		`json:"store_jenis_usaha"`
	StoreAlamat         	string		`json:"store_alamat"`
	StoreKelurahan			string		`json:"store_kelurahan"`
	StoreKecamatan			string		`json:"store_kecamatan"`
	StoreJamOperasional	    string		`json:"store_jam_operasional"`

	StoreJenisLokasiBisnis  string		`json:"store_jenis_lokasi_bisnis"`
	StoreKabupatenKota      string		`json:"store_kabupaten_kota"`
	StorePostalCode         string		`json:"store_postal_code"`
	StoreProvince			string		`json:"store_provinsi"`
	StoreLatitude			string		`json:"store_latitude"`
	StoreLongitude		    string		`json:"store_longitude"`

	StoreLokasiBisnis       string		`json:"store_lokasi_bisnis"`
	StorePhoneNumber        string		`json:"store_phone_number"`
	AgentId		         	string		`json:"agent_id"`
	AgentName				string		`json:"agent_name"`
	AgentCompanyID			string		`json:"agent_company_id"`
	AgentPhoneNumber		string		`json:"agent_phone_number"`
	OwnerAddress			string		`json:"owner_address"`
	OwnerFirstName		    string		`json:"owner_firstname"`

	OwnerJenisKelamin       string		`json:"owner_jenis_kelamin"`
	OwnerKabupatenKota      string		`json:"owner_kabupaten_kota"`
	OwnerKecamatan         	string		`json:"owner_kecamatan"`
	OwnerKelurahan			string		`json:"owner_kelurahan"`
	OwnerKodePos			string		`json:"owner_kode_pos"`
	OwnerLastName		    string		`json:"owner_lastname"`

	OwnerNamaGadisIbuKandung	string		`json:"owner_nama_gadis_ibu_kandung"`
	OwnerNoId		        string		`json:"owner_no_id"`
	OwnerNoTelf         	string		`json:"owner_no_telf"`
	OwnerNoTelfLainnya		string		`json:"owner_no_telf_lainnya"`
	OwnerPekerjaan			string		`json:"owner_pekerjaan"`
	OwnerProvinsi		    string		`json:"owner_provinsi"`

	OwnerRt	            	string		`json:"owner_rt"`
	OwnerRW         		string		`json:"owner_rw"`
	OwnerTanggalLahir       string		`json:"owner_tanggal_lahir"`
	OwnerTempatTinggal		string		`json:"owner_tempat_lahir"`
	OwnerTglWxpiredId		string		`json:"owner_tgl_expired_id"`
	OwnerTipeId			    string		`json:"owner_tipe_id"`

	OwnerTitle            	string		`json:"owner_title"`
	DeviceType		        string		`json:"device_type"`
	MetodePembayaran      	string		`json:"metode_pembayaran"`
	DeviceGroup				string		`json:"device_group"`
	DeviceBrand				string		`json:"device_brand"`
	OutletId			    string		`json:"outlet_id"`

	TerminalPhoneNumber     string		`json:"terminal_phone_number"`
	TerminalProvider		string		`json:"terminal_provider"`
	InstitutionId			string		`json:"institution_id"`
	Notes				    string		`json:"notes"`
	MPAN				    string		`json:"mpan"`
	MerchantPAN				string		`json:"merchant_pan"`
	MerchantOutletId		string		`json:"merchant_outlet_id"`

	SalesID					string 		`json:"sales_id"`
	KategoriBisnis		    string		`json:"kategori_bisnis"`
	NMID		    		string		`json:"nmid"`
	Level		    		string		`json:"level"`
	ExistingQRValue	   		string		`json:"existing_qr_value"`
	Category	   			string		`json:"category"`
	StoreNamePreprinted	   	string		`json:"store_name_preprinted"`
	PhotoLocationLeft	   	string		`json:"photo_location_left"`
	PhotoLocationRight	   	string		`json:"photo_location_right"`
	PhotoQrPreprinted	   	string		`json:"foto_preprinted"`
	PartnerCustomerId	   	string		`json:"partner_customer_id"`
	Patokan				   	string		`json:"patokan"`

	AcquitisionDate			time.Time	`json:"acquitision_date"`
	SrId                    string      `json:"sr_id"`
}
