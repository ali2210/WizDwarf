/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */
package piplines

import (
	"net/http"
	"testing"
	"time"

	skynet "github.com/SkynetLabs/go-skynet/v2"
	"github.com/ali2210/wizdwarf/other/users"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// piplines tests suits
const pipline_describe = "Pipline testing started"
const date_describe = "piplines testing started now"
const date_infix = "date tests started"
const month_infix = "month tests started"
const year_infix = "year tests started"
const time_key = "time tests started"
const skynet_protocol = "skynet tests started"
const time_spec_desc = "time utc tests started"
const img_file_desc = "pictures tests started"
const skynet_object = "skynet tests started"
const image_format = "image tests case insensitive"
const bytes2alpha = "bytes tests case insensitive"
const anscitest = "anscitest tests started"
const server_file = "server_file tests started"
const server_files = "file paths tests "
const updateinfo = "update info test started"
const updateinfotest = "user updated request proceed"
const addinfo = "add user test started"
const adduser = "new user created"
const db_credientials = "db credentials test started"
const db_credientials_test = "db credentials test started"

// pipline test suit for specific tests
const specify = " must not be empty"
const specify_state = " validation ...."

func PiplinesBDDTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, pipline_describe)
}

var _ = Describe(date_describe, func() {
	Context(date_infix, func() {
		s := "20"
		It(date_infix+specify, func() {
			// empty date
			Expect(Date(s)).Should(BeEmpty())
		})
		It(date_infix+specify_state, func() {
			// date have some numeric values
			Expect(Date(s)).ShouldNot(BeEmpty())
		})
	})
	Context(month_infix, func() {

		s := "09"
		It(month_infix+specify, func() {
			// empty month
			Expect(Date(s)).Should(BeEmpty())
		})
		It(month_infix+specify_state, func() {
			// month have some numeric values
			Expect(Date(s)).ShouldNot(BeEmpty())
		})
	})
	Context(year_infix, func() {

		s := "2021"
		It(year_infix+specify, func() {
			// empty year
			Expect(Date(s)).Should(BeEmpty())
		})
		It(year_infix+specify_state, func() {
			// year have some numeric values
			Expect(Date(s)).ShouldNot(BeEmpty())
		})
	})
	Context(time_key, func() {
		It(time_spec_desc, func() {
			// time according utc format
			Expect(GetToday(2021, time.Month(9), 20)).Should(BeEmpty())
		})
	})
	Context(image_format, func() {
		s := "mickeymouse.jpeg"
		It(img_file_desc, func() {
			//  image format
			Expect(ParseTags(s)).Should(BeEmpty())
		})
	})
	Context(skynet_protocol, func() {
		It(skynet_object, func() {
			Expect(SiaObjectStorage(skynet.New(), "mickymouse.jpeg")).Should(BeFalse())
		})
	})
	Context(anscitest, func() {
		It(bytes2alpha, func() {
			Expect(ToRunes('a')).Should(BeEmpty())
		})
	})
	Context(server_file, func() {
		It(server_files, func() {
			Expect(Open_SFiles("app_data/", "mickeymouse.jpeg")).ShouldNot(BeNil())
		})
	})
	Context(updateinfo, func() {
		It(updateinfotest, func() {
			Expect(UpdateProfileInfo(&users.Visitors{})).Should(BeTrue())
		})
	})
	Context(addinfo, func() {
		It(adduser, func() {
			var w http.ResponseWriter
			r := &http.Request{}
			Expect(AddNewProfile(w, r, users.Visitors{}, "0x1223...")).Should(BeTrue())
		})
	})
	Context(db_credientials, func() {
		It(db_credientials_test, func() {
			var w http.ResponseWriter
			r := &http.Request{}
			Expect(Firebase_Gatekeeper(w, r, users.Visitors{})).Should(BeEmpty())
		})
	})
})
