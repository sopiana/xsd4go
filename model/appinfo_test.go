package model

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnmarshallAppInfoUnknownAttr(t *testing.T) {
	data:=[]byte(`<xs:appinfo ns:source="test" source1="test2">Application Information</xs:appinfo>`)
	appInfo := New_AppInfo(&Schema{})
	err:=xml.Unmarshal(data, &appInfo)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(appInfo.root.errors))
}

func Test_UnmarshallAppInfoSyntaxError(t *testing.T){
	data:=[]byte(`<xs:appinfo>Application Information</xs:appinfo`)
	appInfo := New_AppInfo(&Schema{})
	err:=xml.Unmarshal(data, &appInfo)
	assert.NotNil(t,  err)
	assert.Equal(t, 1,len(appInfo.root.errors))
}

func Test_UnmarshallAppInfo(t *testing.T){
	data:=[]byte(`<xs:appinfo source="test">Application Information <otherElem>My Data</otherElem> Saja</xs:appinfo>`)
	appInfo := New_AppInfo(&Schema{})
	err:=xml.Unmarshal(data, &appInfo)
	assert.Nil(t,  err)
	assert.Equal(t, 0,len(appInfo.root.errors))
	assert.Equal(t, "test", appInfo.source)
	assert.Equal(t, "Application Information My Data Saja", appInfo.content)

	appInfo1 := New_AppInfo(&Schema{})
	err = xml.Unmarshal([]byte(appInfo.X(0)), &appInfo1)
	assert.Nil(t, err)
	assert.Equal(t, 0,len(appInfo1.root.errors))
	assert.Equal(t, "test", appInfo1.source)
	assert.Equal(t, "Application Information My Data Saja", appInfo1.content)
}