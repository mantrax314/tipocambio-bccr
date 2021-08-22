package tipocambiobccr

import "encoding/xml"

type indicadorEcoXML struct {
	XMLName      xml.Name `xml:"Datos_de_INGC011_CAT_INDICADORECONOMIC"`
	CodIndicador struct {
		CodIndInterno string  `xml:"COD_INDICADORINTERNO"`
		NumValor      float64 `xml:"NUM_VALOR"`
	} `xml:"INGC011_CAT_INDICADORECONOMIC"`
}
