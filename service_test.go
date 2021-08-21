package tipocambiobccr

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type serviceSuite struct {
	email    string
	token    string
	name     string
	sBCCRSvc BCCRSvc
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
