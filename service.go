package tipocambiobccr

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	nurl "net/url"
	"strings"
	"time"
)

const (
	svcURL = "https://gee.bccr.fi.cr/Indicadores/Suscripciones/WS/wsindicadoreseconomicos.asmx/ObtenerIndicadoresEconomicosXML"

	dollarSellCode = 318
	dollarBuyCode  = 318
	euroPriceCode  = 333
)

// BCCRSvc BCCR Service struct
type BCCRSvc struct {
	email string
	token string
	name  string
}

// NewBCCRSvc Return a new BCCR Service
func NewBCCRSvc(email, token, name string) (BCCRSvc, error) {

	// TODO: Add Errors

	return BCCRSvc{
		email: email,
		token: token,
		name:  name,
	}, nil

}

func fixXML(xmlString string) string {
	fixedXML := strings.ReplaceAll(xmlString, "<string xmlns=\"http://ws.sdde.bccr.fi.cr\">", "")
	fixedXML = strings.ReplaceAll(fixedXML, "&lt;", "<")
	fixedXML = strings.ReplaceAll(fixedXML, "&gt;", ">")
	fixedXML = strings.ReplaceAll(fixedXML, "</string>", "")
	return fixedXML
}

// GetCurrentDollarSell Get Current Dol	lar sell value
func (svc BCCRSvc) GetCurrentDollarSell() (float64, error) {
	numValor, err := svc.getIndicadorNumValor(dollarSellCode)
	return numValor, err
}

// GetCurrentDollarBuy Get Current Dol	lar sell value
func (svc BCCRSvc) GetCurrentDollarBuy() (float64, error) {
	numValor, err := svc.getIndicadorNumValor(dollarBuyCode)
	return numValor, err
}

// GetCurrentEuroPrice Get Current Euro Price
func (svc BCCRSvc) GetCurrentEuroPrice() (float64, error) {
	numValor, err := svc.getIndicadorNumValor(euroPriceCode)
	return numValor, err
}

// getIndicadorNumValor connects to indicadores service to pull response value
func (svc BCCRSvc) getIndicadorNumValor(indicadorCode int) (float64, error) {
	location, err := time.LoadLocation("America/Costa_Rica")
	if err != nil {
		return 0, err
	}
	currentDay := time.Now().In(location).Format("02/01/2006")
	getParams := fmt.Sprintf("?Indicador=%d&FechaInicio=%s&FechaFinal=%s&Nombre=%s&SubNiveles=N&CorreoElectronico=%s&Token=%s",
		indicadorCode,
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(svc.name),
		nurl.QueryEscape(svc.email),
		nurl.QueryEscape(svc.token),
	)
	url := fmt.Sprintf("%s%s", svcURL, getParams)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return 0, err
	}
	res, err := client.Do(req)
	if err != nil || (res != nil && res.StatusCode != http.StatusOK) {
		if res.StatusCode != http.StatusOK {
			return 0, fmt.Errorf("error http status %d", res.StatusCode)
		}
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	stXML := string(body)
	if strings.Contains(stXML, "error") {
		return 0, fmt.Errorf(stXML)
	}
	resSt := fixXML(stXML)
	var ieco indicadorEcoXML
	err = xml.Unmarshal([]byte(resSt), &ieco)
	if err != nil {
		return 0, err
	}

	return ieco.CodIndicador.NumValor, nil

}
