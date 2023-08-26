// @ts-nocheck
// This file is generated by Umi automatically
// DO NOT CHANGE IT MANUALLY!
import React from 'react';

export async function getRoutes() {
  const routes = {"1":{"path":"/","redirect":"/home","parentId":"ant-design-pro-layout","id":"1"},"2":{"path":"/login","layout":false,"id":"2"},"3":{"name":"首页","hideInMenu":true,"path":"/home","parentId":"ant-design-pro-layout","id":"3"},"4":{"path":"/admin","name":"系统管理","parentId":"ant-design-pro-layout","id":"4"},"5":{"path":"/admin","redirect":"/admin/user","parentId":"4","id":"5"},"6":{"path":"/admin/user","name":"管理员","parentId":"4","id":"6"},"7":{"path":"/admin/user","redirect":"/admin/user/list","parentId":"6","id":"7"},"8":{"path":"/admin/user/list","name":"管理员列表","hideInMenu":true,"parentId":"6","id":"8"},"9":{"path":"/admin/user/add","name":"添加管理员","hideInMenu":true,"parentId":"6","id":"9"},"10":{"path":"/admin/user/edit/:id","name":"修改管理员","hideInMenu":true,"parentId":"6","id":"10"},"11":{"path":"/admin/role","name":"角色管理","parentId":"4","id":"11"},"12":{"path":"/admin/role","redirect":"/admin/role/list","parentId":"11","id":"12"},"13":{"path":"/admin/role/list","name":"角色列表","hideInMenu":true,"parentId":"11","id":"13"},"14":{"path":"/admin/role/add","name":"添加管理员","hideInMenu":true,"parentId":"11","id":"14"},"15":{"path":"/admin/role/edit/:id","name":"修改管理员","hideInMenu":true,"parentId":"11","id":"15"},"16":{"path":"/admin/department","name":"部门管理","parentId":"4","id":"16"},"17":{"path":"/admin/department","redirect":"/admin/department/list","parentId":"16","id":"17"},"18":{"path":"/admin/department/list","name":"部门列表","hideInMenu":true,"parentId":"16","id":"18"},"19":{"path":"/admin/department/add","name":"添加部门","hideInMenu":true,"parentId":"16","id":"19"},"20":{"path":"/admin/department/edit/:id","name":"修改部门","hideInMenu":true,"parentId":"16","id":"20"},"21":{"path":"/admin/post","name":"岗位管理","parentId":"4","id":"21"},"22":{"path":"/admin/post","redirect":"/admin/post/list","parentId":"21","id":"22"},"23":{"path":"/admin/post/list","name":"岗位列表","hideInMenu":true,"parentId":"21","id":"23"},"24":{"path":"/admin/post/add","name":"添加岗位","hideInMenu":true,"parentId":"21","id":"24"},"25":{"path":"/admin/post/edit/:id","name":"修改岗位","hideInMenu":true,"parentId":"21","id":"25"},"26":{"name":"菜单管理","path":"/admin/menu","hideInMenu":true,"parentId":"4","id":"26"},"27":{"path":"/admin/menu","redirect":"/admin/menu/list","parentId":"26","id":"27"},"28":{"path":"/admin/menu/list","name":"菜单列表","parentId":"26","id":"28"},"29":{"name":"应用管理","path":"/apps","parentId":"ant-design-pro-layout","id":"29"},"ant-design-pro-layout":{"id":"ant-design-pro-layout","path":"/","isLayout":true}} as const;
  return {
    routes,
    routeComponents: {
'1': React.lazy(() => import( './EmptyRoute')),
'2': React.lazy(() => import(/* webpackChunkName: "p__Login__index" */'@/pages/Login/index.tsx')),
'3': React.lazy(() => import(/* webpackChunkName: "p__Home__index" */'@/pages/Home/index.tsx')),
'4': React.lazy(() => import( './EmptyRoute')),
'5': React.lazy(() => import( './EmptyRoute')),
'6': React.lazy(() => import( './EmptyRoute')),
'7': React.lazy(() => import( './EmptyRoute')),
'8': React.lazy(() => import(/* webpackChunkName: "p__Admin__index" */'@/pages/Admin/index.tsx')),
'9': React.lazy(() => import( './EmptyRoute')),
'10': React.lazy(() => import( './EmptyRoute')),
'11': React.lazy(() => import( './EmptyRoute')),
'12': React.lazy(() => import( './EmptyRoute')),
'13': React.lazy(() => import( './EmptyRoute')),
'14': React.lazy(() => import( './EmptyRoute')),
'15': React.lazy(() => import( './EmptyRoute')),
'16': React.lazy(() => import( './EmptyRoute')),
'17': React.lazy(() => import( './EmptyRoute')),
'18': React.lazy(() => import( './EmptyRoute')),
'19': React.lazy(() => import( './EmptyRoute')),
'20': React.lazy(() => import( './EmptyRoute')),
'21': React.lazy(() => import( './EmptyRoute')),
'22': React.lazy(() => import( './EmptyRoute')),
'23': React.lazy(() => import( './EmptyRoute')),
'24': React.lazy(() => import( './EmptyRoute')),
'25': React.lazy(() => import( './EmptyRoute')),
'26': React.lazy(() => import( './EmptyRoute')),
'27': React.lazy(() => import( './EmptyRoute')),
'28': React.lazy(() => import(/* webpackChunkName: "p__Menu__index" */'@/pages/Menu/index.tsx')),
'29': React.lazy(() => import( './EmptyRoute')),
'ant-design-pro-layout': React.lazy(() => import(/* webpackChunkName: "umi__plugin-layout__Layout" */'/Users/return/workspace/mygo/zerocmf-micro/client/antd-admin/src/.umi/plugin-layout/Layout.tsx')),
},
  };
}
