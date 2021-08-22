package tipocambiobccr

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	nurl "net/url"
	"testing"
	"time"
)

type serviceSuite struct {
	email                            string
	token                            string
	name                             string
	sBCCRSvc                         BCCRSvc
	TestGetCurrentDollarSellXML      string
	TestGetCurrentDollarBuyXML       string
	TestGetCurrentEuroPriceXML       string
	TestGetIndicadorNumValorErrorXML string
	suite.Suite
}

func (suite *serviceSuite) SetupTest() {
	suite.email = "email@gmail.com"
	suite.token = "tokensample"
	suite.name = "name sample"
	suite.sBCCRSvc = BCCRSvc{
		email: suite.email,
		token: suite.token,
		name:  suite.name,
	}
	suite.TestGetCurrentDollarSellXML = `<?xml version="1.0" encoding="utf-8"?>
<Datos_de_INGC011_CAT_INDICADORECONOMIC>
  <INGC011_CAT_INDICADORECONOMIC>
    <COD_INDICADORINTERNO>318</COD_INDICADORINTERNO>
    <DES_FECHA>2021-08-21T00:00:00-06:00</DES_FECHA>
    <NUM_VALOR>122.85000000</NUM_VALOR>
  </INGC011_CAT_INDICADORECONOMIC>
</Datos_de_INGC011_CAT_INDICADORECONOMIC>`
	suite.TestGetCurrentDollarBuyXML = `<?xml version="1.0" encoding="utf-8"?>
<Datos_de_INGC011_CAT_INDICADORECONOMIC>
  <INGC011_CAT_INDICADORECONOMIC>
    <COD_INDICADORINTERNO>318</COD_INDICADORINTERNO>
    <DES_FECHA>2021-08-21T00:00:00-06:00</DES_FECHA>
    <NUM_VALOR>123.55000000</NUM_VALOR>
  </INGC011_CAT_INDICADORECONOMIC>
</Datos_de_INGC011_CAT_INDICADORECONOMIC>`
	suite.TestGetCurrentEuroPriceXML = `<?xml version="1.0" encoding="utf-8"?>
<Datos_de_INGC011_CAT_INDICADORECONOMIC>
  <INGC011_CAT_INDICADORECONOMIC>
    <COD_INDICADORINTERNO>318</COD_INDICADORINTERNO>
    <DES_FECHA>2021-08-21T00:00:00-06:00</DES_FECHA>
    <NUM_VALOR>1.17700000</NUM_VALOR>
  </INGC011_CAT_INDICADORECONOMIC>
</Datos_de_INGC011_CAT_INDICADORECONOMIC>`
	suite.TestGetIndicadorNumValorErrorXML = `<?xml version="1.0" encoding="utf-8"?>
<string xmlns="http://ws.sdde.bccr.fi.cr">Ocurrió un error: Formato incorrecto en la fecha de inicio. / An error occurred: The begin date format is not correct.</string>`

}

func TestBCCRSvc(t *testing.T) {
	s := new(serviceSuite)
	suite.Run(t, s)
}

func (suite *serviceSuite) TestNewBCCRSvc() {

	bccrSvc, err := NewBCCRSvc(suite.email, suite.token, suite.name)
	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), suite.sBCCRSvc, bccrSvc)

}

func (suite *serviceSuite) TestGetCurrentDollarSell() {
	currentDay := time.Now().Format("02/01/2006")
	getParams := fmt.Sprintf("?Indicador=%d&FechaInicio=%s&FechaFinal=%s&Nombre=%s&SubNiveles=N&CorreoElectronico=%s&Token=%s",
		dollarSellCode,
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(suite.name),
		nurl.QueryEscape(suite.email),
		nurl.QueryEscape(suite.token),
	)
	url := fmt.Sprintf("%s%s", svcURL, getParams)
	httpmock.Activate()
	httpmock.RegisterResponder(http.MethodGet, url, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusOK, suite.TestGetCurrentDollarSellXML), nil
	})
	bccrSvc, err := NewBCCRSvc(suite.email, suite.token, suite.name)
	assert.NoError(suite.T(), err)
	dollarSell, err := bccrSvc.GetCurrentDollarSell()
	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), 122.85, dollarSell)
}

