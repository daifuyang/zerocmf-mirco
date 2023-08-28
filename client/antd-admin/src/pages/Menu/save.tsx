import {
  ModalForm,
  ProFormDigit,
  ProFormRadio,
  ProFormSwitch,
  ProFormText,
  ProFormTreeSelect,
} from '@ant-design/pro-components';
import { Form, message } from 'antd';
import { forwardRef, useImperativeHandle, useState } from 'react';
import { useImmer } from 'use-immer';

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

const SaveModal = forwardRef((props, ref) => {
  const [open, setOpen] = useState(false);
  const [state, setState] = useImmer({
    title: '添加菜单',
    treeData: [
      {
        name: '作为顶级菜单',
        id: '0',
      },
    ],
  });

  const [form] = Form.useForm<{ name: string; company: string }>();

  useImperativeHandle(ref, () => ({
    open(data: any = {}) {
      const { title } = data;
      if (title) {
        setState((prevState) => ({
          ...prevState,
          title,
        }));
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
        width: 520,
        destroyOnClose: true,
        onCancel: () => {
          setOpen(false)
        },
      }}
      submitTimeout={30000}
      layout="horizontal"
      grid={true}
      rowProps={{
        gutter: [16, 0],
      }}
      onFinish={async (values) => {
        await waitTime(2000);
        console.log(values.name);
        message.success('提交成功');
        return true;
      }}
    >
      <ProFormTreeSelect
        name="parentId"
        colProps={{
          span: 24,
        }}
        label="上级菜单"
        request={ async () => {
          
        } }
      />

      <ProFormRadio.Group
        name="menuType"
        label="菜单类型"
        initialValue="group"
        options={[
          {
            label: '分组',
            value: 'group',
          },
          {
            label: '菜单',
            value: 'menu',
          },
          {
            label: '表单',
            value: 'form',
          },
          {
            label: '按钮',
            value: 'button',
          },
        ]}
      />

      <ProFormText
        name="name"
        colProps={{
          span: 12,
        }}
        label="菜单名称"
      />

      <ProFormDigit
        name="order"
        colProps={{
          span: 12,
        }}
        label="显示排序"
      />

      <ProFormText
        colProps={{
          span: 12,
        }}
        name="path"
        label="路由地址"
      />

      <ProFormText
        colProps={{
          span: 12,
        }}
        name="access"
        label="权限字符"
      />

      <ProFormSwitch
        colProps={{
          span: 12,
        }}
        name="blank"
        label="是否外链"
      />

      <ProFormSwitch
        colProps={{
          span: 12,
        }}
        name="status"
        label="显示状态"
        initialValue={true}
      />
    </ModalForm>
  );
});

export default SaveModal;
