FROM {{.repo_base}}/centos:8

RUN yum install -y make automake diffutils gcc gcc-c++ kernel-devel && \
  curl -sSLo redis.tar.gz "{{.redis_download_url}}" && \
  mkdir -p /opt/src/redis && \
  tar -C /opt/src/redis --strip-components 1 -xf redis.tar.gz && \
  rm -f redis.tar.gz && \
  cd /opt/src/redis && \
  make && \
  make PREFIX=/opt install && \
  yum remove -y make automake diffutils gcc gcc-c++ kernel-devel && \
  yum autoremove -y && \
  yum clean all

WORKDIR /data

#ENV MINIT_THP never
#ENV MINIT_SYSCTL "vm.overcommit_memory=1"

ADD minit.d /etc/minit.d
ADD redis.conf /etc/redis.conf

ENV REDISCFG_bind 0.0.0.0
ENV REDISCFG_protected_mode no
ENV REDISCFG_port 6379
ENV REDISCFG_tcp_backlog 511
ENV REDISCFG_timeout 0
ENV REDISCFG_tcp_keepalive 300
ENV REDISCFG_daemonize no
ENV REDISCFG_supervised no
ENV REDISCFG_loglevel notice
ENV REDISCFG_logfile '""'
ENV REDISCFG_databases 256
ENV REDISCFG_always_show_logo no
ENV REDISCFG_save__1 900 1
ENV REDISCFG_save__2 300 10
ENV REDISCFG_save__3 60 10000
ENV REDISCFG_stop_writes_on_bgsave_error yes
ENV REDISCFG_rdbcompression yes
ENV REDISCFG_rdbchecksum yes
ENV REDISCFG_dbfilename dump.rdb
ENV REDISCFG_dir ./
ENV REDISCFG_replica_serve_stale_data yes
ENV REDISCFG_replica_read_only yes
ENV REDISCFG_repl_diskless_sync no
ENV REDISCFG_repl_diskless_sync_delay 5
ENV REDISCFG_repl_disable_tcp_nodelay no
ENV REDISCFG_replica_priority 100
ENV REDISCFG_lazyfree_lazy_eviction no
ENV REDISCFG_lazyfree_lazy_expire no
ENV REDISCFG_lazyfree_lazy_server_del no
ENV REDISCFG_replica_lazy_flush no
ENV REDISCFG_appendonly no
ENV REDISCFG_appendfilename '"appendonly.aof"'
ENV REDISCFG_appendfsync everysec
ENV REDISCFG_no_appendfsync_on_rewrite no
ENV REDISCFG_auto_aof_rewrite_percentage 100
ENV REDISCFG_auto_aof_rewrite_min_size 64mb
ENV REDISCFG_aof_load_truncated yes
ENV REDISCFG_aof_use_rdb_preamble yes
ENV REDISCFG_lua_time_limit 5000
ENV REDISCFG_slowlog_log_slower_than 10000
ENV REDISCFG_slowlog_max_len 128
ENV REDISCFG_latency_monitor_threshold 0
ENV REDISCFG_notify_keyspace_events '""'
ENV REDISCFG_hash_max_ziplist_entries 512
ENV REDISCFG_hash_max_ziplist_value 64
ENV REDISCFG_list_max_ziplist_size -2
ENV REDISCFG_list_compress_depth 0
ENV REDISCFG_set_max_intset_entries 512
ENV REDISCFG_zset_max_ziplist_entries 128
ENV REDISCFG_zset_max_ziplist_value 64
ENV REDISCFG_hll_sparse_max_bytes 3000
ENV REDISCFG_stream_node_max_bytes 4096
ENV REDISCFG_stream_node_max_entries 100
ENV REDISCFG_activerehashing yes
ENV REDISCFG_client_output_buffer_limit normal 0 0 0
ENV REDISCFG_client_output_buffer_limit replica 256mb 64mb 60
ENV REDISCFG_client_output_buffer_limit pubsub 32mb 8mb 60
ENV REDISCFG_hz 10
ENV REDISCFG_dynamic_hz yes
ENV REDISCFG_aof_rewrite_incremental_fsync yes
ENV REDISCFG_rdb_save_incremental_fsync yes