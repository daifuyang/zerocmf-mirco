import { authRequest } from "@/utils/request";

export async function getMenusTree() {
    return authRequest('/api/v1/admin/menu/trees', {
      method: 'GET',
    });
  }
  