package commands

import (
	"errors"
	"github.com/mortedecai/gbb/commands/mocks"
	"github.com/mortedecai/gbb/gbberror"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

const version = "v0.0.0-test"

var _ = Describe("Root", func() {
	It("should create a root command object", func() {
		cmd, err := Root(version)
		Expect(cmd).ToNot(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})
})

var _ = Describe("Root Option", func() {
	It("should be possible to create and use a root option", func() {
		ro := &rootOption{
			host:      "example.com",
			port:      1234,
			authToken: "abc",
		}
		Expect(ro.Host()).To(Equal(ro.host))
		Expect(ro.Port()).To(Equal(ro.port))
		Expect(ro.AuthToken()).To(Equal(ro.authToken))
	})
})
var _ = Describe("Root Option", func() {
	var (
		baseRo *rootOption
	)
	BeforeEach(func() {
		baseRo = &rootOption{
			host:      "example.com",
			port:      1234,
			authToken: "abc",
		}
	})
	It("should be possible to create and use a download option", func() {
		ro := baseRo
		Expect(ro.Host()).To(Equal(ro.host))
		Expect(ro.Port()).To(Equal(ro.port))
		Expect(ro.AuthToken()).To(Equal(ro.authToken))
		Expect(ro.Valid()).To(BeTrue())
	})
	It("should not be valid if the root option isn't valid", func() {
		ro := baseRo
		ro.host = ""
		Expect(ro.Valid()).To(BeFalse())
	})
	It("should not be valid if the root option isn't valid", func() {
		ro := baseRo
		ro.authToken = ""
		Expect(ro.Valid()).To(BeFalse())
	})
	It("should not be valid if the root option isn't valid", func() {
		ro := baseRo
		ro.port = 0
		Expect(ro.Valid()).To(BeFalse())
	})
	It("should return the correct server address", func() {
		ro := baseRo
		Expect(ro.Server()).To(Equal("http://example.com:1234"))
	})
})
var _ = Describe("handleCommonFlags", func() {
	var (
		ctrl           *gomock.Controller
		origFlagReader Flagger
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		origFlagReader = flagReader
		flagReader = mocks.NewMockFlagger(ctrl)
	})
	AfterEach(func() {
		ctrl.Finish()
		flagReader = origFlagReader
	})
	It("should return an error if GetString fails for host", func() {
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), gomock.Any()).Return("", errors.New("bad host")).Times(1)
		_, _, _, err := handleCommonFlags(nil)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("bad host"))
	})
	It("should return an error if GetInt fails for port", func() {
		const localhost = "localhost"
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), gomock.Any()).Return(localhost, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetInt(gomock.Any(), gomock.Any()).Return(0, errors.New("bad port"))
		host, _, _, err := handleCommonFlags(nil)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("bad port"))
		Expect(host).To(Equal(localhost))
	})
	It("should return an error if GetString fails for authToken", func() {
		const localhost = "localhost"
		const expPort = 1234
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "host").Return(localhost, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetInt(gomock.Any(), "port").Return(expPort, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "authToken").Return("", errors.New("bad token")).Times(1)
		host, port, _, err := handleCommonFlags(nil)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("bad token"))
		Expect(host).To(Equal(localhost))
		Expect(port).To(Equal(expPort))
	})
	It("should return an error if host is empty", func() {
		const expHost = ""
		const expPort = 1234
		const expToken = "abc"
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "host").Return(expHost, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetInt(gomock.Any(), "port").Return(expPort, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "authToken").Return(expToken, nil).Times(1)
		host, port, token, err := handleCommonFlags(nil)
		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(gbberror.ErrBadHost))
		Expect(host).To(Equal(expHost))
		Expect(port).To(Equal(expPort))
		Expect(token).To(Equal(expToken))
	})
	It("should return an error if port is out of range", func() {
		const expHost = "localhost"
		const expPort = -1
		const expToken = "abc"
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "host").Return(expHost, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetInt(gomock.Any(), "port").Return(expPort, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "authToken").Return(expToken, nil).Times(1)
		host, port, token, err := handleCommonFlags(nil)
		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(gbberror.ErrBadPort))
		Expect(host).To(Equal(expHost))
		Expect(port).To(Equal(expPort))
		Expect(token).To(Equal(expToken))
	})
	It("should return an error if token is empty", func() {
		const expHost = "localhost"
		const expPort = 1234
		const expToken = ""
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "host").Return(expHost, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetInt(gomock.Any(), "port").Return(expPort, nil).Times(1)
		flagReader.(*mocks.MockFlagger).EXPECT().GetString(gomock.Any(), "authToken").Return(expToken, nil).Times(1)
		host, port, token, err := handleCommonFlags(nil)
		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(gbberror.ErrNoAuthToken))
		Expect(host).To(Equal(expHost))
		Expect(port).To(Equal(expPort))
		Expect(token).To(Equal(expToken))
	})
})
