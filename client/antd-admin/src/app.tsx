import NotFound from '@/pages/404';
import * as icons from '@ant-design/icons';
import type { Settings as LayoutSettings } from '@ant-design/pro-components';
import { ProBreadcrumb } from '@ant-design/pro-components';
import type { RunTimeLayoutConfig } from '@umijs/max';
import { RequestConfig, createGlobalStyle, history } from '@umijs/max';
import { Spin, Tooltip, message, notification } from 'antd';
import defaultSettings from '../config/defaultSettings';
import {
  AvatarDropdown,
  AvatarName,
} from './components/RightContent/AvatarDropdown';

import Footer from '@/components/Footer';

import { currentUser as queryCurrentUser } from '@/services/user';
import { AppstoreOutlined } from '@ant-design/icons';

import { Navigate } from '@umijs/max';
import React from 'react';
import { getMenusTree } from './services/menu';
import styles from './style.less';

const codeMessage: any = {
  200: '服务器成功返回请求的数据。',
  201: '新建或修改数据成功。',
  202: '一个请求已经进入后台排队（异步任务）。',
  204: '删除数据成功。',
  400: '发出的请求有错误，服务器没有进行新建或修改数据的操作。',
  401: '用户没有权限（令牌、用户名、密码错误）。',
  403: '用户得到授权，但是访问是被禁止的。',
  404: '发出的请求针对的是不存在的记录，服务器没有进行操作。',
  406: '请求的格式不可得。',
  410: '请求的资源被永久删除，且不会再得到的。',
  422: '当创建一个对象时，发生一个验证错误。',
  500: '服务器发生错误，请检查服务器。',
  502: '网关错误。',
  503: '服务不可用，服务器暂时过载或维护。',
  504: '网关超时。',
};

const loginPath = '/login';

export const styledComponents = {
  GlobalStyle: createGlobalStyle`
    .ant-pro-form-login-container  {
      padding-top: 100px
    }
  `,
};

function processRoutes(routes = []) {
  routes.forEach((route: any) => {
    let Component: React.ComponentType<any> | null = route.element || null;
    let Icon: React.ComponentType<any> | null = route.element || null;
    if (route.component) {
      let { component } = route;
      if (component.startsWith('/pages')) {
        component = component.replace('/pages', '');
      }
      const pattern = /\.+.*$/;
      if (!pattern.test(component)) {
        component = component.replace(/\/+$/, '') + '/index.tsx';
      }
      Component = React.lazy(() => import(`./pages${component}`));
      route.element = (
        <React.Suspense
          fallback={
            <Spin spinning={true}>
              <div style={{ width: '100%', paddingTop: '30%' }}></div>
            </Spin>
          }
        >
          {Component && <Component />}
        </React.Suspense>
      );
    }

    if (route.icon) {
      Icon = (icons as any)[route.icon];
      if (Icon) {
        route.icon = <Icon />;
      }
    }

    if (route.routes?.length > 0) {
      route.routes.unshift({
        path: route.path,
        element: <Navigate to={route.routes[0].path} replace />,
      });
      route.children = route.routes;
      processRoutes(route.routes);
    }
  });
}

let extraRoutes: any;
export function patchClientRoutes({ routes }: any) {
  // 根据 extraRoutes 对 routes 做一些修改
  const routerIndex = routes.findIndex((item: any) => item.path === '/');
  const childRoutes = routes[routerIndex].children;
  const childIndex =
    childRoutes.findIndex((item: any) => item.path === '/') || 0;
  if (!childRoutes[childIndex].children) {
    childRoutes[childIndex].children = [];
  }

  processRoutes(extraRoutes);
  if (extraRoutes) {
    extraRoutes.push({
      path: '/*',
      element: <NotFound />,
    });

    childRoutes[childIndex].children.push(...extraRoutes);
  }
}

