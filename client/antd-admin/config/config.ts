import { defineConfig } from '@umijs/max';
import routes from './routes'
import defaultSettings from './defaultSettings';
import proxy from './proxy';

const { REACT_APP_ENV } = process.env;

export default defineConfig({
  antd: {},
  access: {},
  model: {},
  initialState: {},
  layout: {
    locale: false,
    ...defaultSettings,
  },
  routes,
  request: {},
  styledComponents: {},
  proxy: proxy[REACT_APP_ENV || 'dev'],
  npmClient: 'pnpm',
});

