/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"encoding/json"
	"testing"

	"github.com/opentelekomcloud/provider-opentelekomcloud/apis/namespaced/v1beta1"
)

const (
	testProjectID   = "p-1"
	testProjectName = "my-project"
)

func TestBuildScopeMap(t *testing.T) {
	cases := []struct {
		name    string
		scope   *v1beta1.WebIdentityScope
		want    map[string]any
		wantErr bool
	}{
		{
			name:  "nil scope is unscoped",
			scope: nil,
			want:  nil,
		},
		{
			name:  "empty scope is unscoped",
			scope: &v1beta1.WebIdentityScope{},
			want:  nil,
		},
		{
			name: "project by id",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Id: testProjectID},
			},
			want: map[string]any{"project": map[string]any{"id": testProjectID}},
		},
		{
			name: "project by name",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Name: testProjectName},
			},
			want: map[string]any{"project": map[string]any{"name": testProjectName}},
		},
		{
			name: "domain by name",
			scope: &v1beta1.WebIdentityScope{
				Domain: &v1beta1.WebIdentityScopeResource{Name: "my-domain"},
			},
			want: map[string]any{"domain": map[string]any{"name": "my-domain"}},
		},
		{
			name: "domain by id",
			scope: &v1beta1.WebIdentityScope{
				Domain: &v1beta1.WebIdentityScopeResource{Id: "d-1"},
			},
			want: map[string]any{"domain": map[string]any{"id": "d-1"}},
		},
		{
			name: "both project and domain is invalid",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Id: testProjectID},
				Domain:  &v1beta1.WebIdentityScopeResource{Id: "d-1"},
			},
			wantErr: true,
		},
		{
			name: "project with both id and name is invalid",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Id: testProjectID, Name: testProjectName},
			},
			wantErr: true,
		},
		{
			name: "project with neither id nor name is invalid",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{},
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := buildScopeMap(tc.scope)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil (scope=%+v)", tc.scope)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.want == nil {
				if got != nil {
					t.Fatalf("expected nil scope, got %+v", got)
				}
				return
			}
			key, ok := tc.want["project"]
			if ok {
				if g, ok := got["project"]; !ok || !mapsEqual(g, key) {
					t.Fatalf("project mismatch: got=%+v want=%+v", got["project"], key)
				}
			}
			key, ok = tc.want["domain"]
			if ok {
				if g, ok := got["domain"]; !ok || !mapsEqual(g, key) {
					t.Fatalf("domain mismatch: got=%+v want=%+v", got["domain"], key)
				}
			}
		})
	}
}

func mapsEqual(a, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	am, okA := a.(map[string]any)
	bm, okB := b.(map[string]any)
	if !okA || !okB {
		return a == b
	}
	if len(am) != len(bm) {
		return false
	}
	for k, v := range am {
		bv, ok := bm[k]
		if !ok || v != bv {
			return false
		}
	}
	return true
}

func TestProjectTenantFields(t *testing.T) {
	cases := []struct {
		name  string
		scope *v1beta1.WebIdentityScope
		want  map[string]any
	}{
		{
			name:  "nil scope",
			scope: nil,
			want:  nil,
		},
		{
			name:  "empty scope",
			scope: &v1beta1.WebIdentityScope{},
			want:  nil,
		},
		{
			name: "project by id",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Id: testProjectID},
			},
			want: map[string]any{"tenant_id": testProjectID},
		},
		{
			name: "project by name",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Name: testProjectName},
			},
			want: map[string]any{"tenant_name": testProjectName},
		},
		{
			name: "id preferred over name",
			scope: &v1beta1.WebIdentityScope{
				Project: &v1beta1.WebIdentityScopeResource{Id: testProjectID, Name: testProjectName},
			},
			want: map[string]any{"tenant_id": testProjectID},
		},
		{
			name: "domain scope emits no tenant",
			scope: &v1beta1.WebIdentityScope{
				Domain: &v1beta1.WebIdentityScopeResource{Id: "d-1"},
			},
			want: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := projectTenantFields(tc.scope)
			if !mapsEqual(got, tc.want) {
				t.Fatalf("projectTenantFields(%+v) = %+v, want %+v", tc.scope, got, tc.want)
			}
		})
	}
}

const testAuthURLBase = "https://iam.eu-de.otc.t-systems.com"

func TestNormalizeAuthURLV3(t *testing.T) {
	want := testAuthURLBase + "/v3"
	for _, in := range []string{
		testAuthURLBase,
		testAuthURLBase + "/",
		testAuthURLBase + "/v3",
		testAuthURLBase + "/v3.0",
		testAuthURLBase + "/v3/",
	} {
		if got := normalizeAuthURLV3(in); got != want {
			t.Errorf("normalizeAuthURLV3(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestNormalizeAuthURLBase(t *testing.T) {
	for _, in := range []string{
		testAuthURLBase,
		testAuthURLBase + "/",
		testAuthURLBase + "/v3",
		testAuthURLBase + "/v3.0",
	} {
		if got := normalizeAuthURLBase(in); got != testAuthURLBase {
			t.Errorf("normalizeAuthURLBase(%q) = %q, want %q", in, got, testAuthURLBase)
		}
	}
}

func TestTempCredsEnabled(t *testing.T) {
	// nil defaults to enabled
	if !tempCredsEnabled(nil) {
		t.Errorf("expected tempCredsEnabled(nil) to be true")
	}
	f := false
	tr := true
	if tempCredsEnabled(&f) {
		t.Errorf("expected tempCredsEnabled(&false) to be false")
	}
	if !tempCredsEnabled(&tr) {
		t.Errorf("expected tempCredsEnabled(&true) to be true")
	}
}

// TestBuildIdTokenBody locks in the exact request body for the
// OS-AUTH/id-token/tokens endpoint. The ID token must be carried under
// auth.id_token.id (not auth.identity), with an optional auth.scope.
func TestBuildIdTokenBody(t *testing.T) {
	const token = "eyJhbGciOiJSUzI1NiJ9.header.sig"

	t.Run("unscoped", func(t *testing.T) {
		got, err := buildIdTokenBody(token, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(got, &m); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}
		auth, ok := m["auth"].(map[string]any)
		if !ok {
			t.Fatalf("missing auth object: %s", string(got))
		}
		idToken, ok := auth["id_token"].(map[string]any)
		if !ok {
			t.Fatalf("auth.id_token must be an object: %s", string(got))
		}
		if idToken["id"] != token {
			t.Errorf("auth.id_token.id = %v, want %q", idToken["id"], token)
		}
		if _, present := auth["scope"]; present {
			t.Errorf("scope should be omitted when unscoped: %s", string(got))
		}
	})

	t.Run("project scoped", func(t *testing.T) {
		got, err := buildIdTokenBody(token, map[string]any{"project": map[string]any{"id": testProjectID}})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		var m map[string]any
		if err := json.Unmarshal(got, &m); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}
		auth, ok := m["auth"].(map[string]any)
		if !ok {
			t.Fatalf("missing auth object: %s", string(got))
		}
		scope, ok := auth["scope"].(map[string]any)
		if !ok {
			t.Fatalf("auth.scope must be an object: %s", string(got))
		}
		project, ok := scope["project"].(map[string]any)
		if !ok {
			t.Fatalf("auth.scope.project must be an object: %s", string(got))
		}
		if project["id"] != testProjectID {
			t.Errorf("auth.scope.project.id = %v, want %s", project["id"], testProjectID)
		}
	})
}
