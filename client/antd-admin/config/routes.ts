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
    path: '/home',
    component: './Home',
  },
  {
    name: '权限演示',
    path: '/access',
    component: './Access',
  },
];
