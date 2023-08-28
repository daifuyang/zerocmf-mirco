// treeToList
export const treeToList = (array = [], key = 'children') => {
    const list:any = [];
    array.forEach((item) => {
      list.push(item);
      if (item[key]) {
        const children = treeToList(item[key], key);
        list.push(...children);
      }
    });
    return list;
  };
  