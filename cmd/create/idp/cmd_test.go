package idp_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mrshah-at-ibm/rosa/cmd/create/idp"
	"github.com/mrshah-at-ibm/rosa/cmd/create/idp/mocks"
)

var _ = Describe("Cmd", func() {
	var (
		mockCtrl *gomock.Controller
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("GenerateIdpName", func() {

		var (
			idpType string
			idps    []idp.IdentityProvider
		)
		BeforeEach(func() {
			idpType = "github"
			idps = []idp.IdentityProvider{}
		})
		Context("when no IDP exists", func() {
			It("generates a idp name name-1", func() {
				name := idp.GenerateIdpName(idpType, idps)
				Expect(name).To(Equal(idpType + "-1"))
			})
		})

		Context("when an IDP with the name of the type already exists", func() {
			BeforeEach(func() {
				mockIdp := mocks.NewMockIdentityProvider(mockCtrl)
				mockIdp.EXPECT().Name().Return("github").AnyTimes()
				idps = append(idps, mockIdp)
			})
			It("generates a unique idp name", func() {
				name := idp.GenerateIdpName(idpType, idps)
				expectUnique(name, idps)
			})
		})

		Context("when an IDP with a generated name already exists", func() {
			BeforeEach(func() {
				mockIdp := mocks.NewMockIdentityProvider(mockCtrl)
				mockIdp.EXPECT().Name().Return(idp.GenerateIdpName(idpType, idps)).AnyTimes()
				idps = append(idps, mockIdp)
			})
			It("generates a unique idp name", func() {
				name := idp.GenerateIdpName(idpType, idps)
				expectUnique(name, idps)
			})
		})
	})
})

func expectUnique(name string, idps []idp.IdentityProvider) {
	for _, idp := range idps {
		Expect(name).NotTo(Equal(idp.Name()))
	}
}
