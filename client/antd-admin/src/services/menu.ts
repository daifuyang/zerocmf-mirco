import { authRequest } from '@/utils/request';

export async function getMenusTree() {
  return authRequest('/api/v1/admin/menu/trees', {
    method: 'GET',
  });
}

export async function getMenu(id: number) {
  return authRequest(`/api/v1/admin/menu/${id}`, {
    method: 'GET',
  });
}

export async function saveMenu(data: any) {
  const { id = '' } = data;
  return authRequest(`/api/v1/admin/menu/${id}`, {
    method: 'POST',
    data,
  });
}
