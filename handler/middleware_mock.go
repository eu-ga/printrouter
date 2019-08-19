package handler

import (
	"context"
	"net/http"

	"github.com/rockspoon/go-common/middleware"
	s "github.com/rockspoon/rs.cor.middleware/model"
	soajsgo "github.com/soajs/soajs.golang"
)

func soajsTest(addContext bool, tenantID, uracID, eKey string) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			if addContext {
				ctx = context.WithValue(r.Context(), soajsgo.SoajsKey, soajsgo.ContextData{
					Tenant: soajsgo.Tenant{
						ID:          tenantID,
						Code:        "",
						Locked:      false,
						Key:         soajsgo.Key{EKey: eKey},
						Roaming:     nil,
						Application: soajsgo.Application{},
					},
					Urac: soajsgo.Urac{
						ID:          uracID,
						Username:    "",
						FirstName:   "",
						LastName:    "",
						Email:       "",
						Groups:      nil,
						SocialLogin: nil,
						Tenant: soajsgo.Tenant{
							ID:          tenantID,
							Code:        "",
							Locked:      false,
							Key:         soajsgo.Key{},
							Roaming:     nil,
							Application: soajsgo.Application{},
						},
						Profile:   nil,
						ACL:       nil,
						ACLAllEnv: nil,
					},
					ServicesConfig: nil,
					Device:         "",
					Geo:            nil,
					Awareness:      soajsgo.Host{},
					Reg:            nil,
				})
			}
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getContextData(tenantID, uracID, eKey string) *s.ContextData {
	return &s.ContextData{
		//nolint:staticcheck
		Tenant: s.Tenant{
			ID:   tenantID,
			Code: "",
			Key:  eKey,
		},
		User: s.User{
			ID:        uracID,
			Username:  "",
			FirstName: "",
			LastName:  "",
			Email:     "",
		},
		Paths: map[string]string{"urac": ":0/urac/", "venue": ":0/venue/"},
	}
}
