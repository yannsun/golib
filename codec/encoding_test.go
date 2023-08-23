package codec

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type ServiceResponse struct {
	XMLName               xml.Name `xml:"serviceResponse"`
	Text                  string   `xml:",chardata"`
	Cas                   string   `xml:"cas,attr"`
	AuthenticationSuccess struct {
		Text       string `xml:",chardata"`
		User       string `xml:"user"`
		Attributes []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name"`
			UserType string `xml:"userType"`
			UserName string `xml:"userName"`
			UserId   string `xml:"userId"`
			OrgId    string `xml:"orgId"`
			Ticket   string `xml:"ticket"`
		} `xml:"attributes"`
	} `xml:"authenticationSuccess"`
	AuthenticationFailure struct {
		Text string `xml:",chardata"`
		Code string `xml:"code,attr"`
	} `xml:"authenticationFailure"`
}

func TestXmlDecode(t *testing.T)  {
	s := `
<cas:serviceResponse xmlns:cas='http://dcloud.fj.sgcc.com.cn'>
<cas:authenticationSuccess>
<cas:user>ssotest</cas:user>
<!-- 新加 start--> 
<cas:attributes>
<cas:name>测试用户</cas:name>
<cas:userType>1</cas:userType>
<cas:userName>ssotest</cas:userName>
<cas:userId>123a0fddce02407488734f82fa88d138</cas:userId>
<cas:orgId>0021990100</cas:orgId>
</cas:attributes>
<!-- 新加 end--> 
<!-- 新加 start--> 
<cas:attributes> 
<cas:ticket></cas:ticket> 
</cas:attributes> 
<!-- 新加 end--> 
</cas:authenticationSuccess>
</cas:serviceResponse>
`
	serviceResponse := ServiceResponse{}
	err := xml.Unmarshal([]byte(s), &serviceResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v",serviceResponse.XMLName)
	t.Logf("%+v",serviceResponse.Text)
	t.Logf("%+v",serviceResponse.Cas)
	t.Logf("%+v",serviceResponse.AuthenticationSuccess.User)
	t.Logf("%+v",serviceResponse.AuthenticationSuccess.Attributes[0].Name)

	s = `
<cas:serviceResponse xmlns:cas='http://www.yale.edu/tp/cas'>
            <cas:authenticationFailure code='INVALID_TICKET'>
                    Ticket &#039;ST-1-4bPYjlIXxkkUBZdbRUWO-cas01.example.org&#039; not recognized
            </cas:authenticationFailure>
        </cas:serviceResponse>
`
	serviceResponse = ServiceResponse{}
	err = xml.Unmarshal([]byte(s), &serviceResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v",serviceResponse.XMLName)
	t.Logf("%+v",serviceResponse.Text)
	t.Logf("%+v",serviceResponse.Cas)
	t.Logf("%+v",serviceResponse.AuthenticationSuccess)
	t.Logf("%+v",serviceResponse.AuthenticationFailure)
	t.Logf("%+v",serviceResponse.AuthenticationFailure.Code)
	t.Logf("%+v",strings.TrimSpace(serviceResponse.AuthenticationFailure.Text))

	url := "http://182.92.159.121:8081/cas/serviceValidate?ticket=ST-1-4bPYjlIXxkkUBZdbRUWO-cas01.example.org&service=https://www.baidu.com/"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	serviceResponse = ServiceResponse{}
	err = xml.Unmarshal(body, &serviceResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v",serviceResponse.XMLName)
	t.Logf("%+v",serviceResponse.Text)
	t.Logf("%+v",serviceResponse.Cas)
	t.Logf("%+v",serviceResponse.AuthenticationSuccess)
	t.Logf("%+v",serviceResponse.AuthenticationFailure)
	t.Logf("%+v",serviceResponse.AuthenticationFailure.Code)
	t.Logf("%+v",strings.TrimSpace(serviceResponse.AuthenticationFailure.Text))

}
