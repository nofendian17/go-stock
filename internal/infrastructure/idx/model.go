package idx

type StockListResponse struct {
	Draw            int             `json:"draw"`
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	StockListData   []StockListData `json:"data"`
}

type StockListData struct {
	Code        string  `json:"Code"`
	Name        string  `json:"Name"`
	Share       float64 `json:"Shares"`
	ListingDate string  `json:"ListingDate"`
	Board       string  `json:"ListingBoard"`
}

type BrokerListResponse struct {
	Draw            int              `json:"draw"`
	RecordsTotal    int              `json:"recordsTotal"`
	RecordsFiltered int              `json:"recordsFiltered"`
	BrokerListData  []BrokerListData `json:"data"`
}

type BrokerListData struct {
	Code    string `json:"Code"`
	Name    string `json:"Name"`
	License string `json:"License"`
}

type CompanyProfileResponse struct {
	ResultCount    int              `json:"ResultCount"`
	Search         Search           `json:"Search"`
	Profiles       []Profiles       `json:"Profiles"`
	Sekretaris     []Sekretaris     `json:"Sekretaris"`
	Direktur       []Direktur       `json:"Direktur"`
	Komisaris      []Komisaris      `json:"Komisaris"`
	KomiteAudit    []KomiteAudit    `json:"KomiteAudit"`
	PemegangSaham  []PemegangSaham  `json:"PemegangSaham"`
	AnakPerusahaan []AnakPerusahaan `json:"AnakPerusahaan"`
	Kap            []any            `json:"KAP"`
	Dividen        []Dividen        `json:"Dividen"`
	BondsAndSukuk  []any            `json:"BondsAndSukuk"`
	IssuedBond     []any            `json:"IssuedBond"`
}
type Search struct {
	ReportType string `json:"ReportType"`
	KodeEmiten string `json:"KodeEmiten"`
	Year       string `json:"Year"`
	SortColumn string `json:"SortColumn"`
	SortOrder  string `json:"SortOrder"`
	EmitenType any    `json:"EmitenType"`
	Periode    string `json:"Periode"`
	Indexfrom  int    `json:"indexfrom"`
	Pagesize   int    `json:"pagesize"`
}
type Profiles struct {
	Alamat             string `json:"Alamat"`
	Bae                string `json:"BAE"`
	DataID             int    `json:"DataID"`
	Divisi             any    `json:"Divisi"`
	EfekEmitenEBA      bool   `json:"EfekEmiten_EBA"`
	EfekEmitenETF      bool   `json:"EfekEmiten_ETF"`
	EfekEmitenObligasi bool   `json:"EfekEmiten_Obligasi"`
	EfekEmitenSaham    bool   `json:"EfekEmiten_Saham"`
	EfekEmitenSPEI     bool   `json:"EfekEmiten_SPEI"`
	Industri           string `json:"Industri"`
	SubIndustri        string `json:"SubIndustri"`
	Email              string `json:"Email"`
	Fax                string `json:"Fax"`
	ID                 int    `json:"id"`
	JenisEmiten        any    `json:"JenisEmiten"`
	KegiatanUsahaUtama string `json:"KegiatanUsahaUtama"`
	KodeDivisi         any    `json:"KodeDivisi"`
	KodeEmiten         string `json:"KodeEmiten"`
	NamaEmiten         string `json:"NamaEmiten"`
	Npkp               string `json:"NPKP"`
	Npwp               string `json:"NPWP"`
	PapanPencatatan    string `json:"PapanPencatatan"`
	Sektor             string `json:"Sektor"`
	SubSektor          string `json:"SubSektor"`
	TanggalPencatatan  string `json:"TanggalPencatatan"`
	Telepon            string `json:"Telepon"`
	Website            string `json:"Website"`
	Status             int    `json:"Status"`
	Logo               string `json:"Logo"`
}
type Sekretaris struct {
	Nama    string `json:"Nama"`
	Telepon string `json:"Telepon"`
	Website string `json:"Website"`
	Email   string `json:"Email"`
	Fax     string `json:"Fax"`
	Hp      string `json:"HP"`
}
type Direktur struct {
	Nama     string `json:"Nama"`
	Jabatan  string `json:"Jabatan"`
	Afiliasi bool   `json:"Afiliasi"`
}
type Komisaris struct {
	Nama       string `json:"Nama"`
	Jabatan    string `json:"Jabatan"`
	Independen bool   `json:"Independen"`
}
type KomiteAudit struct {
	Jabatan string `json:"Jabatan"`
	Nama    string `json:"Nama"`
}
type PemegangSaham struct {
	Jumlah     float64 `json:"Jumlah"`
	Kategori   string  `json:"Kategori"`
	Nama       string  `json:"Nama"`
	Pengendali bool    `json:"Pengendali"`
	Persentase float64 `json:"Persentase"`
}
type AnakPerusahaan struct {
	BidangUsaha   string  `json:"BidangUsaha"`
	JumlahAset    float64 `json:"JumlahAset"`
	Lokasi        string  `json:"Lokasi"`
	MataUang      string  `json:"MataUang"`
	Nama          string  `json:"Nama"`
	Persentase    float64 `json:"Persentase"`
	Satuan        string  `json:"Satuan"`
	StatusOperasi string  `json:"StatusOperasi"`
	TahunKomersil string  `json:"TahunKomersil"`
}
type Dividen struct {
	Nama                         string  `json:"Nama"`
	Jenis                        string  `json:"Jenis"`
	TahunBuku                    string  `json:"TahunBuku"`
	TotalSahamBonus              float64 `json:"TotalSahamBonus"`
	CashDividenPerSahamMU        string  `json:"CashDividenPerSahamMU"`
	CashDividenPerSaham          float64 `json:"CashDividenPerSaham"`
	TanggalCum                   string  `json:"TanggalCum"`
	TanggalExRegulerDanNegosiasi string  `json:"TanggalExRegulerDanNegosiasi"`
	TanggalDPS                   string  `json:"TanggalDPS"`
	TanggalPembayaran            string  `json:"TanggalPembayaran"`
	Rasio1                       int     `json:"Rasio1"`
	Rasio2                       int     `json:"Rasio2"`
	CashDividenTotalMU           string  `json:"CashDividenTotalMU"`
	CashDividenTotal             float64 `json:"CashDividenTotal"`
}

