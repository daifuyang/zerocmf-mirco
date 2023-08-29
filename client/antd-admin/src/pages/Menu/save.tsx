import { getMenu, getMenusTree, saveMenu } from '@/services/menu';
import { CloseCircleOutlined, CopyOutlined } from '@ant-design/icons';
import {
  ModalForm,
  ProFormDependency,
  ProFormDigit,
  ProFormGroup,
  ProFormList,
  ProFormRadio,
  ProFormSelect,
  ProFormSwitch,
  ProFormText,
  ProFormTreeSelect,
} from '@ant-design/pro-components';
import { Form, message } from 'antd';
import { forwardRef, useImperativeHandle, useState } from 'react';
import { useImmer } from 'use-immer';

const SaveModal = forwardRef((props: any, ref) => {
  const { onFinish } = props;

  const [open, setOpen] = useState(false);
  const [state, setState] = useImmer({
    title: '添加菜单',
  });

  const [form] = Form.useForm<any>();

  useImperativeHandle(ref, () => ({
    open(data: any = {}) {
      const { id, title } = data;
      if (title) {
        setState((prevState) => ({
          ...prevState,
          title,
        }));
      }
      if (id > 0) {
        getMenu(id).then((res) => {
          if (res.code != 1) {
            message.error(res.msg);
            return;
          }
          // message.success(res.msg);
          form.setFieldsValue(res.data);
        });
      }
      setOpen(true);
    },
  }));

  return (
    <ModalForm<any>
      title={state.title}
      form={form}
      autoFocusFirstInput
      open={open}
      modalProps={{
        wrapClassName: 'zero-modal',
        width: 650,
        destroyOnClose: true,
        forceRender: true,
        onCancel: () => {
          setOpen(false);
        },
      }}
      submitTimeout={5000}
      layout="horizontal"
      grid={true}
      rowProps={{
        gutter: [16, 0],
      }}
      onFinish={async (values) => {
        const res = await saveMenu(values);
        if (res.code === 1) {
          if (onFinish) {
            onFinish();
          }
          message.success(res.msg);
          setOpen(false)
          return true;
        }
        message.error(res.msg);
        return false;
      }}
    >
      <ProFormDigit name="id" hidden />
      <ProFormTreeSelect
        name="parentId"
        colProps={{
          span: 24,
        }}
        label="上级菜单"
        fieldProps={{
          fieldNames: {
            label: 'name',
            value: 'id',
            children: 'routes',
          },
          treeDefaultExpandAll: true,
        }}
        request={async () => {
          const res = await getMenusTree();
          if (res.code !== 1) {
            return [];
          }
          return [
            {
              name: '作为顶级菜单显示',
              id: 0,
            },
            ...res.data,
          ];
        }}
      />

      <ProFormRadio.Group
        name="menuType"
        label="菜单类型"
        initialValue={0}
        options={[
          {
            label: '分组',
            value: 0,
          },
          {
            label: '菜单',
            value: 1,
          },
          {
            label: '表单',
            value: 2,
          },
          {
            label: '按钮',
            value: 3,
          },
        ]}
      />

      <ProFormText
        name="name"
        colProps={{
          span: 24,
        }}
        label="菜单名称"
      />

      <ProFormDependency name={['menuType']}>
        {({ menuType }) => {
          const formItem = [];
          if (menuType !== 3) {
            formItem.push(
              <ProFormText
                key="icon"
                name="icon"
                colProps={{
                  span: 12,
                }}
                label="菜单图标"
              />,
            );
          }

          if (menuType < 2) {
            formItem.push(
              <ProFormSwitch
                colProps={{
                  span: 12,
                }}
                key="link"
                name="link"
                label="是否外链"
              />,
            );
          }
          return formItem;
        }}
      </ProFormDependency>

      <ProFormDependency name={['menuType']}>
        {({ menuType }) => {
          const path = (
            <ProFormText
              colProps={{
                span: 12,
              }}
              name="path"
              label="路由地址"
            />
          );

          const access = (
            <ProFormText
              colProps={{
                span: 12,
              }}
              name="access"
              label="权限字符"
            />
          );
          if (menuType === 0) {
            return (
              <>
              {path}
              {access}
              </>
            )
          } else if (menuType === 1) {
            return (
              <>
                {path}
                <ProFormText
                  colProps={{
                    span: 12,
                  }}
                  name="component"
                  label="组件地址"
                />
                {access}
              </>
            );
          } else if (menuType === 2) {
            return (
              <>
                {path}
                <ProFormTreeSelect
                  name="formId"
                  colProps={{
                    span: 24,
                  }}
                  label="关联表单"
                  request={async () => {
                    return await [];
                  }}
                />
                {access}
              </>
            );
          } else if (menuType === 3) {
            return <>{access}</>;
          }
        }}
      </ProFormDependency>

      <ProFormDigit
        name="order"
        initialValue={10000}
        colProps={{
          span: 12,
        }}
        label="显示排序"
      />

      <ProFormDependency name={['menuType']}>
        {({ menuType }) => {
          let render = null;
          let apiRender = null;

          if (menuType !== 2) {
            apiRender = (
              <ProFormList
                name="apiList"
                label="关联api"
                rowProps={{
                  style: {
                    margin: 0,
                  },
                }}
                colProps={{
                  span: 24,
                }}
                initialValue={[]}
                copyIconProps={{
                  Icon: CopyOutlined,
                  tooltipText: '复制此项到末尾',
                }}
                deleteIconProps={{
                  Icon: CloseCircleOutlined,
                  tooltipText: '不需要这行了',
                }}
              >
                <ProFormGroup key="group">
                  <ProFormText
                    colProps={{
                      span: 6,
                    }}
                    labelCol={{
                      span: 24,
                    }}
                    name="name"
                    label="名称"
                  />
                  <ProFormText
                    colProps={{
                      span: 6,
                    }}
                    labelCol={{
                      span: 24,
                    }}
                    name="api"
                    label="接口"
                  />
                  <ProFormSelect
                    colProps={{
                      span: 6,
                    }}
                    labelCol={{
                      span: 24,
                    }}
                    name="method"
                    label="方法"
                    options={[
                      { label: 'GET', value: 'get' },
                      { label: 'POST', value: 'post' },
                      { label: 'PUT', value: 'put' },
                      { label: 'DELETE', value: 'delete' },
                    ]}
                  />
                  <ProFormText
                    colProps={{
                      span: 6,
                    }}
                    labelCol={{
                      span: 24,
                    }}
                    name="desc"
                    label="描述"
                  />
                </ProFormGroup>
              </ProFormList>
            );
          }

          if (menuType !== 3) {
            render = (
              <>
                <ProFormRadio.Group
                  name="hideInMenu"
                  initialValue={1}
                  label="显示状态"
                  colProps={{
                    span: 12,
                  }}
                  options={[
                    {
                      label: '显示',
                      value: 1,
                    },
                    {
                      label: '隐藏',
                      value: 0,
                    },
                  ]}
                />

                <ProFormRadio.Group
                  name="status"
                  initialValue={1}
                  label="菜单状态"
                  colProps={{
                    span: 12,
                  }}
                  options={[
                    {
                      label: '启用',
                      value: 1,
                    },
                    {
                      label: '停用',
                      value: 0,
                    },
                  ]}
                />
              </>
            );
          }

          return (
            <>
              {render}
              {apiRender}
            </>
          );
        }}
      </ProFormDependency>
    </ModalForm>
  );
});

export default SaveModal;