func (suite *serviceSuite) TestGetCurrentDollarBuy() {
	currentDay := time.Now().Format("02/01/2006")
	getParams := fmt.Sprintf("?Indicador=%d&FechaInicio=%s&FechaFinal=%s&Nombre=%s&SubNiveles=N&CorreoElectronico=%s&Token=%s",
		dollarBuyCode,
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(suite.name),
		nurl.QueryEscape(suite.email),
		nurl.QueryEscape(suite.token),
	)
	url := fmt.Sprintf("%s%s", svcURL, getParams)
	httpmock.Activate()
	httpmock.RegisterResponder(http.MethodGet, url, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusOK, suite.TestGetCurrentDollarBuyXML), nil
	})
	bccrSvc, err := NewBCCRSvc(suite.email, suite.token, suite.name)
	assert.NoError(suite.T(), err)
	dollarBuy, err := bccrSvc.GetCurrentDollarBuy()
	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), 123.55, dollarBuy)
}

func (suite *serviceSuite) TestGetCurrentEuroPrice() {
	currentDay := time.Now().Format("02/01/2006")
	getParams := fmt.Sprintf("?Indicador=%d&FechaInicio=%s&FechaFinal=%s&Nombre=%s&SubNiveles=N&CorreoElectronico=%s&Token=%s",
		euroPriceCode,
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(suite.name),
		nurl.QueryEscape(suite.email),
		nurl.QueryEscape(suite.token),
	)
	url := fmt.Sprintf("%s%s", svcURL, getParams)
	httpmock.Activate()
	httpmock.RegisterResponder(http.MethodGet, url, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusOK, suite.TestGetCurrentEuroPriceXML), nil
	})
	bccrSvc, err := NewBCCRSvc(suite.email, suite.token, suite.name)
	assert.NoError(suite.T(), err)
	euroPrice, err := bccrSvc.GetCurrentEuroPrice()
	assert.NoError(suite.T(), err)
	assert.EqualValues(suite.T(), 1.177, euroPrice)
}

func (suite *serviceSuite) TestGetIndicadorNumValorError() {
	currentDay := time.Now().Format("02/01/2006")
	getParams := fmt.Sprintf("?Indicador=%d&FechaInicio=%s&FechaFinal=%s&Nombre=%s&SubNiveles=N&CorreoElectronico=%s&Token=%s",
		euroPriceCode,
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(suite.name),
		nurl.QueryEscape(suite.email),
		nurl.QueryEscape(suite.token),
	)
	url := fmt.Sprintf("%s%s", svcURL, getParams)
	httpmock.Activate()
	httpmock.RegisterResponder(http.MethodGet, url, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusOK, suite.TestGetIndicadorNumValorErrorXML), nil
	})
	bccrSvc, err := NewBCCRSvc(suite.email, suite.token, suite.name)
	assert.NoError(suite.T(), err)
	euroPrice, err := bccrSvc.getIndicadorNumValor(euroPriceCode)
	assert.Error(suite.T(), err)
	assert.EqualValues(suite.T(), 0, euroPrice)
}

func (suite *serviceSuite) TestGetIndicadorNumValorErrorCode() {
	currentDay := time.Now().Format("02/01/2006")
	getParams := fmt.Sprintf("?Indicador=%d&FechaInicio=%s&FechaFinal=%s&Nombre=%s&SubNiveles=N&CorreoElectronico=%s&Token=%s",
		euroPriceCode,
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(currentDay),
		nurl.QueryEscape(suite.name),
		nurl.QueryEscape(suite.email),
		nurl.QueryEscape(suite.token),
	)
	url := fmt.Sprintf("%s%s", svcURL, getParams)
	httpmock.Activate()
	httpmock.RegisterResponder(http.MethodGet, url, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusInternalServerError, suite.TestGetIndicadorNumValorErrorXML), nil
	})
	bccrSvc, err := NewBCCRSvc(suite.email, suite.token, suite.name)
	assert.NoError(suite.T(), err)
	euroPrice, err := bccrSvc.getIndicadorNumValor(euroPriceCode)
	assert.Errorf(suite.T(), err, "error http status 500")
	assert.EqualValues(suite.T(), 0, euroPrice)
}