export function render(oldRender: any) {
  getInitialState().then((res) => {
    extraRoutes = res.menus;
    oldRender();
  });
}

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: any;
  tabRoutes: any[];
  menus: any[];
  fetchUserInfo?: () => Promise<any | undefined>;
  fetchMenusTree?: () => Promise<any | undefined>;
}> {
  const fetchUserInfo = async () => {
    const res = await queryCurrentUser();
    if (res.code != 1) {
      message.error('用户身份已失效!');
      history.push(loginPath);
      return undefined;
    }
    return res.data;
  };

  const fetchMenusTree = async () => {
    const res = await getMenusTree();
    if (res.code != 1) {
      message.error('用户身份已失效!');
      history.push(loginPath);
      return undefined;
    }
    return res.data;
  };

  const tabRoutes = [
    {
      name: '首页',
      path: '/home',
      closable: false,
    },
  ];

  // 如果不是登录页面，执行
  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    const menus = await fetchMenusTree();
    return {
      fetchUserInfo,
      fetchMenusTree,
      currentUser,
      tabRoutes,
      settings: defaultSettings as Partial<LayoutSettings>,
      menus,
    };
  }

  return {
    fetchUserInfo,
    fetchMenusTree,
    tabRoutes,
    settings: defaultSettings as Partial<LayoutSettings>,
    menus: [],
  };
}

export const layout: RunTimeLayoutConfig = ({
  initialState,
  setInitialState,
}) => {
  return {
    logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
    /* menu: {
      locale: false,
      params: {
        userId: initialState?.currentUser?.userId,
      },
      request: async (params, defaultMenuData) => {
        // initialState.currentUser 中包含了所有用户信息
        const res = await getMenusTree()
        if(res.code != 1) {
          message.error(res.msg)
          return [];
        }

        const list = treeToList(res.data,'routes')
        setInitialState((prevState: any) => ({ ...prevState, menus: list }));
        return res.data;
      },
    }, */
    avatarProps: {
      src: 'https://gw.alipayobjects.com/zos/antfincdn/efFD%24IOql2/weixintupian_20170331104822.jpg',
      size: 'small',
      title: <AvatarName />,
      render: (_, avatarChildren) => {
        return <AvatarDropdown>{avatarChildren}</AvatarDropdown>;
      },
    },
    style: {
      height: '100vh',
    },
    headerContentRender: () => {
      return (
        <div style={{ display: 'flex' }}>
          <div
            onClick={() => {
              history.push('/admin/menu/list');
            }}
            style={{
              cursor: 'pointer',
              fontSize: '20px',
              color: '#fff',
            }}
          >
            <Tooltip title="菜单管理">
              <AppstoreOutlined />
            </Tooltip>
          </div>
          <div className={styles.breadcrumb}>
            <ProBreadcrumb />
          </div>
        </div>
      );
    },
    footerRender: () => <Footer />,
    childrenRender: (children) => {
      return (
        <>
          {children}
          {/* <SettingDrawer
            disableUrlParams
            enableDarkTheme
            settings={initialState?.settings}
            onSettingChange={(settings) => {
              setInitialState((preInitialState: any) => ({
                ...preInitialState,
                settings,
              }));
            }}
          /> */}
        </>
      );
    },
    ...initialState?.settings,
  };
};

const authHeaderInterceptor = (
  url: string,
  options: RequestConfig & { token: string },
) => {
  if (options.token) {
    const token: any = localStorage.getItem('token');
    if (token) {
      const data = JSON.parse(token);
      options.headers = {
        Authorization: `Bearer ${data.access_token}`,
      };
    } else {
      history.push(loginPath);
    }
  }

  return {
    url,
    options,
  };
};

export const request: RequestConfig = {
  requestInterceptors: [authHeaderInterceptor],
  errorConfig: {
    errorHandler(error: any, opts: any) {
      const { response } = error;
      if (response && response.status) {
        const errorText = codeMessage[response.status] || response.statusText;
        const { status, url } = response;

        if (status == 401) {
          history.push(loginPath);
        }

        notification.error({
          message: `请求错误 ${status}: ${url}`,
          description: errorText,
        });
      } else if (!response) {
        notification.error({
          description: '您的网络发生异常，无法连接服务器',
          message: '网络异常',
        });
      }

      return response;
    },
  },
};
