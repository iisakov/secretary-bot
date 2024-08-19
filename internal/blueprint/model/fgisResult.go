package model

type EtaMI struct {
	RegNumber       string `json:"regNumber"`
	MitypeNumber    string `json:"mitypeNumber"`
	MitypeURL       string `json:"mitypeURL"`
	MitypeTitle     string `json:"mitypeTitle"`
	MitypeType      string `json:"mitypeType"`
	Modification    string `json:"modification"`
	ManufactureNum  string `json:"manufactureNum"`
	ManufactureYear int    `json:"manufactureYear"`
	RankCode        string `json:"rankCode"`
	RankTitle       string `json:"rankTitle"`
	SchemaTitle     string `json:"schemaTitle"`
}

type Applicable struct {
	CertNum    string `json:"certNum"`
	StickerNum string `json:"stickerNum"`
	SignPass   bool   `json:"signPass"`
	SignMi     bool   `json:"signMi"`
}

type Inapplicable struct {
	NoticeNum string `json:"noticeNum"`
}

type VriInfo struct {
	Organization string       `json:"organization"`
	SignCipher   string       `json:"signCipher"`
	MiOwner      string       `json:"miOwner"`
	VrfDate      string       `json:"vrfDate"`
	ValidDate    string       `json:"validDate"`
	DocTitle     string       `json:"docTitle"`
	VriType      string       `json:"vriType"`
	Applicable   Applicable   `json:"applicable"`
	Inapplicable Inapplicable `json:"inapplicable"`
}

type SingleMI struct {
	MitypeNumber   string `json:"mitypeNumber"`
	MitypeType     string `json:"mitypeType"`
	MitypeTitle    string `json:"mitypeTitle"`
	MitypeURL      string `json:"mitypeURL"`
	ManufactureNum string `json:"manufactureNum"`
	Modification   string `json:"modification"`
}

type MiInfo struct {
	EtaMI    EtaMI    `json:"etaMI"`
	SingleMI SingleMI `json:"singleMI"`
}

type Mieta struct {
	RegNumber       string `json:"regNumber"`
	MietaURL        string `json:"mietaURL"`
	MitypeNumber    string `json:"mitypeNumber"`
	MitypeURL       string `json:"mitypeURL"`
	MitypeTitle     string `json:"mitypeTitle"`
	Notation        string `json:"notation"`
	Modification    string `json:"modification"`
	ManufactureNum  string `json:"manufactureNum"`
	ManufactureYear int    `json:"manufactureYear"`
	RankCode        string `json:"rankCode"`
	RankTitle       string `json:"rankTitle"`
	SchemaTitle     string `json:"schemaTitle"`
}

type Mis struct {
	MitypeNumber string `json:"mitypeNumber"`
	MitypeTitle  string `json:"mitypeTitle"`
	MitypeURL    string `json:"mitypeURL"`
	Number       string `json:"number"`
}

type Uve struct {
	Number string `json:"number"`
	Title  string `json:"title"`
	UveURL string `json:"uveURL"`
}

type Means struct {
	Mieta []Mieta `json:"mieta"`
	Mis   []Mis   `json:"mis"`
	Uve   []Uve   `json:"uve"`
}

type Protocol struct {
	Title    string `json:"title"`
	Mimetype string `json:"mimetype"`
	Filename string `json:"filename"`
	Doc_id   int    `json:"doc_id"`
}

type Info struct {
	BriefIndicator       bool     `json:"briefIndicator"`
	BriefCharacteristics string   `json:"briefCharacteristics"`
	Protocol             Protocol `json:"protocol"`
}

type Result struct {
	MiInfo  MiInfo  `json:"miInfo"`
	VriInfo VriInfo `json:"vriInfo"`
	Means   Means   `json:"means"`
	Info    Info    `json:"info"`
	Url     string
}

type Response struct {
	Result Result `json:"result"`
}

type Cert struct {
	Organization         string   `json:"organization"`
	CertNum              string   `json:"certNum"`
	ValidDate            string   `json:"validDate"`
	MI                   string   `json:"mi"`
	ManufactureNum       string   `json:"manufactureNum"`
	DocTitle             string   `json:"docTitle"`
	Mieta                []string `json:"mieta"`
	Url                  string   `json:"url"`
	BriefCharacteristics string   `json:"riefCharacteristics"`
	RecordsNum           string   `json:"recordsNum"`
	VrfDate              string   `json:"vrfDate"`
}
