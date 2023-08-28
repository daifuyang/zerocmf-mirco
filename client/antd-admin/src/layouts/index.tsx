import { useAppData, useSelectedRoutes } from '@umijs/max';
import { history, useModel, useOutlet, useRouteProps } from '@umijs/max';
import { ConfigProvider, Tabs } from 'antd';
import { useEffect } from 'react';

const TabLayout = (props: any) => {
  const routePros = useRouteProps();
  const outlet = useOutlet();
  const { initialState, setInitialState } = useModel('@@initialState');
  // 获取当前路由信息
  const routesArr: any[] = initialState?.tabRoutes || [];

  useEffect(() => {
    const index = routesArr.findIndex((item) => item?.path === routePros?.path);
    if (routePros?.name) {
      let canSet = false;
      const tabRoutes = [...routesArr];
      if (index > -1) {
        if (!tabRoutes[index]?.outlet) {
          canSet = true;
          tabRoutes[index] = { ...tabRoutes[index], outlet };
        }
      } else {
        canSet = true;
        tabRoutes.push({ ...routePros, outlet });
      }

      if (canSet) {
        setInitialState((preInitialState: any) => ({
          ...preInitialState,
          tabRoutes,
        }));
      }
    }
  }, [routesArr, routePros]);

  const defaultPanes = routesArr.map((item) => {
    return {
      label: item.name,
      key: item.path,
      closable: item.closable,
      children: item.outlet,
    };
  });

  return (
    <>
      <ConfigProvider
        theme={{
          components: {
            Tabs: {
              horizontalMargin: '8px 0px',
            },
          },
        }}
      >
        <Tabs
          size="small"
          hideAdd
          activeKey={routePros.path}
          type="editable-card"
          items={defaultPanes}
          onTabClick={(key) => {
            history.push(key);
          }}
          onEdit={(targetKey: any, action) => {
            if (action === 'remove') {
              const tabRoutes: any = [...routesArr];
              const index = tabRoutes.findIndex(
                (item: any) => item.path === targetKey,
              );
              const total = tabRoutes?.length;
              if (targetKey == routePros.path) {
                let redirectPath = '/';
                // 不是最后一个
                if (total - 1 > index) {
                  redirectPath = tabRoutes[index + 1]?.path;
                } else if (total - 1 === index && index > 0) {
                  redirectPath = tabRoutes[index - 1]?.path;
                }
                history.push(redirectPath);
              }

              tabRoutes.splice(index, 1);
              setInitialState((preInitialState: any) => ({
                ...preInitialState,
                tabRoutes,
              }));
            }
          }}
        />
      </ConfigProvider>

      {!routePros?.name && outlet}
    </>
  );
};
export default TabLayout;
