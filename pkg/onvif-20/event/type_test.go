package event

import (
	"encoding/xml"
	"testing"

	"test-func/pkg/onvif-20/xsd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var eventPropertiesData = []byte(`
        <tev:GetEventPropertiesResponse>
            <tev:TopicNamespaceLocation>http://www.onvif.org/onvif/ver10/topics/topicns.xml</tev:TopicNamespaceLocation>
            <wsnt:FixedTopicSet>true</wsnt:FixedTopicSet>
            <wstop:TopicSet>
                <tns1:UserAlarm wstop:topic="true">
                    <tnshik:IllegalAccess wstop:topic="true"/>
                </tns1:UserAlarm>
                <tns1:RuleEngine wstop:topic="true">
                    <CellMotionDetector wstop:topic="true">
                        <Motion wstop:topic="true">
                            <tt:MessageDescription IsProperty="true">
                                <tt:Source>
                                    <tt:SimpleItemDescription Name="VideoSourceConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="VideoAnalyticsConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="Rule" Type="xs:string"/>
                                </tt:Source>
                                <tt:Data>
                                    <tt:SimpleItemDescription Name="IsMotion" Type="xs:boolean"/>
                                </tt:Data>
                            </tt:MessageDescription>
                        </Motion>
                    </CellMotionDetector>
                    <TamperDetector wstop:topic="true">
                        <Tamper wstop:topic="true">
                            <tt:MessageDescription IsProperty="true">
                                <tt:Source>
                                    <tt:SimpleItemDescription Name="VideoSourceConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="VideoAnalyticsConfigurationToken" Type="tt:ReferenceToken"/>
                                    <tt:SimpleItemDescription Name="Rule" Type="xs:string"/>
                                </tt:Source>
                                <tt:Data>
                                    <tt:SimpleItemDescription Name="IsTamper" Type="xs:boolean"/>
                                </tt:Data>
                            </tt:MessageDescription>
                        </Tamper>
                    </TamperDetector>
                </tns1:RuleEngine>
            </wstop:TopicSet>
            <tev:MessageContentFilterDialect>http://www.onvif.org/ver10/tev/messageContentFilter/ItemFilter</tev:MessageContentFilterDialect>
            <tev:MessageContentSchemaLocation>http://www.onvif.org/onvif/ver10/schema/onvif.xsd</tev:MessageContentSchemaLocation>
        </tev:GetEventPropertiesResponse>
`)

func TestEventPropertiesUnmarshalXML(t *testing.T) {
	res := GetEventPropertiesResponse{}
	err := xml.Unmarshal(eventPropertiesData, &res)
	require.NoError(t, err)
	assert.Equal(t, FixedTopicSet(true), *res.FixedTopicSet)
	assert.Equal(t, xsd.AnyURI("http://www.onvif.org/ver10/tev/messageContentFilter/ItemFilter"), *res.MessageContentFilterDialect)
	assert.Equal(t, xsd.AnyURI("http://www.onvif.org/onvif/ver10/schema/onvif.xsd"), *res.MessageContentSchemaLocation)
	userAlarm, exists := map[string]interface{}(*res.TopicSet)["tns1:UserAlarm"]
	assert.True(t, exists)
	_, exists = (userAlarm).(map[string]interface{})["tnshik:IllegalAccess"]
	assert.True(t, exists)
	ruleEngine, exists := map[string]interface{}(*res.TopicSet)["tns1:RuleEngine"]
	assert.True(t, exists)
	tamperDetector, exists := (ruleEngine).(map[string]interface{})["TamperDetector"]
	assert.True(t, exists)
	tamper, exists := (tamperDetector).(map[string]interface{})["Tamper"]
	assert.True(t, exists)
	_, exists = (tamper).(map[string]interface{})["tt:MessageDescription"]
	assert.True(t, exists)
}
