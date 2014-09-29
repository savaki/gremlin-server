package accounts

import (
	"fmt"
	"github.com/savaki/stormpath-go/auth"
	. "github.com/savaki/stormpath-go/types"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"time"
)

func TestAccounts(t *testing.T) {
	if _, err := auth.EnvAuth(); err != nil {
		return
	}
	if os.Getenv("STORMPATH_URL") == "" {
		return
	}

	var apiKey auth.ApiKey
	var accounts *Accounts
	var err error
	var account *Account
	var saved *Account
	var email string
	var accountUrl Url

	Convey("Given a StormPath client", t, func() {
		apiKey, _ = auth.EnvAuth()
		accounts = FromUrl(apiKey, os.Getenv("STORMPATH_URL"))

		So(err, ShouldBeNil)
		So(accounts, ShouldNotBeNil)

		Convey("When I create an account", func() {
			email = fmt.Sprintf("joe.public-%d@gmail.com", time.Now().Unix())
			account = &Account{
				Username:  email,
				Email:     email,
				GivenName: "Joe",
				Surname:   "Public",
				Password:  "Pass123!",
			}
			saved, err = accounts.Create(*account)

			Convey("Then I expect no errors", func() {
				So(err, ShouldBeNil)
			})

			Convey("And I expect a valid saved account back", func() {
				So(saved, ShouldNotBeNil)
				So(saved.Href, ShouldNotEqual, "")
			})

			Convey("When I attempt to login with this account", func() {
				accountUrl, err = accounts.Login(account.Username, account.Password)

				Convey("Then I expect no errors", func() {
					So(err, ShouldBeNil)
				})

				Convey("And I expect a valid account url back", func() {
					So(accountUrl, ShouldNotBeNil)
				})
			})

			Convey("When I attempt to search for this account", func() {
				results, err := accounts.Search(account.Email)
				So(err, ShouldBeNil)
				So(len(results), ShouldEqual, 1)
				So(results[0].Email, ShouldEqual, account.Email)
				So(results[0].Username, ShouldEqual, account.Username)
			})
		})

		Reset(func() {
			if accounts != nil && saved != nil {
				err = accounts.Delete(saved)
				if err != nil {
					fmt.Println(err)
				}
			}
		})
	})
}
