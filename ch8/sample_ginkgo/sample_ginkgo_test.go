package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "cabos.io/gwp/ch8/sample_ginkgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SampleGinkgo", func() {
	var mux *http.ServeMux
	var post *FakePost
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		mux.HandleFunc("/post/", HandleRequest(post))
		writer = httptest.NewRecorder()
	})

	Context("GEt a post using an id", func() {
		It("should get a post", func() {
			request, _ := http.NewRequest("GET", "/post/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)
			Expect(post.Id).To(Equal(1))

		})
	})
})
