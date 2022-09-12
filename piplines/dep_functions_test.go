/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */
package piplines

import (
	"net/http"
	"strings"
	"testing"

	skynet "github.com/SkynetLabs/go-skynet/v2"
	"github.com/ali2210/wizdwarf/other/users"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// piplines tests suits
const pipline_describe = "Pipline testing started"
const date_describe = "piplines testing started now"
const skynet_protocol = "skynet tests started"
const skynet_object = "skynet tests started"
const bytes2alpha = "bytes insensitive"
const anscitest = "anscitest tests started"
const server_file = "server_file tests started"
const server_files = "file paths "
const updateinfo = "update info test started"
const updateinfotest = "user updated request proceed"
const addinfo = "add user test started"
const adduser = "new user created"
const db_credientials = "db credentials test started"
const db_credientials_test = "db credentials validation"
const gene_validation = "gene_validation test started"
const gene_chain = "genome have valid sequence"
const protein_valid_seq = "protein generated with correct pair"
const Aminochain_valid = "Generate valid amino acid with valid sequence"

func PiplinesBDDTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, pipline_describe)
}

var _ = Describe(date_describe, func() {
	// decentralize content storage
	Context(skynet_protocol, func() {
		It(skynet_object, func() {
			Expect(SkyDataCenter(skynet.New(), "mickymouse.jpeg", "jpeg")).Should(BeFalse())
		})
	})
	Context(anscitest, func() {
		It(bytes2alpha, func() {
			Expect(ToRunes('a')).Should(BeEmpty())
		})
	})
	// user content store on server
	Context(server_file, func() {
		It(server_files, func() {
			Expect(Open_SFiles("app_data/", "mickeymouse.jpeg")).ShouldNot(BeNil())
		})
	})
	// update user profile
	Context(updateinfo, func() {
		It(updateinfotest, func() {
			Expect(UpdateProfileInfo(&users.Visitors{})).Should(BeTrue())
		})
	})

	// create new profile
	Context(addinfo, func() {
		It(adduser, func() {
			var w http.ResponseWriter
			r := &http.Request{}
			Expect(AddNewProfile(w, r, users.Visitors{}, "0x1223...")).Should(BeTrue())
		})
	})

	// validate user profile
	Context(db_credientials, func() {
		It(db_credientials_test, func() {
			var w http.ResponseWriter
			r := &http.Request{}
			Expect(Firebase_Gatekeeper(w, r, users.Visitors{})).Should(BeEmpty())
		})
	})

	// genes sequences are valid sequence
	Context(gene_validation, func() {
		SetGenes([]string{""})
		It(gene_chain, func() {
			Expect(GetGenes()).Should(BeEmpty())
		})
		It(protein_valid_seq, func() {
			Expect(Active_Proteins(strings.Join(GetGenes(), ""))).Should(BeEmpty())
		})
	})

	// amino acid generate with same valid genes
	Context(Aminochain_valid, func() {
		It(Aminochain_valid, func() {
			Expect(AminoChains("")).Should(BeEmpty())
		})
	})

	// three pair codons should be 3
	Context("three pair family validation", func() {
		It("Should be valid", func() {
			Expect(threepairs("", 0)).Should(BeFalse())
		})
	})

	// message had proof to be sign
	Context("message verification", func() {
		It("Should be message have crypto proof of signed ", func() {
			Expect(TrustRequest("hello world", "0x1666", "")).Should(BeEmpty())
		})
	})
})
