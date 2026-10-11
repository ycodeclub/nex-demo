// Package dao 持久化边界：自行选择 Ent / GORM / sqlc / database/sql 等实现。
// ORM 与 SQL 只在此包及子包；与 domain 的转换放在 repository。
package dao
