import { treeToList } from "./utils/util";

export default (initialState: any) => {
  // 在这里按照初始化数据定义项目中的权限，统一管理
  // 参考文档 https://umijs.org/docs/max/access

  const access:any = {}

  const list = treeToList(initialState?.menus,'routes')
  list?.forEach( (item:any) => {
    access[item.access] = true
  } )  

  return {
    ...access,
    can: (str: string) => 
     {
      return access[str] ? true : false
     }
  };
};
