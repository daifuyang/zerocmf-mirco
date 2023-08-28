import { getMenusTree } from '@/services/menu';
import * as icons from '@ant-design/icons';
import { PlusOutlined } from '@ant-design/icons';
import type { ActionType, ProColumns } from '@ant-design/pro-components';
import { ProTable } from '@ant-design/pro-components';
import { Button, Popconfirm } from 'antd';
import { useRef } from 'react';
import Save from './save';

const columns: ProColumns<any>[] = [
  {
    title: '菜单名称',
    dataIndex: 'name',
  },
  {
    title: '图标',
    dataIndex: 'icon',
    hideInSearch: true,
    render: (_, record) => {
      if (record.icon) {
        const Icon = (icons as any)[record.icon];
        return <Icon />;
      }
      return;
    },
  },
  {
    title: '排序',
    dataIndex: 'order',
    hideInSearch: true,
  },
  {
    title: '权限标识',
    dataIndex: 'access',
    hideInSearch: true,
  },
  {
    title: '组件路径',
    dataIndex: 'component',
    hideInSearch: true,
  },
  {
    title: '状态',
    dataIndex: 'status',
    valueType: 'select',
    request: () =>
      Promise.resolve([
        {
          label: '启用',
          value: 1,
        },
        {
          label: '停用',
          value: 0,
        },
      ]),
  },
  {
    title: '创建时间',
    key: 'showTime',
    dataIndex: 'createdAt',
    valueType: 'dateTime',
    renderText: (_, record) => {
      return record.createTime;
    },
    sorter: true,
    hideInSearch: true,
  },
  {
    title: '操作',
    valueType: 'option',
    key: 'option',
    render: (text, record, _, action) => [
      <a key="editable" onClick={() => {}}>
        编辑
      </a>,
      <Popconfirm
        key="delete"
        title="您确定删除吗？"
        //  onConfirm={confirm}
        //  onCancel={cancel}
        okText="确定"
        cancelText="取消"
      >
        <a
          style={{
            color: '#ff4d4f',
          }}
        >
          删除
        </a>
      </Popconfirm>,
    ],
  },
];

export default () => {
  const actionRef = useRef<ActionType>();
  const modalRef = useRef<any>();
  return (
    <>
      <Save ref={modalRef} />
      <ProTable<any>
        columns={columns}
        actionRef={actionRef}
        cardBordered
        expandable={{
          childrenColumnName: 'routes',
          defaultExpandAllRows: true,
          indentSize: 20,
        }}
        request={async (params = {}, sort, filter) => {
          const res = await getMenusTree();
          if (res.code !== 1) {
            return {
              data: [],
              success: false,
            };
          }
          return {
            data: res.data,
            success: true,
          };
        }}
        // editable={{
        //   type: 'multiple',
        // }}
        /*       columnsState={{
        persistenceKey: 'pro-table-singe-demos',
        persistenceType: 'localStorage',
        onChange(value) {
          console.log('value: ', value);
        },
      }} */
        rowKey="id"
        search={{
          labelWidth: 'auto',
        }}
        options={{
          setting: {
            listsHeight: 400,
          },
        }}
        form={{
          // 由于配置了 transform，提交的参与与定义的不同这里需要转化一下
          syncToUrl: (values, type) => {
            if (type === 'get') {
              return {
                ...values,
                created_at: [values.startTime, values.endTime],
              };
            }
            return values;
          },
        }}
        pagination={false}
        headerTitle="菜单列表"
        dateFormatter="number"
        toolBarRender={() => [
          <Button
            key="button"
            icon={<PlusOutlined />}
            onClick={() => {
              modalRef.current.open();
            }}
            type="primary"
          >
            新建
          </Button>,
        ]}
      />
    </>
  );
};
