package users

import (
	"testing"
	// "cloud.google.com/go/firestore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)




var customer Visitors = Visitors{
	Id:       "ab324k",
	Name:     "john",
	Email:    "john@gmail.com",
	Password: "1234",
	LastName:    "dvy",
	City:     "calforina",
	Zip:      "94300",
	Address:  "west moon",
	Apparment: "block 24, street 1",
	Country:  "usa",
	Eve:      false,
}
var client DBFirestore = &FirestoreClient{}

func profileDBTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Profile DB Test cases")
}

var _ = Describe("Profile Data store interface test..", func() {
	Context("Data store Objects state", func() {
		It("Data store Object", func() {
			client = NewCloudInstance()
			// Expect(NewCloudInstance()).Should(BeNil())
		})
		It("Data storage", func() {
			//  Expect(client.AddUser(Firestore_Reference(), customer)).Should(BeNil())
		})
		It("Data find", func() {

			// Expect(client.GetDocumentById(GetDBClientRef(), customer)).Should(BeNil())
		})
		It("Data group", func() {
			// Expect(client.ToFindByGroupSet(customer.Id, customer.Email, &firebase.App{})).Should(BeNil())
		})
		
	})

})
