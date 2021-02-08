package users

import (
	"testing"

	firebase "firebase.google.com/go"
	"github.com/ali2210/wizdwarf/structs/users/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var customer model.Vistors = model.Vistors{
	Id:       "ab324k",
	Name:     "john",
	Email:    "john@gmail.com",
	Password: "1234",
	FName:    "dvy",
	City:     "calforina",
	Zip:      "94300",
	Address:  "west moon",
	LAddress: "block 24, street 1",
	Country:  "usa",
	Eve:      false,
}
var client DBFirestore = &cloud_data{}

func profileDBTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Profile DB Test cases")
}

var _ = Describe("Profile Data store interface test..", func() {
	Context("Data store Objects state", func() {
		It("Data store Object", func() {
			client = NewCloudInstance()
			Expect(NewCloudInstance()).Should(BeNil())
		})
		It("Data storage", func() {
			Expect(client.SaveData(&customer, &firebase.App{})).Should(BeNil())
		})
		It("Data find", func() {

			Expect(client.FindAllData(&firebase.App{}, customer.Email, customer.Password)).Should(BeNil())
		})
		It("Data group", func() {
			Expect(client.ToFindByGroupSet(customer.Id, customer.Email, &firebase.App{})).Should(BeNil())
		})
		It("Data update", func() {
			Profile := model.UpdateProfile{
				Id:           "ab324k",
				FirstName:    "john",
				LastName:     "Dvy",
				Phone:        "+01-26531",
				HouseAddress: "west moon",
				SubAddress:   "block 24, street 1",
				Country:      "usa",
				Zip:          "94300",
				Male:         true,
				Email:        "johnDevy@coke.com",
				Twitter:      "https://twitter.com/john",
			}
			Expect(client.UpdateProfiles(&firebase.App{}, &Profile)).Should(BeNil())
		})
		It("Data get", func() {
			Expect(client.GetProfile(&firebase.App{}, customer.Id)).Should(BeNil())
		})
	})

})
