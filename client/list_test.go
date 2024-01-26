package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"

	"github.com/mortedecai/gbb/client/mocks"
)

var _ = Describe("List", func() {
	It("should call to the server and list the files", func() {
		var req *http.Request
		var err error
		mockCtrl := gomock.NewController(GinkgoT())
		defer mockCtrl.Finish()

		client := mocks.NewMockGBBClient(mockCtrl)
		opt := mocks.NewMockCommandOption(mockCtrl)
		opt.EXPECT().Host().Return(localhost).Times(1)
		opt.EXPECT().Port().Return(9990).Times(1)
		opt.EXPECT().AuthToken().Return(token).Times(1)
		addAuth := opt.EXPECT().AddAuth(gomock.AssignableToTypeOf(req))
		addAuth.Do(func(r *http.Request) {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", opt.AuthToken()))
			addAuth.Return(r)
		})
		Client = client

		req, err = http.NewRequest(http.MethodGet, "http://localhost", bytes.NewBuffer([]byte("{}")))
		Expect(err).ToNot(HaveOccurred())
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		data, err := os.ReadFile("testdata/sample_download_files.json")
		Expect(err).ToNot(HaveOccurred())
		mockResp := httptest.NewRecorder()
		mockResp.WriteHeader(http.StatusOK)
		mockResp.Write(data)

		reader, writer, err := os.Pipe()
		Expect(err).ToNot(HaveOccurred())
		stderr := os.Stderr
		stdout := os.Stdout

		os.Stderr = writer
		os.Stdout = writer

		out := make(chan string)

		wg := new(sync.WaitGroup)
		wg.Add(1)
		client.EXPECT().Do(gomock.AssignableToTypeOf(req)).MaxTimes(1).Return(mockResp.Result(), nil)
		Expect(HandleList(opt)).ToNot(HaveOccurred())

		go func() {
			var buf bytes.Buffer
			wg.Done()
			io.Copy(&buf, reader)
			out <- buf.String()
		}()

		wg.Wait()

		writer.Close()
		os.Stdout = stdout
		os.Stderr = stderr

		str := strings.TrimSpace(<-out)

		builder := strings.Builder{}
		builder.WriteString("BitBurner Files\n")
		builder.WriteString("===============\n")
		builder.WriteString("\n")
		builder.WriteString("n00dles.js")

		Expect(str).To(ContainSubstring(builder.String()))
	})
})
