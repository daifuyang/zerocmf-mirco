import { authRequest, request } from '@/utils/request';

export async function login(params: any) {
  return request('/api/authn/admin/login', {
    method: 'POST',
    data: params,
  });
}

export async function currentUser() {
  return authRequest('/api/v1/user/admin/user_info', {
    method: 'GET',
  });
}
