import request from './request';

// 角色类型
export interface Role {
  id: number;
  name: string;
  code: string;
  description: string;
  status: number;
  created_at: string;
}

// 权限类型
export interface Permission {
  id: number;
  name: string;
  code: string;
  type: string;
  parent_id: number;
  path: string;
  icon: string;
  status: number;
  children?: Permission[];
}

// 管理员类型
export interface AdminUser {
  id: number;
  username: string;
  realname: string;
  email: string;
  phone: string;
  avatar: string;
  status: number;
  last_login_at: string;
  created_at: string;
  roles?: Role[];
}

// ========== 角色管理 ==========
export const getRoleList = (params?: { page?: number; page_size?: number }) => {
  return request.get<any, { total: number; list: Role[] }>('/role/list', { params });
};

export const createRole = (data: Partial<Role>) => {
  return request.post('/role', data);
};

export const updateRole = (id: number, data: Partial<Role>) => {
  return request.put(`/role/${id}`, data);
};

export const deleteRole = (id: number) => {
  return request.delete(`/role/${id}`);
};

// ========== 权限管理 ==========
export const getPermissionList = () => {
  return request.get<any, Permission[]>('/permission/list');
};

export const getPermissionTree = () => {
  return request.get<any, Permission[]>('/permission/tree');
};

export const createPermission = (data: Partial<Permission>) => {
  return request.post('/permission', data);
};

export const updatePermission = (id: number, data: Partial<Permission>) => {
  return request.put(`/permission/${id}`, data);
};

export const deletePermission = (id: number) => {
  return request.delete(`/permission/${id}`);
};

// ========== 管理员管理 ==========
export const getAdminUserList = (params?: { page?: number; page_size?: number; keyword?: string }) => {
  return request.get<any, { total: number; list: AdminUser[] }>('/admin-user/list', { params });
};

export const createAdminUser = (data: Partial<AdminUser> & { password: string; role_ids?: number[] }) => {
  return request.post('/admin-user', data);
};

export const updateAdminUser = (id: number, data: Partial<AdminUser> & { role_ids?: number[] }) => {
  return request.put(`/admin-user/${id}`, data);
};

export const deleteAdminUser = (id: number) => {
  return request.delete(`/admin-user/${id}`);
};

// ========== 权限分配 ==========
export const assignRolePermissions = (roleId: number, permissionIds: number[]) => {
  return request.post(`/role/${roleId}/permissions`, { permission_ids: permissionIds });
};

export const getRolePermissions = (roleId: number) => {
  return request.get<any, number[]>(`/role/${roleId}/permissions`);
};

export const assignAdminRoles = (adminId: number, roleIds: number[]) => {
  return request.post(`/admin-user/${adminId}/roles`, { role_ids: roleIds });
};

export const getAdminRoles = (adminId: number) => {
  return request.get<any, number[]>(`/admin-user/${adminId}/roles`);
};
