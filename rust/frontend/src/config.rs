use crate::{
    executor::config::{ExecutorConfig, LocalExecutorConfig},
    CollectionsWithSegmentsProviderConfig,
};
use chroma_log::config::LogConfig;
use chroma_segment::local_segment_manager::LocalSegmentManagerConfig;
use chroma_sqlite::config::SqliteDBConfig;
use chroma_sysdb::SysDbConfig;
use figment::providers::{Env, Format, Yaml};
use mdac::CircuitBreakerConfig;
use rust_embed::Embed;
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct ScorecardRule {
    pub patterns: Vec<String>,
    pub score: u32,
}

fn default_sysdb_config() -> SysDbConfig {
    SysDbConfig::Sqlite(Default::default())
}

fn default_log_config() -> LogConfig {
    LogConfig::Sqlite(Default::default())
}

fn default_executor_config() -> ExecutorConfig {
    ExecutorConfig::Local(LocalExecutorConfig {})
}

fn default_sqlitedb() -> Option<SqliteDBConfig> {
    Some(SqliteDBConfig::default())
}

fn default_segment_manager_config() -> Option<LocalSegmentManagerConfig> {
    Some(LocalSegmentManagerConfig {
        hnsw_index_pool_cache_config: chroma_cache::CacheConfig::Memory(
            chroma_cache::FoyerCacheConfig {
                capacity: 65536,
                ..Default::default()
            },
        ),
        persist_path: None,
    })
}

#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct FrontendConfig {
    #[serde(default)]
    pub allow_reset: bool,
    #[serde(default = "default_sqlitedb")]
    pub sqlitedb: Option<SqliteDBConfig>,
    #[serde(default = "default_segment_manager_config")]
    pub segment_manager: Option<LocalSegmentManagerConfig>,
    #[serde(default = "default_sysdb_config")]
    pub sysdb: SysDbConfig,
    #[serde(default)]
    pub collections_with_segments_provider: CollectionsWithSegmentsProviderConfig,
    #[serde(default = "default_log_config")]
    pub log: LogConfig,
    #[serde(default = "default_executor_config")]
    pub executor: ExecutorConfig,
}

#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct OpenTelemetryConfig {
    pub endpoint: String,
    pub service_name: String,
}

fn default_port() -> u16 {
    8000
}

fn default_listen_address() -> String {
    "0.0.0.0".to_string()
}

fn default_max_payload_size_bytes() -> usize {
    40 * 1024 * 1024 // 40 MB
}

#[derive(Deserialize, Serialize, Clone, Debug)]
pub struct FrontendServerConfig {
    #[serde(flatten)]
    pub frontend: FrontendConfig,
    #[serde(default = "default_port")]
    pub port: u16,
    #[serde(default = "default_listen_address")]
    pub listen_address: String,
    #[serde(default = "default_max_payload_size_bytes")]
    pub max_payload_size_bytes: usize,
    #[serde(default = "CircuitBreakerConfig::default")]
    pub circuit_breaker: CircuitBreakerConfig,
    #[serde(default)]
    pub scorecard_enabled: bool,
    #[serde(default)]
    pub scorecard: Vec<ScorecardRule>,
    pub open_telemetry: Option<OpenTelemetryConfig>,
    #[serde(default)]
    pub persist_path: Option<String>,
}

const DEFAULT_CONFIG_PATH: &str = "./frontend_config.yaml";
const DEFAULT_SINGLE_NODE_CONFIG_FILENAME: &str = "single_node_frontend_config.yaml";

#[derive(Embed)]
#[folder = "./"]
#[include = "*.yaml"]
struct DefaultConfigurationsFolder;

impl FrontendServerConfig {
    pub fn load() -> Self {
        Self::load_from_path(DEFAULT_CONFIG_PATH)
    }

    pub fn load_from_path(path: &str) -> Self {
        // Unfortunately, figment doesn't support environment variables with underscores. So we have to map and replace them.
        // Excluding our own environment variables, which are prefixed with CHROMA_.
        let mut f = figment::Figment::from(
            Env::prefixed("CHROMA_").map(|k| k.as_str().replace("__", ".").into()),
        );
        if std::path::Path::new(path).exists() {
            f = figment::Figment::from(Yaml::file(path)).merge(f);
        }
        let res = f.extract();
        match res {
            Ok(config) => config,
            Err(e) => panic!("Error loading config: {}", e),
        }
    }

    pub fn single_node_default() -> Self {
        // TOOD: unify this with load_from_path to get the env overrides
        let config = DefaultConfigurationsFolder::get(DEFAULT_SINGLE_NODE_CONFIG_FILENAME)
            .expect("Failed to load default single node frontend config");
        let config_data = config.data;
        let config_str = std::str::from_utf8(&config_data).expect("Failed to parse config data");
        let f = figment::Figment::from(Yaml::string(config_str));
        let res = f.extract();
        match res {
            Ok(config) => config,
            Err(e) => panic!("Error loading config: {}", e),
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::config::FrontendServerConfig;
    use chroma_cache::CacheConfig;

    #[test]
    fn test_load_config() {
        let config = FrontendServerConfig::load();
        let sysdb_config = config.frontend.sysdb;
        let sysdb_config = match sysdb_config {
            chroma_sysdb::SysDbConfig::Grpc(grpc_sys_db_config) => grpc_sys_db_config,
            chroma_sysdb::SysDbConfig::Sqlite(_) => {
                panic!("Expected grpc sysdb config, got sqlite sysdb config")
            }
        };
        assert_eq!(sysdb_config.host, "sysdb.chroma");
        assert_eq!(sysdb_config.port, 50051);
        assert_eq!(sysdb_config.connect_timeout_ms, 60000);
        assert_eq!(sysdb_config.request_timeout_ms, 60000);
        assert_eq!(sysdb_config.num_channels, 5);
        assert_eq!(
            config
                .frontend
                .collections_with_segments_provider
                .permitted_parallelism,
            180
        );
        match config.frontend.collections_with_segments_provider.cache {
            CacheConfig::Memory(c) => {
                assert_eq!(c.capacity, 1000);
            }
            CacheConfig::Disk(c) => {
                assert_eq!(c.capacity, 1000);
            }
            CacheConfig::Nop => {}
            _ => {}
        }
    }
}
