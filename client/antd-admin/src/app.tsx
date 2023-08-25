import Footer from '@/components/Footer';
import type { Settings as LayoutSettings } from '@ant-design/pro-components';
import { SettingDrawer } from '@ant-design/pro-components';
import type { RunTimeLayoutConfig } from '@umijs/max';
import { RequestConfig, createGlobalStyle, history } from '@umijs/max';
import { message, notification } from 'antd';
import defaultSettings from '../config/defaultSettings';
import {
  AvatarDropdown,
  AvatarName,
} from './components/RightContent/AvatarDropdown';

import { currentUser as queryCurrentUser } from '@/services/user';

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

// 全局初始化数据配置，用于 Layout 用户信息和权限初始化
// 更多信息见文档：https://umijs.org/docs/api/runtime-config#getinitialstate
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: any;
  fetchUserInfo?: () => Promise<any | undefined>;
}> {
  const fetchUserInfo = async () => {
    try {
      const res = await queryCurrentUser();
      if (res.code != 1) {
        message.error('用户身份已失效!');
        history.push(loginPath);
        return;
      }
      return res.data;
    } catch (error) {
      history.push(loginPath);
    }
    return undefined;
  };

  // 如果不是登录页面，执行
  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    console.log('currentUser', currentUser)
    return {
      fetchUserInfo,
      currentUser,
      settings: defaultSettings as Partial<LayoutSettings>,
    };
  }

  return {
    fetchUserInfo,
    settings: defaultSettings as Partial<LayoutSettings>,
  };
}

export const layout: RunTimeLayoutConfig = ({
  initialState,
  setInitialState,
}) => {
  return {
    logo: 'https://img.alicdn.com/tfs/TB1YHEpwUT1gK0jSZFhXXaAtVXa-28-27.svg',
    menu: {
      locale: false,
    },
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
    footerRender: () => <Footer />,
    childrenRender: (children) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
          <SettingDrawer
            disableUrlParams
            enableDarkTheme
            settings={initialState?.settings}
            onSettingChange={(settings) => {
              setInitialState((preInitialState) => ({
                ...preInitialState,
                settings,
              }));
            }}
          />
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
