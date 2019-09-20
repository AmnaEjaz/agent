/****************************************************************************
 * Copyright 2019, Optimizely, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/

// Package api //
package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	// "github.com/optimizely/sidedoor/pkg/optimizely"
	"github.com/optimizely/sidedoor/pkg/optimizelytest"

	"github.com/stretchr/testify/suite"
)

type MockOptlyMiddleware struct{}

func (m *MockOptlyMiddleware) ClientCtx(next http.Handler) http.Handler {
	return next
}

type MockFeatureAPI struct{}

func (m *MockFeatureAPI) ListFeatures(w http.ResponseWriter, r *http.Request)    {}
func (m *MockFeatureAPI) GetFeature(w http.ResponseWriter, r *http.Request)      {}
func (m *MockFeatureAPI) ActivateFeature(w http.ResponseWriter, r *http.Request) {}

type MockUserEventAPI struct{}

func (m *MockUserEventAPI) AddUserEvent(w http.ResponseWriter, r *http.Request) {}

type RouterTestSuite struct {
	suite.Suite
	tc  *optimizelytest.TestClient
	mux *chi.Mux
	// opt *RouterOptions
}

func (suite *RouterTestSuite) SetupTest() {
	testClient := optimizelytest.NewClient()
	suite.tc = testClient
	// client := &optimizely.OptlyClient{testClient.OptimizelyClient}

	opts := &RouterOptions{
		featureAPI:   new(MockFeatureAPI),
		userEventAPI: new(MockUserEventAPI),
		middleware:   new(MockOptlyMiddleware),
	}

	suite.mux = NewRouter(opts)
}

func (suite *RouterTestSuite) TestRouter() {
	r := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	suite.mux.ServeHTTP(rr, r)
}

func TestRouter(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}