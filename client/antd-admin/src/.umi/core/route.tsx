// @ts-nocheck
// This file is generated by Umi automatically
// DO NOT CHANGE IT MANUALLY!
import React from 'react';

export async function getRoutes() {
  const routes = {"1":{"path":"/login","layout":false,"id":"1"},"2":{"path":"/","redirect":"/home","parentId":"@@/global-layout","id":"2"},"@@/global-layout":{"id":"@@/global-layout","path":"/","parentId":"ant-design-pro-layout","isLayout":true},"ant-design-pro-layout":{"id":"ant-design-pro-layout","path":"/","isLayout":true}} as const;
  return {
    routes,
    routeComponents: {
'1': React.lazy(() => import(/* webpackChunkName: "p__Login__index" */'@/pages/Login/index.tsx')),
'2': React.lazy(() => import( './EmptyRoute')),
'@@/global-layout': React.lazy(() => import(/* webpackChunkName: "layouts__index" */'/Users/return/workspace/mygo/zerocmf-micro/client/antd-admin/src/layouts/index.tsx')),
'ant-design-pro-layout': React.lazy(() => import(/* webpackChunkName: "umi__plugin-layout__Layout" */'/Users/return/workspace/mygo/zerocmf-micro/client/antd-admin/src/.umi/plugin-layout/Layout.tsx')),
},
  };
}
