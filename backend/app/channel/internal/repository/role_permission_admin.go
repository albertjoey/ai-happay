package repository

import (
	"context"
	"happy/app/channel/internal/types"
	"time"

	"gorm.io/gorm"
)

// ==================== 角色仓储实现 ====================

type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository 创建角色仓储
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) List(ctx context.Context, req *types.RoleListRequest) ([]types.Role, int64, error) {
	var total int64
	var roles []types.Role

	r.db.WithContext(ctx).Raw("SELECT COUNT(*) FROM role WHERE deleted_at IS NULL").Scan(&total)

	offset := (req.Page - 1) * req.PageSize
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, name, code, description, status, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at
		FROM role
		WHERE deleted_at IS NULL
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`, req.PageSize, offset).Scan(&roles).Error

	return roles, total, err
}

func (r *roleRepository) FindByID(ctx context.Context, id uint) (*types.Role, error) {
	var role types.Role
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, name, code, description, status
		FROM role
		WHERE id = ? AND deleted_at IS NULL
	`, id).Scan(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *types.Role) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO role (name, code, description, status)
		VALUES (?, ?, ?, 1)
	`, role.Name, role.Code, role.Description).Error
}

func (r *roleRepository) Update(ctx context.Context, role *types.Role) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE role SET name = ?, description = ?, status = ?
		WHERE id = ? AND deleted_at IS NULL
	`, role.Name, role.Description, role.Status, role.ID).Error
}

func (r *roleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE role SET deleted_at = NOW() WHERE id = ?", id).Error
}

func (r *roleRepository) GetPermissions(ctx context.Context, roleID uint) ([]uint, error) {
	var permissionIDs []uint
	err := r.db.WithContext(ctx).Raw(`
		SELECT permission_id FROM role_permission 
		WHERE role_id = ?
	`, roleID).Scan(&permissionIDs).Error
	return permissionIDs, err
}

func (r *roleRepository) AssignPermissions(ctx context.Context, roleID uint, permissionIDs []uint) error {
	// 使用事务
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除旧的关联
		if err := tx.Exec("DELETE FROM role_permission WHERE role_id = ?", roleID).Error; err != nil {
			return err
		}

		// 插入新的关联
		for _, permID := range permissionIDs {
			if err := tx.Exec(`
				INSERT INTO role_permission (role_id, permission_id)
				VALUES (?, ?)
			`, roleID, permID).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ==================== 权限仓储实现 ====================

type permissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository 创建权限仓储
func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) List(ctx context.Context) ([]types.Permission, error) {
	// 使用中间结构体避免GORM映射Children字段的问题
	type permissionRow struct {
		ID       uint
		Name     string
		Code     string
		Type     string
		ParentID uint
		Path     string
		Icon     string
		Status   int8
	}

	var rows []permissionRow
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, name, code, type, parent_id, path, icon, status
		FROM permission
		WHERE deleted_at IS NULL
		ORDER BY id ASC
	`).Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	// 转换为types.Permission
	permissions := make([]types.Permission, len(rows))
	for i, row := range rows {
		permissions[i] = types.Permission{
			ID:       row.ID,
			Name:     row.Name,
			Code:     row.Code,
			Type:     row.Type,
			ParentID: row.ParentID,
			Path:     row.Path,
			Icon:     row.Icon,
			Status:   row.Status,
		}
	}

	return permissions, nil
}

func (r *permissionRepository) Create(ctx context.Context, permission *types.Permission) error {
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO permission (name, code, type, parent_id, path, icon)
		VALUES (?, ?, ?, ?, ?, ?)
	`, permission.Name, permission.Code, permission.Type, permission.ParentID, permission.Path, permission.Icon).Error
}

func (r *permissionRepository) Update(ctx context.Context, permission *types.Permission) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE permission SET name = ?, code = ?, type = ?, parent_id = ?, path = ?, icon = ?
		WHERE id = ? AND deleted_at IS NULL
	`, permission.Name, permission.Code, permission.Type, permission.ParentID, permission.Path, permission.Icon, permission.ID).Error
}

func (r *permissionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE permission SET deleted_at = NOW() WHERE id = ?", id).Error
}

// ==================== 管理员仓储实现 ====================

type adminUserRepository struct {
	db *gorm.DB
}

// NewAdminUserRepository 创建管理员仓储
func NewAdminUserRepository(db *gorm.DB) AdminUserRepository {
	return &adminUserRepository{db: db}
}

func (r *adminUserRepository) List(ctx context.Context, req *types.AdminUserListRequest) ([]types.AdminUser, int64, error) {
	var total int64
	var list []types.AdminUser

	r.db.WithContext(ctx).Raw("SELECT COUNT(*) FROM admin_user WHERE deleted_at IS NULL").Scan(&total)

	offset := (req.Page - 1) * req.PageSize
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, username, realname, email, phone, status, created_at
		FROM admin_user
		WHERE deleted_at IS NULL
		ORDER BY id ASC
		LIMIT ? OFFSET ?
	`, req.PageSize, offset).Scan(&list).Error

	return list, total, err
}

func (r *adminUserRepository) FindByID(ctx context.Context, id uint) (*types.AdminUser, error) {
	var admin types.AdminUser
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, username, realname, email, phone, status
		FROM admin_user
		WHERE id = ? AND deleted_at IS NULL
	`, id).Scan(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminUserRepository) FindByUsername(ctx context.Context, username string) (*types.AdminUser, error) {
	var admin types.AdminUser
	err := r.db.WithContext(ctx).Raw(`
		SELECT id, username, password, realname, email, phone, status
		FROM admin_user
		WHERE username = ? AND deleted_at IS NULL
	`, username).Scan(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminUserRepository) Create(ctx context.Context, admin *types.AdminUser) error {
	now := time.Now()
	// 注意：密码应该在Logic层加密后传入
	return r.db.WithContext(ctx).Exec(`
		INSERT INTO admin_user (username, realname, email, phone, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, admin.Username, admin.Realname, admin.Email, admin.Phone, admin.Status, now, now).Error
}

func (r *adminUserRepository) Update(ctx context.Context, admin *types.AdminUser) error {
	return r.db.WithContext(ctx).Exec(`
		UPDATE admin_user SET realname = ?, email = ?, phone = ?, status = ?, updated_at = ?
		WHERE id = ? AND deleted_at IS NULL
	`, admin.Realname, admin.Email, admin.Phone, admin.Status, time.Now(), admin.ID).Error
}

func (r *adminUserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Exec("UPDATE admin_user SET deleted_at = ? WHERE id = ?", time.Now(), id).Error
}

func (r *adminUserRepository) GetRoles(ctx context.Context, adminUserID uint) ([]uint, error) {
	var roleIDs []uint
	err := r.db.WithContext(ctx).Raw(`
		SELECT role_id FROM admin_user_roles 
		WHERE admin_user_id = ?
	`, adminUserID).Scan(&roleIDs).Error
	return roleIDs, err
}

func (r *adminUserRepository) AssignRoles(ctx context.Context, adminUserID uint, roleIDs []uint) error {
	// 使用事务
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 先删除旧的关联
		if err := tx.Exec("DELETE FROM admin_user_roles WHERE admin_user_id = ?", adminUserID).Error; err != nil {
			return err
		}

		// 插入新的关联
		for _, roleID := range roleIDs {
			if err := tx.Exec(`
				INSERT INTO admin_user_roles (admin_user_id, role_id)
				VALUES (?, ?)
			`, adminUserID, roleID).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
