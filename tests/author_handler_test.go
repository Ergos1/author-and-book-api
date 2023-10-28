//go:build integration
// +build integration

package tests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/ergossteam/homework-3/tests/states"
)

func TestAuthorHandler_GetAuthorByID(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		err := db.SetUp(t)
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		err = srv.SetUp(t)
		if err != nil {
			t.Fatal(err)
		}
		defer srv.TearDown()

		// arrange
		newRecorder := httptest.NewRecorder()

		// act
		req, _ := http.NewRequest("GET", "/authors/1", nil)
		srv.Server.Handler.ServeHTTP(newRecorder, req)
		body, err := io.ReadAll(newRecorder.Result().Body)

		// assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, newRecorder.Result().StatusCode)
		assert.Equal(t, "{\"data\":null,\"error\":\"author not found\"}\n", string(body))
	})

	t.Run("create", func(t *testing.T) {
		err := db.SetUp(t)
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		err = srv.SetUp(t)
		if err != nil {
			t.Fatal(err)
		}
		defer srv.TearDown()

		// arrange
		newRecorder := httptest.NewRecorder()
		reqBody := []byte(fmt.Sprintf("{\"id\": %v, \"name\": \"%v\"}", states.Author1ID, states.AuthorName1))

		// act
		req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(reqBody))
		srv.Server.Handler.ServeHTTP(newRecorder, req)
		body, err := io.ReadAll(newRecorder.Result().Body)

		// assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, newRecorder.Result().StatusCode)
		assert.Equal(t, fmt.Sprintf("{\"data\":{\"id\":%v,\"name\":\"%v\"},\"error\":\"\"}\n", states.Author1ID, states.AuthorName1), string(body))
	})

	t.Run("create & get", func(t *testing.T) {
		err := db.SetUp(t)
		if err != nil {
			t.Fatal(err)
		}
		defer db.TearDown()

		err = srv.SetUp(t)
		if err != nil {
			t.Fatal(err)
		}
		defer srv.TearDown()

		t.Run("create", func(t *testing.T) {
			// arrange
			newRecorder := httptest.NewRecorder()
			reqBody := []byte(fmt.Sprintf("{\"id\": %v, \"name\": \"%v\"}", states.Author1ID, states.AuthorName1))

			// act
			req, _ := http.NewRequest("POST", "/authors", bytes.NewBuffer(reqBody))
			srv.Server.Handler.ServeHTTP(newRecorder, req)
			body, err := io.ReadAll(newRecorder.Result().Body)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, http.StatusCreated, newRecorder.Result().StatusCode)
			assert.Equal(t, fmt.Sprintf("{\"data\":{\"id\":%v,\"name\":\"%v\"},\"error\":\"\"}\n", states.Author1ID, states.AuthorName1), string(body))
		})

		t.Run("get", func(t *testing.T) {
			// arrange
			newRecorder := httptest.NewRecorder()

			// act
			req, _ := http.NewRequest("GET", fmt.Sprintf("/authors/%v", states.Author1ID), nil)
			srv.Server.Handler.ServeHTTP(newRecorder, req)
			body, err := io.ReadAll(newRecorder.Result().Body)

			// // assert
			assert.Nil(t, err)
			assert.Equal(t, http.StatusOK, newRecorder.Result().StatusCode)
			assert.Equal(t, fmt.Sprintf("{\"data\":{\"id\":%v,\"name\":\"%v\",\"books\":[]},\"error\":\"\"}\n", states.Author1ID, states.AuthorName1), string(body))
		})

	})
}
