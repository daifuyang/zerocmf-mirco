export default [
  {
    path: '/',
    redirect: '/home',
  },
  {
    path: '/login',
    layout: false,
    component: './Login',
  },
  {
    name: '首页',
    hideInMenu: true,
    path: '/home',
    component: './Home',
  },
  {
    path: '/admin',
    name: '系统管理',
    routes: [
      {
        path: '/admin',
        redirect: '/admin/user',
      },
      {
        path: '/admin/user',
        name: '管理员',
        routes: [
          {
            path: '/admin/user',
            redirect: '/admin/user/list',
          },
          {
            path: '/admin/user/list',
            name: '管理员列表',
            hideInMenu: true,
            component: './Admin'
          },
          {
            path: '/admin/user/add',
            name: '添加管理员',
            hideInMenu: true,
          },
          {
            path: '/admin/user/edit/:id',
            name: '修改管理员',
            hideInMenu: true,
          },
        ],
      },
      {
        path: '/admin/role',
        name: '角色管理',
        routes: [
          {
            path: '/admin/role',
            redirect: '/admin/role/list',
          },
          {
            path: '/admin/role/list',
            name: '角色列表',
            hideInMenu: true,
          },
          {
            path: '/admin/role/add',
            name: '添加管理员',
            hideInMenu: true,
          },
          {
            path: '/admin/role/edit/:id',
            name: '修改管理员',
            hideInMenu: true,
          },
        ],
      },
      {
        path: '/admin/department',
        name: '部门管理',
        routes: [
          {
            path: '/admin/department',
            redirect: '/admin/department/list',
          },
          {
            path: '/admin/department/list',
            name: '部门列表',
            hideInMenu: true,
          },
          {
            path: '/admin/department/add',
            name: '添加部门',
            hideInMenu: true,
          },
          {
            path: '/admin/department/edit/:id',
            name: '修改部门',
            hideInMenu: true,
          },
        ],
      },
      {
        path: '/admin/post',
        name: '岗位管理',
        routes: [
          {
            path: '/admin/post',
            redirect: '/admin/post/list',
          },
          {
            path: '/admin/post/list',
            name: '岗位列表',
            hideInMenu: true,
          },
          {
            path: '/admin/post/add',
            name: '添加岗位',
            hideInMenu: true,
          },
          {
            path: '/admin/post/edit/:id',
            name: '修改岗位',
            hideInMenu: true,
          },
        ],
      },
      {
        name: '菜单管理',
        path: '/admin/menu',
        hideInMenu: true,
        routes: [
          {
            path: '/admin/menu',
            redirect: '/admin/menu/list',
          },
          {
            path: '/admin/menu/list',
            name: '菜单列表',
            component: './Menu/index',
          },
        ],
      },
    ],
  },
  {
    name: '应用管理',
    path: '/apps'
  },

  // {
  //   name: '权限演示',
  //   path: '/access',
  //   component: './Access',
  // },
];
