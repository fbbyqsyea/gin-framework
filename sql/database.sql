-- 创建数据库
CREATE DATABASE operation;
-- 生成主库账户并授权
GRANT SELECT,INSERT,UPDATE ON operation.* TO operation_master@'%' IDENTIFIED BY 'OperationMaster!23456';
-- 生成从库账户并授权
GRANT SELECT ON operation.* TO operation_replica@'%' IDENTIFIED BY 'OperationReplica!23456';