type StockSummaryListResponse struct {
	Draw                 int                    `json:"draw"`
	RecordsTotal         int                    `json:"recordsTotal"`
	RecordsFiltered      int                    `json:"recordsFiltered"`
	StockSummaryListData []StockSummaryListData `json:"data"`
}

type StockSummaryListData struct {
	No                  int     `json:"No"`
	IDStockSummary      int     `json:"IDStockSummary"`
	Date                string  `json:"Date"`
	StockCode           string  `json:"StockCode"`
	StockName           string  `json:"StockName"`
	Remarks             string  `json:"Remarks"`
	Previous            float64 `json:"Previous"`
	OpenPrice           float64 `json:"OpenPrice"`
	FirstTrade          float64 `json:"FirstTrade"`
	High                float64 `json:"High"`
	Low                 float64 `json:"Low"`
	Close               float64 `json:"Close"`
	Change              float64 `json:"Change"`
	Volume              float64 `json:"Volume"`
	Value               float64 `json:"Value"`
	Frequency           float64 `json:"Frequency"`
	IndexIndividual     float64 `json:"IndexIndividual"`
	Offer               float64 `json:"Offer"`
	OfferVolume         float64 `json:"OfferVolume"`
	Bid                 float64 `json:"Bid"`
	BidVolume           float64 `json:"BidVolume"`
	ListedShares        float64 `json:"ListedShares"`
	TradebleShares      float64 `json:"TradebleShares"`
	WeightForIndex      float64 `json:"WeightForIndex"`
	ForeignSell         float64 `json:"ForeignSell"`
	ForeignBuy          float64 `json:"ForeignBuy"`
	DelistingDate       string  `json:"DelistingDate"`
	NonRegularVolume    float64 `json:"NonRegularVolume"`
	NonRegularValue     float64 `json:"NonRegularValue"`
	NonRegularFrequency float64 `json:"NonRegularFrequency"`
	Persen              any     `json:"persen"`
	Percentage          any     `json:"percentage"`
}

type FinancialReportResponse struct {
	Search      Search    `json:"Search"`
	ResultCount int       `json:"ResultCount"`
	Results     []Results `json:"Results"`
}

type Attachments struct {
	EmitenCode   string `json:"Emiten_Code"`
	FileID       string `json:"File_ID"`
	FileModified string `json:"File_Modified"`
	FileName     string `json:"File_Name"`
	FilePath     string `json:"File_Path"`
	FileSize     int    `json:"File_Size"`
	FileType     string `json:"File_Type"`
	ReportPeriod string `json:"Report_Period"`
	ReportType   string `json:"Report_Type"`
	ReportYear   string `json:"Report_Year"`
	NamaEmiten   string `json:"NamaEmiten"`
}
type Results struct {
	KodeEmiten   string        `json:"KodeEmiten"`
	FileModified string        `json:"File_Modified"`
	ReportPeriod string        `json:"Report_Period"`
	ReportYear   string        `json:"Report_Year"`
	NamaEmiten   string        `json:"NamaEmiten"`
	Attachments  []Attachments `json:"Attachments"`
}
