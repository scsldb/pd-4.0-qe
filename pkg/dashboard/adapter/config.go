// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"github.com/pingcap-incubator/tidb-dashboard/pkg/config"

	"github.com/pingcap/pd/v4/server"
)

// GenDashboardConfig generates a configuration for Dashboard Server.
func GenDashboardConfig(srv *server.Server) (*config.Config, error) {
	cfg := srv.GetConfig()

	etcdCfg, err := cfg.GenEmbedEtcdConfig()
	if err != nil {
		return nil, err
	}

	dashboardCfg := &config.Config{
		DataDir:            cfg.DataDir,
		PDEndPoint:         etcdCfg.ACUrls[0].String(),
		PublicPathPrefix:   cfg.Dashboard.PublicPathPrefix,
		EnableTelemetry:    cfg.Dashboard.EnableTelemetry,
		EnableExperimental: true,
	}

	if dashboardCfg.ClusterTLSConfig, err = cfg.Security.ToTLSConfig(); err != nil {
		return nil, err
	}
	if dashboardCfg.TiDBTLSConfig, err = cfg.Dashboard.ToTiDBTLSConfig(); err != nil {
		return nil, err
	}

	dashboardCfg.NormalizePublicPathPrefix()

	return dashboardCfg, nil
}
