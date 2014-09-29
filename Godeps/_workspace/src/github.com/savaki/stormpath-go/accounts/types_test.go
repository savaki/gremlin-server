package accounts

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func LiveJsonUnmarshal(t *testing.T) {
	Convey("Should unmarshall Application json", t, func() {
		content := `{
    "href": "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2",
    "username": "joe.public-1411656582@gmail.com",
    "email": "joe.public-1411656582@gmail.com",
    "givenName": "Joe",
    "middleName": null,
    "surname": "Public",
    "fullName": "Joe Public",
    "status": "ENABLED",
    "emailVerificationToken": null,
    "customData": {
        "href": "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2/customData"
    },
    "providerData": {
        "href": "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2/providerData"
    },
    "directory": {
        "href": "https://api.stormpath.com/v1/directories/CfJokgltvyHOPsPGCXqQq"
    },
    "tenant": {
        "href": "https://api.stormpath.com/v1/tenants/4OnjPr5YItrYEYht3qFKPD"
    },
    "groups": {
        "href": "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2/groups"
    },
    "groupMemberships": {
        "href": "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2/groupMemberships"
    },
    "apiKeys": {
        "href": "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2/apiKeys"
    }
}`

		account := &Account{}
		err := json.NewDecoder(strings.NewReader(content)).Decode(account)

		So(err, ShouldBeNil)
		So(account, ShouldNotBeNil)
		So(account.Href, ShouldEqual, "https://api.stormpath.com/v1/accounts/1Th2zTIbFr8KshmSCaIRB2")
	})
}
