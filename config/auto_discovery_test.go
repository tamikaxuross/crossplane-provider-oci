/*
 * Copyright (c) 2025 Oracle and/or its affiliates
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import "testing"

func TestGenerateKindNameCompactsOverlongKinds(t *testing.T) {
	testCases := map[string]struct {
		resourceName string
		group        string
		want         string
	}{
		"autonomous database dbm features management": {
			resourceName: "oci_database_management_autonomous_database_autonomous_database_dbm_features_management",
			group:        "database",
			want:         "ManagementAutonomousDatabaseDbmFeaturesManagement",
		},
		"external container database dbm features management": {
			resourceName: "oci_database_management_externalcontainerdatabase_external_container_dbm_features_management",
			group:        "database",
			want:         "ManagementExternalcontainerdatabaseDbmFeaturesManagement",
		},
		"external mysql database management": {
			resourceName: "oci_database_management_external_my_sql_database_external_mysql_databases_management",
			group:        "database",
			want:         "ManagementExternalMySqlDatabaseExternalMysqlDatabases",
		},
		"external non container dbm features management": {
			resourceName: "oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management",
			group:        "database",
			want:         "ManagementExternalnoncontainerdatabaseDbmFeaturesManagement",
		},
		"external pluggable dbm features management": {
			resourceName: "oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management",
			group:        "database",
			want:         "ManagementExternalpluggabledatabaseDbmFeaturesManagement",
		},
		"pluggable database dbm features management": {
			resourceName: "oci_database_management_pluggabledatabase_pluggable_database_dbm_features_management",
			group:        "database",
			want:         "ManagementPluggabledatabaseDbmFeaturesManagement",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := generateKindName(tc.resourceName, tc.group)
			if got != tc.want {
				t.Fatalf("generateKindName(%q, %q) = %q, want %q", tc.resourceName, tc.group, got, tc.want)
			}
			if len(got) > maxGeneratedKindNameLength {
				t.Fatalf("generateKindName(%q, %q) produced %q with length %d, want <= %d", tc.resourceName, tc.group, got, len(got), maxGeneratedKindNameLength)
			}
			if len(got+"List") > maxKubernetesNameLength {
				t.Fatalf("generateKindName(%q, %q) produced List kind %q with length %d, want <= %d", tc.resourceName, tc.group, got+"List", len(got+"List"), maxKubernetesNameLength)
			}
		})
	}
}

func TestGenerateKindNamePreservesShortKinds(t *testing.T) {
	resourceName := "oci_identity_compartment"
	group := "identity"

	got := generateKindName(resourceName, group)
	if got != "Compartment" {
		t.Fatalf("generateKindName(%q, %q) = %q, want %q", resourceName, group, got, "Compartment")
	}
}
