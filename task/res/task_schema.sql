CREATE TABLE IF NOT EXISTS task
(
    id                   BIGINT(20) auto_increment NOT NULL COMMENT '标识符',
    resource_id          VARCHAR(64)               NOT NULL COMMENT '资源id',
    resource_type        VARCHAR(32)               NOT NULL COMMENT '资源类型',
    task_type            VARCHAR(32)               NOT NULL COMMENT '任务类型',
    status               VARCHAR(32)               NOT NULL COMMENT '任务状态',
    interrupt_signal     VARCHAR(32)               NOT NULL DEFAULT 'NORMAL' COMMENT '中断信号 NORMAL 未中断 SUCCESS FAILED下次执行时中断',
    request_id           VARCHAR(64)               NOT NULL DEFAULT '' COMMENT 'request id',
    next_executable_time TIMESTAMP                 NOT NULL COMMENT '下一个可执行时间',
    retry_count          INT                       NOT NULL DEFAULT 0 COMMENT '重试次数',
    message              VARCHAR(256)              NOT NULL DEFAULT '' COMMENT '成功或失败的原因',
    content              TEXT                      NOT NULL DEFAULT '' COMMENT 'json字符串，存储任务内容',
    create_time          TIMESTAMP                 NOT NULL COMMENT '创建时间',
    update_time          TIMESTAMP                 NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id),
    UNIQUE KEY unq_task_resource_id_type_task_type (resource_id, resource_type, task_type),
    KEY idx_task_status_next_executable_time (status, next_executable_time)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT '任务表';

CREATE TABLE IF NOT EXISTS task_finished
(
    id                   BIGINT(20) auto_increment NOT NULL COMMENT '标识符',
    resource_id          VARCHAR(64)               NOT NULL COMMENT '资源id',
    resource_type        VARCHAR(32)               NOT NULL COMMENT '资源类型',
    task_type            VARCHAR(32)               NOT NULL COMMENT '任务类型',
    status               VARCHAR(32)               NOT NULL COMMENT '任务状态',
    interrupt_signal     VARCHAR(32)               NOT NULL DEFAULT 'NORMAL' COMMENT '中断信号 NORMAL 未中断 SUCCESS FAILED中断',
    request_id           VARCHAR(64)               NOT NULL DEFAULT '' COMMENT 'request id',
    next_executable_time TIMESTAMP                 NOT NULL COMMENT '下一个可执行时间',
    retry_count          INT                       NOT NULL DEFAULT 0 COMMENT '重试次数',
    message              VARCHAR(256)              NOT NULL DEFAULT '' COMMENT '成功或失败的原因',
    content              TEXT                      NOT NULL DEFAULT '' COMMENT 'json字符串，存储任务内容',
    last_update_time     TIMESTAMP                 NOT NULL COMMENT '最后更新时间',
    create_time          TIMESTAMP                 NOT NULL COMMENT '创建时间',
    finish_time          TIMESTAMP                 NOT NULL COMMENT '完成时间',
    PRIMARY KEY (id)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8 COMMENT '任务完成表';